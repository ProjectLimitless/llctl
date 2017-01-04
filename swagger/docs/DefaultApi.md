# \DefaultApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminModulesGet**](DefaultApi.md#AdminModulesGet) | **Get** /admin/modules | Returns a list of loaded modules.
[**AdminRoutesGet**](DefaultApi.md#AdminRoutesGet) | **Get** /admin/routes | Returns a list of available API routes.
[**LoginPost**](DefaultApi.md#LoginPost) | **Post** /login | Logs a user in


# **AdminModulesGet**
> []IModule AdminModulesGet()

Returns a list of loaded modules.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]IModule**](IModule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminRoutesGet**
> []ApiRoute AdminRoutesGet()

Returns a list of available API routes.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]ApiRoute**](APIRoute.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoginPost**
> LoginResponse LoginPost($credentials)

Logs a user in


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**LoginCredentials**](LoginCredentials.md)| The login credentials to use. | 

### Return type

[**LoginResponse**](Login.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

