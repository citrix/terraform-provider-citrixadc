---
subcategory: "Application Firewall"
---

# Data Source: appfwxmlcontenttype

The `appfwxmlcontenttype` data source allows you to retrieve information about an Application Firewall XML content type configuration.

## Example usage

```terraform
data "citrixadc_appfwxmlcontenttype" "tf_appfwxmlcontenttype" {
  xmlcontenttypevalue = "application/xml"
}

output "isregex" {
  value = data.citrixadc_appfwxmlcontenttype.tf_appfwxmlcontenttype.isregex
}

output "xmlcontenttypevalue" {
  value = data.citrixadc_appfwxmlcontenttype.tf_appfwxmlcontenttype.xmlcontenttypevalue
}
```

## Argument Reference

* `xmlcontenttypevalue` - (Required) Content type to be classified as XML.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwxmlcontenttype. It has the same value as the `xmlcontenttypevalue` attribute.
* `isregex` - Is field name a regular expression?. Possible values: [ REGEX, NOTREGEX ]

## Import

A appfwxmlcontenttype can be imported using its xmlcontenttypevalue, e.g.

```shell
terraform import citrixadc_appfwxmlcontenttype.tf_appfwxmlcontenttype application/xml
```
