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

type ServerMerchantCreateRequest struct {
	Description string `json:"description,omitempty"`
	Industry string `json:"industry,omitempty"`
	Name string `json:"name,omitempty"`
	VerificationType int32 `json:"verification_type,omitempty"`
}
