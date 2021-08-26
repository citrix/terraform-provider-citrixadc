---
subcategory: "SSL"
---

# Resource: sslocspresponder

The sslocspresponder resource is used to create a OCSP responder.


## Example usage

```hcl
resource "citrixadc_sslocspresponder" "tf_sslocspresponder" {
	name = "tf_sslocspresponder"
	url = "http://www.google.com"
	batchingdelay = 6
	batchingdepth = 3
	cache = "DISABLED"
	httpmethod = "POST"
	insertclientcert = "NO"
	ocspurlresolvetimeout = 101
	producedattimeskew = 301
	respondercert = "ns-server-certificate"
	resptimeout = 101
	signingcert = "ns-server-certificate"
	trustresponder = false
	usenonce = "YES"
}
```


## Argument Reference

* `name` - (Required) Name for the OCSP responder. Cannot begin with a hash (#) or space character and must contain only ASCII alphanumeric, underscore (_), hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the responder is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my responder" or 'my responder').
* `url` - (Required) URL of the OCSP responder.
* `cache` - (Optional) Enable caching of responses. Caching of responses received from the OCSP responder enables faster responses to the clients and reduces the load on the OCSP responder. Possible values: [ ENABLED, DISABLED ]
* `cachetimeout` - (Optional) Timeout for caching the OCSP response. After the timeout, the Citrix ADC sends a fresh request to the OCSP responder for the certificate status. If a timeout is not specified, the timeout provided in the OCSP response applies.
* `batchingdepth` - (Optional) Number of client certificates to batch together into one OCSP request. Batching avoids overloading the OCSP responder. A value of 1 signifies that each request is queried independently. For a value greater than 1, specify a timeout (batching delay) to avoid inordinately delaying the processing of a single certificate.
* `batchingdelay` - (Optional) 
* `resptimeout` - (Optional) Time, in milliseconds, to wait for an OCSP response. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server. Includes Batching Delay time.
* `ocspurlresolvetimeout` - (Optional) Time, in milliseconds, to wait for an OCSP URL Resolution. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server.
* `respondercert` - (Optional) .
* `trustresponder` - (Optional) A certificate to use to validate OCSP responses.  Alternatively, if -trustResponder is specified, no verification will be done on the reponse.  If both are omitted, only the response times (producedAt, lastUpdate, nextUpdate) will be verified.
* `producedattimeskew` - (Optional) Time, in seconds, for which the Citrix ADC waits before considering the response as invalid. The response is considered invalid if the Produced At time stamp in the OCSP response exceeds or precedes the current Citrix ADC clock time by the amount of time specified.
* `signingcert` - (Optional) Certificate-key pair that is used to sign OCSP requests. If this parameter is not set, the requests are not signed.
* `usenonce` - (Optional) Enable the OCSP nonce extension, which is designed to prevent replay attacks. Possible values: [ YES, NO ]
* `insertclientcert` - (Optional) Include the complete client certificate in the OCSP request. Possible values: [ YES, NO ]
* `httpmethod` - (Optional) HTTP method used to send ocsp request. POST is the default httpmethod. If request length is > 255, POST wil be used even if GET is set as httpMethod. Possible values: [ GET, POST ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslocspresponder. It has the same value as the `name` attribute.


## Import

A sslocspresponder can be imported using its name, e.g.

```shell
terraform import citrixadc_sslocspresponder.tf_sslocspresponder tf_sslocspresponder
```
