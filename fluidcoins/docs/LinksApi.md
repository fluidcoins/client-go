# {{classname}}

All URIs are relative to *//api.fluidcoins.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LinksGet**](LinksApi.md#LinksGet) | **Get** /links | List payment links
[**LinksIdEnableDelete**](LinksApi.md#LinksIdEnableDelete) | **Delete** /links/{id}/enable | Disable a payment link for collection
[**LinksIdEnablePost**](LinksApi.md#LinksIdEnablePost) | **Post** /links/{id}/enable | Enable a payment link for collection
[**LinksIdGet**](LinksApi.md#LinksIdGet) | **Get** /links/{id} | Fetch a payment link
[**LinksIdPatch**](LinksApi.md#LinksIdPatch) | **Patch** /links/{id} | Edit payment link
[**LinksIdTransactionsGet**](LinksApi.md#LinksIdTransactionsGet) | **Get** /links/{id}/transactions | List transactions that belong to a payment link
[**LinksPost**](LinksApi.md#LinksPost) | **Post** /links | Create a new payment link

# **LinksGet**
> ServerCustomerListResponse LinksGet(ctx, optional)
List payment links

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LinksApiLinksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LinksApiLinksGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page to query data from. Defaults to 1 | 
 **perPage** | **optional.Int32**| Number to items to return. Defaults to 10 items | 

### Return type

[**ServerCustomerListResponse**](server.customerListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinksIdEnableDelete**
> ServerPaymentLinkRetrievalResponse LinksIdEnableDelete(ctx, id)
Disable a payment link for collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| payment link unique identifier | 

### Return type

[**ServerPaymentLinkRetrievalResponse**](server.paymentLinkRetrievalResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinksIdEnablePost**
> ServerPaymentLinkRetrievalResponse LinksIdEnablePost(ctx, id)
Enable a payment link for collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| payment link unique identifier | 

### Return type

[**ServerPaymentLinkRetrievalResponse**](server.paymentLinkRetrievalResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinksIdGet**
> ServerPaymentLinkRetrievalResponse LinksIdGet(ctx, id)
Fetch a payment link

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Link unique identifier | 

### Return type

[**ServerPaymentLinkRetrievalResponse**](server.paymentLinkRetrievalResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinksIdPatch**
> ServerPaymentLinkRetrievalResponse LinksIdPatch(ctx, body, id)
Edit payment link

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerLinkEditRequest**](ServerLinkEditRequest.md)| link data | 
  **id** | **string**| Link unique identifier | 

### Return type

[**ServerPaymentLinkRetrievalResponse**](server.paymentLinkRetrievalResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinksIdTransactionsGet**
> ServerCustomerListResponse LinksIdTransactionsGet(ctx, id)
List transactions that belong to a payment link

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Link unique identifier | 

### Return type

[**ServerCustomerListResponse**](server.customerListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinksPost**
> ServerPaymentLinkCreateResponse LinksPost(ctx, body)
Create a new payment link

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerLinkCreateRequest**](ServerLinkCreateRequest.md)| Payment link data | 

### Return type

[**ServerPaymentLinkCreateResponse**](server.paymentLinkCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

