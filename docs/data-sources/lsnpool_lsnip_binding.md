---
subcategory: "LSN"
---

# Data Source: lsnpool_lsnip_binding

The lsnpool_lsnip_binding data source allows you to retrieve information about an LSN IP address (or range) bound to a Large Scale NAT (LSN) pool.


## Example usage

```terraform
data "citrixadc_lsnpool_lsnip_binding" "tf_lsnpool_lsnip_binding" {
  poolname = "lsnpool1"
  lsnip    = "10.20.30.40-10.20.30.50"
}

output "bound_lsnip" {
  value = data.citrixadc_lsnpool_lsnip_binding.tf_lsnpool_lsnip_binding.lsnip
}
```


## Argument Reference

* `poolname` - (Required) Name for the LSN pool whose binding you want to look up.
* `lsnip` - (Required) IPv4 address or range of IPv4 addresses bound to the LSN pool as NAT IP address(es).
* `ownernode` - (Optional) ID(s) of cluster node(s) on which the command is to be executed.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnpool_lsnip_binding. It is the concatenation of the `poolname` and `lsnip` attributes separated by a comma.
* `lsnip` - IPv4 address or range of IPv4 addresses used as NAT IP address(es) for LSN.
* `ownernode` - ID(s) of cluster node(s) on which the command is to be executed.
* `poolname` - Name of the LSN pool.
