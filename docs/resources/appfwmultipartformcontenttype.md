---
subcategory: "Application Firewall"
---

# Resource: appfwmultipartformcontenttype

The appfwmultipartformcontenttype resource is used to create appfw multipart form content type Resource.


## Example usage

```hcl
resource "citrixadc_appfwmultipartformcontenttype" "tf_multipartform" {
  multipartformcontenttypevalue = "data/tf_multipartform"
  isregex                       = "REGEX"
}
```


## Argument Reference

* `multipartformcontenttypevalue` - (Required) Content type to be classified as multipart form
* `isregex` - (Optional) Is multipart_form content type a regular expression?


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwmultipartformcontenttype. It has the same value as the `multipartformcontenttypevalue` attribute.


## Import

A appfwmultipartformcontenttype can be imported using its multipartformcontenttypevalue, e.g.

```shell
terraform import citrixadc_appfwmultipartformcontenttype.tf_multipartform data/tf_multipartform
```
