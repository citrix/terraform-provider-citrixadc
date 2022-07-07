---
subcategory: "DNS"
---

# Resource: dnssuffix

The dnssuffix resource is used to create DNS suffix.


## Example usage

```hcl
resource "citrixadc_dnssuffix" "tf_dnssuffix" {
		dnssuffix = "example.com"
	}
```


## Argument Reference

* `Dnssuffix` - (Required) Suffix to be appended when resolving domain names that are not fully qualified.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnssuffix. It has the same value as the `dnssuffix` attribute.


## Import

A dnssuffix can be imported using its name, e.g.

```shell
terraform import citrixadc_dnssuffix.tf_dnssuffix tf_dnssuffix
```
