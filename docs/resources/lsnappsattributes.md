---
subcategory: "LSN"
---

# Resource: lsnappsattributes

The lsnappsattributes resource is used to create lsnappsattributes.


## Example usage

```hcl
resource "citrixadc_lsnappsattributes" "tf_lsnappsattributes" {
  name              = "my_lsn_appattributes"
  transportprotocol = "TCP"
  port              = 90
  sessiontimeout    = 40
}
```


## Argument Reference

* `name` - (Required) Name for the LSN Application Port ATTRIBUTES. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN application profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn application profile1" or 'lsn application profile1').
* `transportprotocol` - (Required) Name of the protocol(TCP,UDP) for which the parameters of this LSN application port ATTRIBUTES applies
* `port` - (Optional) This is used for Displaying Port/Port range in CLI/Nitro.Lowport, Highport values are populated and used for displaying.Port numbers or range of port numbers to match against the destination port of the incoming packet from a subscriber. When the destination port is matched, the LSN application profile is applied for the LSN session. Separate a range of ports with a hyphen. For example, 40-90.
* `sessiontimeout` - (Optional) Timeout, in seconds, for an idle LSN session. If an LSN session is idle for a time that exceeds this value, the Citrix ADC removes the session.This timeout does not apply for a TCP LSN session when a FIN or RST message is received from either of the endpoints.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnappsattributes. It has the same value as the `name` attribute.


## Import

A lsnappsattributes can be imported using its name, e.g.

```shell
terraform import citrixadc_lsnappsattributes.tf_lsnappsattributes my_lsn_appattributes
```
