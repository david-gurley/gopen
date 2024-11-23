# \NetworkV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIPAMPolicy**](NetworkV1Api.md#AddIPAMPolicy) | **Post** /configs/network/v1/tenant/{O.Tenant}/ipam-policies | Create IPAMPolicy object
[**AddIPAMPolicy1**](NetworkV1Api.md#AddIPAMPolicy1) | **Post** /configs/network/v1/ipam-policies | Create IPAMPolicy object
[**AddNetwork**](NetworkV1Api.md#AddNetwork) | **Post** /configs/network/v1/tenant/{O.Tenant}/networks | Create Network object
[**AddNetwork1**](NetworkV1Api.md#AddNetwork1) | **Post** /configs/network/v1/networks | Create Network object
[**AddPolicerProfile**](NetworkV1Api.md#AddPolicerProfile) | **Post** /configs/network/v1/tenant/{O.Tenant}/policer-profile | Create PolicerProfile object
[**AddPolicerProfile1**](NetworkV1Api.md#AddPolicerProfile1) | **Post** /configs/network/v1/policer-profile | Create PolicerProfile object
[**AddRoutingConfig**](NetworkV1Api.md#AddRoutingConfig) | **Post** /configs/network/v1/routing-config | Create RoutingConfig object
[**AddStaticBindings**](NetworkV1Api.md#AddStaticBindings) | **Post** /configs/network/v1/tenant/{O.Tenant}/ipam-policies/{O.Name}/AddStaticBindings | Add static bindings
[**AddStaticBindings1**](NetworkV1Api.md#AddStaticBindings1) | **Post** /configs/network/v1/ipam-policies/{O.Name}/AddStaticBindings | Add static bindings
[**AddVirtualRouter**](NetworkV1Api.md#AddVirtualRouter) | **Post** /configs/network/v1/tenant/{O.Tenant}/virtualrouters | Create VirtualRouter object
[**AddVirtualRouter1**](NetworkV1Api.md#AddVirtualRouter1) | **Post** /configs/network/v1/virtualrouters | Create VirtualRouter object
[**AddVirtualRouterPeeringGroup**](NetworkV1Api.md#AddVirtualRouterPeeringGroup) | **Post** /configs/network/v1/tenant/{O.Tenant}/virtual-router-peering-groups | Create VirtualRouterPeeringGroup object
[**AddVirtualRouterPeeringGroup1**](NetworkV1Api.md#AddVirtualRouterPeeringGroup1) | **Post** /configs/network/v1/virtual-router-peering-groups | Create VirtualRouterPeeringGroup object
[**DeleteIPAMPolicy**](NetworkV1Api.md#DeleteIPAMPolicy) | **Delete** /configs/network/v1/tenant/{O.Tenant}/ipam-policies/{O.Name} | Delete IPAMPolicy object
[**DeleteIPAMPolicy1**](NetworkV1Api.md#DeleteIPAMPolicy1) | **Delete** /configs/network/v1/ipam-policies/{O.Name} | Delete IPAMPolicy object
[**DeleteNetwork**](NetworkV1Api.md#DeleteNetwork) | **Delete** /configs/network/v1/tenant/{O.Tenant}/networks/{O.Name} | Delete Network object
[**DeleteNetwork1**](NetworkV1Api.md#DeleteNetwork1) | **Delete** /configs/network/v1/networks/{O.Name} | Delete Network object
[**DeletePolicerProfile**](NetworkV1Api.md#DeletePolicerProfile) | **Delete** /configs/network/v1/tenant/{O.Tenant}/policer-profile/{O.Name} | Delete PolicerProfile object
[**DeletePolicerProfile1**](NetworkV1Api.md#DeletePolicerProfile1) | **Delete** /configs/network/v1/policer-profile/{O.Name} | Delete PolicerProfile object
[**DeleteRoutingConfig**](NetworkV1Api.md#DeleteRoutingConfig) | **Delete** /configs/network/v1/routing-config/{O.Name} | Delete RoutingConfig object
[**DeleteVirtualRouter**](NetworkV1Api.md#DeleteVirtualRouter) | **Delete** /configs/network/v1/tenant/{O.Tenant}/virtualrouters/{O.Name} | Delete VirtualRouter object
[**DeleteVirtualRouter1**](NetworkV1Api.md#DeleteVirtualRouter1) | **Delete** /configs/network/v1/virtualrouters/{O.Name} | Delete VirtualRouter object
[**DeleteVirtualRouterPeeringGroup**](NetworkV1Api.md#DeleteVirtualRouterPeeringGroup) | **Delete** /configs/network/v1/tenant/{O.Tenant}/virtual-router-peering-groups/{O.Name} | Delete VirtualRouterPeeringGroup object
[**DeleteVirtualRouterPeeringGroup1**](NetworkV1Api.md#DeleteVirtualRouterPeeringGroup1) | **Delete** /configs/network/v1/virtual-router-peering-groups/{O.Name} | Delete VirtualRouterPeeringGroup object
[**GetIPAMPolicy**](NetworkV1Api.md#GetIPAMPolicy) | **Get** /configs/network/v1/tenant/{O.Tenant}/ipam-policies/{O.Name} | Get IPAMPolicy object
[**GetIPAMPolicy1**](NetworkV1Api.md#GetIPAMPolicy1) | **Get** /configs/network/v1/ipam-policies/{O.Name} | Get IPAMPolicy object
[**GetNetwork**](NetworkV1Api.md#GetNetwork) | **Get** /configs/network/v1/tenant/{O.Tenant}/networks/{O.Name} | Get Network object
[**GetNetwork1**](NetworkV1Api.md#GetNetwork1) | **Get** /configs/network/v1/networks/{O.Name} | Get Network object
[**GetNetworkInterface**](NetworkV1Api.md#GetNetworkInterface) | **Get** /configs/network/v1/networkinterfaces/{O.Name} | Get NetworkInterface object
[**GetPolicerProfile**](NetworkV1Api.md#GetPolicerProfile) | **Get** /configs/network/v1/tenant/{O.Tenant}/policer-profile/{O.Name} | Get PolicerProfile object
[**GetPolicerProfile1**](NetworkV1Api.md#GetPolicerProfile1) | **Get** /configs/network/v1/policer-profile/{O.Name} | Get PolicerProfile object
[**GetRouteTable**](NetworkV1Api.md#GetRouteTable) | **Get** /configs/network/v1/tenant/{O.Tenant}/route-tables/{O.Name} | Get RouteTable object
[**GetRouteTable1**](NetworkV1Api.md#GetRouteTable1) | **Get** /configs/network/v1/route-tables/{O.Name} | Get RouteTable object
[**GetRoutingConfig**](NetworkV1Api.md#GetRoutingConfig) | **Get** /configs/network/v1/routing-config/{O.Name} | Get RoutingConfig object
[**GetVirtualRouter**](NetworkV1Api.md#GetVirtualRouter) | **Get** /configs/network/v1/tenant/{O.Tenant}/virtualrouters/{O.Name} | Get VirtualRouter object
[**GetVirtualRouter1**](NetworkV1Api.md#GetVirtualRouter1) | **Get** /configs/network/v1/virtualrouters/{O.Name} | Get VirtualRouter object
[**GetVirtualRouterPeeringGroup**](NetworkV1Api.md#GetVirtualRouterPeeringGroup) | **Get** /configs/network/v1/tenant/{O.Tenant}/virtual-router-peering-groups/{O.Name} | Get VirtualRouterPeeringGroup object
[**GetVirtualRouterPeeringGroup1**](NetworkV1Api.md#GetVirtualRouterPeeringGroup1) | **Get** /configs/network/v1/virtual-router-peering-groups/{O.Name} | Get VirtualRouterPeeringGroup object
[**LabelIPAMPolicy**](NetworkV1Api.md#LabelIPAMPolicy) | **Post** /configs/network/v1/tenant/{O.Tenant}/ipam-policies/{O.Name}/label | Label IPAMPolicy object
[**LabelIPAMPolicy1**](NetworkV1Api.md#LabelIPAMPolicy1) | **Post** /configs/network/v1/ipam-policies/{O.Name}/label | Label IPAMPolicy object
[**LabelNetwork**](NetworkV1Api.md#LabelNetwork) | **Post** /configs/network/v1/tenant/{O.Tenant}/networks/{O.Name}/label | Label Network object
[**LabelNetwork1**](NetworkV1Api.md#LabelNetwork1) | **Post** /configs/network/v1/networks/{O.Name}/label | Label Network object
[**LabelNetworkInterface**](NetworkV1Api.md#LabelNetworkInterface) | **Post** /configs/network/v1/networkinterfaces/{O.Name}/label | Label NetworkInterface object
[**LabelPolicerProfile**](NetworkV1Api.md#LabelPolicerProfile) | **Post** /configs/network/v1/tenant/{O.Tenant}/policer-profile/{O.Name}/label | Label PolicerProfile object
[**LabelPolicerProfile1**](NetworkV1Api.md#LabelPolicerProfile1) | **Post** /configs/network/v1/policer-profile/{O.Name}/label | Label PolicerProfile object
[**LabelRoutingConfig**](NetworkV1Api.md#LabelRoutingConfig) | **Post** /configs/network/v1/routing-config/{O.Name}/label | Label RoutingConfig object
[**LabelVirtualRouter**](NetworkV1Api.md#LabelVirtualRouter) | **Post** /configs/network/v1/tenant/{O.Tenant}/virtualrouters/{O.Name}/label | Label VirtualRouter object
[**LabelVirtualRouter1**](NetworkV1Api.md#LabelVirtualRouter1) | **Post** /configs/network/v1/virtualrouters/{O.Name}/label | Label VirtualRouter object
[**LabelVirtualRouterPeeringGroup**](NetworkV1Api.md#LabelVirtualRouterPeeringGroup) | **Post** /configs/network/v1/tenant/{O.Tenant}/virtual-router-peering-groups/{O.Name}/label | Label VirtualRouterPeeringGroup object
[**LabelVirtualRouterPeeringGroup1**](NetworkV1Api.md#LabelVirtualRouterPeeringGroup1) | **Post** /configs/network/v1/virtual-router-peering-groups/{O.Name}/label | Label VirtualRouterPeeringGroup object
[**ListIPAMPolicy**](NetworkV1Api.md#ListIPAMPolicy) | **Get** /configs/network/v1/tenant/{O.Tenant}/ipam-policies | List IPAMPolicy objects
[**ListIPAMPolicy1**](NetworkV1Api.md#ListIPAMPolicy1) | **Get** /configs/network/v1/ipam-policies | List IPAMPolicy objects
[**ListNetwork**](NetworkV1Api.md#ListNetwork) | **Get** /configs/network/v1/tenant/{O.Tenant}/networks | List Network objects
[**ListNetwork1**](NetworkV1Api.md#ListNetwork1) | **Get** /configs/network/v1/networks | List Network objects
[**ListNetworkInterface**](NetworkV1Api.md#ListNetworkInterface) | **Get** /configs/network/v1/networkinterfaces | List NetworkInterface objects
[**ListPolicerProfile**](NetworkV1Api.md#ListPolicerProfile) | **Get** /configs/network/v1/tenant/{O.Tenant}/policer-profile | List PolicerProfile objects
[**ListPolicerProfile1**](NetworkV1Api.md#ListPolicerProfile1) | **Get** /configs/network/v1/policer-profile | List PolicerProfile objects
[**ListRouteTable**](NetworkV1Api.md#ListRouteTable) | **Get** /configs/network/v1/tenant/{O.Tenant}/route-tables | List RouteTable objects
[**ListRouteTable1**](NetworkV1Api.md#ListRouteTable1) | **Get** /configs/network/v1/route-tables | List RouteTable objects
[**ListRoutingConfig**](NetworkV1Api.md#ListRoutingConfig) | **Get** /configs/network/v1/routing-config | List RoutingConfig objects
[**ListVirtualRouter**](NetworkV1Api.md#ListVirtualRouter) | **Get** /configs/network/v1/tenant/{O.Tenant}/virtualrouters | List VirtualRouter objects
[**ListVirtualRouter1**](NetworkV1Api.md#ListVirtualRouter1) | **Get** /configs/network/v1/virtualrouters | List VirtualRouter objects
[**ListVirtualRouterPeeringGroup**](NetworkV1Api.md#ListVirtualRouterPeeringGroup) | **Get** /configs/network/v1/tenant/{O.Tenant}/virtual-router-peering-groups | List VirtualRouterPeeringGroup objects
[**ListVirtualRouterPeeringGroup1**](NetworkV1Api.md#ListVirtualRouterPeeringGroup1) | **Get** /configs/network/v1/virtual-router-peering-groups | List VirtualRouterPeeringGroup objects
[**UpdateIPAMPolicy**](NetworkV1Api.md#UpdateIPAMPolicy) | **Put** /configs/network/v1/tenant/{O.Tenant}/ipam-policies/{O.Name} | Update IPAMPolicy object
[**UpdateIPAMPolicy1**](NetworkV1Api.md#UpdateIPAMPolicy1) | **Put** /configs/network/v1/ipam-policies/{O.Name} | Update IPAMPolicy object
[**UpdateNetwork**](NetworkV1Api.md#UpdateNetwork) | **Put** /configs/network/v1/tenant/{O.Tenant}/networks/{O.Name} | Update Network object
[**UpdateNetwork1**](NetworkV1Api.md#UpdateNetwork1) | **Put** /configs/network/v1/networks/{O.Name} | Update Network object
[**UpdateNetworkInterface**](NetworkV1Api.md#UpdateNetworkInterface) | **Put** /configs/network/v1/networkinterfaces/{O.Name} | Update NetworkInterface object
[**UpdatePolicerProfile**](NetworkV1Api.md#UpdatePolicerProfile) | **Put** /configs/network/v1/tenant/{O.Tenant}/policer-profile/{O.Name} | Update PolicerProfile object
[**UpdatePolicerProfile1**](NetworkV1Api.md#UpdatePolicerProfile1) | **Put** /configs/network/v1/policer-profile/{O.Name} | Update PolicerProfile object
[**UpdateRoutingConfig**](NetworkV1Api.md#UpdateRoutingConfig) | **Put** /configs/network/v1/routing-config/{O.Name} | Update RoutingConfig object
[**UpdateVirtualRouter**](NetworkV1Api.md#UpdateVirtualRouter) | **Put** /configs/network/v1/tenant/{O.Tenant}/virtualrouters/{O.Name} | Update VirtualRouter object
[**UpdateVirtualRouter1**](NetworkV1Api.md#UpdateVirtualRouter1) | **Put** /configs/network/v1/virtualrouters/{O.Name} | Update VirtualRouter object
[**UpdateVirtualRouterPeeringGroup**](NetworkV1Api.md#UpdateVirtualRouterPeeringGroup) | **Put** /configs/network/v1/tenant/{O.Tenant}/virtual-router-peering-groups/{O.Name} | Update VirtualRouterPeeringGroup object
[**UpdateVirtualRouterPeeringGroup1**](NetworkV1Api.md#UpdateVirtualRouterPeeringGroup1) | **Put** /configs/network/v1/virtual-router-peering-groups/{O.Name} | Update VirtualRouterPeeringGroup object
[**WatchIPAMPolicy**](NetworkV1Api.md#WatchIPAMPolicy) | **Get** /configs/network/v1/watch/tenant/{O.Tenant}/ipam-policies | Watch IPAMPolicy objects. Supports WebSockets or HTTP long poll
[**WatchIPAMPolicy1**](NetworkV1Api.md#WatchIPAMPolicy1) | **Get** /configs/network/v1/watch/ipam-policies | Watch IPAMPolicy objects. Supports WebSockets or HTTP long poll
[**WatchNetwork**](NetworkV1Api.md#WatchNetwork) | **Get** /configs/network/v1/watch/tenant/{O.Tenant}/networks | Watch Network objects. Supports WebSockets or HTTP long poll
[**WatchNetwork1**](NetworkV1Api.md#WatchNetwork1) | **Get** /configs/network/v1/watch/networks | Watch Network objects. Supports WebSockets or HTTP long poll
[**WatchNetworkInterface**](NetworkV1Api.md#WatchNetworkInterface) | **Get** /configs/network/v1/watch/networkinterfaces | Watch NetworkInterface objects. Supports WebSockets or HTTP long poll
[**WatchPolicerProfile**](NetworkV1Api.md#WatchPolicerProfile) | **Get** /configs/network/v1/watch/tenant/{O.Tenant}/policer-profile | Watch PolicerProfile objects. Supports WebSockets or HTTP long poll
[**WatchPolicerProfile1**](NetworkV1Api.md#WatchPolicerProfile1) | **Get** /configs/network/v1/watch/policer-profile | Watch PolicerProfile objects. Supports WebSockets or HTTP long poll
[**WatchRouteTable**](NetworkV1Api.md#WatchRouteTable) | **Get** /configs/network/v1/watch/tenant/{O.Tenant}/route-tables | Watch RouteTable objects. Supports WebSockets or HTTP long poll
[**WatchRouteTable1**](NetworkV1Api.md#WatchRouteTable1) | **Get** /configs/network/v1/watch/route-tables | Watch RouteTable objects. Supports WebSockets or HTTP long poll
[**WatchRoutingConfig**](NetworkV1Api.md#WatchRoutingConfig) | **Get** /configs/network/v1/watch/routing-config | Watch RoutingConfig objects. Supports WebSockets or HTTP long poll
[**WatchVirtualRouter**](NetworkV1Api.md#WatchVirtualRouter) | **Get** /configs/network/v1/watch/tenant/{O.Tenant}/virtualrouters | Watch VirtualRouter objects. Supports WebSockets or HTTP long poll
[**WatchVirtualRouter1**](NetworkV1Api.md#WatchVirtualRouter1) | **Get** /configs/network/v1/watch/virtualrouters | Watch VirtualRouter objects. Supports WebSockets or HTTP long poll
[**WatchVirtualRouterPeeringGroup**](NetworkV1Api.md#WatchVirtualRouterPeeringGroup) | **Get** /configs/network/v1/watch/tenant/{O.Tenant}/virtual-router-peering-groups | Watch VirtualRouterPeeringGroup objects. Supports WebSockets or HTTP long poll
[**WatchVirtualRouterPeeringGroup1**](NetworkV1Api.md#WatchVirtualRouterPeeringGroup1) | **Get** /configs/network/v1/watch/virtual-router-peering-groups | Watch VirtualRouterPeeringGroup objects. Supports WebSockets or HTTP long poll



## AddIPAMPolicy

> NetworkIPAMPolicy AddIPAMPolicy(ctx, oTenant).Body(body).Execute()

Create IPAMPolicy object

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
    body := *openapiclient.NewNetworkIPAMPolicy() // NetworkIPAMPolicy | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.AddIPAMPolicy(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.AddIPAMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIPAMPolicy`: NetworkIPAMPolicy
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.AddIPAMPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddIPAMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkIPAMPolicy**](NetworkIPAMPolicy.md) |  | 

### Return type

[**NetworkIPAMPolicy**](networkIPAMPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddIPAMPolicy1

> NetworkIPAMPolicy AddIPAMPolicy1(ctx).Body(body).Execute()

Create IPAMPolicy object

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
    body := *openapiclient.NewNetworkIPAMPolicy() // NetworkIPAMPolicy | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.AddIPAMPolicy1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.AddIPAMPolicy1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIPAMPolicy1`: NetworkIPAMPolicy
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.AddIPAMPolicy1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddIPAMPolicy1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NetworkIPAMPolicy**](NetworkIPAMPolicy.md) |  | 

### Return type

[**NetworkIPAMPolicy**](networkIPAMPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddNetwork

> NetworkNetwork AddNetwork(ctx, oTenant).Body(body).Execute()

Create Network object

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
    body := *openapiclient.NewNetworkNetwork() // NetworkNetwork | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.AddNetwork(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.AddNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddNetwork`: NetworkNetwork
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.AddNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkNetwork**](NetworkNetwork.md) |  | 

### Return type

[**NetworkNetwork**](networkNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddNetwork1

> NetworkNetwork AddNetwork1(ctx).Body(body).Execute()

Create Network object

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
    body := *openapiclient.NewNetworkNetwork() // NetworkNetwork | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.AddNetwork1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.AddNetwork1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddNetwork1`: NetworkNetwork
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.AddNetwork1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddNetwork1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NetworkNetwork**](NetworkNetwork.md) |  | 

### Return type

[**NetworkNetwork**](networkNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddPolicerProfile

> NetworkPolicerProfile AddPolicerProfile(ctx, oTenant).Body(body).Execute()

Create PolicerProfile object

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
    body := *openapiclient.NewNetworkPolicerProfile() // NetworkPolicerProfile | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.AddPolicerProfile(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.AddPolicerProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPolicerProfile`: NetworkPolicerProfile
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.AddPolicerProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPolicerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkPolicerProfile**](NetworkPolicerProfile.md) |  | 

### Return type

[**NetworkPolicerProfile**](networkPolicerProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddPolicerProfile1

> NetworkPolicerProfile AddPolicerProfile1(ctx).Body(body).Execute()

Create PolicerProfile object

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
    body := *openapiclient.NewNetworkPolicerProfile() // NetworkPolicerProfile | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.AddPolicerProfile1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.AddPolicerProfile1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPolicerProfile1`: NetworkPolicerProfile
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.AddPolicerProfile1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPolicerProfile1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NetworkPolicerProfile**](NetworkPolicerProfile.md) |  | 

### Return type

[**NetworkPolicerProfile**](networkPolicerProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRoutingConfig

> NetworkRoutingConfig AddRoutingConfig(ctx).Body(body).Execute()

Create RoutingConfig object

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
    body := *openapiclient.NewNetworkRoutingConfig() // NetworkRoutingConfig | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.AddRoutingConfig(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.AddRoutingConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRoutingConfig`: NetworkRoutingConfig
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.AddRoutingConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRoutingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NetworkRoutingConfig**](NetworkRoutingConfig.md) |  | 

### Return type

[**NetworkRoutingConfig**](networkRoutingConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddStaticBindings

> NetworkIPAMPolicy AddStaticBindings(ctx, oTenant).Body(body).Execute()

Add static bindings

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
    body := *openapiclient.NewNetworkAddStaticBindingsRequest() // NetworkAddStaticBindingsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.AddStaticBindings(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.AddStaticBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddStaticBindings`: NetworkIPAMPolicy
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.AddStaticBindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddStaticBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkAddStaticBindingsRequest**](NetworkAddStaticBindingsRequest.md) |  | 

### Return type

[**NetworkIPAMPolicy**](networkIPAMPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddStaticBindings1

> NetworkIPAMPolicy AddStaticBindings1(ctx, oName).Body(body).Execute()

Add static bindings

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
    body := *openapiclient.NewNetworkAddStaticBindingsRequest() // NetworkAddStaticBindingsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.AddStaticBindings1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.AddStaticBindings1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddStaticBindings1`: NetworkIPAMPolicy
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.AddStaticBindings1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddStaticBindings1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkAddStaticBindingsRequest**](NetworkAddStaticBindingsRequest.md) |  | 

### Return type

[**NetworkIPAMPolicy**](networkIPAMPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVirtualRouter

> NetworkVirtualRouter AddVirtualRouter(ctx, oTenant).Body(body).Execute()

Create VirtualRouter object

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
    body := *openapiclient.NewNetworkVirtualRouter() // NetworkVirtualRouter | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.AddVirtualRouter(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.AddVirtualRouter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVirtualRouter`: NetworkVirtualRouter
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.AddVirtualRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVirtualRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkVirtualRouter**](NetworkVirtualRouter.md) |  | 

### Return type

[**NetworkVirtualRouter**](networkVirtualRouter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVirtualRouter1

> NetworkVirtualRouter AddVirtualRouter1(ctx).Body(body).Execute()

Create VirtualRouter object

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
    body := *openapiclient.NewNetworkVirtualRouter() // NetworkVirtualRouter | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.AddVirtualRouter1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.AddVirtualRouter1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVirtualRouter1`: NetworkVirtualRouter
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.AddVirtualRouter1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddVirtualRouter1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NetworkVirtualRouter**](NetworkVirtualRouter.md) |  | 

### Return type

[**NetworkVirtualRouter**](networkVirtualRouter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVirtualRouterPeeringGroup

> NetworkVirtualRouterPeeringGroup AddVirtualRouterPeeringGroup(ctx, oTenant).Body(body).Execute()

Create VirtualRouterPeeringGroup object

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
    body := *openapiclient.NewNetworkVirtualRouterPeeringGroup() // NetworkVirtualRouterPeeringGroup | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.AddVirtualRouterPeeringGroup(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.AddVirtualRouterPeeringGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVirtualRouterPeeringGroup`: NetworkVirtualRouterPeeringGroup
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.AddVirtualRouterPeeringGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVirtualRouterPeeringGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkVirtualRouterPeeringGroup**](NetworkVirtualRouterPeeringGroup.md) |  | 

### Return type

[**NetworkVirtualRouterPeeringGroup**](networkVirtualRouterPeeringGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVirtualRouterPeeringGroup1

> NetworkVirtualRouterPeeringGroup AddVirtualRouterPeeringGroup1(ctx).Body(body).Execute()

Create VirtualRouterPeeringGroup object

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
    body := *openapiclient.NewNetworkVirtualRouterPeeringGroup() // NetworkVirtualRouterPeeringGroup | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.AddVirtualRouterPeeringGroup1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.AddVirtualRouterPeeringGroup1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVirtualRouterPeeringGroup1`: NetworkVirtualRouterPeeringGroup
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.AddVirtualRouterPeeringGroup1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddVirtualRouterPeeringGroup1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NetworkVirtualRouterPeeringGroup**](NetworkVirtualRouterPeeringGroup.md) |  | 

### Return type

[**NetworkVirtualRouterPeeringGroup**](networkVirtualRouterPeeringGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIPAMPolicy

> NetworkIPAMPolicy DeleteIPAMPolicy(ctx, oTenant).Execute()

Delete IPAMPolicy object

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
    resp, r, err := api_client.NetworkV1Api.DeleteIPAMPolicy(context.Background(), oTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.DeleteIPAMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIPAMPolicy`: NetworkIPAMPolicy
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.DeleteIPAMPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIPAMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkIPAMPolicy**](networkIPAMPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIPAMPolicy1

> NetworkIPAMPolicy DeleteIPAMPolicy1(ctx, oName).Execute()

Delete IPAMPolicy object

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
    resp, r, err := api_client.NetworkV1Api.DeleteIPAMPolicy1(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.DeleteIPAMPolicy1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIPAMPolicy1`: NetworkIPAMPolicy
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.DeleteIPAMPolicy1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIPAMPolicy1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkIPAMPolicy**](networkIPAMPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetwork

> NetworkNetwork DeleteNetwork(ctx, oTenant).Execute()

Delete Network object

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
    resp, r, err := api_client.NetworkV1Api.DeleteNetwork(context.Background(), oTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.DeleteNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetwork`: NetworkNetwork
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.DeleteNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkNetwork**](networkNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetwork1

> NetworkNetwork DeleteNetwork1(ctx, oName).Execute()

Delete Network object

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
    resp, r, err := api_client.NetworkV1Api.DeleteNetwork1(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.DeleteNetwork1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetwork1`: NetworkNetwork
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.DeleteNetwork1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetwork1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkNetwork**](networkNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicerProfile

> NetworkPolicerProfile DeletePolicerProfile(ctx, oTenant).Execute()

Delete PolicerProfile object

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
    resp, r, err := api_client.NetworkV1Api.DeletePolicerProfile(context.Background(), oTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.DeletePolicerProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePolicerProfile`: NetworkPolicerProfile
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.DeletePolicerProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkPolicerProfile**](networkPolicerProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicerProfile1

> NetworkPolicerProfile DeletePolicerProfile1(ctx, oName).Execute()

Delete PolicerProfile object

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
    resp, r, err := api_client.NetworkV1Api.DeletePolicerProfile1(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.DeletePolicerProfile1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePolicerProfile1`: NetworkPolicerProfile
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.DeletePolicerProfile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicerProfile1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkPolicerProfile**](networkPolicerProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoutingConfig

> NetworkRoutingConfig DeleteRoutingConfig(ctx, oName).Execute()

Delete RoutingConfig object

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
    resp, r, err := api_client.NetworkV1Api.DeleteRoutingConfig(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.DeleteRoutingConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRoutingConfig`: NetworkRoutingConfig
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.DeleteRoutingConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoutingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkRoutingConfig**](networkRoutingConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualRouter

> NetworkVirtualRouter DeleteVirtualRouter(ctx, oTenant).Execute()

Delete VirtualRouter object

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
    resp, r, err := api_client.NetworkV1Api.DeleteVirtualRouter(context.Background(), oTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.DeleteVirtualRouter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualRouter`: NetworkVirtualRouter
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.DeleteVirtualRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkVirtualRouter**](networkVirtualRouter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualRouter1

> NetworkVirtualRouter DeleteVirtualRouter1(ctx, oName).Execute()

Delete VirtualRouter object

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
    resp, r, err := api_client.NetworkV1Api.DeleteVirtualRouter1(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.DeleteVirtualRouter1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualRouter1`: NetworkVirtualRouter
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.DeleteVirtualRouter1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualRouter1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkVirtualRouter**](networkVirtualRouter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualRouterPeeringGroup

> NetworkVirtualRouterPeeringGroup DeleteVirtualRouterPeeringGroup(ctx, oTenant).Execute()

Delete VirtualRouterPeeringGroup object

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
    resp, r, err := api_client.NetworkV1Api.DeleteVirtualRouterPeeringGroup(context.Background(), oTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.DeleteVirtualRouterPeeringGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualRouterPeeringGroup`: NetworkVirtualRouterPeeringGroup
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.DeleteVirtualRouterPeeringGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualRouterPeeringGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkVirtualRouterPeeringGroup**](networkVirtualRouterPeeringGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualRouterPeeringGroup1

> NetworkVirtualRouterPeeringGroup DeleteVirtualRouterPeeringGroup1(ctx, oName).Execute()

Delete VirtualRouterPeeringGroup object

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
    resp, r, err := api_client.NetworkV1Api.DeleteVirtualRouterPeeringGroup1(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.DeleteVirtualRouterPeeringGroup1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualRouterPeeringGroup1`: NetworkVirtualRouterPeeringGroup
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.DeleteVirtualRouterPeeringGroup1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualRouterPeeringGroup1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkVirtualRouterPeeringGroup**](networkVirtualRouterPeeringGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIPAMPolicy

> NetworkIPAMPolicy GetIPAMPolicy(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).SpecType(specType).IpamOptionsLease(ipamOptionsLease).IpamOptionsRouters(ipamOptionsRouters).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()

Get IPAMPolicy object

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
    specType := "specType_example" // string |  (optional)
    ipamOptionsLease := int64(789) // int64 |  (optional)
    ipamOptionsRouters := []string{"Inner_example"} // []string |  (optional)
    propagationStatusPendingDscs := []string{"Inner_example"} // []string | list of DSCs where propagation did not complete. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.GetIPAMPolicy(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).SpecType(specType).IpamOptionsLease(ipamOptionsLease).IpamOptionsRouters(ipamOptionsRouters).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.GetIPAMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIPAMPolicy`: NetworkIPAMPolicy
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.GetIPAMPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIPAMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specType** | **string** |  | 
 **ipamOptionsLease** | **int64** |  | 
 **ipamOptionsRouters** | **[]string** |  | 
 **propagationStatusPendingDscs** | **[]string** | list of DSCs where propagation did not complete. | 

### Return type

[**NetworkIPAMPolicy**](networkIPAMPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIPAMPolicy1

> NetworkIPAMPolicy GetIPAMPolicy1(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecType(specType).IpamOptionsLease(ipamOptionsLease).IpamOptionsRouters(ipamOptionsRouters).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()

Get IPAMPolicy object

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
    specType := "specType_example" // string |  (optional)
    ipamOptionsLease := int64(789) // int64 |  (optional)
    ipamOptionsRouters := []string{"Inner_example"} // []string |  (optional)
    propagationStatusPendingDscs := []string{"Inner_example"} // []string | list of DSCs where propagation did not complete. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.GetIPAMPolicy1(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecType(specType).IpamOptionsLease(ipamOptionsLease).IpamOptionsRouters(ipamOptionsRouters).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.GetIPAMPolicy1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIPAMPolicy1`: NetworkIPAMPolicy
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.GetIPAMPolicy1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIPAMPolicy1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specType** | **string** |  | 
 **ipamOptionsLease** | **int64** |  | 
 **ipamOptionsRouters** | **[]string** |  | 
 **propagationStatusPendingDscs** | **[]string** | list of DSCs where propagation did not complete. | 

### Return type

[**NetworkIPAMPolicy**](networkIPAMPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetwork

> NetworkNetwork GetNetwork(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).AdminValueFormat(adminValueFormat).AdminValueValue(adminValueValue).SpecIngressSecurityPolicy(specIngressSecurityPolicy).Execute()

Get Network object

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
    adminValueFormat := "adminValueFormat_example" // string |  (optional)
    adminValueValue := int64(789) // int64 |  (optional)
    specIngressSecurityPolicy := []string{"Inner_example"} // []string | Security Policy to apply in the ingress direction. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.GetNetwork(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).AdminValueFormat(adminValueFormat).AdminValueValue(adminValueValue).SpecIngressSecurityPolicy(specIngressSecurityPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.GetNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetwork`: NetworkNetwork
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.GetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **adminValueFormat** | **string** |  | 
 **adminValueValue** | **int64** |  | 
 **specIngressSecurityPolicy** | **[]string** | Security Policy to apply in the ingress direction. | 

### Return type

[**NetworkNetwork**](networkNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetwork1

> NetworkNetwork GetNetwork1(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).AdminValueFormat(adminValueFormat).AdminValueValue(adminValueValue).SpecIngressSecurityPolicy(specIngressSecurityPolicy).Execute()

Get Network object

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
    adminValueFormat := "adminValueFormat_example" // string |  (optional)
    adminValueValue := int64(789) // int64 |  (optional)
    specIngressSecurityPolicy := []string{"Inner_example"} // []string | Security Policy to apply in the ingress direction. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.GetNetwork1(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).AdminValueFormat(adminValueFormat).AdminValueValue(adminValueValue).SpecIngressSecurityPolicy(specIngressSecurityPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.GetNetwork1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetwork1`: NetworkNetwork
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.GetNetwork1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetwork1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **adminValueFormat** | **string** |  | 
 **adminValueValue** | **int64** |  | 
 **specIngressSecurityPolicy** | **[]string** | Security Policy to apply in the ingress direction. | 

### Return type

[**NetworkNetwork**](networkNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkInterface

> NetworkNetworkInterface GetNetworkInterface(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecAttachTenant(specAttachTenant).IpConfigDnsServers(ipConfigDnsServers).StatusMirrorSessions(statusMirrorSessions).Execute()

Get NetworkInterface object

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
    specAttachTenant := "specAttachTenant_example" // string |  (optional)
    ipConfigDnsServers := []string{"Inner_example"} // []string | DNSServers contains a list of DNS Servers that can be used on DistributedServiceCard. (optional)
    statusMirrorSessions := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.GetNetworkInterface(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecAttachTenant(specAttachTenant).IpConfigDnsServers(ipConfigDnsServers).StatusMirrorSessions(statusMirrorSessions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.GetNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkInterface`: NetworkNetworkInterface
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.GetNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specAttachTenant** | **string** |  | 
 **ipConfigDnsServers** | **[]string** | DNSServers contains a list of DNS Servers that can be used on DistributedServiceCard. | 
 **statusMirrorSessions** | **[]string** |  | 

### Return type

[**NetworkNetworkInterface**](networkNetworkInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicerProfile

> NetworkPolicerProfile GetPolicerProfile(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).ExceedActionPolicerAction(exceedActionPolicerAction).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()

Get PolicerProfile object

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
    exceedActionPolicerAction := "exceedActionPolicerAction_example" // string |  (optional)
    propagationStatusPendingDscs := []string{"Inner_example"} // []string | list of DSCs where propagation did not complete. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.GetPolicerProfile(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).ExceedActionPolicerAction(exceedActionPolicerAction).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.GetPolicerProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicerProfile`: NetworkPolicerProfile
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.GetPolicerProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **exceedActionPolicerAction** | **string** |  | 
 **propagationStatusPendingDscs** | **[]string** | list of DSCs where propagation did not complete. | 

### Return type

[**NetworkPolicerProfile**](networkPolicerProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicerProfile1

> NetworkPolicerProfile GetPolicerProfile1(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).ExceedActionPolicerAction(exceedActionPolicerAction).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()

Get PolicerProfile object

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
    exceedActionPolicerAction := "exceedActionPolicerAction_example" // string |  (optional)
    propagationStatusPendingDscs := []string{"Inner_example"} // []string | list of DSCs where propagation did not complete. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.GetPolicerProfile1(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).ExceedActionPolicerAction(exceedActionPolicerAction).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.GetPolicerProfile1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicerProfile1`: NetworkPolicerProfile
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.GetPolicerProfile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicerProfile1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **exceedActionPolicerAction** | **string** |  | 
 **propagationStatusPendingDscs** | **[]string** | list of DSCs where propagation did not complete. | 

### Return type

[**NetworkPolicerProfile**](networkPolicerProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteTable

> NetworkRouteTable GetRouteTable(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()

Get RouteTable object

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
    resp, r, err := api_client.NetworkV1Api.GetRouteTable(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.GetRouteTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRouteTable`: NetworkRouteTable
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.GetRouteTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**NetworkRouteTable**](networkRouteTable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteTable1

> NetworkRouteTable GetRouteTable1(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()

Get RouteTable object

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
    resp, r, err := api_client.NetworkV1Api.GetRouteTable1(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.GetRouteTable1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRouteTable1`: NetworkRouteTable
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.GetRouteTable1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteTable1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**NetworkRouteTable**](networkRouteTable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutingConfig

> NetworkRoutingConfig GetRoutingConfig(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).AsNumberASNumber(asNumberASNumber).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()

Get RoutingConfig object

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
    asNumberASNumber := int64(789) // int64 |  (optional)
    propagationStatusPendingDscs := []string{"Inner_example"} // []string | list of DSCs where propagation did not complete. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.GetRoutingConfig(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).AsNumberASNumber(asNumberASNumber).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.GetRoutingConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoutingConfig`: NetworkRoutingConfig
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.GetRoutingConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **asNumberASNumber** | **int64** |  | 
 **propagationStatusPendingDscs** | **[]string** | list of DSCs where propagation did not complete. | 

### Return type

[**NetworkRoutingConfig**](networkRoutingConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualRouter

> NetworkVirtualRouter GetVirtualRouter(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).SpecType(specType).AdminValueValue(adminValueValue).SpecIngressSecurityPolicy(specIngressSecurityPolicy).Execute()

Get VirtualRouter object

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
    specType := "specType_example" // string |  (optional)
    adminValueValue := int64(789) // int64 |  (optional)
    specIngressSecurityPolicy := []string{"Inner_example"} // []string | Security Policy to apply in the ingress direction. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.GetVirtualRouter(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).SpecType(specType).AdminValueValue(adminValueValue).SpecIngressSecurityPolicy(specIngressSecurityPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.GetVirtualRouter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualRouter`: NetworkVirtualRouter
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.GetVirtualRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specType** | **string** |  | 
 **adminValueValue** | **int64** |  | 
 **specIngressSecurityPolicy** | **[]string** | Security Policy to apply in the ingress direction. | 

### Return type

[**NetworkVirtualRouter**](networkVirtualRouter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualRouter1

> NetworkVirtualRouter GetVirtualRouter1(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecType(specType).AdminValueValue(adminValueValue).SpecIngressSecurityPolicy(specIngressSecurityPolicy).Execute()

Get VirtualRouter object

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
    specType := "specType_example" // string |  (optional)
    adminValueValue := int64(789) // int64 |  (optional)
    specIngressSecurityPolicy := []string{"Inner_example"} // []string | Security Policy to apply in the ingress direction. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.GetVirtualRouter1(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecType(specType).AdminValueValue(adminValueValue).SpecIngressSecurityPolicy(specIngressSecurityPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.GetVirtualRouter1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualRouter1`: NetworkVirtualRouter
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.GetVirtualRouter1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualRouter1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specType** | **string** |  | 
 **adminValueValue** | **int64** |  | 
 **specIngressSecurityPolicy** | **[]string** | Security Policy to apply in the ingress direction. | 

### Return type

[**NetworkVirtualRouter**](networkVirtualRouter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualRouterPeeringGroup

> NetworkVirtualRouterPeeringGroup GetVirtualRouterPeeringGroup(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()

Get VirtualRouterPeeringGroup object

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
    propagationStatusPendingDscs := []string{"Inner_example"} // []string | list of DSCs where propagation did not complete. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.GetVirtualRouterPeeringGroup(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.GetVirtualRouterPeeringGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualRouterPeeringGroup`: NetworkVirtualRouterPeeringGroup
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.GetVirtualRouterPeeringGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualRouterPeeringGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **propagationStatusPendingDscs** | **[]string** | list of DSCs where propagation did not complete. | 

### Return type

[**NetworkVirtualRouterPeeringGroup**](networkVirtualRouterPeeringGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualRouterPeeringGroup1

> NetworkVirtualRouterPeeringGroup GetVirtualRouterPeeringGroup1(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()

Get VirtualRouterPeeringGroup object

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
    propagationStatusPendingDscs := []string{"Inner_example"} // []string | list of DSCs where propagation did not complete. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.GetVirtualRouterPeeringGroup1(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.GetVirtualRouterPeeringGroup1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualRouterPeeringGroup1`: NetworkVirtualRouterPeeringGroup
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.GetVirtualRouterPeeringGroup1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualRouterPeeringGroup1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **propagationStatusPendingDscs** | **[]string** | list of DSCs where propagation did not complete. | 

### Return type

[**NetworkVirtualRouterPeeringGroup**](networkVirtualRouterPeeringGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelIPAMPolicy

> NetworkIPAMPolicy LabelIPAMPolicy(ctx, oTenant).Body(body).Execute()

Label IPAMPolicy object

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
    resp, r, err := api_client.NetworkV1Api.LabelIPAMPolicy(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.LabelIPAMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelIPAMPolicy`: NetworkIPAMPolicy
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.LabelIPAMPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelIPAMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**NetworkIPAMPolicy**](networkIPAMPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelIPAMPolicy1

> NetworkIPAMPolicy LabelIPAMPolicy1(ctx, oName).Body(body).Execute()

Label IPAMPolicy object

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
    resp, r, err := api_client.NetworkV1Api.LabelIPAMPolicy1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.LabelIPAMPolicy1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelIPAMPolicy1`: NetworkIPAMPolicy
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.LabelIPAMPolicy1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelIPAMPolicy1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**NetworkIPAMPolicy**](networkIPAMPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelNetwork

> NetworkNetwork LabelNetwork(ctx, oTenant).Body(body).Execute()

Label Network object

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
    resp, r, err := api_client.NetworkV1Api.LabelNetwork(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.LabelNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelNetwork`: NetworkNetwork
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.LabelNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**NetworkNetwork**](networkNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelNetwork1

> NetworkNetwork LabelNetwork1(ctx, oName).Body(body).Execute()

Label Network object

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
    resp, r, err := api_client.NetworkV1Api.LabelNetwork1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.LabelNetwork1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelNetwork1`: NetworkNetwork
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.LabelNetwork1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelNetwork1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**NetworkNetwork**](networkNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelNetworkInterface

> NetworkNetworkInterface LabelNetworkInterface(ctx, oName).Body(body).Execute()

Label NetworkInterface object

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
    resp, r, err := api_client.NetworkV1Api.LabelNetworkInterface(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.LabelNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelNetworkInterface`: NetworkNetworkInterface
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.LabelNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**NetworkNetworkInterface**](networkNetworkInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelPolicerProfile

> NetworkPolicerProfile LabelPolicerProfile(ctx, oTenant).Body(body).Execute()

Label PolicerProfile object

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
    resp, r, err := api_client.NetworkV1Api.LabelPolicerProfile(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.LabelPolicerProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelPolicerProfile`: NetworkPolicerProfile
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.LabelPolicerProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelPolicerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**NetworkPolicerProfile**](networkPolicerProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelPolicerProfile1

> NetworkPolicerProfile LabelPolicerProfile1(ctx, oName).Body(body).Execute()

Label PolicerProfile object

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
    resp, r, err := api_client.NetworkV1Api.LabelPolicerProfile1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.LabelPolicerProfile1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelPolicerProfile1`: NetworkPolicerProfile
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.LabelPolicerProfile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelPolicerProfile1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**NetworkPolicerProfile**](networkPolicerProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelRoutingConfig

> NetworkRoutingConfig LabelRoutingConfig(ctx, oName).Body(body).Execute()

Label RoutingConfig object

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
    resp, r, err := api_client.NetworkV1Api.LabelRoutingConfig(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.LabelRoutingConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelRoutingConfig`: NetworkRoutingConfig
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.LabelRoutingConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelRoutingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**NetworkRoutingConfig**](networkRoutingConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelVirtualRouter

> NetworkVirtualRouter LabelVirtualRouter(ctx, oTenant).Body(body).Execute()

Label VirtualRouter object

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
    resp, r, err := api_client.NetworkV1Api.LabelVirtualRouter(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.LabelVirtualRouter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelVirtualRouter`: NetworkVirtualRouter
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.LabelVirtualRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelVirtualRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**NetworkVirtualRouter**](networkVirtualRouter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelVirtualRouter1

> NetworkVirtualRouter LabelVirtualRouter1(ctx, oName).Body(body).Execute()

Label VirtualRouter object

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
    resp, r, err := api_client.NetworkV1Api.LabelVirtualRouter1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.LabelVirtualRouter1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelVirtualRouter1`: NetworkVirtualRouter
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.LabelVirtualRouter1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelVirtualRouter1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**NetworkVirtualRouter**](networkVirtualRouter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelVirtualRouterPeeringGroup

> NetworkVirtualRouterPeeringGroup LabelVirtualRouterPeeringGroup(ctx, oTenant).Body(body).Execute()

Label VirtualRouterPeeringGroup object

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
    resp, r, err := api_client.NetworkV1Api.LabelVirtualRouterPeeringGroup(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.LabelVirtualRouterPeeringGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelVirtualRouterPeeringGroup`: NetworkVirtualRouterPeeringGroup
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.LabelVirtualRouterPeeringGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelVirtualRouterPeeringGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**NetworkVirtualRouterPeeringGroup**](networkVirtualRouterPeeringGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelVirtualRouterPeeringGroup1

> NetworkVirtualRouterPeeringGroup LabelVirtualRouterPeeringGroup1(ctx, oName).Body(body).Execute()

Label VirtualRouterPeeringGroup object

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
    resp, r, err := api_client.NetworkV1Api.LabelVirtualRouterPeeringGroup1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.LabelVirtualRouterPeeringGroup1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelVirtualRouterPeeringGroup1`: NetworkVirtualRouterPeeringGroup
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.LabelVirtualRouterPeeringGroup1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelVirtualRouterPeeringGroup1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**NetworkVirtualRouterPeeringGroup**](networkVirtualRouterPeeringGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIPAMPolicy

> NetworkIPAMPolicyList ListIPAMPolicy(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List IPAMPolicy objects

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
    resp, r, err := api_client.NetworkV1Api.ListIPAMPolicy(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.ListIPAMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIPAMPolicy`: NetworkIPAMPolicyList
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.ListIPAMPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIPAMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkIPAMPolicyList**](networkIPAMPolicyList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIPAMPolicy1

> NetworkIPAMPolicyList ListIPAMPolicy1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List IPAMPolicy objects

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
    resp, r, err := api_client.NetworkV1Api.ListIPAMPolicy1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.ListIPAMPolicy1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIPAMPolicy1`: NetworkIPAMPolicyList
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.ListIPAMPolicy1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIPAMPolicy1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkIPAMPolicyList**](networkIPAMPolicyList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetwork

> NetworkNetworkList ListNetwork(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Network objects

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
    resp, r, err := api_client.NetworkV1Api.ListNetwork(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.ListNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetwork`: NetworkNetworkList
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.ListNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkNetworkList**](networkNetworkList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetwork1

> NetworkNetworkList ListNetwork1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Network objects

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
    resp, r, err := api_client.NetworkV1Api.ListNetwork1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.ListNetwork1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetwork1`: NetworkNetworkList
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.ListNetwork1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetwork1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkNetworkList**](networkNetworkList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkInterface

> NetworkNetworkInterfaceList ListNetworkInterface(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List NetworkInterface objects

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
    resp, r, err := api_client.NetworkV1Api.ListNetworkInterface(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.ListNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkInterface`: NetworkNetworkInterfaceList
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.ListNetworkInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkNetworkInterfaceList**](networkNetworkInterfaceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicerProfile

> NetworkPolicerProfileList ListPolicerProfile(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List PolicerProfile objects

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
    resp, r, err := api_client.NetworkV1Api.ListPolicerProfile(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.ListPolicerProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicerProfile`: NetworkPolicerProfileList
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.ListPolicerProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPolicerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkPolicerProfileList**](networkPolicerProfileList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicerProfile1

> NetworkPolicerProfileList ListPolicerProfile1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List PolicerProfile objects

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
    resp, r, err := api_client.NetworkV1Api.ListPolicerProfile1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.ListPolicerProfile1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicerProfile1`: NetworkPolicerProfileList
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.ListPolicerProfile1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPolicerProfile1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkPolicerProfileList**](networkPolicerProfileList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRouteTable

> NetworkRouteTableList ListRouteTable(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List RouteTable objects

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
    resp, r, err := api_client.NetworkV1Api.ListRouteTable(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.ListRouteTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRouteTable`: NetworkRouteTableList
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.ListRouteTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkRouteTableList**](networkRouteTableList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRouteTable1

> NetworkRouteTableList ListRouteTable1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List RouteTable objects

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
    resp, r, err := api_client.NetworkV1Api.ListRouteTable1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.ListRouteTable1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRouteTable1`: NetworkRouteTableList
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.ListRouteTable1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRouteTable1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkRouteTableList**](networkRouteTableList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoutingConfig

> NetworkRoutingConfigList ListRoutingConfig(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List RoutingConfig objects

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
    resp, r, err := api_client.NetworkV1Api.ListRoutingConfig(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.ListRoutingConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoutingConfig`: NetworkRoutingConfigList
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.ListRoutingConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoutingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkRoutingConfigList**](networkRoutingConfigList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualRouter

> NetworkVirtualRouterList ListVirtualRouter(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List VirtualRouter objects

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
    resp, r, err := api_client.NetworkV1Api.ListVirtualRouter(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.ListVirtualRouter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualRouter`: NetworkVirtualRouterList
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.ListVirtualRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkVirtualRouterList**](networkVirtualRouterList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualRouter1

> NetworkVirtualRouterList ListVirtualRouter1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List VirtualRouter objects

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
    resp, r, err := api_client.NetworkV1Api.ListVirtualRouter1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.ListVirtualRouter1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualRouter1`: NetworkVirtualRouterList
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.ListVirtualRouter1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualRouter1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkVirtualRouterList**](networkVirtualRouterList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualRouterPeeringGroup

> NetworkVirtualRouterPeeringGroupList ListVirtualRouterPeeringGroup(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List VirtualRouterPeeringGroup objects

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
    resp, r, err := api_client.NetworkV1Api.ListVirtualRouterPeeringGroup(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.ListVirtualRouterPeeringGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualRouterPeeringGroup`: NetworkVirtualRouterPeeringGroupList
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.ListVirtualRouterPeeringGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualRouterPeeringGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkVirtualRouterPeeringGroupList**](networkVirtualRouterPeeringGroupList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualRouterPeeringGroup1

> NetworkVirtualRouterPeeringGroupList ListVirtualRouterPeeringGroup1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List VirtualRouterPeeringGroup objects

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
    resp, r, err := api_client.NetworkV1Api.ListVirtualRouterPeeringGroup1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.ListVirtualRouterPeeringGroup1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualRouterPeeringGroup1`: NetworkVirtualRouterPeeringGroupList
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.ListVirtualRouterPeeringGroup1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualRouterPeeringGroup1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkVirtualRouterPeeringGroupList**](networkVirtualRouterPeeringGroupList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIPAMPolicy

> NetworkIPAMPolicy UpdateIPAMPolicy(ctx, oTenant).Body(body).Execute()

Update IPAMPolicy object

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
    body := *openapiclient.NewNetworkIPAMPolicy() // NetworkIPAMPolicy | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.UpdateIPAMPolicy(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.UpdateIPAMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIPAMPolicy`: NetworkIPAMPolicy
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.UpdateIPAMPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIPAMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkIPAMPolicy**](NetworkIPAMPolicy.md) |  | 

### Return type

[**NetworkIPAMPolicy**](networkIPAMPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIPAMPolicy1

> NetworkIPAMPolicy UpdateIPAMPolicy1(ctx, oName).Body(body).Execute()

Update IPAMPolicy object

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
    body := *openapiclient.NewNetworkIPAMPolicy() // NetworkIPAMPolicy | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.UpdateIPAMPolicy1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.UpdateIPAMPolicy1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIPAMPolicy1`: NetworkIPAMPolicy
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.UpdateIPAMPolicy1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIPAMPolicy1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkIPAMPolicy**](NetworkIPAMPolicy.md) |  | 

### Return type

[**NetworkIPAMPolicy**](networkIPAMPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetwork

> NetworkNetwork UpdateNetwork(ctx, oTenant).Body(body).Execute()

Update Network object

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
    body := *openapiclient.NewNetworkNetwork() // NetworkNetwork | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.UpdateNetwork(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.UpdateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetwork`: NetworkNetwork
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.UpdateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkNetwork**](NetworkNetwork.md) |  | 

### Return type

[**NetworkNetwork**](networkNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetwork1

> NetworkNetwork UpdateNetwork1(ctx, oName).Body(body).Execute()

Update Network object

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
    body := *openapiclient.NewNetworkNetwork() // NetworkNetwork | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.UpdateNetwork1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.UpdateNetwork1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetwork1`: NetworkNetwork
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.UpdateNetwork1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetwork1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkNetwork**](NetworkNetwork.md) |  | 

### Return type

[**NetworkNetwork**](networkNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkInterface

> NetworkNetworkInterface UpdateNetworkInterface(ctx, oName).Body(body).Execute()

Update NetworkInterface object

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
    body := *openapiclient.NewNetworkNetworkInterface() // NetworkNetworkInterface | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.UpdateNetworkInterface(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.UpdateNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkInterface`: NetworkNetworkInterface
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.UpdateNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkNetworkInterface**](NetworkNetworkInterface.md) |  | 

### Return type

[**NetworkNetworkInterface**](networkNetworkInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicerProfile

> NetworkPolicerProfile UpdatePolicerProfile(ctx, oTenant).Body(body).Execute()

Update PolicerProfile object

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
    body := *openapiclient.NewNetworkPolicerProfile() // NetworkPolicerProfile | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.UpdatePolicerProfile(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.UpdatePolicerProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicerProfile`: NetworkPolicerProfile
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.UpdatePolicerProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkPolicerProfile**](NetworkPolicerProfile.md) |  | 

### Return type

[**NetworkPolicerProfile**](networkPolicerProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicerProfile1

> NetworkPolicerProfile UpdatePolicerProfile1(ctx, oName).Body(body).Execute()

Update PolicerProfile object

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
    body := *openapiclient.NewNetworkPolicerProfile() // NetworkPolicerProfile | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.UpdatePolicerProfile1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.UpdatePolicerProfile1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicerProfile1`: NetworkPolicerProfile
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.UpdatePolicerProfile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicerProfile1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkPolicerProfile**](NetworkPolicerProfile.md) |  | 

### Return type

[**NetworkPolicerProfile**](networkPolicerProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoutingConfig

> NetworkRoutingConfig UpdateRoutingConfig(ctx, oName).Body(body).Execute()

Update RoutingConfig object

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
    body := *openapiclient.NewNetworkRoutingConfig() // NetworkRoutingConfig | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.UpdateRoutingConfig(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.UpdateRoutingConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoutingConfig`: NetworkRoutingConfig
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.UpdateRoutingConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoutingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkRoutingConfig**](NetworkRoutingConfig.md) |  | 

### Return type

[**NetworkRoutingConfig**](networkRoutingConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualRouter

> NetworkVirtualRouter UpdateVirtualRouter(ctx, oTenant).Body(body).Execute()

Update VirtualRouter object

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
    body := *openapiclient.NewNetworkVirtualRouter() // NetworkVirtualRouter | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.UpdateVirtualRouter(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.UpdateVirtualRouter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualRouter`: NetworkVirtualRouter
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.UpdateVirtualRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkVirtualRouter**](NetworkVirtualRouter.md) |  | 

### Return type

[**NetworkVirtualRouter**](networkVirtualRouter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualRouter1

> NetworkVirtualRouter UpdateVirtualRouter1(ctx, oName).Body(body).Execute()

Update VirtualRouter object

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
    body := *openapiclient.NewNetworkVirtualRouter() // NetworkVirtualRouter | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.UpdateVirtualRouter1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.UpdateVirtualRouter1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualRouter1`: NetworkVirtualRouter
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.UpdateVirtualRouter1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualRouter1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkVirtualRouter**](NetworkVirtualRouter.md) |  | 

### Return type

[**NetworkVirtualRouter**](networkVirtualRouter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualRouterPeeringGroup

> NetworkVirtualRouterPeeringGroup UpdateVirtualRouterPeeringGroup(ctx, oTenant).Body(body).Execute()

Update VirtualRouterPeeringGroup object

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
    body := *openapiclient.NewNetworkVirtualRouterPeeringGroup() // NetworkVirtualRouterPeeringGroup | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.UpdateVirtualRouterPeeringGroup(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.UpdateVirtualRouterPeeringGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualRouterPeeringGroup`: NetworkVirtualRouterPeeringGroup
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.UpdateVirtualRouterPeeringGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualRouterPeeringGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkVirtualRouterPeeringGroup**](NetworkVirtualRouterPeeringGroup.md) |  | 

### Return type

[**NetworkVirtualRouterPeeringGroup**](networkVirtualRouterPeeringGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualRouterPeeringGroup1

> NetworkVirtualRouterPeeringGroup UpdateVirtualRouterPeeringGroup1(ctx, oName).Body(body).Execute()

Update VirtualRouterPeeringGroup object

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
    body := *openapiclient.NewNetworkVirtualRouterPeeringGroup() // NetworkVirtualRouterPeeringGroup | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkV1Api.UpdateVirtualRouterPeeringGroup1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.UpdateVirtualRouterPeeringGroup1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualRouterPeeringGroup1`: NetworkVirtualRouterPeeringGroup
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.UpdateVirtualRouterPeeringGroup1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualRouterPeeringGroup1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkVirtualRouterPeeringGroup**](NetworkVirtualRouterPeeringGroup.md) |  | 

### Return type

[**NetworkVirtualRouterPeeringGroup**](networkVirtualRouterPeeringGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchIPAMPolicy

> NetworkAutoMsgIPAMPolicyWatchHelper WatchIPAMPolicy(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch IPAMPolicy objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.NetworkV1Api.WatchIPAMPolicy(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.WatchIPAMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchIPAMPolicy`: NetworkAutoMsgIPAMPolicyWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.WatchIPAMPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchIPAMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkAutoMsgIPAMPolicyWatchHelper**](networkAutoMsgIPAMPolicyWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchIPAMPolicy1

> NetworkAutoMsgIPAMPolicyWatchHelper WatchIPAMPolicy1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch IPAMPolicy objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.NetworkV1Api.WatchIPAMPolicy1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.WatchIPAMPolicy1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchIPAMPolicy1`: NetworkAutoMsgIPAMPolicyWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.WatchIPAMPolicy1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchIPAMPolicy1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkAutoMsgIPAMPolicyWatchHelper**](networkAutoMsgIPAMPolicyWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchNetwork

> NetworkAutoMsgNetworkWatchHelper WatchNetwork(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Network objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.NetworkV1Api.WatchNetwork(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.WatchNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchNetwork`: NetworkAutoMsgNetworkWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.WatchNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkAutoMsgNetworkWatchHelper**](networkAutoMsgNetworkWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchNetwork1

> NetworkAutoMsgNetworkWatchHelper WatchNetwork1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Network objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.NetworkV1Api.WatchNetwork1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.WatchNetwork1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchNetwork1`: NetworkAutoMsgNetworkWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.WatchNetwork1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchNetwork1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkAutoMsgNetworkWatchHelper**](networkAutoMsgNetworkWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchNetworkInterface

> NetworkAutoMsgNetworkInterfaceWatchHelper WatchNetworkInterface(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch NetworkInterface objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.NetworkV1Api.WatchNetworkInterface(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.WatchNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchNetworkInterface`: NetworkAutoMsgNetworkInterfaceWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.WatchNetworkInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkAutoMsgNetworkInterfaceWatchHelper**](networkAutoMsgNetworkInterfaceWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchPolicerProfile

> NetworkAutoMsgPolicerProfileWatchHelper WatchPolicerProfile(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch PolicerProfile objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.NetworkV1Api.WatchPolicerProfile(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.WatchPolicerProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchPolicerProfile`: NetworkAutoMsgPolicerProfileWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.WatchPolicerProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchPolicerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkAutoMsgPolicerProfileWatchHelper**](networkAutoMsgPolicerProfileWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchPolicerProfile1

> NetworkAutoMsgPolicerProfileWatchHelper WatchPolicerProfile1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch PolicerProfile objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.NetworkV1Api.WatchPolicerProfile1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.WatchPolicerProfile1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchPolicerProfile1`: NetworkAutoMsgPolicerProfileWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.WatchPolicerProfile1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchPolicerProfile1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkAutoMsgPolicerProfileWatchHelper**](networkAutoMsgPolicerProfileWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchRouteTable

> NetworkAutoMsgRouteTableWatchHelper WatchRouteTable(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch RouteTable objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.NetworkV1Api.WatchRouteTable(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.WatchRouteTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchRouteTable`: NetworkAutoMsgRouteTableWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.WatchRouteTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkAutoMsgRouteTableWatchHelper**](networkAutoMsgRouteTableWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchRouteTable1

> NetworkAutoMsgRouteTableWatchHelper WatchRouteTable1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch RouteTable objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.NetworkV1Api.WatchRouteTable1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.WatchRouteTable1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchRouteTable1`: NetworkAutoMsgRouteTableWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.WatchRouteTable1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchRouteTable1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkAutoMsgRouteTableWatchHelper**](networkAutoMsgRouteTableWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchRoutingConfig

> NetworkAutoMsgRoutingConfigWatchHelper WatchRoutingConfig(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch RoutingConfig objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.NetworkV1Api.WatchRoutingConfig(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.WatchRoutingConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchRoutingConfig`: NetworkAutoMsgRoutingConfigWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.WatchRoutingConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchRoutingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkAutoMsgRoutingConfigWatchHelper**](networkAutoMsgRoutingConfigWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchVirtualRouter

> NetworkAutoMsgVirtualRouterWatchHelper WatchVirtualRouter(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch VirtualRouter objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.NetworkV1Api.WatchVirtualRouter(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.WatchVirtualRouter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchVirtualRouter`: NetworkAutoMsgVirtualRouterWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.WatchVirtualRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchVirtualRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkAutoMsgVirtualRouterWatchHelper**](networkAutoMsgVirtualRouterWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchVirtualRouter1

> NetworkAutoMsgVirtualRouterWatchHelper WatchVirtualRouter1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch VirtualRouter objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.NetworkV1Api.WatchVirtualRouter1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.WatchVirtualRouter1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchVirtualRouter1`: NetworkAutoMsgVirtualRouterWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.WatchVirtualRouter1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchVirtualRouter1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkAutoMsgVirtualRouterWatchHelper**](networkAutoMsgVirtualRouterWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchVirtualRouterPeeringGroup

> NetworkAutoMsgVirtualRouterPeeringGroupWatchHelper WatchVirtualRouterPeeringGroup(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch VirtualRouterPeeringGroup objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.NetworkV1Api.WatchVirtualRouterPeeringGroup(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.WatchVirtualRouterPeeringGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchVirtualRouterPeeringGroup`: NetworkAutoMsgVirtualRouterPeeringGroupWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.WatchVirtualRouterPeeringGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchVirtualRouterPeeringGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkAutoMsgVirtualRouterPeeringGroupWatchHelper**](networkAutoMsgVirtualRouterPeeringGroupWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchVirtualRouterPeeringGroup1

> NetworkAutoMsgVirtualRouterPeeringGroupWatchHelper WatchVirtualRouterPeeringGroup1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch VirtualRouterPeeringGroup objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.NetworkV1Api.WatchVirtualRouterPeeringGroup1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkV1Api.WatchVirtualRouterPeeringGroup1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchVirtualRouterPeeringGroup1`: NetworkAutoMsgVirtualRouterPeeringGroupWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `NetworkV1Api.WatchVirtualRouterPeeringGroup1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchVirtualRouterPeeringGroup1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**NetworkAutoMsgVirtualRouterPeeringGroupWatchHelper**](networkAutoMsgVirtualRouterPeeringGroupWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

