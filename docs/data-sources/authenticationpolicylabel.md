---
subcategory: "Authentication"
---

# Data Source `authenticationpolicylabel`

The authenticationpolicylabel data source allows you to retrieve information about authentication policy labels.


## Example usage

```terraform
data "citrixadc_authenticationpolicylabel" "tf_authenticationpolicylabel" {
  labelname = "my_authenticationpolicylabel"
}

output "type" {
  value = data.citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel.type
}

output "comment" {
  value = data.citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel.comment
}
```


## Argument Reference

* `labelname` - (Required) Name for the new authentication policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Any comments to preserve information about this authentication policy label.
* `loginschema` - Login schema associated with authentication policy label. Login schema defines the UI rendering by providing customization option of the fields. If user intervention is not needed for a given factor such as group extraction, a loginSchema whose authentication schema is "noschema" should be used.
* `newname` - The new name of the auth policy label.
* `type` - Type of feature (aaatm or rba) against which to match the policies bound to this policy label.

## Attribute Reference

* `id` - The id of the authenticationpolicylabel. It has the same value as the `labelname` attribute.


## Import

A authenticationpolicylabel can be imported using its labelname, e.g.

```shell
terraform import citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel my_authenticationpolicylabel
```
