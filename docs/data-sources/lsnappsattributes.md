---
subcategory: "LSN"
---

# Data Source `lsnappsattributes`

The lsnappsattributes data source allows you to retrieve information about LSN Application Port ATTRIBUTES.


## Example usage

```terraform
data "citrixadc_lsnappsattributes" "tf_lsnappsattributes" {
  name = "my_lsn_appattributes"
}

output "transportprotocol" {
  value = data.citrixadc_lsnappsattributes.tf_lsnappsattributes.transportprotocol
}

output "sessiontimeout" {
  value = data.citrixadc_lsnappsattributes.tf_lsnappsattributes.sessiontimeout
}

output "port" {
  value = data.citrixadc_lsnappsattributes.tf_lsnappsattributes.port
}
```


## Argument Reference

* `name` - (Required) Name for the LSN Application Port ATTRIBUTES. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `port` - Port numbers or range of port numbers to match against the destination port of the incoming packet from a subscriber. When the destination port is matched, the LSN application profile is applied for the LSN session. Separate a range of ports with a hyphen. For example, 40-90.
* `sessiontimeout` - Timeout, in seconds, for an idle LSN session. If an LSN session is idle for a time that exceeds this value, the Citrix ADC removes the session. This timeout does not apply for a TCP LSN session when a FIN or RST message is received from either of the endpoints.
* `transportprotocol` - Name of the protocol (TCP, UDP) for which the parameters of this LSN application port ATTRIBUTES applies.

## Attribute Reference

* `id` - The id of the lsnappsattributes. It has the same value as the `name` attribute.


## Import

A lsnappsattributes can be imported using its name, e.g.

```shell
terraform import citrixadc_lsnappsattributes.tf_lsnappsattributes my_lsn_appattributes
```
