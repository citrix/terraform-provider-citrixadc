---
subcategory: "Protocol"
---

# Resource: protocolhttpband

Configures the band-size granularity used by the Citrix ADC when it collects HTTP request and response size statistics. Tuning the band sizes lets you control the width of the size buckets that appear in the HTTP band statistics reports, so you can profile the distribution of request and response payload sizes at the resolution you need.

This is a singleton settings resource: a single configuration object always exists on the appliance. Creating the resource sets the values; destroying it leaves the last configured values in effect on the appliance.


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
