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

type ServerCurrenciesListResponse struct {
	Currencies []MerchantMerchantCoin `json:"currencies,omitempty"`
	Message string `json:"message,omitempty"`
	Status bool `json:"status,omitempty"`
}
