---
subcategory: "Network"
---

# Resource: netprofile

The netprofile resource is used to create nework profiles.


## Example usage

```hcl
resource "citrixadc_netprofile" "tf_netprofile" {
    name = "tf_netprofile"
    proxyprotocol = "ENABLED"
    proxyprotocoltxversion = "V1"
}
```


## Argument Reference

* `name` - (Optional) Name for the net profile. Cannot be changed after the profile is created. Choose a name that helps identify the net profile.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `srcip` - (Optional) IP address or the name of an IP set.
* `srcippersistency` - (Optional) When the net profile is associated with a virtual server or its bound services, this option enables the Citrix ADC to use the same  address, specified in the net profile, to communicate to servers for all sessions initiated from a particular client to the virtual server. Possible values: [ ENABLED, DISABLED ]
* `overridelsn` - (Optional) USNIP/USIP settings override LSN settings for configured service/virtual server traffic.. . Possible values: [ ENABLED, DISABLED ]
* `mbf` - (Optional) Response will be sent using learnt info if enabled. When creating a netprofile, if you do not set this parameter, the netprofile inherits the global MBF setting (available in the enable ns mode and disable ns mode CLI commands, or in the System > Settings > Configure modes > Configure Modes dialog box). However, you can override this setting after you create the netprofile. Possible values: [ ENABLED, DISABLED ]
* `proxyprotocol` - (Optional) Proxy Protocol Action (Enabled/Disabled). Possible values: [ ENABLED, DISABLED ]
* `proxyprotocoltxversion` - (Optional) Proxy Protocol Version (V1/V2). Possible values: [ V1, V2 ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the netprofile. It has the same value as the `name` attribute.


## Import

A netprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_netprofile.tf_netprofile tf_netprofile
```
