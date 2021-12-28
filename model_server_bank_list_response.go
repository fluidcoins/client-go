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

type ServerBankListResponse struct {
	Banks []MerchantBank `json:"banks,omitempty"`
	// Generic message that tells you the status of the operation
	Message string `json:"message,omitempty"`
	Status bool `json:"status,omitempty"`
}
