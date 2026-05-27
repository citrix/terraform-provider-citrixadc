---
subcategory: "Application Firewall"
---

# Resource: appfwgrpcwebjsoncontenttype

The `appfwgrpcwebjsoncontenttype` resource is used to create an Application Firewall gRPC-web-json content type configuration on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_appfwgrpcwebjsoncontenttype" "tf_appfwgrpcwebjsoncontenttype" {
  grpcwebjsoncontenttypevalue = "application/grpc-web+json"
  isregex                     = "NOTREGEX"
}
```


## Argument Reference

* `grpcwebjsoncontenttypevalue` - (Required) Content type to be classified as gRPC-web-json. Changing this value forces a new resource to be created.
* `isregex` - (Optional) Is gRPC-web-json content type a regular expression?. Possible values: [ REGEX, NOTREGEX ]. Defaults to `"NOTREGEX"`. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwgrpcwebjsoncontenttype`. It has the same value as the `grpcwebjsoncontenttypevalue` attribute.


## Import

An `appfwgrpcwebjsoncontenttype` resource can be imported using its `grpcwebjsoncontenttypevalue`, e.g.

```shell
terraform import citrixadc_appfwgrpcwebjsoncontenttype.tf_appfwgrpcwebjsoncontenttype application/grpc-web+json
```
