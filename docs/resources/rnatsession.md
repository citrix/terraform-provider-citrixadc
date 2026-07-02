---
subcategory: "Network"
---

# Resource: rnatsession

The rnatsession resource flushes active RNAT (Reverse NAT) sessions on the Citrix ADC. It is an action-only resource: applying it invokes the NITRO flush action, which clears the RNAT session table so that new connections re-establish through the current RNAT configuration. This is useful after changing RNAT rules, for troubleshooting stale mappings, or for forcing sessions off a particular NAT IP or ACL.

This resource does not create, read, or manage a persistent object on the appliance. There is no NITRO GET endpoint for RNAT sessions, so there is no corresponding data source. Each apply performs the flush; changing any filter argument re-triggers the flush.


## Example usage

### Flush all RNAT sessions

With no filter arguments, all RNAT sessions are flushed.

```hcl
resource "citrixadc_rnatsession" "flush_all" {}
```

### Flush RNAT sessions for a subnet

Use `network` and `netmask` together to flush sessions whose traffic matches a subnet.

```hcl
resource "citrixadc_rnatsession" "flush_subnet" {
  network = "10.10.10.0"
  netmask = "255.255.255.0"
}
```

### Flush RNAT sessions by NAT IP

```hcl
resource "citrixadc_rnatsession" "flush_natip" {
  natip = "192.0.2.100"
}
```

### Flush RNAT sessions by ACL

```hcl
resource "citrixadc_rnatsession" "flush_acl" {
  aclname = "rnat_allow_acl"
}
```


## Argument Reference

The filter arguments are mutually exclusive: use either the `network` + `netmask` pair, or `natip`, or `aclname`. If no filter is specified, all RNAT sessions are flushed. Changing any of these arguments forces a new flush action to be performed.

* `network` - (Optional) IPv4 network address on whose traffic you want the Citrix ADC to do RNAT processing. Must be used together with `netmask`. Changing this attribute re-triggers the flush.
* `netmask` - (Optional) Subnet mask associated with the network address. Must be used together with `network`. Changing this attribute re-triggers the flush.
* `natip` - (Optional) The NAT IP address defined for the RNAT entry. Flushes sessions associated with this NAT IP. Changing this attribute re-triggers the flush.
* `aclname` - (Optional) Name of any configured extended ACL whose action is ALLOW. Flushes sessions associated with this ACL. Maximum length = 127. Changing this attribute re-triggers the flush.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `rnatsession-config`. It does not correspond to any object on the Citrix ADC.
