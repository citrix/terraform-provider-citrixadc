---
subcategory: "Content Switching"
---

# Data Source: csvserver_domain_binding

The csvserver_domain_binding data source allows you to retrieve information about a DNS domain bound to a content switching virtual server, including its TTL and backup service IP address.


## Example usage

```terraform
data "citrixadc_csvserver_domain_binding" "example" {
  name       = "tf_csvserver"
  domainname = "example.com"
}

output "domain_ttl" {
  value = data.citrixadc_csvserver_domain_binding.example.ttl
}
```


## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the domain is bound.
* `domainname` - (Required) Domain name bound to the content switching virtual server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_domain_binding. It is a composite of comma-separated `key:value` pairs in the format `domainname:<domainname>,name:<name>`.
* `ttl` - Time to live (TTL), in seconds, for the domain.
* `backupip` - Backup service IP address for the domain.
* `cookiedomain` - Domain attribute set in the persistence cookie for the domain.
* `cookietimeout` - Persistence cookie timeout, in minutes, for the domain.
* `sitedomainttl` - TTL, in seconds, for all internally registered records that share this domain name as a suffix.
