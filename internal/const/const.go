package teaconst

const (
	Version = "1.4.1"

	ProductName = "Edge Node"
	ProcessName = "edge-node"

	Role = "node"

	EncryptMethod = "aes-256-cfb"

	// SystemdServiceName systemd
	SystemdServiceName = "edge-node"

	AccessLogSockName    = "edge-node.accesslog"
	CacheGarbageSockName = "edge-node.cache.garbage"

	EnableKVCacheStore = true // determine store cache keys in KVStore or sqlite
)
