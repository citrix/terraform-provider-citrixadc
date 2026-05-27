---
subcategory: "Application Firewall"
---

# Data Source: appfwgrpcwebtextcontenttype

The `appfwgrpcwebtextcontenttype` data source allows you to retrieve information about an Application Firewall gRPC-web-text content type configuration on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_appfwgrpcwebtextcontenttype" "tf_appfwgrpcwebtextcontenttype" {
  grpcwebtextcontenttypevalue = "application/grpc-web-text"
}

output "grpcwebtextcontenttypevalue" {
  value = data.citrixadc_appfwgrpcwebtextcontenttype.tf_appfwgrpcwebtextcontenttype.grpcwebtextcontenttypevalue
}

output "isregex" {
  value = data.citrixadc_appfwgrpcwebtextcontenttype.tf_appfwgrpcwebtextcontenttype.isregex
}
```


## Argument Reference

* `grpcwebtextcontenttypevalue` - (Required) Content type to be classified as gRPC-web-text.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwgrpcwebtextcontenttype`. It has the same value as the `grpcwebtextcontenttypevalue` attribute.
* `isregex` - Is gRPC-web-text content type a regular expression?. Possible values: [ REGEX, NOTREGEX ]
