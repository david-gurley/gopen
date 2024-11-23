module github.com/david-gurley/gopen

go 1.18

require (
	github.com/PaloAltoNetworks/pango v0.8.0
	github.com/gocarina/gocsv v0.0.0-20220823132111-71f3a5cb2654
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.4.2
	github.com/pensando/go-psm/psm_dss v0.0.0-00010101000000-000000000000
	github.com/pensando/go-psm/psm_ent v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2
	google.golang.org/grpc v1.37.1
)

require (
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/appengine v1.4.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/pensando/go-psm/psm_dss => ./psm/psm_dss

replace github.com/pensando/go-psm/psm_ent => ./psm/psm_ent
