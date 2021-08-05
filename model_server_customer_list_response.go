/*
 * Fluidcoins Pay api documentation
 *
 * A suite of APIs to create, manage and accept cryptocurrency payments in Africa
 *
 * API version: 0.1.0
 * Contact: developers@fluidcoins.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ServerCustomerListResponse struct {
	Customers []MerchantMerchantCustomer `json:"customers,omitempty"`
	Message string `json:"message,omitempty"`
	Meta *ServerMeta `json:"meta,omitempty"`
	Status bool `json:"status,omitempty"`
}
