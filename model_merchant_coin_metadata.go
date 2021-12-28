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

type MerchantCoinMetadata struct {
	// XRP addresses must have a destination tag
	DestinationTag int32 `json:"destination_tag,omitempty"`
	// XLM requires a memo
	Memo string `json:"memo,omitempty"`
	// can be erc20, trc20 or bsc if empty, then assume the default chain in cases such as btc,ltc and what not
	Network string `json:"network,omitempty"`
}
