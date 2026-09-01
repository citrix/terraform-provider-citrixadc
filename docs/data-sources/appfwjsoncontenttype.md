---
subcategory: "Application Firewall"
---

# Data Source: appfwjsoncontenttype

The `appfwjsoncontenttype` data source allows you to retrieve information about an Application Firewall JSON content type configuration.

## Example usage

```terraform
data "citrixadc_appfwjsoncontenttype" "tf_appfwjsoncontenttype" {
  jsoncontenttypevalue = "application/json"
}

output "isregex" {
  value = data.citrixadc_appfwjsoncontenttype.tf_appfwjsoncontenttype.isregex
}

output "jsoncontenttypevalue" {
  value = data.citrixadc_appfwjsoncontenttype.tf_appfwjsoncontenttype.jsoncontenttypevalue
}
```

## Argument Reference

* `jsoncontenttypevalue` - (Required) Content type to be classified as JSON.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwjsoncontenttype. It has the same value as the `jsoncontenttypevalue` attribute.
* `isregex` - Is json content type a regular expression?. Possible values: [ REGEX, NOTREGEX ]

### Read-only appfwjsoncontenttype metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_appfwjsoncontenttype` resource) and are Computed/GET-only. Any attribute the appliance does not return is `null`.

* `builtin` - Flag to determine if jsoncontenttype is built-in or not. A list of strings. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ]
* `feature` - The feature to be checked while applying this config.

## Import

A appfwjsoncontenttype can be imported using its jsoncontenttypevalue, e.g.

```shell
terraform import citrixadc_appfwjsoncontenttype.tf_appfwjsoncontenttype application/json
```
