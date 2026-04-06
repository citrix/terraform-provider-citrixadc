---
subcategory: "GSLB"
---

# Data Source: gslbservicegroup_gslbservicegroupmember_binding

The gslbservicegroup_gslbservicegroupmember_binding data source allows you to retrieve information about a GSLB service group to GSLB service group member binding.


## Example Usage

```terraform
data "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" "tf_binding" {
  servicegroupname = "test_gslbvservicegroup"
  servername       = "tf_server"
  port             = 60
}

output "servicegroupname" {
  value = data.citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding.servicegroupname
}

output "servername" {
  value = data.citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding.servername
}

output "port" {
  value = data.citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding.port
}
```


## Argument Reference

* `servicegroupname` - (Required) Name of the GSLB service group.
* `ip` - (Optional) IP Address. Either ip or servername is required.
* `port` - (Required) Server port number.
* `servername` - (Optional) Name of the server to which to bind the service group. Either ip or servername is required.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbservicegroup_gslbservicegroupmember_binding. 
* `hashid` - The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.
* `order` - Order number to be assigned to the gslb servicegroup member
* `publicip` - The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.
* `publicport` - The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.
* `siteprefix` - The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.
* `state` - Initial state of the GSLB service group.
* `weight` - Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.
