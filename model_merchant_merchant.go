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

type MerchantMerchant struct {
	BusinessName string `json:"business_name,omitempty"`
	Country *MerchantCountry `json:"country,omitempty"`
	CountryId string `json:"country_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	// This is what we show on the dashboard. Does not determine anything else like if the merchant has gone into live mode or can accept payments
	DashboardDomain string `json:"dashboard_domain,omitempty"`
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"`
	IndustryId string `json:"industry_id,omitempty"`
	Metadata *MerchantMerchantMetadata `json:"metadata,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	VerificationStatus int32 `json:"verification_status,omitempty"`
	VerificationType int32 `json:"verification_type,omitempty"`
}
