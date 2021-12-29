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

type ServerHomeResponseTransactions struct {
	Amount int32 `json:"amount,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
	Date string `json:"date,omitempty"`
	Reference string `json:"reference,omitempty"`
	Status string `json:"status,omitempty"`
}