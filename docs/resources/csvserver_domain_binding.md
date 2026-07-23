---
subcategory: "Content Switching"
---

# Resource: csvserver_domain_binding

This resource is used to bind a DNS domain to a content switching virtual server.


## Example usage

```hcl
resource "citrixadc_csvserver_domain_binding" "tf_csvserver_domain_binding" {
  name          = citrixadc_csvserver.tf_csvserver.name
  domainname    = "example.com"
  ttl           = 3600
  backupip      = "10.20.30.40"
  cookiedomain  = ".example.com"
  cookietimeout = 30
  sitedomainttl = 7200
}

resource "citrixadc_csvserver" "tf_csvserver" {
  name        = "tf_csvserver"
  servicetype = "HTTP"
  ipv46       = "10.10.10.10"
  port        = 80
}
```


## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the domain is bound. Changing this attribute forces a new resource to be created.
* `domainname` - (Required) Domain name for which to change the time to live (TTL) and/or backup service IP address. Changing this attribute forces a new resource to be created.
* `ttl` - (Optional) Time to live (TTL), in seconds, for the domain. Can be updated in place.
* `backupip` - (Optional) Backup service IP address for the domain. Can be updated in place.
* `cookiedomain` - (Optional) Domain attribute to set in the persistence cookie for the domain. Can be updated in place.
* `cookietimeout` - (Optional) Persistence cookie timeout, in minutes, for the domain. Can be updated in place.
* `sitedomainttl` - (Optional) TTL, in seconds, for all internally registered records that share this domain name as a suffix (site domain TTL). Can be updated in place.

Note: `name` and `domainname` form the resource identity, so changing either forces the binding to be replaced. The optional attributes (`ttl`, `backupip`, `cookiedomain`, `cookietimeout`, `sitedomainttl`) are updated in place - the provider re-issues the bind (PUT) request when any of them changes.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_domain_binding. It is a composite of comma-separated `key:value` pairs in the format `domainname:<domainname>,name:<name>` (values are URL-encoded). For example: `domainname:example.com,name:tf_csvserver`.


## Import

A csvserver_domain_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_csvserver_domain_binding.tf_csvserver_domain_binding domainname:example.com,name:tf_csvserver
```
