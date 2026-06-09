---
subcategory: "Cluster"
---

# Resource: clusternodegroup_gslbsite_binding

Associates a GSLB site with a cluster node group on the Citrix ADC. A cluster node group defines a logical subset of cluster nodes that own a set of resources; binding a GSLB site to the node group restricts ownership and processing of that site's GSLB traffic to the nodes in the group. Use this resource to pin a GSLB site to specific cluster nodes for spotted/striped placement and predictable traffic distribution.


## Example usage

```hcl
resource "citrixadc_clusternodegroup" "example" {
  name   = "ng1"
  strict = "YES"
  sticky = "NO"
}

resource "citrixadc_gslbsite" "example" {
  sitename      = "site1"
  siteipaddress = "10.0.0.1"
  sitetype      = "LOCAL"
}

resource "citrixadc_clusternodegroup_gslbsite_binding" "example" {
  name     = citrixadc_clusternodegroup.example.name
  gslbsite = citrixadc_gslbsite.example.sitename
}
```


## Argument Reference

* `name` - (Required) Name of the cluster node group. The name uniquely identifies the node group on the cluster. Changing this value forces a new resource to be created.
* `gslbsite` - (Required) Name of the GSLB site to bind to this node group. Maximum length = 31. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_gslbsite_binding. It is a composite key of the form `gslbsite:<gslbsite>,name:<name>`, with each value URL-encoded.


## Import

A clusternodegroup_gslbsite_binding can be imported using its composite id (the bound `gslbsite` and the node group `name`), e.g.

```shell
terraform import citrixadc_clusternodegroup_gslbsite_binding.example "name:ng1,gslbsite:site1"
```
