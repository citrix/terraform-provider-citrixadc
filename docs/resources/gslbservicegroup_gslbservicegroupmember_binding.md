---
subcategory: "GSLB"
---

# Resource: gslbservicegroup_gslbservicegroupmember_binding

The gslbservicegroup_gslbservicegroupmember_binding resource is used to bind gslbservicegroupmember to gslbservicegroup.


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

resource "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" "name" {
  servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
  servername       = "10.10.10.10"
  port             = 60
}

```


## Argument Reference

* `ip` - (Optional) IP Address.
* `port` - (Optional) Server port number. Range 1 - 65535 * in CLI is represented as 65535 in NITRO API
* `weight` - (Optional) Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service. Minimum value =  1 Maximum value =  100
* `servername` - (Optional) Name of the server to which to bind the service group. Minimum length =  1
* `state` - (Optional) Initial state of the GSLB service group. Possible values: [ ENABLED, DISABLED ]
* `hashid` - (Optional) The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods. Minimum value =  1
* `publicip` - (Optional) The public IP address that a NAT device translates to the GSLB service's private IP address. Optional. Minimum length =  1
* `publicport` - (Optional) The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional. Minimum value =  1
* `siteprefix` - (Optional) The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.
* `servicegroupname` - (Optional) Name of the GSLB service group. Minimum length =  1
* `order` - (Optional) Order number to be assigned to the gslb servicegroup member. Minimum value = 1 | Maximum value = 8192


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbservicegroup_gslbservicegroupmember_binding. It has the same value as the `name` attribute.


## Import

A gslbservicegroup_gslbservicegroupmember_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_bind test_gslbvservicegroup,10.10.10.10,60
```
