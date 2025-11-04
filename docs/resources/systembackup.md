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
* `comment` - (Optional) Comment specified at the time of creation of the backup file(*.tgz).
* `includekernel` - (Optional) Use this option to add kernel in the backup file
* `level` - (Optional) Level of data to be backed up.
* `skipbackup` - (Optional) Use this option to skip taking backup during restore operation
* `uselocaltimezone` - (Optional) This option will create backup file with local timezone timestamp
* `action` - (Optional) Use this option to Create, Add or Restore a backup. Valid values are 'create', 'add' and 'restore'. Default value is 'create'


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systembackup. It is a unique string prefixed with `filename` attribute's value.
