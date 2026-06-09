---
subcategory: "Cluster"
---

# Resource: clusternodegroup_vpnvserver_binding

Binds a VPN virtual server to a cluster node group on the Citrix ADC. A node group is a subset of the cluster nodes; binding a VPN vserver to it constrains that vserver (and its traffic distribution) to the nodes in the group rather than across the entire cluster, which is how you achieve spotted/striped placement of VPN workloads within a cluster.


## Example usage

```hcl
resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
  name = "ng_vpn"
}

resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "vpn_vserver1"
  servicetype = "SSL"
  ipv46       = "10.10.10.10"
  port        = 443
}

resource "citrixadc_clusternodegroup_vpnvserver_binding" "tf_clusternodegroup_vpnvserver_binding" {
  name    = citrixadc_clusternodegroup.tf_clusternodegroup.name
  vserver = citrixadc_vpnvserver.tf_vpnvserver.name
}
```


## Argument Reference

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster. Changing this forces a new resource to be created.
* `vserver` - (Required) Name of the VPN vserver that needs to be bound to this nodegroup. Changing this forces a new resource to be created.

~> **Note** This binding has no NITRO update endpoint and every attribute forces replacement. Any change to `name` or `vserver` recreates the binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_vpnvserver_binding. It is a comma-separated set of `key:value` pairs in the form `name:<name>,vserver:<vserver>`, with each value URL-encoded.


## Import

A clusternodegroup_vpnvserver_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_clusternodegroup_vpnvserver_binding.tf_clusternodegroup_vpnvserver_binding "name:ng_vpn,vserver:vpn_vserver1"
```
