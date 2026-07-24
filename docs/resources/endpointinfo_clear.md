---
subcategory: "Endpoint"
---

# Resource: endpointinfo_clear

This resource is used to clear cached endpoint information on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_endpointinfo_clear" "tf_endpointinfo_clear" {
  endpointkind = "IP"
}
```


## Argument Reference

* `endpointkind` - (Optional) Endpoint kind whose information to clear. Currently, IP endpoints are supported. Possible values: [ IP ]. Changing this value forces the resource to be recreated (re-running the clear action against the new endpoint kind).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the endpointinfo_clear resource. It has the format `endpointinfo_clear-<endpointkind>` (for example, `endpointinfo_clear-IP`).
