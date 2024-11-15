---
subcategory: "GSLB"
---

# Resource: gslbvserver_gslbservicegroup_binding

The gslbvserver_gslbservicegroup_binding resource is used to create gslbvserver_gslbservicegroup_binding.


## Example usage

```hcl
resource "citrixadc_gslbvserver_gslbservicegroup_binding" "tf_gslbvserver_gslbservicegroup_binding" {
  name             = citrixadc_gslbvserver.tf_gslbvserver.name
  servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
}

resource "citrixadc_gslbsite" "site_local" {
  sitename        = "Site-Local"
  siteipaddress   = "172.31.96.234"
  sessionexchange = "DISABLED"
}

resource "citrixadc_gslbvserver" "tf_gslbvserver" {
  dnsrecordtype = "A"
  name          = "GSLB-East-Coast-Vserver"
  servicetype   = "HTTP"
  domain {
    domainname = "www.fooco.co"
    ttl        = "60"
  }
  domain {
    domainname = "www.barco.com"
    ttl        = "65"
  }
}
resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
  servicegroupname = "test_gslbvservicegroup"
  servicetype      = "HTTP"
  cip              = "DISABLED"
  healthmonitor    = "NO"
  sitename         = citrixadc_gslbsite.site_local.sitename
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server on which to perform the binding operation.
* `servicegroupname` - (Required) The GSLB service group name bound to the selected GSLB virtual server.
* `order` - (Required) Order number to be assigned to the service when it is bound to the lb vserver. Minimum value = 1 | Maximum value = 8192


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbvserver_gslbservicegroup_binding is the concatenation of the `name` and `servicegroupname` attributes separated by a comma.


## Import

A gslbvserver_gslbservicegroup_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_gslbvserver_gslbservicegroup_binding.tf_gslbvserver_gslbservicegroup_binding GSLB-East-Coast-Vserver,test_gslbvservicegroup
```
