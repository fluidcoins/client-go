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

type ServerLinkEditRequest struct {
	Amount int32 `json:"amount,omitempty"`
	CollectPhoneNumber bool `json:"collect_phone_number,omitempty"`
	Description string `json:"description,omitempty"`
	Title string `json:"title,omitempty"`
}