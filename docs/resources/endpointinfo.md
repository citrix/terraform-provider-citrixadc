---
subcategory: "Endpoint"
---

# Resource: endpointinfo

The endpointinfo resource registers an endpoint (currently an IP address) with the Citrix ADC and attaches structured metadata and labels to it. These labels and metadata qualifiers let policies and analytics reference the endpoint by descriptive attributes (for example `cluster.default.frontend`) instead of by raw address.


## Example usage

```hcl
resource "citrixadc_endpointinfo" "tf_endpointinfo" {
  endpointkind     = "IP"
  endpointname     = "192.168.10.25"
  endpointmetadata = "cluster.default.frontend"
  endpointlabelsjson = jsonencode({
    app  = "frontend"
    tier = "web"
  })
}
```


## Argument Reference

* `endpointname` - (Required) Name of the endpoint, which depends on the kind. For an IP endpoint this is the IP address.
* `endpointkind` - (Optional) Endpoint kind. Currently, only IP endpoints are supported. Defaults to `"IP"`. Changing this attribute forces a new resource to be created, because it is the key used in the delete URL and cannot be updated in place. Possible values: [ IP ]
* `endpointmetadata` - (Optional) String of qualifiers, in dotted notation, providing structured metadata for the endpoint. Each qualifier is more specific than the one that precedes it, as in `cluster.namespace.service` (for example `cluster.default.frontend`). Note: A qualifier that includes a dot (.) or space ( ) must be enclosed in double quotation marks.
* `endpointlabelsjson` - (Optional) String representing labels for the endpoint in JSON form. Maximum length 16K.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the endpointinfo. It is the concatenation of the `endpointkind` and `endpointname` attributes in `endpointkind:<value>,endpointname:<value>` form.


## Import

An endpointinfo can be imported using its id, which is the concatenation of the `endpointkind` and `endpointname` values, e.g.

```shell
terraform import citrixadc_endpointinfo.tf_endpointinfo endpointkind:IP,endpointname:192.168.10.25
```
