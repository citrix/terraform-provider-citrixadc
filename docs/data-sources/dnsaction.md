---
subcategory: "DNS"
---

# Data Source: citrixadc_dnsaction

This data source is used to retrieve information about an existing DNS action.

## Example Usage

```hcl
data "citrixadc_dnsaction" "example" {
  actionname = "my_dnsaction"
}
```

## Argument Reference

* `actionname` - (Required) Name of the DNS action.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the DNS action (same as `actionname`).
* `actiontype` - The type of DNS action that is being configured.
* `dnsprofilename` - Name of the DNS profile to be associated with the transaction for which the action is chosen.
* `ipaddress` - List of IP address to be returned in case of rewrite_response actiontype. They can be of IPV4 or IPV6 type. In case of set command We will remove all the IP address previously present in the action and will add new once given in set dns action command.
* `preferredloclist` - The location list in priority order used for the given action.
* `ttl` - Time to live, in seconds.
* `viewname` - The view name that must be used for the given action.
