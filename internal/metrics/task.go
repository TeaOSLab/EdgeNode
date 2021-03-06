// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package metrics

import (
	"database/sql"
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeNode/internal/remotelogs"
	"github.com/TeaOSLab/EdgeNode/internal/rpc"
	"github.com/TeaOSLab/EdgeNode/internal/utils"
	"github.com/iwind/TeaGo/Tea"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Task 单个指标任务
// 数据库存储：
//  data/
//     metric.$ID.db
//        stats
//           id, keys, value, time, serverId, hash
//  原理：
//     添加或者有变更时 isUploaded = false
//     上传时检查 isUploaded 状态
//     只上传每个服务中排序最前面的 N 个数据
type Task struct {
	item     *serverconfigs.MetricItemConfig
	isLoaded bool

	db            *sql.DB
	statTableName string
	statsChan     chan *Stat
	isStopped     bool

	cleanTicker  *utils.Ticker
	uploadTicker *utils.Ticker

	cleanVersion int32

	insertStatStmt          *sql.Stmt
	deleteByVersionStmt     *sql.Stmt
	deleteByExpiresTimeStmt *sql.Stmt
	selectTopStmt           *sql.Stmt
	sumStmt                 *sql.Stmt

	serverIdMap       map[int64]bool  // 所有的服务Ids
	timeMap           map[string]bool // time => bool
	serverIdMapLocker sync.Mutex
}

// NewTask 获取新任务
func NewTask(item *serverconfigs.MetricItemConfig) *Task {
	return &Task{
		item:        item,
		statsChan:   make(chan *Stat, 40960),
		serverIdMap: map[int64]bool{},
		timeMap:     map[string]bool{},
	}
}

// Init 初始化
func (this *Task) Init() error {
	this.statTableName = "stats"

	// 检查目录是否存在
	var dir = Tea.Root + "/data"
	_, err := os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}
		remotelogs.Println("METRIC", "create data dir '"+dir+"'")
	}

	db, err := sql.Open("sqlite3", "file:"+dir+"/metric."+strconv.FormatInt(this.item.Id, 10)+".db?cache=shared&mode=rwc&_journal_mode=WAL")
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(1)
	this.db = db

	//创建统计表
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "` + this.statTableName + `" (
  "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  "hash" varchar(32),
  "keys" varchar(1024),
  "value" real DEFAULT 0,
  "time" varchar(32),
  "serverId" integer DEFAULT 0,
  "version" integer DEFAULT 0,
  "isUploaded" integer DEFAULT 0
);

CREATE INDEX IF NOT EXISTS "serverId"
ON "` + this.statTableName + `" (
  "serverId" ASC,
  "version" ASC
);

