package router

import (
	"github.com/tmnhs/admin-mall/server/router/autocode"
	"github.com/tmnhs/admin-mall/server/router/example"
	"github.com/tmnhs/admin-mall/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
