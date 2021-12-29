# {{classname}}

All URIs are relative to *//api.fluidcoins.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PayoutsAccountsBanksGet**](PayoutsApi.md#PayoutsAccountsBanksGet) | **Get** /payouts/accounts/banks | List banks
[**PayoutsAccountsBanksResolveGet**](PayoutsApi.md#PayoutsAccountsBanksResolveGet) | **Get** /payouts/accounts/banks/resolve | Resolve bank accounts
[**PayoutsAccountsGet**](PayoutsApi.md#PayoutsAccountsGet) | **Get** /payouts/accounts | List payouts accounts
[**PayoutsAccountsPost**](PayoutsApi.md#PayoutsAccountsPost) | **Post** /payouts/accounts | Create a new payout account
[**PayoutsGet**](PayoutsApi.md#PayoutsGet) | **Get** /payouts | List payouts
[**PayoutsIdDelete**](PayoutsApi.md#PayoutsIdDelete) | **Delete** /payouts/{id} | Cancel a payout.
[**PayoutsIdGet**](PayoutsApi.md#PayoutsIdGet) | **Get** /payouts/{id} | Fetch the details of a payout
[**PayoutsPost**](PayoutsApi.md#PayoutsPost) | **Post** /payouts | Request for a payout

# **PayoutsAccountsBanksGet**
> ServerBankListResponse PayoutsAccountsBanksGet(ctx, optional)
List banks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PayoutsApiPayoutsAccountsBanksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PayoutsApiPayoutsAccountsBanksGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **optional.String**| Optional. Defaults to NG | 

### Return type

[**ServerBankListResponse**](server.bankListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PayoutsAccountsBanksResolveGet**
> ServerResolveBankAccountResponse PayoutsAccountsBanksResolveGet(ctx, bankCode, account)
Resolve bank accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bankCode** | **string**| Sort code of the bank | 
  **account** | **string**| Bank account number | 

### Return type

[**ServerResolveBankAccountResponse**](server.resolveBankAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PayoutsAccountsGet**
> ServerPayoutAccountsListResponse PayoutsAccountsGet(ctx, optional)
List payouts accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PayoutsApiPayoutsAccountsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PayoutsApiPayoutsAccountsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page to query data from. Defaults to 1 | 
 **perPage** | **optional.Int32**| Number to items to return. Defaults to 20 items | 

### Return type

[**ServerPayoutAccountsListResponse**](server.payoutAccountsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PayoutsAccountsPost**
> ServerCreatePayoutAccountResponse PayoutsAccountsPost(ctx, body)
Create a new payout account

create a transfer recipient and store the account ID so you can send money to the account whenever

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerCreatePayoutAccountRequest**](ServerCreatePayoutAccountRequest.md)| payout data | 

### Return type

[**ServerCreatePayoutAccountResponse**](server.createPayoutAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PayoutsGet**
> ServerPayoutListResponse PayoutsGet(ctx, optional)
List payouts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PayoutsApiPayoutsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PayoutsApiPayoutsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page to query data from. Defaults to 1 | 
 **perPage** | **optional.Int32**| Number to items to return. Defaults to 20 items | 

### Return type

[**ServerPayoutListResponse**](server.payoutListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PayoutsIdDelete**
> ServerPayoutRequestResponse PayoutsIdDelete(ctx, id)
Cancel a payout.

payouts work like a state machine. From requested to pending to processed. While a payout is still in requested state, you can cancel it from being executed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Payout unique identifier | 

### Return type

[**ServerPayoutRequestResponse**](server.payoutRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PayoutsIdGet**
> ServerPayoutRequestResponse PayoutsIdGet(ctx, id)
Fetch the details of a payout

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Payout unique identifier | 

### Return type

[**ServerPayoutRequestResponse**](server.payoutRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PayoutsPost**
> ServerPayoutRequestResponse PayoutsPost(ctx, body)
Request for a payout

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerCreateRequestPayout**](ServerCreateRequestPayout.md)| payout request data | 

### Return type

[**ServerPayoutRequestResponse**](server.payoutRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

