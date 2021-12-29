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

type MerchantTransactionMetadata struct {
	// Data contains key/value pair that gets sent from the SDK
	Data *interface{} `json:"data,omitempty"`
	Fee *MerchantTransactionMetadataFee `json:"fee,omitempty"`
	Ip string `json:"ip,omitempty"`
	// Determines if this payment was done by Fluidcoins widget
	IsWidgetPayment bool `json:"is_widget_payment,omitempty"`
	Onchain *MerchantTransactionMetadataOnchain `json:"onchain,omitempty"`
	Payment *MerchantTransactionMetadataPayment `json:"payment,omitempty"`
	Provider string `json:"provider,omitempty"`
	WidgetCoins []MerchantWidgetCoin `json:"widget_coins,omitempty"`
}