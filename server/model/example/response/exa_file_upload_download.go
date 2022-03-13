package response

import "github.com/tmnhs/admin-mall/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
