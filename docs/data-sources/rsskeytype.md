---
subcategory: "Network"
---

# Data Source: citrixadc_rsskeytype

The rsskeytype data source allows you to retrieve information about RSS key type configuration.

## Example Usage

```terraform
data "citrixadc_rsskeytype" "tf_rsskeytype" {
}

output "rsstype" {
  value = data.citrixadc_rsskeytype.tf_rsskeytype.rsstype
}
```

## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `rsstype` - Type of RSS key. Possible values: `SYMMETRIC`, `ASYMMETRIC`.
* `id` - The id of the rsskeytype. It is a system-generated identifier.
