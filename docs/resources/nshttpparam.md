---
subcategory: "NS"
---

# Resource: nshttpparam

The nshttpparam resource is used to create HTTP parameter resource.


## Example usage

```hcl
resource "citrixadc_nshttpparam" "tf_nshttpparam" {
  dropinvalreqs             = "OFF"
  markconnreqinval          = "OFF"
  maxreusepool              = 1
  markhttp09inval           = "OFF"
  insnssrvrhdr              = "OFF"
  logerrresp                = "OFF"
  conmultiplex              = "ENABLED"
  http2serverside           = "OFF"
  ignoreconnectcodingscheme = "DISABLED"
}
```


## Argument Reference

* `dropinvalreqs` - (Optional) Drop invalid HTTP requests or responses. Possible values: [ on, off ]
* `markhttp09inval` - (Optional) Mark HTTP/0.9 requests as invalid. Possible values: [ on, off ]
* `markconnreqinval` - (Optional) Mark CONNECT requests as invalid. Possible values: [ on, off ]
* `insnssrvrhdr` - (Optional) Enable or disable Citrix ADC server header insertion for Citrix ADC generated HTTP responses. Possible values: [ on, off ]
* `nssrvrhdr` - (Optional) The server header value to be inserted. If no explicit header is specified then NSBUILD.RELEASE is used as default server header. Minimum length =  1
* `logerrresp` - (Optional) Server header value to be inserted. Possible values: [ on, off ]
* `conmultiplex` - (Optional) Reuse server connections for requests from more than one client connections. Possible values: [ ENABLED, DISABLED ]
* `maxreusepool` - (Optional) Maximum limit on the number of connections, from the Citrix ADC to a particular server that are kept in the reuse pool. This setting is helpful for optimal memory utilization and for reducing the idle connections to the server just after the peak time. Minimum value =  0 Maximum value =  360000
* `http2serverside` - (Optional) Enable/Disable HTTP/2 on server side. Possible values: [ on, off ]
* `ignoreconnectcodingscheme` - (Optional) Ignore Coding scheme in CONNECT request. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nshttpparam. It is a unique string prefixed with "tf-nshttpparam-"

