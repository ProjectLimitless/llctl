# \DefaultApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminModulesGet**](DefaultApi.md#AdminModulesGet) | **Get** /admin/modules | Returns a list of loaded modules.
[**AdminRoutesGet**](DefaultApi.md#AdminRoutesGet) | **Get** /admin/routes | Returns a list of available API routes.
[**AdminUsersGet**](DefaultApi.md#AdminUsersGet) | **Get** /admin/users | Returns a list of users for the installation.
[**LoginPost**](DefaultApi.md#LoginPost) | **Post** /login | Logs a user in
[**SkillsGet**](DefaultApi.md#SkillsGet) | **Get** /skills | Returns a list of registered skills for the installation
[**SkillsIdDelete**](DefaultApi.md#SkillsIdDelete) | **Delete** /skills/{id} | Deregister a previously registered skill
[**SkillsPost**](DefaultApi.md#SkillsPost) | **Post** /skills | Registers a new skill


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

# **AdminUsersGet**
> []BaseUser AdminUsersGet()

Returns a list of users for the installation.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]BaseUser**](BaseUser.md)

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

# **SkillsGet**
> []Skill SkillsGet()

Returns a list of registered skills for the installation


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]Skill**](Skill.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SkillsIdDelete**
> SkillsIdDelete($id)

Deregister a previously registered skill


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The Skill&#39;s registered UUID | 

### Return type

void (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SkillsPost**
> SkillResponseSuccess SkillsPost($skill)

Registers a new skill


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skill** | [**Skill**](Skill.md)| The skill to register | 

### Return type

[**SkillResponseSuccess**](Skill.Response.Success.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

