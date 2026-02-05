---
subcategory: "LSN"
---

# Data Source `lsnip6profile`

The lsnip6profile data source allows you to retrieve information about LSN (Large Scale NAT) IPv6 profiles.


## Example usage

```terraform
data "citrixadc_lsnip6profile" "tf_lsnaip6profile_ds" {
  name = "my_lsn_ip6profile_ds"
}

output "type" {
  value = data.citrixadc_lsnip6profile.tf_lsnaip6profile_ds.type
}

output "network6" {
  value = data.citrixadc_lsnip6profile.tf_lsnaip6profile_ds.network6
}
```


## Argument Reference

* `name` - (Required) Name for the LSN ip6 profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN ip6 profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn ip6 profile1" or 'lsn ip6 profile1').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `natprefix` - IPv6 address(es) of the LSN subscriber(s) or subscriber network(s) on whose traffic you want the Citrix ADC to perform Large Scale NAT.
* `network6` - IPv6 address of the Citrix ADC AFTR device
* `type` - IPv6 translation type for which to set the LSN IP6 profile parameters.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnip6profile. It has the same value as the `name` attribute.


## Import

A lsnip6profile can be imported using its name, e.g.

```shell
terraform import citrixadc_lsnip6profile.tf_lsnaip6profile_ds my_lsn_ip6profile_ds
```
