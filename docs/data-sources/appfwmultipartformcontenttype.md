---
subcategory: "Application Firewall"
---

# Data Source: appfwmultipartformcontenttype

The appfwmultipartformcontenttype data source allows you to retrieve information about application firewall multipart form content types.

## Example usage

```terraform
data "citrixadc_appfwmultipartformcontenttype" "tf_multipartform" {
  multipartformcontenttypevalue = "date/tf_multipartform"
}

output "isregex" {
  value = data.citrixadc_appfwmultipartformcontenttype.tf_multipartform.isregex
}

output "multipartformcontenttypevalue" {
  value = data.citrixadc_appfwmultipartformcontenttype.tf_multipartform.multipartformcontenttypevalue
}
```

## Argument Reference

* `multipartformcontenttypevalue` - (Required) Content type to be classified as multipart form.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwmultipartformcontenttype. It has the same value as the `multipartformcontenttypevalue` attribute.
* `isregex` - Is multipart_form content type a regular expression?

## Import

A appfwmultipartformcontenttype can be imported using its multipartformcontenttypevalue, e.g.

```shell
terraform import citrixadc_appfwmultipartformcontenttype.tf_multipartform date/tf_multipartform
```
