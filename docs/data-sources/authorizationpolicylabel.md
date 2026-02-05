---
subcategory: "Authorization"
---

# Data Source `authorizationpolicylabel`

The authorizationpolicylabel data source allows you to retrieve information about an existing authorization policy label.


## Example usage

```terraform
data "citrixadc_authorizationpolicylabel" "tf_authorizationpolicylabel" {
  labelname = "my_authorizationpolicylabel"
}

output "id" {
  value = data.citrixadc_authorizationpolicylabel.tf_authorizationpolicylabel.id
}

output "labelname" {
  value = data.citrixadc_authorizationpolicylabel.tf_authorizationpolicylabel.labelname
}
```


## Argument Reference

* `labelname` - (Required) Name for the new authorization policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the authorization policy is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authorizationpolicylabel. It has the same value as the `labelname` attribute.
* `newname` - The new name of the auth policy label.


## Import

A authorizationpolicylabel can be imported using its labelname, e.g.

```shell
terraform import citrixadc_authorizationpolicylabel.tf_authorizationpolicylabel my_authorizationpolicylabel
```
