---
subcategory: "System"
---

# Data Source `systembackup`

The systembackup data source allows you to retrieve information about system backup files on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_systembackup" "tf_systembackup" {
  filename = "my_backup.tgz"
}

output "backup_level" {
  value = data.citrixadc_systembackup.tf_systembackup.level
}

output "backup_comment" {
  value = data.citrixadc_systembackup.tf_systembackup.comment
}
```

## Argument Reference

* `filename` - (Required) Name of the backup file(*.tgz) to be restored.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Comment specified at the time of creation of the backup file(*.tgz).
* `includekernel` - Use this option to add kernel in the backup file.
* `level` - Level of data to be backed up.
* `skipbackup` - Use this option to skip taking backup during restore operation.
* `uselocaltimezone` - This option will create backup file with local timezone timestamp.

## Attribute Reference

* `id` - The id of the systembackup. It has the same value as the `filename` attribute.

## Import

A systembackup can be imported using its filename, e.g.

```shell
terraform import citrixadc_systembackup.tf_systembackup my_backup.tgz
```
