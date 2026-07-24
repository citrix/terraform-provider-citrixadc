---
subcategory: "LSN"
---

# Resource: lsnpool_lsnip_binding

This resource is used to manage the binding of NAT IP addresses to an LSN pool.


## Example usage

```hcl
resource "citrixadc_lsnpool" "tf_lsnpool" {
  poolname = "lsnpool1"
}

resource "citrixadc_lsnpool_lsnip_binding" "tf_lsnpool_lsnip_binding" {
  poolname = citrixadc_lsnpool.tf_lsnpool.poolname
  lsnip    = "10.20.30.40-10.20.30.50"
}

```


## Argument Reference

* `poolname` - (Required) Name for the LSN pool. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN pool is created. Changing this attribute forces a new resource to be created.
* `lsnip` - (Required) IPv4 address or a range of IPv4 addresses to be used as NAT IP address(es) for LSN. After the pool is created, these IPv4 addresses are added to the Citrix ADC as Citrix ADC owned IP addresses of type LSN. A maximum of 4096 IP addresses can be bound to an LSN pool. An LSN IP address associated with an LSN pool cannot be shared with other LSN pools, and the addresses must not already exist on the Citrix ADC as any Citrix ADC owned IP addresses. Specify a range with a hyphen, for example `10.102.29.30-10.102.29.189`. By default, ARP is enabled on the LSN IP address; you can disable it using the `set ns ip` command. Changing this attribute forces a new resource to be created.
* `ownernode` - (Optional) ID(s) of cluster node(s) on which the command is to be executed. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnpool_lsnip_binding. It is the concatenation of the `poolname` and `lsnip` attributes separated by a comma.


## Import

A lsnpool_lsnip_binding can be imported using its id (the composite key `poolname,lsnip`), e.g.

```shell
terraform import citrixadc_lsnpool_lsnip_binding.tf_lsnpool_lsnip_binding lsnpool1,10.20.30.40-10.20.30.50
```
