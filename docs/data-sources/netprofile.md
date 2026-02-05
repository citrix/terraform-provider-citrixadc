---
subcategory: "Network"
---

# Data Source `netprofile`

The netprofile data source allows you to retrieve information about a network profile configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_netprofile" "tf_netprofile" {
  name = "my_netprofile"
}

output "proxyprotocol" {
  value = data.citrixadc_netprofile.tf_netprofile.proxyprotocol
}

output "proxyprotocoltxversion" {
  value = data.citrixadc_netprofile.tf_netprofile.proxyprotocoltxversion
}
```


## Argument Reference

* `name` - (Required) Name for the net profile. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created. Choose a name that helps identify the net profile.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `mbf` - Response will be sent using learnt info if enabled. When creating a netprofile, if you do not set this parameter, the netprofile inherits the global MBF setting (available in the enable ns mode and disable ns mode CLI commands, or in the System > Settings > Configure modes > Configure Modes dialog box). However, you can override this setting after you create the netprofile
* `overridelsn` - USNIP/USIP settings override LSN settings for configured service/virtual server traffic.
* `proxyprotocol` - Proxy Protocol Action (Enabled/Disabled)
* `proxyprotocolaftertlshandshake` - ADC doesnt look for proxy header before TLS handshake, if enabled. Proxy protocol parsed after TLS handshake
* `proxyprotocoltxversion` - Proxy Protocol Version (V1/V2)
* `srcip` - IP address or the name of an IP set.
* `srcippersistency` - When the net profile is associated with a virtual server or its bound services, this option enables the Citrix ADC to use the same address, specified in the net profile, to communicate to servers for all sessions initiated from a particular client to the virtual server.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.

## Attribute Reference

* `id` - The id of the netprofile. It has the same value as the `name` attribute.


## Import

A netprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_netprofile.tf_netprofile my_netprofile
```
