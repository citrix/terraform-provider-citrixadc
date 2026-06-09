---
subcategory: "Cluster"
---

# Data Source: clusternodegroup_gslbsite_binding

The clusternodegroup_gslbsite_binding data source allows you to retrieve information about a binding between a cluster node group and a GSLB site on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_clusternodegroup_gslbsite_binding" "example" {
  name     = "ng1"
  gslbsite = "site1"
}

output "bound_gslbsite" {
  value = data.citrixadc_clusternodegroup_gslbsite_binding.example.gslbsite
}
```


## Argument Reference

* `name` - (Required) Name of the cluster node group. The name uniquely identifies the node group on the cluster.
* `gslbsite` - (Required) Name of the GSLB site bound to this node group.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_gslbsite_binding. It is a composite key of the form `name:<name>,gslbsite:<gslbsite>`, with each value URL-encoded.
