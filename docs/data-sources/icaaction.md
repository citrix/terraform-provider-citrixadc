---
subcategory: "ICA"
---

# Data Source `icaaction`

The icaaction data source allows you to retrieve information about ICA actions that are used to configure settings for ICA connections.


## Example usage

```terraform
data "citrixadc_icaaction" "tf_icaaction" {
  name = "my_ica_action"
}

output "accessprofilename" {
  value = data.citrixadc_icaaction.tf_icaaction.accessprofilename
}

output "latencyprofilename" {
  value = data.citrixadc_icaaction.tf_icaaction.latencyprofilename
}
```


## Argument Reference

* `name` - (Required) Name for the ICA action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `accessprofilename` - Name of the ica accessprofile to be associated with this action.
* `latencyprofilename` - Name of the ica latencyprofile to be associated with this action.
* `newname` - New name for the ICA action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#),period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `id` - The id of the icaaction. It has the same value as the `name` attribute.


## Import

An icaaction can be imported using its name, e.g.

```shell
terraform import citrixadc_icaaction.tf_icaaction my_ica_action
```
