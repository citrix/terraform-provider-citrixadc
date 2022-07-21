---
subcategory: "GSLB"
---

# Resource: gslbvserver_gslbservice_binding

The gslbvserver_gslbservice_binding resource is used to create gslbvserver_gslbservice_binding.


## Example usage

```hcl
resource "citrixadc_gslbvserver_gslbservice_binding" "tf_gslbvserver_gslbservice_binding"{
  name = citrixadc_gslbvserver.tf_gslbvserver.name
  servicename = citrixadc_gslbservice.gslb_svc1.servicename
}

resource "citrixadc_gslbsite" "site_local" {
  sitename        = "Site-Local"
  siteipaddress   = "172.31.96.234"
  sessionexchange = "DISABLED"
}

resource "citrixadc_gslbservice" "gslb_svc1" {
  ip          = "172.16.1.121"
  port        = "80"
  servicename = "gslb1vservice"
  servicetype = "HTTP"
  sitename    = citrixadc_gslbsite.site_local.sitename
}

resource "citrixadc_gslbvserver" "tf_gslbvserver" {
  dnsrecordtype = "A"
  name          = "gslb_vserver"
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


```


## Argument Reference

* `name` - (Required) Name of the virtual server on which to perform the binding operation.
* `servicename` - (Required) Name of the GSLB service for which to change the weight.
* `domainname` - (Optional) Domain name for which to change the time to live (TTL) and/or backup service IP address.
* `weight` - (Optional) Weight to assign to the GSLB service.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbvserver_gslbservice_binding is the concatenation of the `name` and `servicename` attributes separated by a comma.


## Import

A gslbvserver_gslbservice_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_gslbvserver_gslbservice_binding.tf_gslbvserver_gslbservice_binding gslb_vserver,gslb1vservice
```
