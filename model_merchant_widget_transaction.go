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

type MerchantWidgetTransaction struct {
	// Value of transaction in Merchant =currency Usually Naira
	Amount int32 `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	Id string `json:"id,omitempty"`
	Merchant *MerchantWidgetTransactionMerchant `json:"merchant,omitempty"`
	Reference string `json:"reference,omitempty"`
}
