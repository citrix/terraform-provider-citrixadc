---
subcategory: "DNS"
---

# Data Source: citrixadc_dnspolicy

This data source is used to retrieve information about an existing DNS policy.

## Example Usage

```hcl
data "citrixadc_dnspolicy" "example" {
  name = "my_dnspolicy"
}
```

## Argument Reference

* `name` - (Required) Name for the DNS policy.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the DNS policy (same as `name`).
* `actionname` - Name of the DNS action to perform when the rule evaluates to TRUE. The built in actions function as follows: dns_default_act_Drop (Drop the DNS request), dns_default_act_Cachebypass (Bypass the DNS cache and forward the request to the name server).
* `cachebypass` - By pass dns cache for this.
* `drop` - The dns packet must be dropped.
* `logaction` - Name of the messagelog action to use for requests that match this policy.
* `preferredlocation` - The location used for the given policy. This is deprecated attribute. Please use preferredloclist.
* `preferredloclist` - The location list in priority order used for the given policy.
* `rule` - Expression against which DNS traffic is evaluated.
* `viewname` - The view name that must be used for the given policy.
