---
subcategory: "DNS"
---

# Resource: dnsview

The dnsview resource is used to create DNS view.


## Example usage

```hcl
resource "citrixadc_dnsview" "tf_dnsview" {
		viewname = "view3"
		
	}
```


## Argument Reference

* `viewname` - (Required) Name for the DNS view.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the tf_dnsview. It has the same value as the `viewname` attribute.


## Import

A tf_dnsview can be imported using its name, e.g.

```shell
terraform import citrixadc_dnsview.tf_dnsview view3
```
