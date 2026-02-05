---
subcategory: "Application Firewall"
---

# Data Source `appfwconfidfield`

The appfwconfidfield data source allows you to retrieve information about an existing appfwconfidfield.


## Example usage

```terraform
data "citrixadc_appfwconfidfield" "tf_confidfield" {
  fieldname = "tf_confidfield"
  url       = "www.example.com/"
}

output "fieldname" {
  value = data.citrixadc_appfwconfidfield.tf_confidfield.fieldname
}

output "url" {
  value = data.citrixadc_appfwconfidfield.tf_confidfield.url
}

output "comment" {
  value = data.citrixadc_appfwconfidfield.tf_confidfield.comment
}

output "isregex" {
  value = data.citrixadc_appfwconfidfield.tf_confidfield.isregex
}

output "state" {
  value = data.citrixadc_appfwconfidfield.tf_confidfield.state
}
```


## Argument Reference

* `fieldname` - (Required) Name of the form field to designate as confidential.
* `url` - (Required) URL of the web page that contains the web form.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwconfidfield. It is the concatenation of `fieldname` and `url` attribute.
* `comment` - Any comments to preserve information about the form field designation.
* `isregex` - Method of specifying the form field name. Available settings function as follows: * REGEX. Form field is a regular expression. * NOTREGEX. Form field is a literal string.
* `state` - Enable or disable the confidential field designation.
