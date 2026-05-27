---
subcategory: "Application Firewall"
---

# Resource: appfwgrpcwebtextcontenttype

The `appfwgrpcwebtextcontenttype` resource is used to create an Application Firewall gRPC-web-text content type configuration on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_appfwgrpcwebtextcontenttype" "tf_appfwgrpcwebtextcontenttype" {
  grpcwebtextcontenttypevalue = "application/grpc-web-text"
  isregex                     = "NOTREGEX"
}
```


## Argument Reference

* `grpcwebtextcontenttypevalue` - (Required) Content type to be classified as gRPC-web-text. Changing this value forces a new resource to be created.
* `isregex` - (Optional) Is gRPC-web-text content type a regular expression?. Possible values: [ REGEX, NOTREGEX ]. Defaults to `"NOTREGEX"`. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwgrpcwebtextcontenttype`. It has the same value as the `grpcwebtextcontenttypevalue` attribute.


## Import

An `appfwgrpcwebtextcontenttype` resource can be imported using its `grpcwebtextcontenttypevalue`, e.g.

```shell
terraform import citrixadc_appfwgrpcwebtextcontenttype.tf_appfwgrpcwebtextcontenttype application/grpc-web-text
```
