# {{classname}}

All URIs are relative to *//api.fluidcoins.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CurrenciesGet**](CurrenciesApi.md#CurrenciesGet) | **Get** /currencies | List currencies
[**RatesGet**](CurrenciesApi.md#RatesGet) | **Get** /rates | fetch fiat exchange rates

# **CurrenciesGet**
> ServerCurrenciesListResponse CurrenciesGet(ctx, optional)
List currencies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CurrenciesApiCurrenciesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CurrenciesApiCurrenciesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testNetOnly** | **optional.Bool**| Retrieve only coins that have a test-net network | 

### Return type

[**ServerCurrenciesListResponse**](server.currenciesListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RatesGet**
> ServerRatesResponse RatesGet(ctx, optional)
fetch fiat exchange rates

fetch a list of the current exchange rate of all supported fiat currencies on Fluidcoins. If you provide both to and from query params, we will return only that currency pair.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CurrenciesApiRatesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CurrenciesApiRatesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **optional.Bool**| base currency to convert from | 
 **to** | **optional.Bool**| currency you want to covert to | 

### Return type

[**ServerRatesResponse**](server.ratesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

