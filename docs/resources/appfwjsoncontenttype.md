---
subcategory: "Application Firewall"
---

# Resource: appfwjsoncontenttype

The `appfwjsoncontenttype` resource is used to create Application Firewall ContentType resource.

## Example usage

``` hcl
resource "citrixadc_appfwjsoncontenttype" "demo_appfwjsoncontenttype" {
  jsoncontenttypevalue = "demo.*test"
  isregex = "REGEX"
}
```

## Argument Reference

* `jsoncontenttypevalue` - Content type to be classified as JSON.
* `isregex` - (Optional) Is json content type a regular expression?. Possible values: [ REGEX, NOTREGEX ]

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwjsoncontenttype`. It has the same value as the `jsoncontenttypevalue` attribute.
