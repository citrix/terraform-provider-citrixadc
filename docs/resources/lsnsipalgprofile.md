---
subcategory: "Lsn"
---

# Resource: lsnsipalgprofile

The lsnsipalgprofile resource is used to create lsnsipalgprofile.


## Example usage

```hcl
resource "citrixadc_lsnsipalgprofile" "tf_lsnsipalgprofile" {
  sipalgprofilename      = "my_lsn_sipalgprofile"
  datasessionidletimeout = 150
  sipsessiontimeout      = 150
  registrationtimeout    = 150
  sipsrcportrange        = "4200"
  siptransportprotocol   = "TCP"
}
```


## Argument Reference

* `sipalgprofilename` - (Required) The name of the SIPALG Profile.
* `siptransportprotocol` - (Required) SIP ALG Profile transport protocol type.
* `datasessionidletimeout` - (Optional) Idle timeout for the data channel sessions in seconds.
* `opencontactpinhole` - (Optional) ENABLE/DISABLE ContactPinhole creation.
* `openrecordroutepinhole` - (Optional) ENABLE/DISABLE RecordRoutePinhole creation.
* `openregisterpinhole` - (Optional) ENABLE/DISABLE RegisterPinhole creation.
* `openroutepinhole` - (Optional) ENABLE/DISABLE RoutePinhole creation.
* `openviapinhole` - (Optional) ENABLE/DISABLE ViaPinhole creation.
* `registrationtimeout` - (Optional) SIP registration timeout in seconds.
* `rport` - (Optional) ENABLE/DISABLE rport.
* `sipdstportrange` - (Optional) Destination port range for SIP_UDP and SIP_TCP.
* `sipsessiontimeout` - (Optional) SIP control channel session timeout in seconds.
* `sipsrcportrange` - (Optional) Source port range for SIP_UDP and SIP_TCP.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnsipalgprofile. It has the same value as the `name` attribute.


## Import

A lsnsipalgprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile my_lsn_sipalgprofile
```
