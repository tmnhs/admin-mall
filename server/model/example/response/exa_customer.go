package response

import "github.com/tmnhs/admin-mall/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
