---
subcategory: "Network"
---

# Data Source: vlan\_linkset\_binding

The vlan\_linkset\_binding data source allows you to retrieve information about an
interface bound to a VLAN. Despite the resource name, this binding references an
interface (`ifnum`) bound to a VLAN identified by `vlanid`, not a linkset object.

## Example usage

```terraform
data "citrixadc_vlan_linkset_binding" "tf_bind" {
  vlanid = 40
  ifnum  = "1/3"
}

output "vlanid" {
  value = data.citrixadc_vlan_linkset_binding.tf_bind.vlanid
}

output "ifnum" {
  value = data.citrixadc_vlan_linkset_binding.tf_bind.ifnum
}

output "tagged" {
  value = data.citrixadc_vlan_linkset_binding.tf_bind.tagged
}
```

## Argument Reference

* `vlanid` - (Required) Specifies the virtual LAN ID.
* `ifnum` - (Required) The interface to be bound to the VLAN, specified in slot/port notation (for example, `1/3`).

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vlan\_linkset\_binding. It is a composite key of the form `vlanid:<vlanid>,ifnum:<ifnum>`, where the `ifnum` value is URL-encoded.
* `ownergroup` - The owner node group in a Cluster for this VLAN.
* `tagged` - Whether the interface is an 802.1q tagged interface.
