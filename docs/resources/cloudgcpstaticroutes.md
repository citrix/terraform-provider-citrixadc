---
subcategory: "Cloud"
---

# Resource: cloudgcpstaticroutes

This resource is used to manage the Cloud API configuration that pushes static routes to GCP.


## Example usage

```hcl
resource "citrixadc_cloudgcpstaticroutes" "tf_cloudgcpstaticroutes" {
  status  = "ENABLED"
  project = "my-gcp-project"
}
```


## Argument Reference

* `status` - (Optional) status to push routes or not. Possible values: [ ENABLED, DISABLED ]
* `project` - (Optional) GCP project name for which static routes functionality is enabled. Minimum length 1, maximum length 127.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cloudgcpstaticroutes. It is set to `cloudgcpstaticroutes-config`.


## Import

A cloudgcpstaticroutes can be imported using its id, e.g.

```shell
terraform import citrixadc_cloudgcpstaticroutes.tf_cloudgcpstaticroutes cloudgcpstaticroutes-config
```
