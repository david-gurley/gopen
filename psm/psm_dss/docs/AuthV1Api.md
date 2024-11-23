# \AuthV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAuthenticationPolicy**](AuthV1Api.md#AddAuthenticationPolicy) | **Post** /configs/auth/v1/authn-policy | Create AuthenticationPolicy object
[**AddRole**](AuthV1Api.md#AddRole) | **Post** /configs/auth/v1/tenant/{O.Tenant}/roles | Create Role object
[**AddRole1**](AuthV1Api.md#AddRole1) | **Post** /configs/auth/v1/roles | Create Role object
[**AddRoleBinding**](AuthV1Api.md#AddRoleBinding) | **Post** /configs/auth/v1/tenant/{O.Tenant}/role-bindings | Create RoleBinding object
[**AddRoleBinding1**](AuthV1Api.md#AddRoleBinding1) | **Post** /configs/auth/v1/role-bindings | Create RoleBinding object
[**AddUser**](AuthV1Api.md#AddUser) | **Post** /configs/auth/v1/tenant/{O.Tenant}/users | Create User object
[**AddUser1**](AuthV1Api.md#AddUser1) | **Post** /configs/auth/v1/users | Create User object
[**DeleteRole**](AuthV1Api.md#DeleteRole) | **Delete** /configs/auth/v1/tenant/{O.Tenant}/roles/{O.Name} | Delete Role object
[**DeleteRole1**](AuthV1Api.md#DeleteRole1) | **Delete** /configs/auth/v1/roles/{O.Name} | Delete Role object
[**DeleteRoleBinding**](AuthV1Api.md#DeleteRoleBinding) | **Delete** /configs/auth/v1/tenant/{O.Tenant}/role-bindings/{O.Name} | Delete RoleBinding object
[**DeleteRoleBinding1**](AuthV1Api.md#DeleteRoleBinding1) | **Delete** /configs/auth/v1/role-bindings/{O.Name} | Delete RoleBinding object
[**DeleteUser**](AuthV1Api.md#DeleteUser) | **Delete** /configs/auth/v1/tenant/{O.Tenant}/users/{O.Name} | Delete User object
[**DeleteUser1**](AuthV1Api.md#DeleteUser1) | **Delete** /configs/auth/v1/users/{O.Name} | Delete User object
[**GetAuthenticationPolicy**](AuthV1Api.md#GetAuthenticationPolicy) | **Get** /configs/auth/v1/authn-policy | Get AuthenticationPolicy object
[**GetRole**](AuthV1Api.md#GetRole) | **Get** /configs/auth/v1/tenant/{O.Tenant}/roles/{O.Name} | Get Role object
[**GetRole1**](AuthV1Api.md#GetRole1) | **Get** /configs/auth/v1/roles/{O.Name} | Get Role object
[**GetRoleBinding**](AuthV1Api.md#GetRoleBinding) | **Get** /configs/auth/v1/tenant/{O.Tenant}/role-bindings/{O.Name} | Get RoleBinding object
[**GetRoleBinding1**](AuthV1Api.md#GetRoleBinding1) | **Get** /configs/auth/v1/role-bindings/{O.Name} | Get RoleBinding object
[**GetUser**](AuthV1Api.md#GetUser) | **Get** /configs/auth/v1/tenant/{O.Tenant}/users/{O.Name} | Get User object
[**GetUser1**](AuthV1Api.md#GetUser1) | **Get** /configs/auth/v1/users/{O.Name} | Get User object
[**GetUserPreference**](AuthV1Api.md#GetUserPreference) | **Get** /configs/auth/v1/tenant/{O.Tenant}/user-preferences/{O.Name} | Get UserPreference object
[**GetUserPreference1**](AuthV1Api.md#GetUserPreference1) | **Get** /configs/auth/v1/user-preferences/{O.Name} | Get UserPreference object
[**IsAuthorized**](AuthV1Api.md#IsAuthorized) | **Post** /configs/auth/v1/tenant/{O.Tenant}/users/{O.Name}/IsAuthorized | Review authorization for user
[**IsAuthorized1**](AuthV1Api.md#IsAuthorized1) | **Post** /configs/auth/v1/users/{O.Name}/IsAuthorized | Review authorization for user
[**LabelAuthenticationPolicy**](AuthV1Api.md#LabelAuthenticationPolicy) | **Post** /configs/auth/v1/authn-policy/label | Label AuthenticationPolicy object
[**LabelRole**](AuthV1Api.md#LabelRole) | **Post** /configs/auth/v1/tenant/{O.Tenant}/roles/{O.Name}/label | Label Role object
[**LabelRole1**](AuthV1Api.md#LabelRole1) | **Post** /configs/auth/v1/roles/{O.Name}/label | Label Role object
[**LabelRoleBinding**](AuthV1Api.md#LabelRoleBinding) | **Post** /configs/auth/v1/tenant/{O.Tenant}/role-bindings/{O.Name}/label | Label RoleBinding object
[**LabelRoleBinding1**](AuthV1Api.md#LabelRoleBinding1) | **Post** /configs/auth/v1/role-bindings/{O.Name}/label | Label RoleBinding object
[**LabelUser**](AuthV1Api.md#LabelUser) | **Post** /configs/auth/v1/tenant/{O.Tenant}/users/{O.Name}/label | Label User object
[**LabelUser1**](AuthV1Api.md#LabelUser1) | **Post** /configs/auth/v1/users/{O.Name}/label | Label User object
[**LabelUserPreference**](AuthV1Api.md#LabelUserPreference) | **Post** /configs/auth/v1/tenant/{O.Tenant}/user-preferences/{O.Name}/label | Label UserPreference object
[**LabelUserPreference1**](AuthV1Api.md#LabelUserPreference1) | **Post** /configs/auth/v1/user-preferences/{O.Name}/label | Label UserPreference object
[**LdapBindCheck**](AuthV1Api.md#LdapBindCheck) | **Post** /configs/auth/v1/authn-policy/LdapBindCheck | Test LDAP bind operation
[**LdapConnectionCheck**](AuthV1Api.md#LdapConnectionCheck) | **Post** /configs/auth/v1/authn-policy/LdapConnectionCheck | Test LDAP connection
[**ListRole**](AuthV1Api.md#ListRole) | **Get** /configs/auth/v1/tenant/{O.Tenant}/roles | List Role objects
[**ListRole1**](AuthV1Api.md#ListRole1) | **Get** /configs/auth/v1/roles | List Role objects
[**ListRoleBinding**](AuthV1Api.md#ListRoleBinding) | **Get** /configs/auth/v1/tenant/{O.Tenant}/role-bindings | List RoleBinding objects
[**ListRoleBinding1**](AuthV1Api.md#ListRoleBinding1) | **Get** /configs/auth/v1/role-bindings | List RoleBinding objects
[**ListUser**](AuthV1Api.md#ListUser) | **Get** /configs/auth/v1/tenant/{O.Tenant}/users | List User objects
[**ListUser1**](AuthV1Api.md#ListUser1) | **Get** /configs/auth/v1/users | List User objects
[**PasswordChange**](AuthV1Api.md#PasswordChange) | **Post** /configs/auth/v1/tenant/{O.Tenant}/users/{O.Name}/PasswordChange | Change user password
[**PasswordChange1**](AuthV1Api.md#PasswordChange1) | **Post** /configs/auth/v1/users/{O.Name}/PasswordChange | Change user password
[**PasswordReset**](AuthV1Api.md#PasswordReset) | **Post** /configs/auth/v1/tenant/{O.Tenant}/users/{O.Name}/PasswordReset | Reset user password
[**PasswordReset1**](AuthV1Api.md#PasswordReset1) | **Post** /configs/auth/v1/users/{O.Name}/PasswordReset | Reset user password
[**TokenSecretGenerate**](AuthV1Api.md#TokenSecretGenerate) | **Post** /configs/auth/v1/authn-policy/TokenSecretGenerate | Generate secret for token signing
[**Unlock**](AuthV1Api.md#Unlock) | **Post** /configs/auth/v1/tenant/{O.Tenant}/users/{O.Name}/Unlock | Unlock user
[**Unlock1**](AuthV1Api.md#Unlock1) | **Post** /configs/auth/v1/users/{O.Name}/Unlock | Unlock user
[**UpdateAuthenticationPolicy**](AuthV1Api.md#UpdateAuthenticationPolicy) | **Put** /configs/auth/v1/authn-policy | Update AuthenticationPolicy object
[**UpdateRole**](AuthV1Api.md#UpdateRole) | **Put** /configs/auth/v1/tenant/{O.Tenant}/roles/{O.Name} | Update Role object
[**UpdateRole1**](AuthV1Api.md#UpdateRole1) | **Put** /configs/auth/v1/roles/{O.Name} | Update Role object
[**UpdateRoleBinding**](AuthV1Api.md#UpdateRoleBinding) | **Put** /configs/auth/v1/tenant/{O.Tenant}/role-bindings/{O.Name} | Update RoleBinding object
[**UpdateRoleBinding1**](AuthV1Api.md#UpdateRoleBinding1) | **Put** /configs/auth/v1/role-bindings/{O.Name} | Update RoleBinding object
[**UpdateUser**](AuthV1Api.md#UpdateUser) | **Put** /configs/auth/v1/tenant/{O.Tenant}/users/{O.Name} | Update User object
[**UpdateUser1**](AuthV1Api.md#UpdateUser1) | **Put** /configs/auth/v1/users/{O.Name} | Update User object
[**UpdateUserPreference**](AuthV1Api.md#UpdateUserPreference) | **Put** /configs/auth/v1/tenant/{O.Tenant}/user-preferences/{O.Name} | Update UserPreference object
[**UpdateUserPreference1**](AuthV1Api.md#UpdateUserPreference1) | **Put** /configs/auth/v1/user-preferences/{O.Name} | Update UserPreference object
[**WatchAuthenticationPolicy**](AuthV1Api.md#WatchAuthenticationPolicy) | **Get** /configs/auth/v1/watch/authn-policy | Watch AuthenticationPolicy objects. Supports WebSockets or HTTP long poll
[**WatchRole**](AuthV1Api.md#WatchRole) | **Get** /configs/auth/v1/watch/tenant/{O.Tenant}/roles | Watch Role objects. Supports WebSockets or HTTP long poll
[**WatchRole1**](AuthV1Api.md#WatchRole1) | **Get** /configs/auth/v1/watch/roles | Watch Role objects. Supports WebSockets or HTTP long poll
[**WatchRoleBinding**](AuthV1Api.md#WatchRoleBinding) | **Get** /configs/auth/v1/watch/tenant/{O.Tenant}/role-bindings | Watch RoleBinding objects. Supports WebSockets or HTTP long poll
[**WatchRoleBinding1**](AuthV1Api.md#WatchRoleBinding1) | **Get** /configs/auth/v1/watch/role-bindings | Watch RoleBinding objects. Supports WebSockets or HTTP long poll
[**WatchUser**](AuthV1Api.md#WatchUser) | **Get** /configs/auth/v1/watch/tenant/{O.Tenant}/users | Watch User objects. Supports WebSockets or HTTP long poll
[**WatchUser1**](AuthV1Api.md#WatchUser1) | **Get** /configs/auth/v1/watch/users | Watch User objects. Supports WebSockets or HTTP long poll
[**WatchUserPreference**](AuthV1Api.md#WatchUserPreference) | **Get** /configs/auth/v1/watch/tenant/{O.Tenant}/user-preferences | Watch UserPreference objects. Supports WebSockets or HTTP long poll
[**WatchUserPreference1**](AuthV1Api.md#WatchUserPreference1) | **Get** /configs/auth/v1/watch/user-preferences | Watch UserPreference objects. Supports WebSockets or HTTP long poll



## AddAuthenticationPolicy

> AuthAuthenticationPolicy AddAuthenticationPolicy(ctx).Body(body).Execute()

Create AuthenticationPolicy object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAuthAuthenticationPolicy() // AuthAuthenticationPolicy | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.AddAuthenticationPolicy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.AddAuthenticationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAuthenticationPolicy`: AuthAuthenticationPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.AddAuthenticationPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAuthenticationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthAuthenticationPolicy**](AuthAuthenticationPolicy.md) |  | 

### Return type

[**AuthAuthenticationPolicy**](authAuthenticationPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRole

> AuthRole AddRole(ctx, oTenant).Body(body).Execute()

Create Role object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewAuthRole() // AuthRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.AddRole(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.AddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRole`: AuthRole
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.AddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthRole**](AuthRole.md) |  | 

### Return type

[**AuthRole**](authRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRole1

> AuthRole AddRole1(ctx).Body(body).Execute()

Create Role object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAuthRole() // AuthRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.AddRole1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.AddRole1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRole1`: AuthRole
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.AddRole1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRole1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthRole**](AuthRole.md) |  | 

### Return type

[**AuthRole**](authRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRoleBinding

> AuthRoleBinding AddRoleBinding(ctx, oTenant).Body(body).Execute()

Create RoleBinding object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewAuthRoleBinding() // AuthRoleBinding | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.AddRoleBinding(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.AddRoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRoleBinding`: AuthRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.AddRoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthRoleBinding**](AuthRoleBinding.md) |  | 

### Return type

[**AuthRoleBinding**](authRoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRoleBinding1

> AuthRoleBinding AddRoleBinding1(ctx).Body(body).Execute()

Create RoleBinding object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAuthRoleBinding() // AuthRoleBinding | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.AddRoleBinding1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.AddRoleBinding1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRoleBinding1`: AuthRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.AddRoleBinding1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRoleBinding1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthRoleBinding**](AuthRoleBinding.md) |  | 

### Return type

[**AuthRoleBinding**](authRoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUser

> AuthUser AddUser(ctx, oTenant).Body(body).Execute()

Create User object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewAuthUser() // AuthUser | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.AddUser(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.AddUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUser`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.AddUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthUser**](AuthUser.md) |  | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUser1

> AuthUser AddUser1(ctx).Body(body).Execute()

Create User object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAuthUser() // AuthUser | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.AddUser1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.AddUser1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUser1`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.AddUser1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUser1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthUser**](AuthUser.md) |  | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> AuthRole DeleteRole(ctx, oTenant).Execute()

Delete Role object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.DeleteRole(context.Background(), oTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRole`: AuthRole
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.DeleteRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthRole**](authRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole1

> AuthRole DeleteRole1(ctx, oName).Execute()

Delete Role object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.DeleteRole1(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.DeleteRole1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRole1`: AuthRole
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.DeleteRole1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRole1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthRole**](authRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoleBinding

> AuthRoleBinding DeleteRoleBinding(ctx, oTenant).Execute()

Delete RoleBinding object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.DeleteRoleBinding(context.Background(), oTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.DeleteRoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRoleBinding`: AuthRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.DeleteRoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthRoleBinding**](authRoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoleBinding1

> AuthRoleBinding DeleteRoleBinding1(ctx, oName).Execute()

Delete RoleBinding object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.DeleteRoleBinding1(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.DeleteRoleBinding1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRoleBinding1`: AuthRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.DeleteRoleBinding1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleBinding1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthRoleBinding**](authRoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> AuthUser DeleteUser(ctx, oTenant).Execute()

Delete User object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.DeleteUser(context.Background(), oTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUser`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser1

> AuthUser DeleteUser1(ctx, oName).Execute()

Delete User object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.DeleteUser1(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.DeleteUser1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUser1`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.DeleteUser1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUser1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticationPolicy

> AuthAuthenticationPolicy GetAuthenticationPolicy(ctx).TKind(tKind).MetaCreationTime(metaCreationTime).AuthenticatorsAuthenticatorOrder(authenticatorsAuthenticatorOrder).Execute()

Get AuthenticationPolicy object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    authenticatorsAuthenticatorOrder := []string{"Inner_example"} // []string | Order in which authenticators are applied. If an authenticator returns success, others are skipped. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.GetAuthenticationPolicy(context.Background()).TKind(tKind).MetaCreationTime(metaCreationTime).AuthenticatorsAuthenticatorOrder(authenticatorsAuthenticatorOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.GetAuthenticationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationPolicy`: AuthAuthenticationPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.GetAuthenticationPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **authenticatorsAuthenticatorOrder** | **[]string** | Order in which authenticators are applied. If an authenticator returns success, others are skipped. | 

### Return type

[**AuthAuthenticationPolicy**](authAuthenticationPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> AuthRole GetRole(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()

Get Role object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.GetRole(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: AuthRole
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**AuthRole**](authRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole1

> AuthRole GetRole1(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()

Get Role object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.GetRole1(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.GetRole1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole1`: AuthRole
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.GetRole1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRole1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**AuthRole**](authRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleBinding

> AuthRoleBinding GetRoleBinding(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).SpecUsers(specUsers).SpecRole(specRole).Execute()

Get RoleBinding object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    specUsers := []string{"Inner_example"} // []string |  (optional)
    specRole := "specRole_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.GetRoleBinding(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).SpecUsers(specUsers).SpecRole(specRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.GetRoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleBinding`: AuthRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.GetRoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specUsers** | **[]string** |  | 
 **specRole** | **string** |  | 

### Return type

[**AuthRoleBinding**](authRoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleBinding1

> AuthRoleBinding GetRoleBinding1(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecUsers(specUsers).SpecRole(specRole).Execute()

Get RoleBinding object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    specUsers := []string{"Inner_example"} // []string |  (optional)
    specRole := "specRole_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.GetRoleBinding1(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecUsers(specUsers).SpecRole(specRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.GetRoleBinding1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleBinding1`: AuthRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.GetRoleBinding1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleBinding1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specUsers** | **[]string** |  | 
 **specRole** | **string** |  | 

### Return type

[**AuthRoleBinding**](authRoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> AuthUser GetUser(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).SpecFullname(specFullname).StatusRoles(statusRoles).Execute()

Get User object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    specFullname := "specFullname_example" // string |  (optional)
    statusRoles := []string{"Inner_example"} // []string | Roles assigned to user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.GetUser(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).SpecFullname(specFullname).StatusRoles(statusRoles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specFullname** | **string** |  | 
 **statusRoles** | **[]string** | Roles assigned to user. | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser1

> AuthUser GetUser1(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecFullname(specFullname).StatusRoles(statusRoles).Execute()

Get User object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    specFullname := "specFullname_example" // string |  (optional)
    statusRoles := []string{"Inner_example"} // []string | Roles assigned to user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.GetUser1(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecFullname(specFullname).StatusRoles(statusRoles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.GetUser1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser1`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.GetUser1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUser1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specFullname** | **string** |  | 
 **statusRoles** | **[]string** | Roles assigned to user. | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPreference

> AuthUserPreference GetUserPreference(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).SpecOptions(specOptions).Execute()

Get UserPreference object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    specOptions := "specOptions_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.GetUserPreference(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).SpecOptions(specOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.GetUserPreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserPreference`: AuthUserPreference
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.GetUserPreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specOptions** | **string** |  | 

### Return type

[**AuthUserPreference**](authUserPreference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPreference1

> AuthUserPreference GetUserPreference1(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecOptions(specOptions).Execute()

Get UserPreference object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    specOptions := "specOptions_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.GetUserPreference1(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecOptions(specOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.GetUserPreference1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserPreference1`: AuthUserPreference
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.GetUserPreference1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPreference1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specOptions** | **string** |  | 

### Return type

[**AuthUserPreference**](authUserPreference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsAuthorized

> AuthUser IsAuthorized(ctx, oTenant).Body(body).Execute()

Review authorization for user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewAuthSubjectAccessReviewRequest() // AuthSubjectAccessReviewRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.IsAuthorized(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.IsAuthorized``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IsAuthorized`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.IsAuthorized`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsAuthorizedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthSubjectAccessReviewRequest**](AuthSubjectAccessReviewRequest.md) |  | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsAuthorized1

> AuthUser IsAuthorized1(ctx, oName).Body(body).Execute()

Review authorization for user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    body := *openapiclient.NewAuthSubjectAccessReviewRequest() // AuthSubjectAccessReviewRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.IsAuthorized1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.IsAuthorized1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IsAuthorized1`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.IsAuthorized1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsAuthorized1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthSubjectAccessReviewRequest**](AuthSubjectAccessReviewRequest.md) |  | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelAuthenticationPolicy

> AuthAuthenticationPolicy LabelAuthenticationPolicy(ctx).Body(body).Execute()

Label AuthenticationPolicy object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewApiLabel() // ApiLabel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.LabelAuthenticationPolicy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.LabelAuthenticationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelAuthenticationPolicy`: AuthAuthenticationPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.LabelAuthenticationPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLabelAuthenticationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**AuthAuthenticationPolicy**](authAuthenticationPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelRole

> AuthRole LabelRole(ctx, oTenant).Body(body).Execute()

Label Role object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewApiLabel() // ApiLabel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.LabelRole(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.LabelRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelRole`: AuthRole
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.LabelRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**AuthRole**](authRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelRole1

> AuthRole LabelRole1(ctx, oName).Body(body).Execute()

Label Role object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    body := *openapiclient.NewApiLabel() // ApiLabel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.LabelRole1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.LabelRole1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelRole1`: AuthRole
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.LabelRole1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelRole1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**AuthRole**](authRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelRoleBinding

> AuthRoleBinding LabelRoleBinding(ctx, oTenant).Body(body).Execute()

Label RoleBinding object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewApiLabel() // ApiLabel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.LabelRoleBinding(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.LabelRoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelRoleBinding`: AuthRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.LabelRoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelRoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**AuthRoleBinding**](authRoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelRoleBinding1

> AuthRoleBinding LabelRoleBinding1(ctx, oName).Body(body).Execute()

Label RoleBinding object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    body := *openapiclient.NewApiLabel() // ApiLabel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.LabelRoleBinding1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.LabelRoleBinding1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelRoleBinding1`: AuthRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.LabelRoleBinding1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelRoleBinding1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**AuthRoleBinding**](authRoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelUser

> AuthUser LabelUser(ctx, oTenant).Body(body).Execute()

Label User object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewApiLabel() // ApiLabel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.LabelUser(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.LabelUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelUser`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.LabelUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelUser1

> AuthUser LabelUser1(ctx, oName).Body(body).Execute()

Label User object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    body := *openapiclient.NewApiLabel() // ApiLabel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.LabelUser1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.LabelUser1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelUser1`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.LabelUser1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelUser1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelUserPreference

> AuthUserPreference LabelUserPreference(ctx, oTenant).Body(body).Execute()

Label UserPreference object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewApiLabel() // ApiLabel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.LabelUserPreference(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.LabelUserPreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelUserPreference`: AuthUserPreference
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.LabelUserPreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelUserPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**AuthUserPreference**](authUserPreference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelUserPreference1

> AuthUserPreference LabelUserPreference1(ctx, oName).Body(body).Execute()

Label UserPreference object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    body := *openapiclient.NewApiLabel() // ApiLabel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.LabelUserPreference1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.LabelUserPreference1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelUserPreference1`: AuthUserPreference
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.LabelUserPreference1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelUserPreference1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**AuthUserPreference**](authUserPreference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LdapBindCheck

> AuthAuthenticationPolicy LdapBindCheck(ctx).Body(body).Execute()

Test LDAP bind operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAuthAuthenticationPolicy() // AuthAuthenticationPolicy | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.LdapBindCheck(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.LdapBindCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LdapBindCheck`: AuthAuthenticationPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.LdapBindCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLdapBindCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthAuthenticationPolicy**](AuthAuthenticationPolicy.md) |  | 

### Return type

[**AuthAuthenticationPolicy**](authAuthenticationPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LdapConnectionCheck

> AuthAuthenticationPolicy LdapConnectionCheck(ctx).Body(body).Execute()

Test LDAP connection

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAuthAuthenticationPolicy() // AuthAuthenticationPolicy | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.LdapConnectionCheck(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.LdapConnectionCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LdapConnectionCheck`: AuthAuthenticationPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.LdapConnectionCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLdapConnectionCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthAuthenticationPolicy**](AuthAuthenticationPolicy.md) |  | 

### Return type

[**AuthAuthenticationPolicy**](authAuthenticationPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRole

> AuthRoleList ListRole(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Role objects

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.ListRole(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.ListRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRole`: AuthRoleList
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.ListRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**AuthRoleList**](authRoleList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRole1

> AuthRoleList ListRole1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Role objects

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.ListRole1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.ListRole1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRole1`: AuthRoleList
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.ListRole1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRole1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**AuthRoleList**](authRoleList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleBinding

> AuthRoleBindingList ListRoleBinding(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List RoleBinding objects

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.ListRoleBinding(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.ListRoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleBinding`: AuthRoleBindingList
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.ListRoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**AuthRoleBindingList**](authRoleBindingList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleBinding1

> AuthRoleBindingList ListRoleBinding1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List RoleBinding objects

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.ListRoleBinding1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.ListRoleBinding1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleBinding1`: AuthRoleBindingList
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.ListRoleBinding1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoleBinding1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**AuthRoleBindingList**](authRoleBindingList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUser

> AuthUserList ListUser(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List User objects

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.ListUser(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.ListUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUser`: AuthUserList
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.ListUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**AuthUserList**](authUserList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUser1

> AuthUserList ListUser1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List User objects

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.ListUser1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.ListUser1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUser1`: AuthUserList
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.ListUser1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUser1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**AuthUserList**](authUserList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PasswordChange

> AuthUser PasswordChange(ctx, oTenant).Body(body).Execute()

Change user password

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewAuthPasswordChangeRequest() // AuthPasswordChangeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.PasswordChange(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.PasswordChange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PasswordChange`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.PasswordChange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPasswordChangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthPasswordChangeRequest**](AuthPasswordChangeRequest.md) |  | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PasswordChange1

> AuthUser PasswordChange1(ctx, oName).Body(body).Execute()

Change user password

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    body := *openapiclient.NewAuthPasswordChangeRequest() // AuthPasswordChangeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.PasswordChange1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.PasswordChange1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PasswordChange1`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.PasswordChange1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPasswordChange1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthPasswordChangeRequest**](AuthPasswordChangeRequest.md) |  | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PasswordReset

> AuthUser PasswordReset(ctx, oTenant).Body(body).Execute()

Reset user password

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewAuthPasswordResetRequest() // AuthPasswordResetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.PasswordReset(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.PasswordReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PasswordReset`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.PasswordReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPasswordResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthPasswordResetRequest**](AuthPasswordResetRequest.md) |  | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PasswordReset1

> AuthUser PasswordReset1(ctx, oName).Body(body).Execute()

Reset user password

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    body := *openapiclient.NewAuthPasswordResetRequest() // AuthPasswordResetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.PasswordReset1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.PasswordReset1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PasswordReset1`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.PasswordReset1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPasswordReset1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthPasswordResetRequest**](AuthPasswordResetRequest.md) |  | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenSecretGenerate

> AuthAuthenticationPolicy TokenSecretGenerate(ctx).Body(body).Execute()

Generate secret for token signing

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAuthTokenSecretRequest() // AuthTokenSecretRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.TokenSecretGenerate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.TokenSecretGenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenSecretGenerate`: AuthAuthenticationPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.TokenSecretGenerate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenSecretGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthTokenSecretRequest**](AuthTokenSecretRequest.md) |  | 

### Return type

[**AuthAuthenticationPolicy**](authAuthenticationPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unlock

> AuthUser Unlock(ctx, oTenant).Body(body).Execute()

Unlock user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewAuthUserUnlockRequest() // AuthUserUnlockRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.Unlock(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.Unlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Unlock`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.Unlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthUserUnlockRequest**](AuthUserUnlockRequest.md) |  | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unlock1

> AuthUser Unlock1(ctx, oName).Body(body).Execute()

Unlock user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    body := *openapiclient.NewAuthUserUnlockRequest() // AuthUserUnlockRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.Unlock1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.Unlock1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Unlock1`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.Unlock1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlock1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthUserUnlockRequest**](AuthUserUnlockRequest.md) |  | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthenticationPolicy

> AuthAuthenticationPolicy UpdateAuthenticationPolicy(ctx).Body(body).Execute()

Update AuthenticationPolicy object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAuthAuthenticationPolicy() // AuthAuthenticationPolicy | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.UpdateAuthenticationPolicy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.UpdateAuthenticationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthenticationPolicy`: AuthAuthenticationPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.UpdateAuthenticationPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthenticationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthAuthenticationPolicy**](AuthAuthenticationPolicy.md) |  | 

### Return type

[**AuthAuthenticationPolicy**](authAuthenticationPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> AuthRole UpdateRole(ctx, oTenant).Body(body).Execute()

Update Role object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewAuthRole() // AuthRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.UpdateRole(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.UpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRole`: AuthRole
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthRole**](AuthRole.md) |  | 

### Return type

[**AuthRole**](authRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole1

> AuthRole UpdateRole1(ctx, oName).Body(body).Execute()

Update Role object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    body := *openapiclient.NewAuthRole() // AuthRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.UpdateRole1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.UpdateRole1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRole1`: AuthRole
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.UpdateRole1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRole1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthRole**](AuthRole.md) |  | 

### Return type

[**AuthRole**](authRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleBinding

> AuthRoleBinding UpdateRoleBinding(ctx, oTenant).Body(body).Execute()

Update RoleBinding object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewAuthRoleBinding() // AuthRoleBinding | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.UpdateRoleBinding(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.UpdateRoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoleBinding`: AuthRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.UpdateRoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthRoleBinding**](AuthRoleBinding.md) |  | 

### Return type

[**AuthRoleBinding**](authRoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleBinding1

> AuthRoleBinding UpdateRoleBinding1(ctx, oName).Body(body).Execute()

Update RoleBinding object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    body := *openapiclient.NewAuthRoleBinding() // AuthRoleBinding | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.UpdateRoleBinding1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.UpdateRoleBinding1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoleBinding1`: AuthRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.UpdateRoleBinding1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleBinding1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthRoleBinding**](AuthRoleBinding.md) |  | 

### Return type

[**AuthRoleBinding**](authRoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> AuthUser UpdateUser(ctx, oTenant).Body(body).Execute()

Update User object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewAuthUser() // AuthUser | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.UpdateUser(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthUser**](AuthUser.md) |  | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser1

> AuthUser UpdateUser1(ctx, oName).Body(body).Execute()

Update User object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    body := *openapiclient.NewAuthUser() // AuthUser | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.UpdateUser1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.UpdateUser1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser1`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.UpdateUser1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUser1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthUser**](AuthUser.md) |  | 

### Return type

[**AuthUser**](authUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPreference

> AuthUserPreference UpdateUserPreference(ctx, oTenant).Body(body).Execute()

Update UserPreference object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewAuthUserPreference() // AuthUserPreference | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.UpdateUserPreference(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.UpdateUserPreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserPreference`: AuthUserPreference
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.UpdateUserPreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthUserPreference**](AuthUserPreference.md) |  | 

### Return type

[**AuthUserPreference**](authUserPreference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPreference1

> AuthUserPreference UpdateUserPreference1(ctx, oName).Body(body).Execute()

Update UserPreference object

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    body := *openapiclient.NewAuthUserPreference() // AuthUserPreference | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.UpdateUserPreference1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.UpdateUserPreference1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserPreference1`: AuthUserPreference
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.UpdateUserPreference1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPreference1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthUserPreference**](AuthUserPreference.md) |  | 

### Return type

[**AuthUserPreference**](authUserPreference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchAuthenticationPolicy

> AuthAutoMsgAuthenticationPolicyWatchHelper WatchAuthenticationPolicy(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch AuthenticationPolicy objects. Supports WebSockets or HTTP long poll

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.WatchAuthenticationPolicy(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.WatchAuthenticationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchAuthenticationPolicy`: AuthAutoMsgAuthenticationPolicyWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.WatchAuthenticationPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchAuthenticationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**AuthAutoMsgAuthenticationPolicyWatchHelper**](authAutoMsgAuthenticationPolicyWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchRole

> AuthAutoMsgRoleWatchHelper WatchRole(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Role objects. Supports WebSockets or HTTP long poll

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.WatchRole(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.WatchRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchRole`: AuthAutoMsgRoleWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.WatchRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**AuthAutoMsgRoleWatchHelper**](authAutoMsgRoleWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchRole1

> AuthAutoMsgRoleWatchHelper WatchRole1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Role objects. Supports WebSockets or HTTP long poll

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.WatchRole1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.WatchRole1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchRole1`: AuthAutoMsgRoleWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.WatchRole1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchRole1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**AuthAutoMsgRoleWatchHelper**](authAutoMsgRoleWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchRoleBinding

> AuthAutoMsgRoleBindingWatchHelper WatchRoleBinding(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch RoleBinding objects. Supports WebSockets or HTTP long poll

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.WatchRoleBinding(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.WatchRoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchRoleBinding`: AuthAutoMsgRoleBindingWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.WatchRoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchRoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**AuthAutoMsgRoleBindingWatchHelper**](authAutoMsgRoleBindingWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchRoleBinding1

> AuthAutoMsgRoleBindingWatchHelper WatchRoleBinding1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch RoleBinding objects. Supports WebSockets or HTTP long poll

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.WatchRoleBinding1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.WatchRoleBinding1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchRoleBinding1`: AuthAutoMsgRoleBindingWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.WatchRoleBinding1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchRoleBinding1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**AuthAutoMsgRoleBindingWatchHelper**](authAutoMsgRoleBindingWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchUser

> AuthAutoMsgUserWatchHelper WatchUser(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch User objects. Supports WebSockets or HTTP long poll

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.WatchUser(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.WatchUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchUser`: AuthAutoMsgUserWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.WatchUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**AuthAutoMsgUserWatchHelper**](authAutoMsgUserWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchUser1

> AuthAutoMsgUserWatchHelper WatchUser1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch User objects. Supports WebSockets or HTTP long poll

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.WatchUser1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.WatchUser1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchUser1`: AuthAutoMsgUserWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.WatchUser1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchUser1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**AuthAutoMsgUserWatchHelper**](authAutoMsgUserWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchUserPreference

> AuthAutoMsgUserPreferenceWatchHelper WatchUserPreference(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch UserPreference objects. Supports WebSockets or HTTP long poll

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.WatchUserPreference(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.WatchUserPreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchUserPreference`: AuthAutoMsgUserPreferenceWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.WatchUserPreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchUserPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**AuthAutoMsgUserPreferenceWatchHelper**](authAutoMsgUserPreferenceWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchUserPreference1

> AuthAutoMsgUserPreferenceWatchHelper WatchUserPreference1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch UserPreference objects. Supports WebSockets or HTTP long poll

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthV1Api.WatchUserPreference1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthV1Api.WatchUserPreference1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchUserPreference1`: AuthAutoMsgUserPreferenceWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `AuthV1Api.WatchUserPreference1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchUserPreference1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**AuthAutoMsgUserPreferenceWatchHelper**](authAutoMsgUserPreferenceWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

