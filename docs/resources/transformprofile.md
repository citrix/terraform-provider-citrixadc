---
subcategory: "Transform"
---

# Resource: transformprofile

The transformprofile resource is used to create transform profiles.


## Example usage

```hcl
resource "citrixadc_transformprofile" "tf_trans_profile" {
  name    = "tf_trans_profile"
  comment = "Some comment"
}
```


## Argument Reference

* `name` - (Optional) Name for the URL transformation profile. Must begin with a letter, number, or the underscore character (\_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the URL transformation profile is added.
* `type` - (Optional) Type of transformation. Always URL for URL Transformation profiles. Possible values: [ URL ]
* `onlytransformabsurlinbody` - (Optional) In the HTTP body, transform only absolute URLs. Relative URLs are ignored. Possible values: [ on, off ]
* `comment` - (Optional) Any comments to preserve information about this URL Transformation profile.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the transformprofile. It has the same value as the `name` attribute.


## Import

A transformprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_transformprofile.tf_trans_profile tf_trans_profile
```
