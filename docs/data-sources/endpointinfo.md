---
subcategory: "Endpoint"
---

# Data Source: endpointinfo

The endpointinfo data source allows you to retrieve information about an endpoint registered on the Citrix ADC, looked up by its kind and name.


## Example usage

```terraform
data "citrixadc_endpointinfo" "tf_endpointinfo" {
  endpointkind = "IP"
  endpointname = "192.168.10.25"
}

output "endpoint_metadata" {
  value = data.citrixadc_endpointinfo.tf_endpointinfo.endpointmetadata
}

output "endpoint_labels" {
  value = data.citrixadc_endpointinfo.tf_endpointinfo.endpointlabelsjson
}
```


## Argument Reference

* `endpointkind` - (Required) Endpoint kind. Currently, only IP endpoints are supported.
* `endpointname` - (Required) Name of the endpoint, which depends on the kind. For an IP endpoint this is the IP address.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `endpointmetadata` - String of qualifiers, in dotted notation, providing structured metadata for the endpoint. Each qualifier is more specific than the one that precedes it, as in `cluster.namespace.service` (for example `cluster.default.frontend`). Note: A qualifier that includes a dot (.) or space ( ) must be enclosed in double quotation marks.
* `endpointlabelsjson` - String representing labels for the endpoint in JSON form. Maximum length 16K.
* `id` - The id of the endpointinfo. It is the concatenation of the `endpointkind` and `endpointname` attributes in `endpointkind:<value>,endpointname:<value>` form.
