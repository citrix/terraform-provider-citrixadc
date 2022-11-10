---
subcategory: "Network"
---

# Resource: rnat_clear

The rnat_clear resource is used to create and clear the RNAT rule.

~> 
This resource supports create and clear operation. To use add, update and delete operation that supports latest ADC version(from 13.0 and above) refer [`citrixadc_rnat'](./rnat.md) resource.

## Example usage

```hcl
resource "citrixadc_rnat" "allrnat" {
  rnatsname = "rnatsall"

  rnat {
    network = "192.168.20.0"
    netmask = "255.255.255.0"
  }

  rnat {
    network = "192.168.88.0"
    netmask = "255.255.255.0"
    natip   = "172.17.0.2"
  }

  rnat {
    aclname = "RNAT_ACL_1"
    natip   = "172.17.0.2"
  }
}
```


## Argument Reference

* `rnatsname` - (Optional) the name for the rnat rules.
* `rnat` - (Optional) blocks of rnat rules. Documented below.

A rnat block supports the following:

* `network` - (Optional) The network address defined for the RNAT entry.
* `netmask` - (Optional) The subnet mask for the network address.
* `aclname` - (Optional) An extended ACL defined for the RNAT entry.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `redirectport` - (Optional) Port number to which the IPv4 packets are redirected. Applicable to TCP and UDP protocols.
* `natip` - (Optional) Any NetScaler-owned IPv4 address except the NSIP address. The NetScaler appliance replaces the source IP addresses of server-generated packets with the IP address specified. The IP address must be a public NetScaler-owned IP address. If you specify multiple addresses for this field, NATIP selection uses the round robin algorithm for each session. By specifying a range of IP addresses, you can specify all NetScaler-owned IP addresses, except the NSIP, that fall within the specified range.
* `natip2` - (Optional)


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rnat_clear rules. It has the same value as the `rnatsname` attribute.
