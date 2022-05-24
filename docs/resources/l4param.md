---
subcategory: "Network"
---

# Resource: l4param

The l4param resource is used to create Layer 4 related parameter resource.


## Example usage

```hcl
resource "citrixadc_l4param" "tf_l4param" {
  l2connmethod = "MacVlanChannel"
  l4switch     = "DISABLED"
}
```


## Argument Reference

* `l2connmethod` - (Optional) Layer 2 connection method based on the combination of  channel number, MAC address and VLAN. It is tuned with l2conn param of lb vserver. If l2conn of lb vserver is ON then method specified here will be used to identify a connection in addition to the 4-tuple (<source IP>:<source port>::<destination IP>:<destination port>). Possible values: [ Channel, Vlan, VlanChannel, Mac, MacChannel, MacVlan, MacVlanChannel ]
* `l4switch` - (Optional) In L4 switch topology, always clients and servers are on the same side. Enable l4switch to allow such connections. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the l4param. It is a unique string prefixed with "tf-l4param-"

