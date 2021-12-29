# {{classname}}

All URIs are relative to *//api.fluidcoins.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressGet**](AddressApi.md#AddressGet) | **Get** /address | List addresses
[**AddressIdGet**](AddressApi.md#AddressIdGet) | **Get** /address/{id} | Fetch an address
[**AddressIdTransactionsGet**](AddressApi.md#AddressIdTransactionsGet) | **Get** /address/{id}/transactions | List transactions for a given address
[**AddressPost**](AddressApi.md#AddressPost) | **Post** /address | Create a new address
[**AddressTransactionsIdGet**](AddressApi.md#AddressTransactionsIdGet) | **Get** /address/transactions/{id} | Fetch a single transaction that occurred on a given address

# **AddressGet**
> ServerAddressListResponse AddressGet(ctx, optional)
List addresses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddressApiAddressGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressApiAddressGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page to query data from. Defaults to 1 | 
 **perPage** | **optional.Int32**| Number to items to return. Defaults to 20 items | 
 **coinId** | **optional.String**| fetch addresses for a specific coin. Must be a uuid and you can fetch the id of the coin by using the v1/currencies endpoint | 

### Return type

[**ServerAddressListResponse**](server.addressListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressIdGet**
> ServerAddressResponse AddressIdGet(ctx, id)
Fetch an address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| address unique identifier | 

### Return type

[**ServerAddressResponse**](server.addressResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressIdTransactionsGet**
> ServerAddressTransactionsListResponse AddressIdTransactionsGet(ctx, id, optional)
List transactions for a given address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| address unique identifier | 
 **optional** | ***AddressApiAddressIdTransactionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressApiAddressIdTransactionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page to query data from. Defaults to 1 | 
 **perPage** | **optional.Int32**| Number to items to return. Defaults to 20 items | 

### Return type

[**ServerAddressTransactionsListResponse**](server.addressTransactionsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressPost**
> ServerAddressResponse AddressPost(ctx, body)
Create a new address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerGenerateCoinRequest**](ServerGenerateCoinRequest.md)| Coin data | 

### Return type

[**ServerAddressResponse**](server.addressResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressTransactionsIdGet**
> ServerAddressTransactionResponse AddressTransactionsIdGet(ctx, id)
Fetch a single transaction that occurred on a given address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| transaction reference | 

### Return type

[**ServerAddressTransactionResponse**](server.addressTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

