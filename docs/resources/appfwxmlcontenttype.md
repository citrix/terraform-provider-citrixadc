---
subcategory: "Application Firewall"
---

# Resource: appfwxmlcontenttype

The `appfwxmlcontenttype` resource is used to crate Application Firewall XML ContentType.

## Example usage

``` hcl
resource "citrixadc_appfwxmlcontenttype" "demo_appfwxmlcontenttype" {
  xmlcontenttypevalue = "demo.*test"
  isregex = "REGEX"
}
```

## Argument Reference

* `xmlcontenttypevalue` - Content type to be classified as XML.
* `isregex` - (Optional) Is field name a regular expression?. Possible values: [ REGEX, NOTREGEX ]

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwxmlcontenttype`. It has the same value as the `xmlcontenttypevalue` attribute.
