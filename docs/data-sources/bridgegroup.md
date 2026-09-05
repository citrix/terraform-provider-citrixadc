---
subcategory: "Network"
---

# Data Source: bridgegroup

The bridgegroup data source allows you to retrieve information about bridge groups.


## Example usage

```terraform
data "citrixadc_bridgegroup" "tf_bridgegroup" {
  bridgegroup_id = 2
}

output "dynamicrouting" {
  value = data.citrixadc_bridgegroup.tf_bridgegroup.dynamicrouting
}

output "ipv6dynamicrouting" {
  value = data.citrixadc_bridgegroup.tf_bridgegroup.ipv6dynamicrouting
}
```


## Argument Reference

* `bridgegroup_id` - (Required) An integer that uniquely identifies the bridge group.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `dynamicrouting` - Enable dynamic routing for this bridgegroup.
* `ipv6dynamicrouting` - Enable all IPv6 dynamic routing protocols on all VLANs bound to this bridgegroup. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.
* `id` - The id of the bridgegroup. It has the same value as the `bridgegroup_id` attribute.

### Read-only bridgegroup metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_bridgegroup` resource) and are Computed-only. Any attribute the appliance does not return is `null`.

* `flags` - Temporary flag used for internal purpose.
* `portbitmap` - Member interfaces of this bridge group.
* `tagbitmap` - Tagged members of this bridge group.
* `ifaces` - Names of all member interfaces of this bridge group.
* `tagifaces` - Names of all tagged member interfaces of this bridge group.
* `rnat` - Temporary flag used for internal purpose.
* `partitionname` - Name of the Partition to which this vlan bound to.
