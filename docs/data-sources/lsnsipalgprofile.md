---
subcategory: "LSN"
---

# Data Source `lsnsipalgprofile`

The lsnsipalgprofile data source allows you to retrieve information about LSN (Large Scale NAT) SIP Application Layer Gateway profiles.


## Example usage

```terraform
data "citrixadc_lsnsipalgprofile" "tf_lsnsipalgprofile_ds" {
  sipalgprofilename = "my_lsn_sipalgprofile_ds"
}

output "datasessionidletimeout" {
  value = data.citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile_ds.datasessionidletimeout
}

output "sipsessiontimeout" {
  value = data.citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile_ds.sipsessiontimeout
}

output "siptransportprotocol" {
  value = data.citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile_ds.siptransportprotocol
}
```


## Argument Reference

* `sipalgprofilename` - (Required) The name of the SIPALG Profile.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `datasessionidletimeout` - Idle timeout for the data channel sessions in seconds.
* `opencontactpinhole` - ENABLE/DISABLE ContactPinhole creation.
* `openrecordroutepinhole` - ENABLE/DISABLE RecordRoutePinhole creation.
* `openregisterpinhole` - ENABLE/DISABLE RegisterPinhole creation.
* `openroutepinhole` - ENABLE/DISABLE RoutePinhole creation.
* `openviapinhole` - ENABLE/DISABLE ViaPinhole creation.
* `registrationtimeout` - SIP registration timeout in seconds.
* `rport` - ENABLE/DISABLE rport.
* `sipdstportrange` - Destination port range for SIP_UDP and SIP_TCP.
* `sipsessiontimeout` - SIP control channel session timeout in seconds.
* `sipsrcportrange` - Source port range for SIP_UDP and SIP_TCP.
* `siptransportprotocol` - SIP ALG Profile transport protocol type.

## Attribute Reference

* `id` - The id of the lsnsipalgprofile. It has the same value as the `sipalgprofilename` attribute.


## Import

A lsnsipalgprofile can be imported using its sipalgprofilename, e.g.

```shell
terraform import citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile_ds my_lsn_sipalgprofile_ds
```
