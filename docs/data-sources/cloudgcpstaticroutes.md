---
subcategory: "Cloud"
---

# Data Source: cloudgcpstaticroutes

The cloudgcpstaticroutes data source allows you to retrieve information about the Cloud API configuration that pushes static routes to GCP.


## Example usage

```terraform
data "citrixadc_cloudgcpstaticroutes" "tf_cloudgcpstaticroutes" {
}

output "status" {
  value = data.citrixadc_cloudgcpstaticroutes.tf_cloudgcpstaticroutes.status
}

output "project" {
  value = data.citrixadc_cloudgcpstaticroutes.tf_cloudgcpstaticroutes.project
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `status` - Whether pushing routes to GCP is enabled. Possible values: [ ENABLED, DISABLED ]
* `project` - GCP project name for which static routes functionality is enabled.
* `id` - The id of the cloudgcpstaticroutes. It is set to `cloudgcpstaticroutes-config`.
