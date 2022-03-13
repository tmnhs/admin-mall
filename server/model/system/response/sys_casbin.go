package response

import (
	"github.com/tmnhs/admin-mall/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