CREATE UNIQUE INDEX IF NOT EXISTS "hash"
ON "` + this.statTableName + `" (
  "hash" ASC
);`)
	if err != nil {
		return err
	}

	// insert stat stmt
	this.insertStatStmt, err = db.Prepare(`INSERT INTO "stats" ("serverId", "hash", "keys", "value", "time", "version", "isUploaded") VALUES (?, ?, ?, ?, ?, ?, 0) ON CONFLICT("hash") DO UPDATE SET "value"="value"+?, "isUploaded"=0`)
	if err != nil {
		return err
	}

	// delete by version
	this.deleteByVersionStmt, err = db.Prepare(`DELETE FROM "` + this.statTableName + `" WHERE "version"<?`)
	if err != nil {
		return err
	}

	// delete by expires time
	this.deleteByExpiresTimeStmt, err = db.Prepare(`DELETE FROM "` + this.statTableName + `" WHERE "time"<?`)
	if err != nil {
		return err
	}

	// select topN stmt
	this.selectTopStmt, err = db.Prepare(`SELECT "id", "hash", "keys", "value", "isUploaded" FROM "` + this.statTableName + `" WHERE "serverId"=? AND "version"=? AND time=? ORDER BY "value" DESC LIMIT 100`)
	if err != nil {
		return err
	}

	// sum stmt
	this.sumStmt, err = db.Prepare(`SELECT COUNT(*), IFNULL(SUM(value), 0) FROM "` + this.statTableName + `" WHERE "serverId"=? AND "version"=? AND time=?`)
	if err != nil {
		return err
	}

	// 所有的服务IDs
	err = this.loadServerIdMap()
	if err != nil {
		return err
	}

	this.isLoaded = true

	return nil
}

// Start 启动任务
func (this *Task) Start() error {
	// 读取数据
	go func() {
		for stat := range this.statsChan {
			if stat == nil {
				return
			}
			err := this.InsertStat(stat)
			if err != nil {
				remotelogs.Error("METRIC", "insert stat failed: "+err.Error())
			}
		}
	}()

	// 清理
	this.cleanTicker = utils.NewTicker(24 * time.Hour)
	go func() {
		for this.cleanTicker.Next() {
			err := this.CleanExpired()
			if err != nil {
				remotelogs.Error("METRIC", "clean expired stats failed: "+err.Error())
			}
		}
	}()

	// 上传
	this.uploadTicker = utils.NewTicker(this.item.UploadDuration())
	go func() {
		for this.uploadTicker.Next() {
			err := this.Upload(1 * time.Second)
			if err != nil {
				remotelogs.Error("METRIC", "upload stats failed: "+err.Error())
			}
		}
	}()

	return nil
}

// Add 添加数据
func (this *Task) Add(obj MetricInterface) {
	if this.isStopped || !this.isLoaded {
		return
	}

	var keys = []string{}
	for _, key := range this.item.Keys {
		k := obj.MetricKey(key)
		keys = append(keys, k)
	}

	v, ok := obj.MetricValue(this.item.Value)
	if !ok {
		return
	}

	var stat = &Stat{
		ServerId: obj.MetricServerId(),
		Keys:     keys,
		Value:    v,
		Time:     this.item.CurrentTime(),
	}

	select {
	case this.statsChan <- stat:
	default:
		// 丢弃
	}
}

// Stop 停止任务
func (this *Task) Stop() error {
	this.isStopped = true

	if this.cleanTicker != nil {
		this.cleanTicker.Stop()
	}
	if this.uploadTicker != nil {
		this.uploadTicker.Stop()
	}

	_ = this.insertStatStmt.Close()
	_ = this.deleteByVersionStmt.Close()
	_ = this.deleteByExpiresTimeStmt.Close()
	_ = this.selectTopStmt.Close()
	_ = this.sumStmt.Close()

	if this.statsChan != nil {
		go func() {
			// 延时关闭，防止关闭时写入
			time.Sleep(5 * time.Second)
			close(this.statsChan)
		}()
	}

	if this.db != nil {
		_ = this.db.Close()
	}

	return nil
}

// InsertStat 写入数据
func (this *Task) InsertStat(stat *Stat) error {
	if this.isStopped {
		return nil
	}
	if stat == nil {
		return nil
	}

	this.serverIdMapLocker.Lock()
	this.serverIdMap[stat.ServerId] = true
	this.timeMap[stat.Time] = true
	this.serverIdMapLocker.Unlock()

	keyData, err := json.Marshal(stat.Keys)
	if err != nil {
		return err
	}
	stat.keysData = keyData
	stat.Sum(this.item.Version, this.item.Id)

	_, err = this.insertStatStmt.Exec(stat.ServerId, stat.Hash, stat.keysData, stat.Value, stat.Time, this.item.Version, stat.Value)
	if err != nil {
		return err
	}
	return nil
}

// CleanExpired 清理数据
func (this *Task) CleanExpired() error {
	if this.isStopped {
		return nil
	}

	// 清除低版本数据
	if this.cleanVersion < this.item.Version {
		_, err := this.deleteByVersionStmt.Exec(this.item.Version)
		if err != nil {
			return err
		}
		this.cleanVersion = this.item.Version
	}

	// 清除过期的数据
	_, err := this.deleteByExpiresTimeStmt.Exec(this.item.LocalExpiresTime())
	if err != nil {
		return err
	}

	return nil
}

// Upload 上传数据
func (this *Task) Upload(pauseDuration time.Duration) error {
	if this.isStopped {
		return nil
	}

	this.serverIdMapLocker.Lock()

	// 服务IDs
	var serverIds []int64
	for serverId := range this.serverIdMap {
		serverIds = append(serverIds, serverId)
	}
	this.serverIdMap = map[int64]bool{} // 清空数据

	// 时间
	var times = []string{}
	for t := range this.timeMap {
		times = append(times, t)
	}
	this.timeMap = map[string]bool{} // 清空数据

	this.serverIdMapLocker.Unlock()

	rpcClient, err := rpc.SharedRPC()
	if err != nil {
		return err
	}

	for _, serverId := range serverIds {
		for _, currentTime := range times {
			idStrings, err := func(serverId int64, currentTime string) (ids []string, err error) {
				rows, err := this.selectTopStmt.Query(serverId, this.item.Version, currentTime)
				if err != nil {
					return nil, err
				}
				var isClosed bool
				defer func() {
					if isClosed {
						return
					}
					_ = rows.Close()
				}()

				var pbStats []*pb.UploadingMetricStat
				for rows.Next() {
					var pbStat = &pb.UploadingMetricStat{
					}
					// "id", "hash", "keys", "value", "isUploaded"
					var isUploaded int
					var keysData []byte
					err = rows.Scan(&pbStat.Id, &pbStat.Hash, &keysData, &pbStat.Value, &isUploaded)
					if err != nil {
						return nil, err
					}
					if isUploaded == 1 {
						continue
					}
					if len(keysData) > 0 {
						err = json.Unmarshal(keysData, &pbStat.Keys)
						if err != nil {
							return nil, err
						}
					}
					pbStats = append(pbStats, pbStat)
					ids = append(ids, strconv.FormatInt(pbStat.Id, 10))
				}

				// 提前关闭
				_ = rows.Close()
				isClosed = true

				// 上传
				if len(pbStats) > 0 {
					// 计算总和
					count, total, err := this.sum(serverId, currentTime)
					if err != nil {
						return nil, err
					}

					_, err = rpcClient.MetricStatRPC().UploadMetricStats(rpcClient.Context(), &pb.UploadMetricStatsRequest{
						MetricStats: pbStats,
						Time:        currentTime,
						ServerId:    serverId,
						ItemId:      this.item.Id,
						Version:     this.item.Version,
						Count:       count,
						Total:       float32(total),
					})
					if err != nil {
						return nil, err
					}
				}

				return
			}(serverId, currentTime)
			if err != nil {
				return err
			}

			if len(idStrings) > 0 {
				// 设置为已上传
				_, err = this.db.Exec(`UPDATE "` + this.statTableName + `" SET isUploaded=1 WHERE id IN (` + strings.Join(idStrings, ",") + `)`)
				if err != nil {
					return err
				}
			}
		}

		// 休息一下，防止短时间内上传数据过多
		if pauseDuration > 0 {
			time.Sleep(pauseDuration)
		}
	}

	return nil
}

// 加载服务ID
func (this *Task) loadServerIdMap() error {
	{
		rows, err := this.db.Query(`SELECT DISTINCT "serverId" FROM `+this.statTableName+" WHERE version=?", this.item.Version)
		if err != nil {
			return err
		}
		defer func() {
			_ = rows.Close()
		}()

		var serverId int64
		for rows.Next() {
			err = rows.Scan(&serverId)
			if err != nil {
				return err
			}
			this.serverIdMapLocker.Lock()
			this.serverIdMap[serverId] = true
			this.serverIdMapLocker.Unlock()
		}
	}

	{
		rows, err := this.db.Query(`SELECT DISTINCT "time" FROM `+this.statTableName+" WHERE version=?", this.item.Version)
		if err != nil {
			return err
		}
		defer func() {
			_ = rows.Close()
		}()

		var timeString string
		for rows.Next() {
			err = rows.Scan(&timeString)
			if err != nil {
				return err
			}
			this.serverIdMapLocker.Lock()
			this.timeMap[timeString] = true
			this.serverIdMapLocker.Unlock()
		}
	}

	return nil
}

// 计算数量和综合
func (this *Task) sum(serverId int64, time string) (count int64, total float64, err error) {
	rows, err := this.sumStmt.Query(serverId, this.item.Version, time)
	if err != nil {
		return 0, 0, err
	}
	defer func() {
		_ = rows.Close()
	}()
	if rows.Next() {
		err = rows.Scan(&count, &total)
		if err != nil {
			return 0, 0, err
		}
	}
	return
}
