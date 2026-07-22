---
subcategory: "Cloud"
---

# Resource: cloudservice_check

The cloudservice_check resource performs the NITRO `cloudservice` `check` action, which checks the cloud service configuration on the Citrix ADC. It is an action-only, zero-attribute resource: applying it triggers the check, and there are no configurable arguments.


## Example usage

```hcl
resource "citrixadc_cloudservice_check" "tf_cloudservice_check" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cloudservice_check resource. It is set to `cloudservice_check`.
