---
subcategory: "SNMP"
---

# Data Source: citrixadc_snmpmanager

The `citrixadc_snmpmanager` data source is used to retrieve information about an SNMP manager configured on the Citrix ADC.

## Example usage

```hcl
data "citrixadc_snmpmanager" "example" {
  ipaddress = "192.168.1.100"
  netmask   = "255.255.255.255"
}
```

## Argument Reference

* `ipaddress` - (Required) IP address of the SNMP manager. Can be an IPv4 or IPv6 address.
* `netmask` - (Required) Subnet mask associated with an IPv4 network address. If the IP address specifies the address or host name of a specific host, accept the default value of 255.255.255.255.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the SNMP manager.
* `domainresolveretry` - Amount of time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the SNMP manager if the last query failed. This parameter is valid for host-name based SNMP managers only. After a query succeeds, the TTL determines the wait time. The minimum and default value is 5.
