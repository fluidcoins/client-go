# {{classname}}

All URIs are relative to *//api.fluidcoins.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusLinksIdGet**](WidgetApi.md#StatusLinksIdGet) | **Get** /status/links/{id} | get basic public data about a payment link
[**StatusTransactionIdGet**](WidgetApi.md#StatusTransactionIdGet) | **Get** /status/transaction/{id} | Fetch the status of a transaction
[**WidgetPost**](WidgetApi.md#WidgetPost) | **Post** /widget | Create a new transaction
[**WidgetPut**](WidgetApi.md#WidgetPut) | **Put** /widget | simulate a widget payment

# **StatusLinksIdGet**
> ServerPublicPaymentLinkResponse StatusLinksIdGet(ctx, id)
get basic public data about a payment link

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Link unique identifier | 

### Return type

[**ServerPublicPaymentLinkResponse**](server.publicPaymentLinkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatusTransactionIdGet**
> ServerTransactionStatusResponse StatusTransactionIdGet(ctx, id, publicKey)
Fetch the status of a transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| transaction unique identifier | 
  **publicKey** | **string**| Public key for the merchant | 

### Return type

[**ServerTransactionStatusResponse**](server.transactionStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WidgetPost**
> ServerWidgetInitalizationResponse WidgetPost(ctx, body)
Create a new transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerInitWidgetTransactionRequest**](ServerInitWidgetTransactionRequest.md)| Payment link data | 

### Return type

[**ServerWidgetInitalizationResponse**](server.widgetInitalizationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WidgetPut**
> ServerApiStatus WidgetPut(ctx, body, publicKey)
simulate a widget payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerSimulateWidgetPayment**](ServerSimulateWidgetPayment.md)| simulation data | 
  **publicKey** | **string**| Public key for the merchant | 

### Return type

[**ServerApiStatus**](server.APIStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

