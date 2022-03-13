package response

import "github.com/tmnhs/admin-mall/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
