---
subcategory: "System"
---

# Resource: systembackup

The systembackup resource is used to add the systembackup to another ADC.


## Example usage

```hcl
resource "citrixadc_systembackup" "tf_systembackup" {
  filename         = "my_restore_file.tgz"
}

```


## Argument Reference

* `filename` - (Required) Name of the backup file to be restored. Maximum length =  63


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systembackup. It is a unique string prefixed with `filename` attribute's value.
