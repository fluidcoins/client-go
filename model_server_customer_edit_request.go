/*
 * Fluidcoins Pay api documentation
 *
 * A suite of APIs to create, manage and accept cryptocurrency payments in Africa
 *
 * API version: 0.1.5
 * Contact: developers@fluidcoins.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ServerCustomerEditRequest struct {
	Email string `json:"email,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Phone *ServerCustomerCreateRequestPhone `json:"phone,omitempty"`
}
