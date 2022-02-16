---
subcategory: "Authentication"
---

# Resource: authenticationpolicylabel

The authenticationpolicylabel resource is used to create authentication policy label.


## Example usage

```hcl
resource "citrixadc_authenticationpolicylabel" "tf_authenticationpolicylabel" {
  labelname = "tf_authenticationpolicylabel"
  type      = "AAATM_REQ"
  comment   = "Testing"
}
```


## Argument Reference

* `labelname` - (Required) Name for the new authentication policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy label" or 'authentication policy label').
* `comment` - (Optional) Any comments to preserve information about this authentication policy label.
* `loginschema` - (Optional) Login schema associated with authentication policy label. Login schema defines the UI rendering by providing customization option of the fields. If user intervention is not needed for a given factor such as group extraction, a loginSchema whose authentication schema is "noschema" should be used.
* `newname` - (Optional) The new name of the auth policy label
* `type` - (Optional) Type of feature (aaatm or rba) against which to match the policies bound to this policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationpolicylabel. It has the same value as the `labelname` attribute.


## Import

A authenticationpolicylabel can be imported using its labelname, e.g.

```shell
terraform import citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel tf_authenticationpolicylabel
```
