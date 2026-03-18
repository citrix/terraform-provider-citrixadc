---
subcategory: "GSLB"
---

# Data Source: gslbservicegroup_lbmonitor_binding

The gslbservicegroup_lbmonitor_binding data source allows you to retrieve information about a monitor binding to a GSLB service group.

## Example Usage

```terraform
data "citrixadc_gslbservicegroup_lbmonitor_binding" "tf_gslbservicegroup_lbmonitor_binding" {
  servicegroupname = "test_gslbvservicegroup"
  monitor_name     = "tf_monitor"
}

output "servicegroupname" {
  value = data.citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding.servicegroupname
}

output "weight" {
  value = data.citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding.weight
}

output "monitor_name" {
  value = data.citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding.monitor_name
}
```

## Argument Reference

* `servicegroupname` - (Required) Name of the GSLB service group.
* `monitor_name` - (Required) Monitor name.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `hashid` - Unique numerical identifier used by hash based load balancing methods to identify a service.
* `port` - Port number of the GSLB service. Each service must have a unique port number. When not specified, the datasource returns the binding that applies to all ports.
* `monstate` - Monitor state.
* `order` - Order number to be assigned to the gslb servicegroup member.
* `passive` - Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.
* `publicip` - The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.
* `publicport` - The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.
* `siteprefix` - The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.
* `state` - Initial state of the service after binding.
* `weight` - Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.

## Attribute Reference

* `id` - The id of the gslbservicegroup_lbmonitor_binding. It is a system-generated identifier.
