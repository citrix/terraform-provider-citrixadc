---
subcategory: "NS"
---

# Data Source `nshttpparam`

The nshttpparam data source allows you to retrieve information about the HTTP parameters configured on the Citrix ADC appliance.


## Example usage

```terraform
data "citrixadc_nshttpparam" "example" {
}

output "http_connection_multiplex" {
  value = data.citrixadc_nshttpparam.example.conmultiplex
}
```


## Argument Reference

This datasource does not require any arguments. All attributes are optional and computed.

## Attribute Reference

The following attributes are available:

* `id` - The id of the nshttpparam datasource.
* `conmultiplex` - Reuse server connections for requests from more than one client connections. Possible values: `ENABLED`, `DISABLED`. Default: `ENABLED`.
* `dropinvalreqs` - Drop invalid HTTP requests or responses. Possible values: `ON`, `OFF`.
* `http2serverside` - Enable/Disable HTTP/2 on server side. Possible values: `ON`, `OFF`.
* `ignoreconnectcodingscheme` - Ignore Coding scheme in CONNECT request. Possible values: `ENABLED`, `DISABLED`. Default: `DISABLED`.
* `insnssrvrhdr` - Enable or disable Citrix ADC server header insertion for Citrix ADC generated HTTP responses. Possible values: `ON`, `OFF`.
* `logerrresp` - Server header value to be inserted. Possible values: `True`, `False`. Default: `True`.
* `markconnreqinval` - Mark CONNECT requests as invalid. Possible values: `ON`, `OFF`.
* `markhttp09inval` - Mark HTTP/0.9 requests as invalid. Possible values: `ON`, `OFF`.
* `maxreusepool` - Maximum limit on the number of connections, from the Citrix ADC to a particular server that are kept in the reuse pool. This setting is helpful for optimal memory utilization and for reducing the idle connections to the server just after the peak time. Minimum value: `0`, Maximum value: `360000`.
* `nssrvrhdr` - The server header value to be inserted. If no explicit header is specified then NSBUILD.RELEASE is used as default server header.
