---
subcategory: "Cmp"
---

# Data Source `cmpaction`

The cmpaction data source allows you to retrieve information about an existing compression action.


## Example usage

```terraform
data "citrixadc_cmpaction" "tf_cmpaction" {
  name = "my_cmpaction"
}

output "cmptype" {
  value = data.citrixadc_cmpaction.tf_cmpaction.cmptype
}

output "addvaryheader" {
  value = data.citrixadc_cmpaction.tf_cmpaction.addvaryheader
}
```


## Argument Reference

* `name` - (Required) Name of the compression action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the action is added.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cmpaction. It has the same value as the `name` attribute.
* `addvaryheader` - Control insertion of the Vary header in HTTP responses compressed by Citrix ADC. Intermediate caches store different versions of the response for different values of the headers present in the Vary response header.
* `cmptype` - Type of compression performed by this action. Available settings function as follows:
  * COMPRESS - Apply GZIP or DEFLATE compression to the response, depending on the request header. Prefer GZIP.
  * GZIP - Apply GZIP compression.
  * DEFLATE - Apply DEFLATE compression.
  * NOCOMPRESS - Do not compress the response if the request matches a policy that uses this action.
* `deltatype` - The type of delta action (if delta type compression action is defined).
* `newname` - New name for the compression action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `varyheadervalue` - The value of the HTTP Vary header for compressed responses.


## Import

A cmpaction can be imported using its name, e.g.

```shell
terraform import citrixadc_cmpaction.tf_cmpaction my_cmpaction
```
