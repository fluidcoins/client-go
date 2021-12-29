# {{classname}}

All URIs are relative to *//api.fluidcoins.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsPasswordPost**](SettingsApi.md#SettingsPasswordPost) | **Post** /settings/password | Update your password
[**SettingsPreferencesGet**](SettingsApi.md#SettingsPreferencesGet) | **Get** /settings/preferences | Fetch preferences
[**SettingsPreferencesPatch**](SettingsApi.md#SettingsPreferencesPatch) | **Patch** /settings/preferences | Update preferences
[**SettingsProfilePatch**](SettingsApi.md#SettingsProfilePatch) | **Patch** /settings/profile | Update your user profile
[**SettingsToggleDomainPost**](SettingsApi.md#SettingsToggleDomainPost) | **Post** /settings/toggle-domain | Switch your dashboard domain from test to live and viceversa

# **SettingsPasswordPost**
> ServerMerchantUserResponse SettingsPasswordPost(ctx, body)
Update your password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerUpdatePasswordRequest**](ServerUpdatePasswordRequest.md)| Password data | 

### Return type

[**ServerMerchantUserResponse**](server.merchantUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsPreferencesGet**
> ServerPreferencesResponse SettingsPreferencesGet(ctx, )
Fetch preferences

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServerPreferencesResponse**](server.preferencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsPreferencesPatch**
> ServerPreferencesResponse SettingsPreferencesPatch(ctx, body)
Update preferences

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerPreferenceUpdateRequest**](ServerPreferenceUpdateRequest.md)| preferences data | 

### Return type

[**ServerPreferencesResponse**](server.preferencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsProfilePatch**
> ServerMerchantUserResponse SettingsProfilePatch(ctx, body)
Update your user profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerUpdateUserProfileRequest**](ServerUpdateUserProfileRequest.md)| user profile request | 

### Return type

[**ServerMerchantUserResponse**](server.merchantUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsToggleDomainPost**
> ServerMerchantCreationResponse SettingsToggleDomainPost(ctx, )
Switch your dashboard domain from test to live and viceversa

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServerMerchantCreationResponse**](server.merchantCreationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

