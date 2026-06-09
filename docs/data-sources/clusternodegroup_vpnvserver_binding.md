---
subcategory: "Cluster"
---

# Data Source: clusternodegroup_vpnvserver_binding

The `clusternodegroup_vpnvserver_binding` data source allows you to retrieve information about a VPN vserver bound to a cluster node group on the Citrix ADC, confirming which VPN vserver is associated with a given node group.


## Example usage

```terraform
data "citrixadc_clusternodegroup_vpnvserver_binding" "tf_clusternodegroup_vpnvserver_binding" {
  name    = "ng_vpn"
  vserver = "vpn_vserver1"
}

output "vserver" {
  value = data.citrixadc_clusternodegroup_vpnvserver_binding.tf_clusternodegroup_vpnvserver_binding.vserver
}
```


## Argument Reference

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.
* `vserver` - (Required) Name of the VPN vserver bound to the nodegroup, used to look up the binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_vpnvserver_binding. It is a comma-separated set of `key:value` pairs in the form `name:<name>,vserver:<vserver>`, with each value URL-encoded.
