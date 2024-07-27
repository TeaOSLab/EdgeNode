module github.com/TeaOSLab/EdgeNode

go 1.22

replace github.com/TeaOSLab/EdgeCommon => ../EdgeCommon

require (
	github.com/aliyun/aliyun-oss-go-sdk v3.0.2+incompatible
	github.com/andybalholm/brotli v1.0.5
	github.com/aws/aws-sdk-go v1.44.279
	github.com/baidubce/bce-sdk-go v0.9.170
	github.com/biessek/golang-ico v0.0.0-20180326222316-d348d9ea4670
	github.com/cespare/xxhash/v2 v2.3.0
	github.com/coreos/go-iptables v0.6.0
	github.com/cockroachdb/pebble v1.1.0
	github.com/dchest/captcha v1.0.0
	github.com/florianl/go-nfqueue v1.3.1
	github.com/fsnotify/fsnotify v1.7.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/google/gopacket v1.1.19
	github.com/google/nftables v0.2.0
	github.com/huaweicloud/huaweicloud-sdk-go-obs v3.23.4+incompatible
	github.com/iwind/TeaGo v0.0.0-20240508072741-7647e70b7070
	github.com/iwind/gofcgi v0.0.0-20210528023741-a92711d45f11
	github.com/iwind/gosock v0.0.0-20220505115348-f88412125a62
	github.com/iwind/gowebp v0.0.0-20240109104518-489f3429f5c5
	github.com/klauspost/compress v1.17.8
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/mdlayher/netlink v1.7.2
	github.com/miekg/dns v1.1.59
	github.com/mssola/useragent v1.0.0
	github.com/panjf2000/ants/v2 v2.7.5
	github.com/pires/go-proxyproto v0.6.1
	github.com/qiniu/go-sdk/v7 v7.16.0
	github.com/quic-go/quic-go v0.44.0
	github.com/shirou/gopsutil/v3 v3.24.2
	github.com/tdewolff/minify/v2 v2.20.24
	github.com/tencentyun/cos-go-sdk-v5 v0.7.41
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842
	golang.org/x/image v0.16.0
	golang.org/x/net v0.25.0
	golang.org/x/sys v0.20.0
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.33.0
	gopkg.in/yaml.v3 v3.0.1
	rogchap.com/v8go v0.8.0
)

require (
	github.com/DataDog/zstd v1.5.6-0.20230824185856-869dae002e5e // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/cockroachdb/errors v1.11.1 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/cockroachdb/tokenbucket v0.0.0-20230807174530-cc333fc44b06 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/getsentry/sentry-go v0.27.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/pprof v0.0.0-20240509144519-723abb6459b7 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/native v1.1.0 // indirect
	github.com/jsummers/gobmp v0.0.0-20151104160322-e2ba15ffa76e // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/mdlayher/socket v0.5.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mozillazg/go-httpheader v0.2.1 // indirect
	github.com/onsi/ginkgo/v2 v2.17.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/prometheus/client_golang v1.19.0 // indirect
	github.com/prometheus/client_model v0.6.0 // indirect
	github.com/prometheus/common v0.51.0 // indirect
	github.com/prometheus/procfs v0.13.0 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/tdewolff/parse/v2 v2.7.14 // indirect
	github.com/tklauser/go-sysconf v0.3.13 // indirect
	github.com/tklauser/numcpus v0.7.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.uber.org/mock v0.4.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240415180920-8c6c420018be // indirect
)
