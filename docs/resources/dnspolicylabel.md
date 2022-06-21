---
subcategory: "DNS"
---

# Resource: dnspolicylabel

The dnspolicylabel resource is used to create DNS policylabel.


## Example usage

```hcl
resource "citrixadc_dnspolicylabel" "dnspolicylabel" {
	labelname = "blue_label"
	transform = "dns_req"
  }
```


## Argument Reference

* `labelname` - (Required) Name of the dns policy label.
* `transform` - (Required) The type of transformations allowed by the policies bound to the label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnspolicylabel. It has the same value as the `labelname` attribute.


## Import

A dnspolicylabel can be imported using its name, e.g.

```shell
terraform import citrixadc_dnspolicylabel.dnspolicylabel label1
```
