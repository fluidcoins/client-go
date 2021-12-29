# ServerRateData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyRate** | **int32** | BuyRate is the rate at which we will convert the from currency to the \&quot;to\&quot; currency | [optional] [default to null]
**From** | [***ServerRateDataFrom**](server.rateData_from.md) |  | [optional] [default to null]
**HumanReadableAmount** | **float64** | Deprecated. please see HumanReadableBuyRate and HumanReadableSaleRate | [optional] [default to null]
**HumanReadableBuyRate** | **float64** |  | [optional] [default to null]
**HumanReadableSaleRate** | **float64** |  | [optional] [default to null]
**Pair** | **string** |  | [optional] [default to null]
**Rate** | **int32** | Deprecated, please use out BuyRate and SaleRate | [optional] [default to null]
**SaleRate** | **int32** | SaleRate is the rate at which we will convert the \&quot;to\&quot; currency to the \&quot;from\&quot; currency | [optional] [default to null]
**To** | [***ServerRateDataFrom**](server.rateData_from.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

