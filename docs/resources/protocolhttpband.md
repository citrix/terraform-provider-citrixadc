---
subcategory: "Protocol"
---

# Resource: protocolhttpband

This resource is used to manage the HTTP band-size settings on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_protocolhttpband" "tf_protocolhttpband" {
  reqbandsize  = 100
  respbandsize = 1024
}
```


## Argument Reference

* `reqbandsize` - (Optional) Band size, in bytes, for HTTP request band statistics. For example, if you specify a band size of 100 bytes, statistics are maintained and displayed for the size ranges 0 - 99 bytes, 100 - 199 bytes, 200 - 299 bytes, and so on. Minimum value = 50. Defaults to `100`.
* `respbandsize` - (Optional) Band size, in bytes, for HTTP response band statistics. For example, if you specify a band size of 100 bytes, statistics are maintained and displayed for the size ranges 0 - 99 bytes, 100 - 199 bytes, 200 - 299 bytes, and so on. Minimum value = 50. Defaults to `1024`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the protocolhttpband resource. It is set to `protocolhttpband-config`.
