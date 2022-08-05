---
subcategory: "<fillme>"
---

# Resource: clusternodegroup_gslbsite_binding

The clusternodegroup_gslbsite_binding resource is used to create clusternodegroup_gslbsite_binding.


## Example usage

```hcl
resource "citrixadc_clusternodegroup_gslbsite_binding" "tf_clusternodegroup_gslbsite_binding" {
  gslbsite = citrixadc_gslbsite.site_remote.sitename
  name     = "my_group"
}

resource "citrixadc_gslbsite" "site_remote" {
  sitename        = "my_local_site"
  siteipaddress   = "10.222.74.169"
  sessionexchange = "DISABLED"
  sitetype        = "LOCAL"
}
```


## Argument Reference

* `gslbsite` - (Optional) vserver that need to be bound to this nodegroup.
* `name` - (Optional) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_gslbsite_binding. It is the concatenation of the `name` and `gslbsite` attributes.


## Import

A clusternodegroup_gslbsite_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_clusternodegroup_gslbsite_binding.tf_clusternodegroup_gslbsite_binding my_group,my_local_site
```
