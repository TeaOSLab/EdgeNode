// Copyright 2023 GoEdge goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package fasttime_test

import (
	"testing"
	"time"

	"github.com/TeaOSLab/EdgeNode/internal/utils/fasttime"
	timeutil "github.com/iwind/TeaGo/utils/time"
)

func TestFastTime_Unix(t *testing.T) {
	for i := 0; i < 5; i++ {
		var now = fasttime.Now()
		t.Log(now.Unix(), now.UnixMilli(), "real:", time.Now().Unix())
		time.Sleep(1 * time.Second)
	}
}

func TestFastTime_UnixMilli(t *testing.T) {
	t.Log(fasttime.Now().UnixMilliString())
}

func TestFastTime_UnixFloor(t *testing.T) {
	var now = fasttime.Now()

	var timestamp = time.Now().Unix()
	t.Log("floor 60:", timestamp, now.UnixFloor(60), timeutil.FormatTime("Y-m-d H:i:s", now.UnixFloor(60)))
	t.Log("ceil 60:", timestamp, now.UnixCell(60), timeutil.FormatTime("Y-m-d H:i:s", now.UnixCell(60)))
	t.Log("floor 300:", timestamp, now.UnixFloor(300), timeutil.FormatTime("Y-m-d H:i:s", now.UnixFloor(300)))
	t.Log("next minute:", now.UnixNextMinute(), timeutil.FormatTime("Y-m-d H:i:s", now.UnixNextMinute()))
	t.Log("day:", now.Ymd())
	t.Log("round 5 minute:", now.Round5Hi())
}

func TestFastTime_Format(t *testing.T) {
	var now = fasttime.Now()
	t.Log(now.Format("Y-m-d H:i:s"))
}

func TestFastTime_Hour(t *testing.T) {
	var now = fasttime.Now()
	t.Log(now.Hour())
}

func BenchmarkNewFastTime(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var now = fasttime.Now()
			_ = now.Ymd()
		}
	})
}

func BenchmarkNewFastTime_Raw(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var now = time.Now()
			_ = timeutil.Format("Ymd", now)
		}
	})
}
