---
subcategory: "Cache"
---

# Data Source: citrixadc_cachepolicylabel

The `citrixadc_cachepolicylabel` data source is used to retrieve information about an existing Cache Policy Label configured on a Citrix ADC appliance.

## Example usage

```hcl
# Retrieve a cache policy label by name
data "citrixadc_cachepolicylabel" "example" {
  labelname = "my_cachepolicylabel"
}

# Use the retrieved data in other resources
output "policy_label_evaluates" {
  value = data.citrixadc_cachepolicylabel.example.evaluates
}

```

## Argument Reference

The following arguments are required:

* `labelname` - (Required) Name for the label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the label is created.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the cache policy label. It has the same value as the `labelname` attribute.
* `evaluates` - When to evaluate policies bound to this label: request-time or response-time. Possible values: `REQ`, `RES`, `MSSQL_REQ`, `MSSQL_RES`, `MYSQL_REQ`, `MYSQL_RES`.
* `newname` - New name for the cache-policy label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
