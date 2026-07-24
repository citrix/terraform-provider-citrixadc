---
subcategory: "Network"
---

# Resource: rnatsession_flush

This resource is used to flush active RNAT sessions on the Citrix ADC.


## Example usage

### Flush all RNAT sessions

With no filter arguments, all RNAT sessions are flushed.

```hcl
resource "citrixadc_rnatsession_flush" "flush_all" {}
```

### Flush RNAT sessions for a subnet

Use `network` and `netmask` together to flush sessions whose traffic matches a subnet.

```hcl
resource "citrixadc_rnatsession_flush" "flush_subnet" {
  network = "10.10.10.0"
  netmask = "255.255.255.0"
}
```

### Flush RNAT sessions by NAT IP

```hcl
resource "citrixadc_rnatsession_flush" "flush_natip" {
  natip = "192.0.2.100"
}
```

### Flush RNAT sessions by ACL

```hcl
resource "citrixadc_rnatsession_flush" "flush_acl" {
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

* `id` - The id of the rnatsession_flush resource. It is set to `rnatsession_flush`.
