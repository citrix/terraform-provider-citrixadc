---
subcategory: "GSLB"
---

# Resource: gslbservicegroup

A GSLB (Global Server Load Balancing) service group groups together a set of GSLB service members that belong to the same GSLB site and share a common protocol and set of settings. Binding a service group to a GSLB virtual server lets the Citrix ADC distribute client traffic across data centers (sites) for disaster recovery, proximity-based routing, and capacity sharing, while managing the member services as a single logical unit. Use this resource to define the service group and its site-wide settings; individual members and monitor bindings are typically managed with the companion `citrixadc_gslbservicegroup_gslbservicegroupmember_binding` and `citrixadc_gslbservicegroup_lbmonitor_binding` resources.


## Example usage

```hcl
resource "citrixadc_gslbsite" "site_local" {
  sitename        = "Site-Local"
  siteipaddress   = "172.31.96.234"
  sessionexchange = "DISABLED"
}

resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
  servicegroupname = "gslb_svcgrp_http"
  servicetype      = "HTTP"
  sitename         = citrixadc_gslbsite.site_local.sitename
  cip              = "DISABLED"
  healthmonitor    = "NO"
  comment          = "GSLB service group for the local site"
}
```


## Argument Reference

* `servicegroupname` - (Required) Name of the GSLB service group. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Changing this value forces a new resource to be created (to rename in place, use `newname` instead).
* `servicetype` - (Required) Protocol used to exchange data with the GSLB service (for example, `HTTP`, `SSL`, `TCP`). Changing this value forces a new resource to be created.
* `sitename` - (Required) Name of the GSLB site to which the service group belongs. This attribute can be updated in place.

The following arguments are **create-only**. Changing them forces the service group to be destroyed and recreated:

* `autoscale` - (Optional) Auto scale option for a GSLB service group. Defaults to `"DISABLED"`. Changing this value forces a new resource to be created.
* `autodelayedtrofs` - (Optional) Indicates graceful movement of the service to TROFS. The system waits for the monitor response time-out before moving the service to TROFS. Changing this value forces a new resource to be created.
* `state` - (Optional) Initial state of the GSLB service group. Defaults to `"ENABLED"`. Changing this value forces a new resource to be created.

The following arguments can be **updated in place**:

* `appflowlog` - (Optional) Enable logging of AppFlow information for the specified GSLB service group. Defaults to `"ENABLED"`.
* `cip` - (Optional) Insert the Client IP header in requests forwarded to the GSLB service.
* `cipheader` - (Optional) Name of the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If client IP insertion is enabled, and the client IP header is not specified, the value of the Client IP Header parameter or the value set by the `set ns config` command is used as the client's IP header name.
* `clttimeout` - (Optional) Time, in seconds, after which to terminate an idle client connection.
* `svrtimeout` - (Optional) Time, in seconds, after which to terminate an idle server connection.
* `maxclient` - (Optional) Maximum number of simultaneous open connections for the GSLB service group.
* `maxbandwidth` - (Optional) Maximum bandwidth, in Kbps, allocated for all the services in the GSLB service group.
* `monthreshold` - (Optional) Minimum sum of weights of the monitors that are bound to this GSLB service. Used to determine whether to mark a GSLB service as UP or DOWN.
* `downstateflush` - (Optional) Flush all active transactions associated with all the services in the GSLB service group whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions. Defaults to `"ENABLED"`.
* `sitepersistence` - (Optional) Use cookie-based site persistence. Applicable only to HTTP and SSL non-autoscale enabled GSLB service groups.
* `healthmonitor` - (Optional) Monitor the health of this GSLB service. Available settings function as follows: `YES` - Send probes to check the health of the GSLB service. `NO` - Do not send probes to check the health of the GSLB service; with the `NO` option, the appliance shows the service as UP at all times. Defaults to `YES`.
* `comment` - (Optional) Any information about the GSLB service group.

### Rename

* `newname` - (Optional) New name for the GSLB service group. Setting or changing this value renames the existing service group in place via the NITRO `rename` action (rather than destroying and recreating it). The resource ID and the live object track the new name; the `servicegroupname` attribute remains pinned to the originally configured value.

### Member-level arguments

The following arguments describe a single service group member. They are optional here; in most configurations members are managed separately with the `citrixadc_gslbservicegroup_gslbservicegroupmember_binding` resource.

* `servername` - (Optional) Name of the server to which to bind the service group.
* `port` - (Optional) Server port number.
* `weight` - (Optional) Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.
* `hashid` - (Optional) The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.
* `publicip` - (Optional) The public IP address that a NAT device translates to the GSLB service's private IP address.
* `publicport` - (Optional) The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service.
* `siteprefix` - (Optional) The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string `NONE` is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.
* `order` - (Optional) Order number to be assigned to the GSLB service group member.
* `monitor_name_svc` - (Optional) Name of the monitor bound to the GSLB service group. Used to assign a weight to the monitor.
* `dup_weight` - (Optional) Weight of the monitor that is bound to the GSLB service group.

### Other arguments

* `delay` - (Optional) The time allowed (in seconds) for a graceful shutdown. During this period, new connections or requests will continue to be sent to this service for clients who already have a persistent session on the system. Connections or requests from fresh or new clients who do not yet have a persistence session on the system will not be sent to the service; instead, they will be load balanced among other available services. After the delay time expires, no new requests or connections will be sent to the service. Note: This argument applies only to the disable action and is not sent on create or update.
* `graceful` - (Optional) Wait for all existing connections to the service to terminate before shutting down the service. Note: This argument applies only to the disable action and is not sent on create or update.
* `includemembers` - (Optional) Display the members of the listed GSLB service groups in addition to their settings. Note: This is a read/GET-only filter; it is not sent on create or update.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbservicegroup. It has the same value as the `servicegroupname` attribute (or the value of `newname` after a rename).


## Import

A gslbservicegroup can be imported using its name, e.g.

```shell
terraform import citrixadc_gslbservicegroup.tf_gslbservicegroup gslb_svcgrp_http
```
