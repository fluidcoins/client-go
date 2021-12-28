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

type MerchantWidgetInitilizationOptions struct {
	Coins []MerchantWidgetCoin `json:"coins,omitempty"`
	PaymentMethods []MerchantWidgetPaymentMethod `json:"payment_methods,omitempty"`
	Transaction *MerchantWidgetTransaction `json:"transaction,omitempty"`
}
