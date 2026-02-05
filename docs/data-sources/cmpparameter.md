---
subcategory: "CMP"
---

# Data Source `cmpparameter`

The cmpparameter data source allows you to retrieve information about compression parameters configuration.


## Example usage

```terraform
data "citrixadc_cmpparameter" "tf_cmpparameter" {
}

output "cmplevel" {
  value = data.citrixadc_cmpparameter.tf_cmpparameter.cmplevel
}

output "quantumsize" {
  value = data.citrixadc_cmpparameter.tf_cmpparameter.quantumsize
}

output "servercmp" {
  value = data.citrixadc_cmpparameter.tf_cmpparameter.servercmp
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `addvaryheader` - Control insertion of the Vary header in HTTP responses compressed by Citrix ADC. Intermediate caches store different versions of the response for different values of the headers present in the Vary response header.
* `cmpbypasspct` - Citrix ADC CPU threshold after which compression is not performed. Range: 0 - 100.
* `cmplevel` - Specify a compression level. Available settings function as follows:
  * Optimal - Corresponds to a gzip GZIP level of 5-7.
  * Best speed - Corresponds to a gzip level of 1.
  * Best compression - Corresponds to a gzip level of 9.
* `cmponpush` - Citrix ADC does not wait for the quantum to be filled before starting to compress data. Upon receipt of a packet with a PUSH flag, the appliance immediately begins compression of the accumulated packets.
* `externalcache` - Enable insertion of Cache-Control: private response directive to indicate response message is intended for a single user and must not be cached by a shared or proxy cache.
* `heurexpiry` - Heuristic basefile expiry.
* `heurexpiryhistwt` - For heuristic basefile expiry, weightage to be given to historical delta compression ratio, specified as percentage. For example, to give 25% weightage to historical ratio (and therefore 75% weightage to the ratio for current delta compression transaction), specify 25.
* `heurexpirythres` - Threshold compression ratio for heuristic basefile expiry, multiplied by 100. For example, to set the threshold ratio to 1.25, specify 125.
* `minressize` - Smallest response size, in bytes, to be compressed.
* `policytype` - Type of the policy. The only possible value is ADVANCED.
* `quantumsize` - Minimum quantum of data to be filled before compression begins.
* `randomgzipfilename` - Control the addition of a random filename of random length in the GZIP header to apply the Heal-the-BREACH mitigation for the BREACH attack.
* `randomgzipfilenamemaxlength` - Maximum length of the random filename to be added in the GZIP header to apply the Heal-the-BREACH mitigation for the BREACH attack.
* `randomgzipfilenameminlength` - Minimum length of the random filename to be added in the GZIP header to apply the Heal-the-BREACH mitigation for the BREACH attack.
* `servercmp` - Allow the server to send compressed data to the Citrix ADC. With the default setting, the Citrix ADC appliance handles all compression.
* `varyheadervalue` - The value of the HTTP Vary header for compressed responses. If this argument is not specified, a default value of "Accept-Encoding" will be used.
* `id` - The id of the cmpparameter. It is a system-generated identifier.
