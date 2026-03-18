---
subcategory: "Application Firewall"
---

# Data Source: appfwurlencodedformcontenttype

The appfwurlencodedformcontenttype data source allows you to retrieve information about application firewall URL-encoded form content types.

## Example usage

```terraform
data "citrixadc_appfwurlencodedformcontenttype" "tf_urlencodedform" {
  urlencodedformcontenttypevalue = "application/x-www-form-urlencoded"
}

output "isregex" {
  value = data.citrixadc_appfwurlencodedformcontenttype.tf_urlencodedform.isregex
}

output "urlencodedformcontenttypevalue" {
  value = data.citrixadc_appfwurlencodedformcontenttype.tf_urlencodedform.urlencodedformcontenttypevalue
}
```

## Argument Reference

* `urlencodedformcontenttypevalue` - (Required) Content type to be classified as urlencoded form.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwurlencodedformcontenttype. It has the same value as the `urlencodedformcontenttypevalue` attribute.
* `isregex` - Is urlencoded form content type a regular expression?

## Import

A appfwurlencodedformcontenttype can be imported using its urlencodedformcontenttypevalue, e.g.

```shell
terraform import citrixadc_appfwurlencodedformcontenttype.tf_urlencodedform application/x-www-form-urlencoded
```
