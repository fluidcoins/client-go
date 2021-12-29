# {{classname}}

All URIs are relative to *//api.fluidcoins.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomersGet**](CustomersApi.md#CustomersGet) | **Get** /customers | List customers
[**CustomersIdBlacklistDelete**](CustomersApi.md#CustomersIdBlacklistDelete) | **Delete** /customers/{id}/blacklist | Whitelists a customer
[**CustomersIdBlacklistPost**](CustomersApi.md#CustomersIdBlacklistPost) | **Post** /customers/{id}/blacklist | Blacklists a customer
[**CustomersIdGet**](CustomersApi.md#CustomersIdGet) | **Get** /customers/{id} | Fetch a customer
[**CustomersIdPatch**](CustomersApi.md#CustomersIdPatch) | **Patch** /customers/{id} | Edit a customer
[**CustomersPost**](CustomersApi.md#CustomersPost) | **Post** /customers | Create a new customer

# **CustomersGet**
> ServerCustomerListResponse CustomersGet(ctx, optional)
List customers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomersApiCustomersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiCustomersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page to query data from. Defaults to 1 | 
 **perPage** | **optional.Int32**| Number to items to return. Defaults to 20 items | 
 **blacklisted** | **optional.Bool**| Fetch blacklisted customers | 

### Return type

[**ServerCustomerListResponse**](server.customerListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomersIdBlacklistDelete**
> ServerCustomerRetrievalResponse CustomersIdBlacklistDelete(ctx, id)
Whitelists a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer unique identifier | 

### Return type

[**ServerCustomerRetrievalResponse**](server.customerRetrievalResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomersIdBlacklistPost**
> ServerCustomerRetrievalResponse CustomersIdBlacklistPost(ctx, id)
Blacklists a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer unique identifier | 

### Return type

[**ServerCustomerRetrievalResponse**](server.customerRetrievalResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomersIdGet**
> ServerCustomerRetrievalResponse CustomersIdGet(ctx, id)
Fetch a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| customer unique identifier | 

### Return type

[**ServerCustomerRetrievalResponse**](server.customerRetrievalResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomersIdPatch**
> ServerCustomerRetrievalResponse CustomersIdPatch(ctx, body, id)
Edit a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerCustomerEditRequest**](ServerCustomerEditRequest.md)| Customer data | 
  **id** | **string**| Customer unique identifier | 

### Return type

[**ServerCustomerRetrievalResponse**](server.customerRetrievalResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomersPost**
> ServerCustomerCreateResponse CustomersPost(ctx, body)
Create a new customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerCustomerCreateRequest**](ServerCustomerCreateRequest.md)| Customer data | 

### Return type

[**ServerCustomerCreateResponse**](server.customerCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

