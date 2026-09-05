---
subcategory: "DNS"
---

# Data Source: dnsaction

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

### Read-only dnsaction metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_dnsaction` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `drop` - The dns packet must be dropped. Possible values: [ YES, NO ].
* `cachebypass` - By pass dns cache for this. Possible values: [ YES, NO ].
* `builtin` - Flag to determine whether DNS action is default or not. A list of strings. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ].
* `feature` - The feature to be checked while applying this config.
