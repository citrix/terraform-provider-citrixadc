---
subcategory: "AAA"
---

# Data Source `aaapreauthenticationaction`

The aaapreauthenticationaction data source allows you to retrieve information about AAA preauthentication actions.


## Example usage

```terraform
data "citrixadc_aaapreauthenticationaction" "tf_aaapreauthenticationaction" {
  name = "tf_aaapreauthenticationaction"
}

output "preauthenticationaction" {
  value = data.citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction.preauthenticationaction
}

output "deletefiles" {
  value = data.citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction.deletefiles
}
```


## Argument Reference

* `name` - (Required) Name for the preauthentication action.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `preauthenticationaction` - Deny or allow login after endpoint analysis (EPA) results. Possible values: [ ALLOW, DENY ]
* `deletefiles` - String specifying the path(s) to the file(s) to be deleted by the endpoint analysis (EPA) tool. Multiple files are specified by using the path as a string with commas separating the different file paths.
* `defaultepagroup` - This is the default group that is chosen when the EPA check succeeds.
* `killprocess` - String specifying the name of a process to be terminated by the endpoint analysis (EPA) tool. Multiple processes are specified by using the process name as a string with commas separating multiple processes.
* `quarantinegroup` - This is the quarantine group that is chosen when the EPA check fails.

## Attribute Reference

* `id` - The id of the aaapreauthenticationaction. It has the same value as the `name` attribute.


## Import

A aaapreauthenticationaction can be imported using its name, e.g.

```shell
terraform import citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction tf_aaapreauthenticationaction
```
