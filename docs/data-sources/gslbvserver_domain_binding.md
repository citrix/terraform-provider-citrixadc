---
subcategory: "GSLB"
---

# Data Source: gslbvserver_domain_binding

The gslbvserver_domain_binding data source allows you to retrieve information about a GSLB virtual server domain binding.


## Example usage

```terraform
data "citrixadc_gslbvserver_domain_binding" "tf_gslbvserver_domain_binding" {
  name       = "GSLB-East-Coast-Vserver"
  domainname = "www.exampledomain.com"
}

output "name" {
  value = data.citrixadc_gslbvserver_domain_binding.tf_gslbvserver_domain_binding.name
}

output "domainname" {
  value = data.citrixadc_gslbvserver_domain_binding.tf_gslbvserver_domain_binding.domainname
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server on which the binding is configured.
* `domainname` - (Required) Domain name for which to retrieve the binding information.
* `backupipflag` - (Optional) Filter by the backup IP flag value. Used when filtering bindings by the backup IP configuration.
* `cookie_domainflag` - (Optional) Filter by the cookie domain flag value. Used when filtering bindings by the cookie domain configuration.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbvserver_domain_binding. It is a system-generated identifier.
* `backupip` - The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.
* `cookie_domain` - The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.
* `cookietimeout` - Timeout, in minutes, for the GSLB site cookie.
* `order` - Order number assigned to the service when it is bound to the lb vserver.
* `sitedomainttl` - TTL, in seconds, for all internally created site domains (created when a site prefix is configured on a GSLB service) that are associated with this virtual server.
* `ttl` - Time to live (TTL) for the domain.
