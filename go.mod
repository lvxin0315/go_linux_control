module github.com/lvxin0315/go_linux_control

go 1.12

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.51.0
	cloud.google.com/go/bigquery => github.com/googleapis/google-cloud-go/bigquery v1.3.0
	cloud.google.com/go/datastore => github.com/googleapis/google-cloud-go/datastore v1.0.0
	cloud.google.com/go/pubsub => github.com/googleapis/google-cloud-go/pubsub v1.1.0
	cloud.google.com/go/storage => github.com/googleapis/google-cloud-go/storage v1.5.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190927123631-a832865fa7ad
	golang.org/x/exp => github.com/golang/exp v0.0.0-20191129062945-2f5052295587
	golang.org/x/image => github.com/golang/image v0.0.0-20191214001246-9130b4cfad52
	golang.org/x/lint => github.com/golang/lint v0.0.0-20191125180803-fdd1cda4f05f
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20200113190325-b9f03b3fa33c
	golang.org/x/mod => github.com/golang/mod v0.2.0
	golang.org/x/net => github.com/golang/net v0.0.0-20190926025831-c00fd9afed17
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20191202225959-858c2ad4c8b6
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190927073244-c990c680b611
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20191024005414-555d28b269f0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190929041059-e7abfedfabcf
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.15.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.5
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20200113173426-e1de0a7b01eb
	google.golang.org/grpc => github.com/grpc/grpc-go v1.26.0
)

require (
	github.com/Unknwon/goconfig v0.0.0-20191126170842-860a72fb44fd
	github.com/gin-gonic/gin v1.5.0
	github.com/jinzhu/gorm v1.9.11
	github.com/kr/pretty v0.2.0 // indirect
	github.com/nats-io/gnatsd v1.4.1 // indirect
	github.com/nats-io/go-nats v1.7.2
	github.com/nats-io/nkeys v0.1.3 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v1.6.4 // indirect
)
