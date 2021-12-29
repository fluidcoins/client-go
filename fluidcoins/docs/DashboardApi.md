# {{classname}}

All URIs are relative to *//api.fluidcoins.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComplianceSetupGet**](DashboardApi.md#ComplianceSetupGet) | **Get** /compliance/setup | Fetch all data used during compliance
[**DashboardGet**](DashboardApi.md#DashboardGet) | **Get** /dashboard | Fetch dashboard data

# **ComplianceSetupGet**
> ServerComplianceDataResponse ComplianceSetupGet(ctx, )
Fetch all data used during compliance

This endpoint provides a way to fetch multiple datasets a business will need to select

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServerComplianceDataResponse**](server.complianceDataResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardGet**
> ServerHomeResponse DashboardGet(ctx, )
Fetch dashboard data

This endpoint is used to retrieve all data for the home page

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServerHomeResponse**](server.homeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

