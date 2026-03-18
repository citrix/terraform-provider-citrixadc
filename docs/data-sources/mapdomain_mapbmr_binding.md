---
subcategory: "Network"
---

# Data Source: mapdomain_mapbmr_binding

The mapdomain_mapbmr_binding data source allows you to retrieve information about a binding between a MAP domain and a Basic Mapping Rule (BMR).


## Example Usage

```terraform
data "citrixadc_mapdomain_mapbmr_binding" "tf_binding" {
  name       = "tf_mapdomain"
  mapbmrname = "tf_mapbmr"
}

output "binding_id" {
  value = data.citrixadc_mapdomain_mapbmr_binding.tf_binding.id
}
```


## Argument Reference

* `name` - (Required) Name for the MAP Domain. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the  MAP Domain is created . The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "add network MapDomain map1").
* `mapbmrname` - (Required) Basic Mapping rule name.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the mapdomain_mapbmr_binding. It is the concatenation of `name` and `mapbmrname` attributes separated by comma.
