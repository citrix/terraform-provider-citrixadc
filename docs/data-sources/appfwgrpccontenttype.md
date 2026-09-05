---
subcategory: "Application Firewall"
---

# Data Source: appfwgrpccontenttype

The `appfwgrpccontenttype` data source allows you to retrieve information about an Application Firewall gRPC content type configuration on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_appfwgrpccontenttype" "tf_appfwgrpccontenttype" {
  grpccontenttypevalue = "application/grpc"
}

output "grpccontenttypevalue" {
  value = data.citrixadc_appfwgrpccontenttype.tf_appfwgrpccontenttype.grpccontenttypevalue
}

output "isregex" {
  value = data.citrixadc_appfwgrpccontenttype.tf_appfwgrpccontenttype.isregex
}
```


## Argument Reference

* `grpccontenttypevalue` - (Required) Content type to be classified as gRPC.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwgrpccontenttype`. It has the same value as the `grpccontenttypevalue` attribute.
* `isregex` - Is gRPC content type a regular expression?. Possible values: [ REGEX, NOTREGEX ]

### Read-only appfwgrpccontenttype metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_appfwgrpccontenttype` resource). They are Computed/GET-only and are `null` when the appliance omits them.

* `builtin` - Flag to determine if the gRPC content type is built-in or not (for example `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`). A list of strings.
* `feature` - The feature to be checked while applying this configuration.
