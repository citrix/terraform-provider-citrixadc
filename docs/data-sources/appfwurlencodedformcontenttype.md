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

### Read-only appfwurlencodedformcontenttype metadata

These attributes are GET-only (Computed) and are returned by the appliance on a read; they are not configurable on the `citrixadc_appfwurlencodedformcontenttype` resource. Any attribute the appliance does not return is `null`.

* `builtin` - Flag to determine if urlencoded form contenttype is built-in or not. A list of strings (for example `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`).
* `feature` - The feature to be checked while applying this config.

## Import

A appfwurlencodedformcontenttype can be imported using its urlencodedformcontenttypevalue, e.g.

```shell
terraform import citrixadc_appfwurlencodedformcontenttype.tf_urlencodedform application/x-www-form-urlencoded
```
