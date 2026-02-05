---
subcategory: "Network"
---

# Data Source `arpparam`

The arpparam data source allows you to retrieve information about ARP parameters configuration.


## Example usage

```terraform
data "citrixadc_arpparam" "tf_arpparam" {
}

output "timeout" {
  value = data.citrixadc_arpparam.tf_arpparam.timeout
}

output "spoofvalidation" {
  value = data.citrixadc_arpparam.tf_arpparam.spoofvalidation
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `timeout` - Time-out value (aging time) for the dynamically learned ARP entries, in seconds. The new value applies only to ARP entries that are dynamically learned after the new value is set. Previously existing ARP entries expire after the previously configured aging time.
* `spoofvalidation` - Enable/disable ARP spoofing validation. Possible values: `ENABLED`, `DISABLED`.

## Attribute Reference

* `id` - The id of the arpparam. It is a system-generated identifier.
