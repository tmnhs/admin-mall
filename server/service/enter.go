package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/mall"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
	MallServiceGroup     mall.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
