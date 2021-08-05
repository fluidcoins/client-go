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

type MerchantCryptoPayoutAccount struct {
	Address string `json:"address,omitempty"`
	Coin string `json:"coin,omitempty"`
	Label string `json:"label,omitempty"`
}
