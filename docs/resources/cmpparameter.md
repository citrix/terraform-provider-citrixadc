---
subcategory: "Compression"
---

# Resource: cmpparameter

The cmpparameter resource is used to update cmpparameter.


## Example usage

```hcl
resource "citrixadc_cmpparameter" "tf_cmpparameter" {
  cmplevel    = "optimal"
  quantumsize = 20
  servercmp   = "OFF"
}
```


## Argument Reference

* `cmplevel` - (Optional) Specify a compression level. Available settings function as follows: * Optimal - Corresponds to a gzip GZIP level of 5-7. * Best speed - Corresponds to a gzip level of 1. * Best compression - Corresponds to a gzip level of 9. Possible values: [ optimal, bestspeed, bestcompression ]
* `quantumsize` - (Optional) Minimum quantum of data to be filled before compression begins. Minimum value =  8 Maximum value =  63488
* `servercmp` - (Optional) Allow the server to send compressed data to the Citrix ADC. With the default setting, the Citrix ADC appliance handles all compression. Possible values: [ on, off ]
* `heurexpiry` - (Optional) Heuristic basefile expiry. Possible values: [ on, off ]
* `heurexpirythres` - (Optional) Threshold compression ratio for heuristic basefile expiry, multiplied by 100. For example, to set the threshold ratio to 1.25, specify 125. Minimum value =  1 Maximum value =  1000
* `heurexpiryhistwt` - (Optional) For heuristic basefile expiry, weightage to be given to historical delta compression ratio, specified as percentage.  For example, to give 25% weightage to historical ratio (and therefore 75% weightage to the ratio for current delta compression transaction), specify 25. Minimum value =  1 Maximum value =  100
* `minressize` - (Optional) Smallest response size, in bytes, to be compressed.
* `cmpbypasspct` - (Optional) Citrix ADC CPU threshold after which compression is not performed. Range: 0 - 100. Minimum value =  0 Maximum value =  100
* `cmponpush` - (Optional) Citrix ADC does not wait for the quantum to be filled before starting to compress data. Upon receipt of a packet with a PUSH flag, the appliance immediately begins compression of the accumulated packets. Possible values: [ ENABLED, DISABLED ]
* `policytype` - (Optional) Type of policy. Available settings function as follows: * Classic -  Classic policies evaluate basic characteristics of traffic and other data. Deprecated. * Advanced -  Advanced policies (which have been renamed as default syntax policies) can perform the same type of evaluations as classic policies. They also enable you to analyze more data (for example, the body of an HTTP request) and to configure more operations in the policy rule (for example, transforming data in the body of a request into an HTTP header). Possible values: [ CLASSIC, ADVANCED ]
* `addvaryheader` - (Optional) Control insertion of the Vary header in HTTP responses compressed by Citrix ADC. Intermediate caches store different versions of the response for different values of the headers present in the Vary response header. Possible values: [ ENABLED, DISABLED ]
* `varyheadervalue` - (Optional) The value of the HTTP Vary header for compressed responses. If this argument is not specified, a default value of "Accept-Encoding" will be used. Minimum length =  1
* `externalcache` - (Optional) Enable insertion of  Cache-Control: private response directive to indicate response message is intended for a single user and must not be cached by a shared or proxy cache. Possible values: [ YES, NO ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cmpparameter. It is a unique string prefixed with `tf-cmpparameter-`.
