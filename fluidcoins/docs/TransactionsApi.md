# {{classname}}

All URIs are relative to *//api.fluidcoins.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionsGet**](TransactionsApi.md#TransactionsGet) | **Get** /transactions | List transactions
[**TransactionsIdGet**](TransactionsApi.md#TransactionsIdGet) | **Get** /transactions/{id} | Fetch a single transaction

# **TransactionsGet**
> ServerTransactionsListResponse TransactionsGet(ctx, optional)
List transactions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TransactionsApiTransactionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionsApiTransactionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page to query data from. Defaults to 1 | 
 **perPage** | **optional.Int32**| Number to items to return. Defaults to 20 items | 
 **status** | **optional.String**| Filter transactions. Defaults to all,  accepted values are pending,failed,success | 

### Return type

[**ServerTransactionsListResponse**](server.transactionsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsIdGet**
> ServerTransactionItemResponse TransactionsIdGet(ctx, id)
Fetch a single transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Unique ID for the transaction | 

### Return type

[**ServerTransactionItemResponse**](server.transactionItemResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

