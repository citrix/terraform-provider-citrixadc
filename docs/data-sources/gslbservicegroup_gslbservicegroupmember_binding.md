---
subcategory: "GSLB"
---

# Data Source: gslbservicegroup_gslbservicegroupmember_binding

Retrieves information about a member bound to a GSLB service group. Look the member up by its parent service group and port, plus the IP address or server name that identifies it within the group.


## Example Usage

```terraform
data "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" "member_by_ip" {
  servicegroupname = "gslbsvcgrp1"
  ip               = "192.0.2.10"
  port             = 80
}

output "member_weight" {
  value = data.citrixadc_gslbservicegroup_gslbservicegroupmember_binding.member_by_ip.weight
}

output "member_state" {
  value = data.citrixadc_gslbservicegroup_gslbservicegroupmember_binding.member_by_ip.state
}
```


## Argument Reference

* `servicegroupname` - (Required) Name of the GSLB service group.
* `port` - (Required) Server port number.
* `ip` - (Optional) IP Address of the member. Use this to look up a member bound by the IP path; supply either `ip` or `servername`.
* `servername` - (Optional) Name of the server bound to the service group. Use this to look up a member bound by the server-name path; supply either `servername` or `ip`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The composite ID of the binding, in the form `ip:<ip>,port:<port>,servername:<servername>,servicegroupname:<servicegroupname>` (values URL-encoded; the unused member field is empty).
* `hashid` - The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.
* `order` - Order number assigned to the GSLB service group member.
* `publicip` - The public IP address that a NAT device translates to the GSLB service's private IP address.
* `publicport` - The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service.
* `siteprefix` - The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string `NONE` is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.
* `state` - Initial state of the member.
* `weight` - Weight assigned to the member. Specifies the capacity of the server relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.

### Read-only gslbservicegroup_gslbservicegroupmember_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_gslbservicegroup_gslbservicegroupmember_binding` resource). They are GET-only / Computed, and any attribute the appliance does not return is `null`.

* `statechangetimesec` - Time when the last state change occurred (seconds part).
* `preferredlocation` - Preferred location.
* `trofsdelay` - Delay before moving to TROFS.
* `delay` - The time allowed (in seconds) for a graceful shutdown.
* `gslbthreshold` - Indicates whether the GSLB service has reached its threshold.
* `orderstr` - Order number in string form assigned to the GSLB service group member.
* `graceful` - Whether to wait for all existing connections to the service to terminate before shutting it down.
* `threshold` - Threshold indicator (`ABOVE`, `BELOW`).
* `svrstate` - The state of the GSLB service (for example `UP`, `DOWN`, `OUT OF SERVICE`).
* `tickssincelaststatechange` - Time, in 10-millisecond ticks, since the last state change.
