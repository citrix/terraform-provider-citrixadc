---
subcategory: "Network"
---

# Resource: mapdomain_mapbmr_binding

The mapdomain_mapbmr_binding resource is used to bind mapbmr to the mapdomain resource.


## Example usage

```hcl
resource "citrixadc_mapbmr" "tf_mapbmr" {
  name           = "tf_mapbmr"
  ruleipv6prefix = "2001:db8:abcd:12::/64"
  psidoffset     = 6
  eabitlength    = 16
  psidlength     = 8
}
resource "citrixadc_mapdmr" "tf_mapdmr" {
  name         = "tf_mapdmr"
  bripv6prefix = "2002:db8::/64"
}
resource "citrixadc_mapdomain" "tf_mapdomain" {
  name       = "tf_mapdomain"
  mapdmrname = citrixadc_mapdmr.tf_mapdmr.name
}
resource "citrixadc_mapdomain_mapbmr_binding" "tf_binding" {
  name       = citrixadc_mapdomain.tf_mapdomain.name
  mapbmrname = citrixadc_mapbmr.tf_mapbmr.name
}
```


## Argument Reference

* `name` - (Required) Name for the MAP Domain. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the  MAP Domain is created . The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "add network MapDomain map1"). Minimum length =  1 Maximum length =  127
* `mapbmrname` - (Required) Basic Mapping rule name.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the mapdomain_mapbmr_binding. It is the concatenation of `name` and `mapbmrname` attributes separated by comma.


## Import

A mapdomain_mapbmr_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_mapdomain_mapbmr_binding.tf_binding tf_mapdomain,tf_mapbmr
```
