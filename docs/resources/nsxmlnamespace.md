---
subcategory: "NS"
---

# Resource: nsxmlnamespace

The nsxmlnamespace resource is used to create XML namespace resource.


## Example usage

```hcl
resource "citrixadc_nsxmlnamespace" "tf_nsxmlnamespace" {
  prefix      = "tf_nsxmlnamespace"
  namespace   = "http://www.w3.org/2001/04/xmlenc#"
  description = "Description1"
}
```


## Argument Reference

* `prefix` - (Required) XML prefix. Minimum length =  1
* `Namespace` - (Required) Expanded namespace for which the XML prefix is provided. Minimum length =  1
* `description` - (Optional) Description for the prefix. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsxmlnamespace. It has the same value as the `prefix` attribute.


## Import

A nsxmlnamespace can be imported using its prefix, e.g.

```shell
terraform import citrixadc_nsxmlnamespace.tf_nsxmlnamespace tf_nsxmlnamespace
```
