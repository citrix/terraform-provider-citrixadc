---
subcategory: "Application Firewall"
---

# Resource: appfwgrpccontenttype

The `appfwgrpccontenttype` resource is used to create an Application Firewall gRPC content type configuration on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_appfwgrpccontenttype" "tf_appfwgrpccontenttype" {
  grpccontenttypevalue = "application/grpc"
  isregex              = "NOTREGEX"
}
```


## Argument Reference

* `grpccontenttypevalue` - (Required) Content type to be classified as gRPC. Changing this value forces a new resource to be created.
* `isregex` - (Optional) Is gRPC content type a regular expression?. Possible values: [ REGEX, NOTREGEX ]. Defaults to `"NOTREGEX"`. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwgrpccontenttype`. It has the same value as the `grpccontenttypevalue` attribute.


## Import

An `appfwgrpccontenttype` resource can be imported using its `grpccontenttypevalue`, e.g.

```shell
terraform import citrixadc_appfwgrpccontenttype.tf_appfwgrpccontenttype application/grpc
```
