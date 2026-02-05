---
subcategory: "Transform"
---

# Data Source `transformprofile`

The transformprofile data source allows you to retrieve information about a URL Transformation profile.


## Example usage

```terraform
data "citrixadc_transformprofile" "tf_trans_profile" {
  name = "tf_trans_profile"
}

output "comment" {
  value = data.citrixadc_transformprofile.tf_trans_profile.comment
}

output "onlytransformabsurlinbody" {
  value = data.citrixadc_transformprofile.tf_trans_profile.onlytransformabsurlinbody
}
```


## Argument Reference

* `name` - (Required) Name for the URL transformation profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Any comments to preserve information about this URL Transformation profile.
* `onlytransformabsurlinbody` - In the HTTP body, transform only absolute URLs. Relative URLs are ignored.
* `type` - Type of transformation. Always URL for URL Transformation profiles.

## Attribute Reference

* `id` - The id of the transformprofile. It has the same value as the `name` attribute.


## Import

A transformprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_transformprofile.tf_trans_profile tf_trans_profile
```
