---
subcategory: "Compression"
---

# Resource: cmpaction

The cmpaction resource is used to create cmpaction.


## Example usage

```hcl
resource "citrixadc_cmpaction" "tf_cmpaction" {
  name    = "my_cmpaction"
  cmptype = "NOCOMPRESS"
}

```


## Argument Reference

* `name` - (Required) Name of the compression action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the action is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cmp action" or 'my cmp action'). Minimum length =  1
* `cmptype` - (Required) Type of compression performed by this action. Available settings function as follows: * COMPRESS - Apply GZIP or DEFLATE compression to the response, depending on the request header. Prefer GZIP. * GZIP - Apply GZIP compression. * DEFLATE - Apply DEFLATE compression. * NOCOMPRESS - Do not compress the response if the request matches a policy that uses this action. Possible values: [ compress, gzip, deflate, nocompress ]
* `addvaryheader` - (Optional) Control insertion of the Vary header in HTTP responses compressed by Citrix ADC. Intermediate caches store different versions of the response for different values of the headers present in the Vary response header. Possible values: [ GLOBAL, DISABLED, ENABLED ]
* `varyheadervalue` - (Optional) The value of the HTTP Vary header for compressed responses. Minimum length =  1
* `deltatype` - (Optional) The type of delta action (if delta type compression action is defined). Possible values: [ PERURL, PERPOLICY ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cmpaction. It has the same value as the `name` attribute.


## Import

A cmpaction can be imported using its name, e.g.

```shell
terraform import citrixadc_cmpaction.tf_cmpaction my_cmpaction
```
