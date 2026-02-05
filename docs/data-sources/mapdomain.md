---
subcategory: "Network"
---

# Data Source `mapdomain`

The mapdomain data source allows you to retrieve information about MAP Domain configurations.


## Example usage

```terraform
data "citrixadc_mapdomain" "tf_mapdomain" {
  name = "domain1"
}

output "name" {
  value = data.citrixadc_mapdomain.tf_mapdomain.name
}

output "mapdmrname" {
  value = data.citrixadc_mapdomain.tf_mapdomain.mapdmrname
}
```


## Argument Reference

* `name` - (Required) Name for the MAP Domain. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the MAP Domain is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `mapdmrname` - Default Mapping rule name.
* `id` - The id of the mapdomain. It has the same value as the `name` attribute.


## Import

A mapdomain can be imported using its `name`, e.g.

```
terraform import citrixadc_mapdomain.tf_mapdomain domain1
```
