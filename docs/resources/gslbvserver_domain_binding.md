---
subcategory: "GSLB"
---

# Resource: gslbvserver_domain_binding

The gslbvserver_domain_binding resource is used to create gslbvserver_domain_binding.


## Example usage

```hcl
resource "citrixadc_gslbvserver_domain_binding" "tf_gslbvserver_domain_binding"{
  name = citrixadc_gslbvserver.tf_gslbvserver.name
  domainname = "www.example.com"
  backupipflag = false
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

```


## Argument Reference

* `name` - (Required) Name of the virtual server on which to perform the binding operation.
* `domainname` - (Required) Domain name for which to change the time to live (TTL) and/or backup service IP address.
* `backupip` - (Optional) The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.
* `backupipflag` - (Optional) The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.
* `cookie_domain` - (Optional) The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.
* `cookie_domainflag` - (Optional) The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.
* `cookietimeout` - (Optional) Timeout, in minutes, for the GSLB site cookie.
* `sitedomainttl` - (Optional) TTL, in seconds, for all internally created site domains (created when a site prefix is configured on a GSLB service) that are associated with this virtual server.
* `ttl` - (Optional) Time to live (TTL) for the domain.
* `order` - (Optional) Order number to be assigned to the service when it is bound to the lb vserver. 


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbvserver_domain_binding is the conatenation of the `name` and `domainname` attributes.


## Import

A gslbvserver_domain_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_gslbvserver_domain_binding.tf_gslbvserver_domain_binding GSLB-East-Coast-Vserver,www.example.com
```
