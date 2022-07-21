---
subcategory: "GSLB"
---

# Resource: gslbservicegroup_lbmonitor_binding


The gslbservicegroup_lbmonitor_binding resource is used to create gslbservicegroup_lbmonitor_binding.


## Example usage

```hcl
resource "citrixadc_gslbservicegroup_lbmonitor_binding" "tf_gslbservicegroup_lbmonitor_binding" {
  weight           = 20
  servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
  monitor_name      = citrixadc_lbmonitor.tfmonitor1.monitorname

}

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

resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "test_monitor"
  type        = "HTTP"
}
```


## Argument Reference

* `servicegroupname` - (Required) Name of the GSLB service group.
* `monitor_name` - (Required) Monitor name.
* `hashid` - (Optional) Unique numerical identifier used by hash based load balancing methods to identify a service.
* `monstate` - (Optional) Monitor state.
* `passive` - (Optional) Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.
* `port` - (Optional) Port number of the GSLB service. Each service must have a unique port number.
* `publicip` - (Optional) The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.
* `publicport` - (Optional) The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.
* `siteprefix` - (Optional) The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.
* `state` - (Optional) Initial state of the service after binding.
* `weight` - (Optional) Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbservicegroup_lbmonitor_binding is the concatenation of the `servicegroupname` and `monitor_name` attributes separated by a comma.



## Import

A gslbservicegroup_lbmonitor_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding test_gslbvservicegroup,test_monitor
```
