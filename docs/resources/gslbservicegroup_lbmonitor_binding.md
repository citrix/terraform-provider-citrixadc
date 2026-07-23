---
subcategory: "GSLB"
---

# Resource: gslbservicegroup_lbmonitor_binding

This resource is used to bind a load balancing monitor to a GSLB service group.


## Example usage

```hcl
resource "citrixadc_gslbservicegroup_lbmonitor_binding" "tf_gslbservicegroup_lbmonitor_binding" {
  servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
  monitor_name     = citrixadc_lbmonitor.tf_lbmonitor.monitorname
  weight           = 20
}

resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
  servicegroupname = "gslb_sg1"
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

resource "citrixadc_lbmonitor" "tf_lbmonitor" {
  monitorname = "http_mon1"
  type        = "HTTP"
}
```


## Argument Reference

* `servicegroupname` - (Required) Name of the GSLB service group. Changing this forces a new resource to be created.
* `monitor_name` - (Required) Monitor name. Changing this forces a new resource to be created.
* `hashid` - (Optional, Computed) Unique numerical identifier used by hash based load balancing methods to identify a service. Changing this forces a new resource to be created.
* `monstate` - (Optional, Computed) Monitor state. Changing this forces a new resource to be created.
* `order` - (Optional, Computed) Order number to be assigned to the GSLB servicegroup member. Changing this forces a new resource to be created.
* `passive` - (Optional, Computed) Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached. Changing this forces a new resource to be created.
* `port` - (Optional, Computed) Port number of the GSLB service. Each service must have a unique port number. Changing this forces a new resource to be created.
* `publicip` - (Optional, Computed) The public IP address that a NAT device translates to the GSLB service's private IP address. Changing this forces a new resource to be created.
* `publicport` - (Optional, Computed) The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Changing this forces a new resource to be created.
* `siteprefix` - (Optional, Computed) The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains. Changing this forces a new resource to be created.
* `state` - (Optional, Computed) Initial state of the service after binding. Defaults to `"ENABLED"`. Changing this forces a new resource to be created.
* `weight` - (Optional, Computed) Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the gslbservicegroup_lbmonitor_binding. It is the concatenation of the `servicegroupname` and `monitor_name` attributes separated by a comma.


## Import

A gslbservicegroup_lbmonitor_binding can be imported using its id, which is the concatenation of `servicegroupname` and `monitor_name` separated by a comma, e.g.

```shell
terraform import citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding gslb_sg1,http_mon1
```
