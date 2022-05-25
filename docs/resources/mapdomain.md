---
subcategory: "Network"
---

# Resource: mapdomain

The mapdomain resource is used to create MAP-T Map Domain resource.


## Example usage

```hcl
resource "citrixadc_mapdmr" "tf_mapdmr" {
  name         = "tf_mapdmr"
  bripv6prefix = "2002:db8::/64"
}
resource "citrixadc_mapdomain" "tf_mapdomain" {
  name       = "tf_mapdomain"
  mapdmrname = citrixadc_mapdmr.tf_mapdmr.name
}
```


## Argument Reference

* `name` - (Required) Name for the MAP Domain. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the  MAP Domain is created . The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "add network MapDomain map1"). Minimum length =  1 Maximum length =  127
* `mapdmrname` - (Required) Default Mapping rule name.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the mapdomain. It has the same value as the `name` attribute.


## Import

A mapdomain can be imported using its name, e.g.

```shell
terraform import citrixadc_mapdomain.tf_mapdomain tf_mapdomain
```
