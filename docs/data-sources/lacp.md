---
subcategory: "Network"
---

# Data Source `lacp`

The lacp data source allows you to retrieve information about Link Aggregation Control Protocol (LACP) configuration.


## Example usage

```terraform
data "citrixadc_lacp" "tf_lacp" {
  ownernode = 255
}

output "syspriority" {
  value = data.citrixadc_lacp.tf_lacp.syspriority
}
```


## Argument Reference

The following arguments are required:

* `ownernode` - (Required) The owner node in a cluster for which we want to set the lacp priority. Owner node can vary from 0 to 31. Ownernode value of 254 is used for Cluster.

## Attribute Reference

The following attributes are available:

* `ownernode` - The owner node in a cluster for which we want to set the lacp priority. Owner node can vary from 0 to 31. Ownernode value of 254 is used for Cluster.
* `syspriority` - Priority number that determines which peer device of an LACP LA channel can have control over the LA channel. This parameter is globally applied to all LACP channels on the Citrix ADC. The lower the number, the higher the priority.
* `id` - The id of the lacp resource. It is a system-generated identifier.
