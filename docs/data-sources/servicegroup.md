---
subcategory: "Basic"
---

# Data Source: servicegroup

The servicegroup data source allows you to retrieve information about a service group configuration.

## Example Usage

```terraform
data "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "test_servicegroup"
}

output "servicetype" {
  value = data.citrixadc_servicegroup.tf_servicegroup.servicetype
}

output "state" {
  value = data.citrixadc_servicegroup.tf_servicegroup.state
}
```

## Argument Reference

* `servicegroupname` - (Required) Name of the service group.

## Attribute Reference

The following attributes are available:

* `id` - The id of the servicegroup. It is a system-generated identifier.
* `servicegroupname` - Name of the service group.
* `servicetype` - Protocol used to exchange data with the service. Example: `HTTP`, `SSL`, `TCP`, `UDP`, `DNS`.
* `state` - Initial state of the service group. Possible values: `ENABLED`, `DISABLED`.
* `cacheable` - Use the transparent cache redirection virtual server to forward requests to the cache server. Possible values: `YES`, `NO`.
* `cip` - Insert the Client IP header in requests forwarded to the service.
* `usip` - Use client's IP address as the source IP address when initiating connection to the server. Possible values: `YES`, `NO`.
* `useproxyport` - Use the proxy port as the source port when initiating connections with the server. Possible values: `YES`, `NO`.
* `sp` - Enable surge protection for the service group. Possible values: `ON`, `OFF`.
* `clttimeout` - Time, in seconds, after which to terminate an idle client connection.
* `svrtimeout` - Time, in seconds, after which to terminate an idle server connection.
* `maxclient` - Maximum number of simultaneous open connections for the service group.
* `maxreq` - Maximum number of requests that can be sent on a persistent connection to the service group.
* `comment` - Any information about the service group.
* `autoscale` - Auto scale option for a servicegroup. Possible values: `DISABLED`, `DNS`, `POLICY`.
* `graceful` - Indicates graceful shutdown of the service. System will wait for all outstanding connections to this service to be closed before disabling the service. Possible values: `YES`, `NO`.

### Read-only servicegroup metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_servicegroup` resource). They are GET-only / Computed. Any attribute the appliance does not return is `null`.

* `numofconnections` - The number of client side connections still open.
* `serviceconftype` - The configuration type of the service group.
* `value` - SSL Status. Possible values = Certkey/Certkeybundle/Vault not bound/Cert-store not usable, SSL feature disabled.
* `svrstate` - The state of the service. Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.
* `ip` - IP Address.
* `monstatcode` - The code indicating the monitor response.
* `monstatparam1` - First parameter for use with message code.
* `monstatparam2` - Second parameter for use with message code.
* `monstatparam3` - Third parameter for use with message code.
* `statechangetimemsec` - Time when last state change occurred. Milliseconds part.
* `stateupdatereason` - Checks state update reason on the secondary node.
* `clmonowner` - Tells the mon owner of the service.
* `clmonview` - Tells the view id of the monitoring owner.
* `groupcount` - Servicegroup Count.
* `serviceipstr` - This field shows the dbs services ip.
* `servicegroupeffectivestate` - Indicates the effective servicegroup state based on the state of the bound service items. Possible values = UP, DOWN, OUT OF SERVICE, PARTIAL-UP, PARTIAL-DOWN.
* `nodefaultbindings` - To determine if the configuration is from stylebooks. Possible values = YES, NO.
* `svcitmactsvcs` - The total active service items for an FQDN for SRV type server binding.
* `svcitmboundsvcs` - The total bound items for an FQDN for SRV type server binding.
* `monuserstatusmesg` - User monitor failure reasons.
