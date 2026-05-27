---
subcategory: "Application Firewall"
---

# Data Source: appfwgrpcwebjsoncontenttype

The `appfwgrpcwebjsoncontenttype` data source allows you to retrieve information about an Application Firewall gRPC-web-json content type configuration on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_appfwgrpcwebjsoncontenttype" "tf_appfwgrpcwebjsoncontenttype" {
  grpcwebjsoncontenttypevalue = "application/grpc-web+json"
}

output "grpcwebjsoncontenttypevalue" {
  value = data.citrixadc_appfwgrpcwebjsoncontenttype.tf_appfwgrpcwebjsoncontenttype.grpcwebjsoncontenttypevalue
}

output "isregex" {
  value = data.citrixadc_appfwgrpcwebjsoncontenttype.tf_appfwgrpcwebjsoncontenttype.isregex
}
```


## Argument Reference

* `grpcwebjsoncontenttypevalue` - (Required) Content type to be classified as gRPC-web-json.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwgrpcwebjsoncontenttype`. It has the same value as the `grpcwebjsoncontenttypevalue` attribute.
* `isregex` - Is gRPC-web-json content type a regular expression?. Possible values: [ REGEX, NOTREGEX ]
