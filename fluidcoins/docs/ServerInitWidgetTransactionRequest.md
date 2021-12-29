# ServerInitWidgetTransactionRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Amount. Must be formatte in Kobo. Minimum of 500 Naira | [optional] [default to null]
**ApiKey** | **string** | Public key only. Never pass your secret key here | [optional] [default to null]
**CustomerName** | **string** | Not required. But the customer name. SDK shoudl allow people set this or not | [optional] [default to null]
**Email** | **string** | Customer email. Required | [optional] [default to null]
**Metadata** | [***interface{}**](interface{}.md) | Key value pair of arbitrary data you want to attach to this request | [optional] [default to null]
**PaymentLinkId** | **string** | If coming from a payment page. Put the id of the payment page here. If this is set, it will use the amount from the payment page instead and disregard the amount field in this request | [optional] [default to null]
**Phone** | **string** |  | [optional] [default to null]
**Reference** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

