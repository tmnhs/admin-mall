package service

import (
	"github.com/tmnhs/admin-mall/server/service/autocode"
	"github.com/tmnhs/admin-mall/server/service/example"
	"github.com/tmnhs/admin-mall/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
