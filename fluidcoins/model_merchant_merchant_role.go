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

type MerchantMerchantRole struct {
	Counter int32 `json:"counter,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Id string `json:"id,omitempty"`
	MerchantId string `json:"merchant_id,omitempty"`
	Role int32 `json:"role,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UserId string `json:"user_id,omitempty"`
}