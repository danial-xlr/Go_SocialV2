module profile

go 1.16

require (
	common v0.0.0-00010101000000-000000000000
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.2 // indirect
	google.golang.org/grpc v1.40.0

)

replace common => ../common/
