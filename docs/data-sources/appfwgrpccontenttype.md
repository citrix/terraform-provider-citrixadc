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
