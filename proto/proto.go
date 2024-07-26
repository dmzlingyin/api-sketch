package proto

import (
	"go-micro.dev/v4"
)

const (
	// ServiceName is the name of this service.
	ServiceName = "service-sketch"
)

var (
	Service micro.Service
)

func init() {
	Service = micro.NewService()
	opts := []micro.Option{
		micro.Name(ServiceName),
	}
	Service.Init(opts...)
}
