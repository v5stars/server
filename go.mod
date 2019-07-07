module github.com/v5stars/server

go 1.13

require (
	github.com/golang/protobuf v1.3.1
	github.com/golang/protobuf/v2 v2.0.0-20181127193627-d7e97bc71bcb // indirect
	github.com/jinzhu/gorm v1.9.8 // indirect
	github.com/jinzhu/inflection v0.0.0-20190603042836-f5c5f50e6090 // indirect
	google.golang.org/grpc v1.22.0
	v2ray.com/core v0.0.0-20190603071532-16e9d39fff74
)

replace v2ray.com/core v4.19.1+incompatible => ./grpc/v2ray.com/core
