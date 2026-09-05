---
subcategory: "SSL"
---

# Data Source: sslzerotouchparam

The sslzerotouchparam data source allows you to retrieve information about the SSL zero touch (OCSP) parameters configuration.


## Example usage

```terraform
data "citrixadc_sslzerotouchparam" "tf_sslzerotouchparam" {
}

output "ocsphttpmethod" {
  value = data.citrixadc_sslzerotouchparam.tf_sslzerotouchparam.ocsphttpmethod
}

output "ocspusenonce" {
  value = data.citrixadc_sslzerotouchparam.tf_sslzerotouchparam.ocspusenonce
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `ocspcachetimeout` - Timeout(in minutes) for caching the OCSP response.
* `ocspbatchingdepth` - Number of certificates to batch together into one OCSP request.
* `ocspbatchingdelay` - Maximum time, in milliseconds, to wait to accumulate OCSP requests to batch.
* `ocspresptimeout` - Time, in milliseconds, to wait for an OCSP response.
* `ocspurlresolvetimeout` - Time, in milliseconds, to wait for an OCSP URL Resolution.
* `ocsptrustresponder` - If trustResponder is set to YES, signature verification will be skipped for the OCSP response. Possible values: [ YES, NO ]
* `ocspproducedattimeskew` - Time, in seconds, for which the Citrix ADC waits before considering the response as invalid.
* `ocspusenonce` - Enable the OCSP nonce extension, which is designed to prevent replay attacks. Possible values: [ ENABLED, DISABLED ]
* `ocsphttpmethod` - HTTP method used to send ocsp request. Possible values: [ GET, POST ]
* `id` - The id of the sslzerotouchparam. It is a system-generated identifier.
