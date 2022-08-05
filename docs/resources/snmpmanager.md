---
subcategory: "SNMP"
---

# Resource: snmpmanager

The snmpmanager resource is used to create snmpmanager.


## Example usage

```hcl
resource "citrixadc_snmpmanager" "tf_snmpmanager" {
  ipaddress          = "192.168.2.2"
  netmask            = "255.255.255.255"
}

```


## Argument Reference

* `ipaddress` - (Required) IP address of the SNMP manager. Can be an IPv4 or IPv6 address. You can instead specify an IPv4 network address or IPv6 network prefix if you want the Citrix ADC to respond to SNMP queries from any device on the specified network. Alternatively, instead of an IPv4 address, you can specify a host name that has been assigned to an SNMP manager. If you do so, you must add a DNS name server that resolves the host name of the SNMP manager to its IP address.  Note: The Citrix ADC does not support host names for SNMP managers that have IPv6 addresses.
* `domainresolveretry` - (Optional) Amount of time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the SNMP manager if the last query failed. This parameter is valid for host-name based SNMP managers only. After a query succeeds, the TTL determines the wait time. The minimum and default value is 5.
* `netmask` - (Optional) Subnet mask associated with an IPv4 network address. If the IP address specifies the address or host name of a specific host, accept the default value of 255.255.255.255.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the snmpmanager. It has the same value as the `ipaddress` attribute.


## Import

A snmpmanager can be imported using its name, e.g.

```shell
terraform import citrixadc_snmpmanager.tf_snmpmanager 192.168.2.2
```
