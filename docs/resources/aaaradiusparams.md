---
subcategory: "AAA"
---

# Resource: aaaradiusparams

The aaaradiusparams resource is used to update aaaradiusparams.


## Example usage

```hcl
resource "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
  radkey             = "sslvpn"
  radnasip           = "ENABLED"
  serverip           = "10.222.74.158"
  authtimeout        = 8
}
```


## Argument Reference

* `radkey` - (Required) The key shared between the RADIUS server and clients. Required for allowing the Citrix ADC to communicate with the RADIUS server. Minimum length =  1
* `serverip` - (Optional) IP address of your RADIUS server. Minimum length =  1
* `serverport` - (Optional) Port number on which the RADIUS server listens for connections. Minimum value =  1
* `authtimeout` - (Optional) Maximum number of seconds that the Citrix ADC waits for a response from the RADIUS server. Minimum value =  1
* `radnasip` - (Optional) Send the Citrix ADC IP (NSIP) address to the RADIUS server as the Network Access Server IP (NASIP) part of the Radius protocol. Possible values: [ ENABLED, DISABLED ]
* `radnasid` - (Optional) Send the Network Access Server ID (NASID) for your Citrix ADC to the RADIUS server as the nasid part of the Radius protocol.
* `radvendorid` - (Optional) Vendor ID for RADIUS group extraction. Minimum value =  1
* `radattributetype` - (Optional) Attribute type for RADIUS group extraction. Minimum value =  1
* `radgroupsprefix` - (Optional) Prefix string that precedes group names within a RADIUS attribute for RADIUS group extraction.
* `radgroupseparator` - (Optional) Group separator string that delimits group names within a RADIUS attribute for RADIUS group extraction.
* `passencoding` - (Optional) Enable password encoding in RADIUS packets that the Citrix ADC sends to the RADIUS server. Possible values: [ pap, chap, mschapv1, mschapv2 ]
* `ipvendorid` - (Optional) Vendor ID attribute in the RADIUS response. If the attribute is not vendor-encoded, it is set to 0.
* `ipattributetype` - (Optional) IP attribute type in the RADIUS response. Minimum value =  1
* `accounting` - (Optional) Configure the RADIUS server state to accept or refuse accounting messages. Possible values: [ on, off ]
* `pwdvendorid` - (Optional) Vendor ID of the password in the RADIUS response. Used to extract the user password. Minimum value =  1
* `pwdattributetype` - (Optional) Attribute type of the Vendor ID in the RADIUS response. Minimum value =  1
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups. Maximum length =  64
* `callingstationid` - (Optional) Send Calling-Station-ID of the client to the RADIUS server. IP Address of the client is sent as its Calling-Station-ID. Possible values: [ ENABLED, DISABLED ]
* `authservretry` - (Optional) Number of retry by the Citrix ADC before getting response from the RADIUS server. Minimum value =  1 Maximum value =  10
* `authentication` - (Optional) Configure the RADIUS server state to accept or refuse authentication messages. Possible values: [ on, off ]
* `tunnelendpointclientip` - (Optional) Send Tunnel Endpoint Client IP address to the RADIUS server. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaaradiusparams. It is a unique string prefixed with `tf-aaaradiusparams-`.