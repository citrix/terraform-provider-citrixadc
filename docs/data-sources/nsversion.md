---
subcategory: "NS"
---

# Data Source `nsversion`

The nsversion data source allows you to retrieve information about the version of the target ADC.

## Example Usage

```terraform
data "citrixadc_nsversion" "nsversion" {
  installedversion = true
}

```

## Argument Reference

- `installedversion` - (Optional) Installed version.

## Attributes Reference

The following attributes are exported.

- `version` - String describing the version of the target ADC.
- `mode` - Kernel mode (KMPE/VMPE).
