---
subcategory: "AAA"
---

# Data Source `aaaradiusparams`

The aaaradiusparams data source allows you to retrieve information about AAA RADIUS parameters configuration.


## Example usage

```terraform
data "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
}

output "serverip" {
  value = data.citrixadc_aaaradiusparams.tf_aaaradiusparams.serverip
}

output "radnasip" {
  value = data.citrixadc_aaaradiusparams.tf_aaaradiusparams.radnasip
}

output "authtimeout" {
  value = data.citrixadc_aaaradiusparams.tf_aaaradiusparams.authtimeout
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `serverip` - IP address of your RADIUS server.
* `serverport` - Port number on which the RADIUS server listens for connections.
* `authtimeout` - Maximum number of seconds that the Citrix ADC waits for a response from the RADIUS server.
* `radkey` - The key shared between the RADIUS server and clients.
* `radnasip` - Send the Citrix ADC IP address to the RADIUS server. Possible values: [ ENABLED, DISABLED ]
* `radnasid` - Send the Network Access Server ID (NASID) for your Citrix ADC to the RADIUS server.
* `messageauthenticator` - Control whether the Message-Authenticator attribute is included in RADIUS packets. Possible values: [ ON, OFF ]
* `accounting` - Configure the RADIUS server state to accept or refuse accounting messages. Possible values: [ ON, OFF ]
* `authentication` - Configure the RADIUS server state to accept or refuse authentication messages. Possible values: [ ON, OFF ]
* `authservretry` - Number of retry by the Citrix ADC before getting response from the RADIUS server.
* `passencoding` - Enable password encoding in RADIUS packets that the Citrix ADC sends to the RADIUS server. Possible values: [ pap, chap, mschapv1, mschapv2 ]
* `ipvendorid` - Vendor ID attribute in the RADIUS response.
* `ipattributetype` - IP attribute type in the RADIUS response.
* `pwdvendorid` - Vendor ID of the password in the RADIUS response.
* `pwdattributetype` - Attribute type of the Vendor ID in the RADIUS response.
* `radvendorid` - Vendor ID for RADIUS group extraction.
* `radattributetype` - Attribute type for RADIUS group extraction.
* `radgroupsprefix` - Prefix string that precedes group names within a RADIUS attribute for RADIUS group extraction.
* `radgroupseparator` - Group separator string that delimits group names within a RADIUS attribute for RADIUS group extraction.
* `defaultauthenticationgroup` - This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `callingstationid` - Send Calling-Station-ID of the client to the RADIUS server. Possible values: [ ENABLED, DISABLED ]
* `tunnelendpointclientip` - Send Tunnel Endpoint Client IP address to the RADIUS server. Possible values: [ ENABLED, DISABLED ]

## Attribute Reference

* `id` - The id of the aaaradiusparams. It is a system-generated identifier.
