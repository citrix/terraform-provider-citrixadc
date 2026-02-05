---
subcategory: "LSN"
---

# Data Source: lsnstatic

The lsnstatic data source allows you to retrieve information about a Large Scale NAT (LSN) static mapping entry.

## Example usage

```terraform
data "citrixadc_lsnstatic" "tf_lsnstatic" {
  name    = "my_lsn_static"
}

output "subscrip" {
  value = data.citrixadc_lsnstatic.tf_lsnstatic.subscrip
}

output "subscrport" {
  value = data.citrixadc_lsnstatic.tf_lsnstatic.subscrport
}

output "transportprotocol" {
  value = data.citrixadc_lsnstatic.tf_lsnstatic.transportprotocol
}
```

## Argument Reference

* `name` - (Required) Name for the LSN static mapping entry. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn static1" or 'lsn static1').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnstatic. It has the same value as the `name` attribute.
* `destip` - Destination IP address for the LSN mapping entry.
* `dsttd` - ID of the traffic domain through which the destination IP address for this LSN mapping entry is reachable from the Citrix ADC. If you do not specify an ID, the destination IP address is assumed to be reachable through the default traffic domain, which has an ID of 0.
* `natip` - IPv4 address, already existing on the Citrix ADC as type LSN, to be used as NAT IP address for this mapping entry.
* `natport` - NAT port for this LSN mapping entry. * represents all ports being used. Used in case of static wildcard
* `network6` - B4 address in DS-Lite setup
* `subscrip` - IPv4(NAT44 & DS-Lite)/IPv6(NAT64) address of an LSN subscriber for the LSN static mapping entry.
* `subscrport` - Port of the LSN subscriber for the LSN mapping entry. * represents all ports being used. Used in case of static wildcard
* `td` - ID of the traffic domain to which the subscriber belongs. If you do not specify an ID, the subscriber is assumed to be a part of the default traffic domain.
* `transportprotocol` - Protocol for the LSN mapping entry.

## Import

A lsnstatic can be imported using its name, e.g.

```shell
terraform import citrixadc_lsnstatic.tf_lsnstatic my_lsn_static
```
