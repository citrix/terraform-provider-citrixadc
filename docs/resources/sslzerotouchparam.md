---
subcategory: "SSL"
---

# Resource: sslzerotouchparam

The sslzerotouchparam resource is used to configure the SSL zero touch (OCSP) parameters. This is a singleton (global) configuration resource.


## Example usage

```hcl
resource "citrixadc_sslzerotouchparam" "tf_sslzerotouchparam" {
  ocspcachetimeout       = 60
  ocspbatchingdepth      = 4
  ocspbatchingdelay      = 100
  ocsptrustresponder     = "YES"
  ocspusenonce           = "DISABLED"
  ocsphttpmethod         = "GET"
  ocspproducedattimeskew = 600
}
```


## Argument Reference

* `ocspcachetimeout` - (Optional) Timeout(in minutes) for caching the OCSP response. Minimum value =  1 Maximum value =  43200
* `ocspbatchingdepth` - (Optional) Number of certificates to batch together into one OCSP request. Batching avoids overloading the OCSP responder. A value of 1 signifies that each request is queried independently. For a value greater than 1, specify a timeout (batching delay) to avoid inordinately delaying the processing of a single certificate. Minimum value =  1 Maximum value =  8
* `ocspbatchingdelay` - (Optional) Maximum time, in milliseconds, to wait to accumulate OCSP requests to batch. Does not apply if the Batching Depth is 1. Minimum value =  1 Maximum value =  10000
* `ocspresptimeout` - (Optional) Time, in milliseconds, to wait for an OCSP response. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server. Includes Batching Delay time. Minimum value =  100 Maximum value =  120000
* `ocspurlresolvetimeout` - (Optional) Time, in milliseconds, to wait for an OCSP URL Resolution. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server. Minimum value =  100 Maximum value =  2000
* `ocsptrustresponder` - (Optional) If trustResponder is set to YES, signature verification will be skipped for the OCSP response. Possible values: [ YES, NO ]
* `ocspproducedattimeskew` - (Optional) Time, in seconds, for which the Citrix ADC waits before considering the response as invalid. The response is considered invalid if the Produced At time stamp in the OCSP response exceeds or precedes the current Citrix ADC clock time by the amount of time specified. Minimum value =  0 Maximum value =  86400
* `ocspusenonce` - (Optional) Enable the OCSP nonce extension, which is designed to prevent replay attacks. Possible values: [ ENABLED, DISABLED ]
* `ocsphttpmethod` - (Optional) HTTP method used to send ocsp request. POST is the default httpmethod. If request length is > 255, POST wil be used even if GET is set as httpMethod. Possible values: [ GET, POST ]


## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The id of the sslzerotouchparam. It is a system-generated identifier.


## Import

A sslzerotouchparam configuration can be imported using its id, e.g.

```shell
terraform import citrixadc_sslzerotouchparam.tf_sslzerotouchparam sslzerotouchparam-config
```
