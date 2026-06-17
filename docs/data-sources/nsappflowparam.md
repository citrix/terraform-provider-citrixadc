---
subcategory: "NS"
---

# Data Source: nsappflowparam

The nsappflowparam data source allows you to retrieve the global AppFlow parameters configured on the Citrix ADC appliance, such as the IPFIX template refresh interval, the UDP packet MTU, and which HTTP fields are included in the exported flow records.


## Example usage

```terraform
data "citrixadc_nsappflowparam" "example" {
}

output "appflow_template_refresh" {
  value = data.citrixadc_nsappflowparam.example.templaterefresh
}
```


## Argument Reference

This datasource is a singleton and does not require any arguments. All attributes are computed.

## Attribute Reference

The following attributes are available:

* `id` - The id of the nsappflowparam datasource. Set to the constant string `nsappflowparam-config`.
* `templaterefresh` - IPFIX template refresh interval, in seconds.
* `udppmtu` - MTU, in bytes, used for IPFIX UDP packets.
* `httpurl` - Whether the HTTP URL is included in the AppFlow records. Possible values: `ON`, `OFF`.
* `httpcookie` - Whether the HTTP cookie is included in the AppFlow records. Possible values: `ON`, `OFF`.
* `httpreferer` - Whether the HTTP referer is included in the AppFlow records. Possible values: `ON`, `OFF`.
* `httpmethod` - Whether the HTTP method is included in the AppFlow records. Possible values: `ON`, `OFF`.
* `httphost` - Whether the HTTP host is included in the AppFlow records. Possible values: `ON`, `OFF`.
* `httpuseragent` - Whether the HTTP user-agent is included in the AppFlow records. Possible values: `ON`, `OFF`.
* `clienttrafficonly` - Whether AppFlow records are generated only for client-side traffic. Possible values: `YES`, `NO`.
