---
subcategory: "System"
---

# Resource: systembackup_restore

The systembackup_restore resource is used to apply the restore operation for systembackup.


## Example usage

```hcl
resource "citrixadc_systembackup_restore" "tf_systembackup_restore" {
  filename   = "my_restore_file.tgz"
  skipbackup = "true"
}
```


## Argument Reference

* `filename` - (Required) Name of the backup file(*.tgz) to be restored. Maximum length =  63
* `skipbackup` - (Optional) Use this option to skip taking backup during restore operation.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systembackup_restore. It is a unique string prefixed with `filename` attribute's value.

