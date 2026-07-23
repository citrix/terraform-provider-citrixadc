---
subcategory: "Cloud"
---

# Data Source: cloudparaminternal

The cloudparaminternal data source allows you to retrieve information about the internal cloud parameters configuration.

~> **Note:** The underlying `show cloud paramInternal` operation is platform-gated; on unsupported platforms the read returns empty attribute values.


## Example usage

```hcl
data "citrixadc_cloudparaminternal" "example" {}

output "cloudparaminternal_details" {
  value = data.citrixadc_cloudparaminternal.example
}
```

## Example usage with Resource

```hcl
data "citrixadc_cloudparaminternal" "tf_cloudparaminternal" {
  depends_on = [citrixadc_cloudparaminternal.tf_cloudparaminternal]
}

output "configured_cloudparaminternal" {
  value = data.citrixadc_cloudparaminternal.tf_cloudparaminternal.nonftumode
}
```

## Argument Reference

This data source is a singleton and does not require any lookup arguments. It retrieves the current internal cloud parameters configuration from the Citrix ADC.

## Attribute Reference

In addition to the above arguments, the following attributes are exported:

* `id` - The id of the cloudparaminternal data source. It is set to `cloudparaminternal-config`.
* `nonftumode` - Indicates whether the management GUI is in first-time-user (FTU) mode or not. Possible values: `YES`, `NO`. May be empty on platforms where the GET operation is not supported.
