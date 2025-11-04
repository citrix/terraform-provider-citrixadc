---
subcategory: "GSLB"
---

# Resource: gslbservicegroup

The gslbservicegroup resource is used to create gslbservicegroup.


## Example usage

```hcl
resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
  servicegroupname = "test_gslbvservicegroup"
  servicetype      = "HTTP"
  cip              = "DISABLED"
  healthmonitor    = "NO"
  sitename         = citrixadc_gslbsite.site_local.sitename
}
resource "citrixadc_gslbsite" "site_local" {
  sitename        = "Site-Local"
  siteipaddress   = "172.31.96.234"
  sessionexchange = "DISABLED"
}
```


## Argument Reference
* `servicegroupname` - (Required) Name of the GSLB service group. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the name is created.
* `appflowlog` - (Optional) Enable logging of AppFlow information for the specified GSLB service group.
* `autoscale` - (Optional) Auto scale option for a GSLB servicegroup
* `cip` - (Optional) Insert the Client IP header in requests forwarded to the GSLB service.
* `cipheader` - (Optional) Name of the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If client IP insertion is enabled, and the client IP header is not specified, the value of Client IP Header parameter or the value set by the set ns config command is used as client's IP header name.
* `clttimeout` - (Optional) Time, in seconds, after which to terminate an idle client connection.
* `comment` - (Optional) Any information about the GSLB service group.
* `delay` - (Optional) The time allowed (in seconds) for a graceful shutdown. During this period, new connections or requests will continue to be sent to this service for clients who already have a persistent session on the system. Connections or requests from fresh or new clients who do not yet have a persistence sessions on the system will not be sent to the service. Instead, they will be load balanced among other available services. After the delay time expires, no new requests or connections will be sent to the service.
* `downstateflush` - (Optional) Flush all active transactions associated with all the services in the GSLB service group whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.
* `dup_weight` - (Optional) weight of the monitor that is bound to GSLB servicegroup.
* `graceful` - (Optional) Wait for all existing connections to the service to terminate before shutting down the service.
* `hashid` - (Optional) The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.
* `healthmonitor` - (Optional) Monitor the health of this GSLB service.Available settings function are as follows: YES - Send probes to check the health of the GSLB service. NO - Do not send probes to check the health of the GSLB service. With the NO option, the appliance shows the service as UP at all times.
* `includemembers` - (Optional) Display the members of the listed GSLB service groups in addition to their settings. Can be specified when no service group name is provided in the command. In that case, the details displayed for each service group are identical to the details displayed when a service group name is provided, except that bound monitors are not displayed.
* `maxbandwidth` - (Optional) Maximum bandwidth, in Kbps, allocated for all the services in the GSLB service group.
* `maxclient` - (Optional) Maximum number of simultaneous open connections for the GSLB service group.
* `monitor_name_svc` - (Optional) Name of the monitor bound to the GSLB service group. Used to assign a weight to the monitor.
* `monthreshold` - (Optional) Minimum sum of weights of the monitors that are bound to this GSLB service. Used to determine whether to mark a GSLB service as UP or DOWN.
* `port` - (Optional) Server port number.
* `publicip` - (Optional) The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.
* `publicport` - (Optional) The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.
* `servername` - (Optional) Name of the server to which to bind the service group.
* `servicetype` - (Optional) Protocol used to exchange data with the GSLB service.
* `sitename` - (Optional) Name of the GSLB site to which the service group belongs.
* `sitepersistence` - (Optional) Use cookie-based site persistence. Applicable only to HTTP and SSL non-autoscale enabled GSLB servicegroups.
* `siteprefix` - (Optional) The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.
* `state` - (Optional) Initial state of the GSLB service group.
* `svrtimeout` - (Optional) Time, in seconds, after which to terminate an idle server connection.
* `weight` - (Optional) Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.
* `autodelayedtrofs` - (Optional) Indicates graceful movement of the service to TROFS. System will wait for monitor response time out before moving to TROFS

gslb
## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbservicegroup. It has the same value as the `servicegroupname` attribute.


## Import

A gslbservicegroup can be imported using its name, e.g.

```shell
terraform import citrixadc_gslbservicegroup.tf_gslbservicegroup test_gslbvservicegroup
```
