# {{classname}}

All URIs are relative to *//api.fluidcoins.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HooksGet**](WebhooksApi.md#HooksGet) | **Get** /hooks | List webhook endpoints. Only 5
[**HooksIdEnableDelete**](WebhooksApi.md#HooksIdEnableDelete) | **Delete** /hooks/{id}/enable | Disable an endpoint from receiving webhooks
[**HooksIdEnablePost**](WebhooksApi.md#HooksIdEnablePost) | **Post** /hooks/{id}/enable | Mark an endpoint as active.
[**HooksIdGet**](WebhooksApi.md#HooksIdGet) | **Get** /hooks/{id} | Fetch an endpoint details
[**HooksIdPatch**](WebhooksApi.md#HooksIdPatch) | **Patch** /hooks/{id} | Edit a webhook
[**HooksPost**](WebhooksApi.md#HooksPost) | **Post** /hooks | Create a new webhook. Limited to only 5 endpoints

# **HooksGet**
> ServerListWebhookResponse HooksGet(ctx, )
List webhook endpoints. Only 5

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServerListWebhookResponse**](server.listWebhookResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HooksIdEnableDelete**
> ServerWebhookCreateResponse HooksIdEnableDelete(ctx, id)
Disable an endpoint from receiving webhooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| hook unique identifier | 

### Return type

[**ServerWebhookCreateResponse**](server.webhookCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HooksIdEnablePost**
> ServerWebhookCreateResponse HooksIdEnablePost(ctx, id)
Mark an endpoint as active.

Mark an endpoint as the default. Please not that this will

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| hook unique identifier | 

### Return type

[**ServerWebhookCreateResponse**](server.webhookCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HooksIdGet**
> ServerWebhookCreateResponse HooksIdGet(ctx, id)
Fetch an endpoint details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| hook unique identifier | 

### Return type

[**ServerWebhookCreateResponse**](server.webhookCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HooksIdPatch**
> ServerWebhookCreateResponse HooksIdPatch(ctx, body, id)
Edit a webhook

This endpoint allows you to update the url for a webhook endpoint while retaining the same secret/signing key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerWebhookCreateRequest**](ServerWebhookCreateRequest.md)| webhook data | 
  **id** | **string**| Webhook unique id | 

### Return type

[**ServerWebhookCreateResponse**](server.webhookCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HooksPost**
> ServerWebhookCreateResponse HooksPost(ctx, body)
Create a new webhook. Limited to only 5 endpoints

You can only create a maximum of 5 endpoints. Each endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerWebhookCreateRequest**](ServerWebhookCreateRequest.md)| hook data | 

### Return type

[**ServerWebhookCreateResponse**](server.webhookCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

