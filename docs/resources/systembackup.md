---
subcategory: "System"
---

# Resource: systembackup

The systembackup resource is used to create the systembackup resource.


## Example usage

```hcl
resource "citrixadc_systembackup_create" "tf_systembackup_create" {
  filename         = "my_restore_file"
  level            = "basic"
  uselocaltimezone = "true"
}
```


## Argument Reference

* `filename` - (Required) Name of the backup file(*.tgz) to be restored. Maximum length =  63
* `uselocaltimezone` - (Optional) This option will create backup file with local timezone timestamp.
* `level` - (Optional) Level of data to be backed up. Possible values: [ basic, full ]
* `includekernel` - (Optional) Use this option to add kernel in the backup file. Possible values: [ NO, YES ]
* `comment` - (Optional) Comment specified at the time of creation of the backup file(*.tgz).
* `skipbackup` - (Optional) Use this option to skip taking backup during restore operation.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systembackup_create. It is a unique string prefixed with `filename` attribute's value.
