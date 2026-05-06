---
subcategory: "AAA"
---

# Resource: aaaradiusparams

The aaaradiusparams resource is used to configure global AAA RADIUS parameters on the Citrix ADC.


## Example usage

### Using radkey (sensitive attribute - persisted in state)

```hcl
variable "aaaradiusparams_radkey" {
  type      = string
  sensitive = true
}

resource "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
  radkey      = var.aaaradiusparams_radkey
  serverip    = "10.0.0.10"
  serverport  = 1812
  authtimeout = 3
  radnasip    = "ENABLED"
}
```

### Using radkey_wo (write-only/ephemeral - NOT persisted in state)

The `radkey_wo` attribute provides an ephemeral path for the RADIUS shared key. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the key value changes, increment `radkey_wo_version`.

```hcl
variable "aaaradiusparams_radkey" {
  type      = string
  sensitive = true
}

resource "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
  radkey_wo         = var.aaaradiusparams_radkey
  radkey_wo_version = 1
  serverip          = "10.0.0.10"
  serverport        = 1812
  authtimeout       = 3
  radnasip          = "ENABLED"
}
```

To rotate the key, update the variable value and bump the version:

```hcl
resource "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
  radkey_wo         = var.aaaradiusparams_radkey
  radkey_wo_version = 2  # Bumped to trigger update
  serverip          = "10.0.0.10"
  serverport        = 1812
  authtimeout       = 3
  radnasip          = "ENABLED"
}
```


## Argument Reference

* `accounting` - (Optional) Configure the RADIUS server state to accept or refuse accounting messages. Possible values: [ on, off ]
* `authentication` - (Optional) Configure the RADIUS server state to accept or refuse authentication messages. Possible values: [ on, off ]
* `authservretry` - (Optional) Number of retry by the Citrix ADC before getting response from the RADIUS server. Defaults to `3`.
* `authtimeout` - (Optional) Maximum number of seconds that the Citrix ADC waits for a response from the RADIUS server. Defaults to `3`.
* `callingstationid` - (Optional) Send Calling-Station-ID of the client to the RADIUS server. IP Address of the client is sent as its Calling-Station-ID. Possible values: [ ENABLED, DISABLED ]
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `ipattributetype` - (Optional) IP attribute type in the RADIUS response.
* `ipvendorid` - (Optional) Vendor ID attribute in the RADIUS response. If the attribute is not vendor-encoded, it is set to 0.
* `messageauthenticator` - (Optional) Control whether the Message-Authenticator attribute is included in a RADIUS Access-Request packet.
* `passencoding` - (Optional) Enable password encoding in RADIUS packets that the Citrix ADC sends to the RADIUS server. Possible values: [ pap, chap, mschapv1, mschapv2 ]
* `pwdattributetype` - (Optional) Attribute type of the Vendor ID in the RADIUS response.
* `pwdvendorid` - (Optional) Vendor ID of the password in the RADIUS response. Used to extract the user password.
* `radattributetype` - (Optional) Attribute type for RADIUS group extraction.
* `radgroupseparator` - (Optional) Group separator string that delimits group names within a RADIUS attribute for RADIUS group extraction.
* `radgroupsprefix` - (Optional) Prefix string that precedes group names within a RADIUS attribute for RADIUS group extraction.
* `radkey` - (Optional, Sensitive) The key shared between the RADIUS server and clients. Required for allowing the Citrix ADC to communicate with the RADIUS server. The value is persisted in Terraform state (encrypted). See also `radkey_wo` for an ephemeral alternative.
* `radkey_wo` - (Optional, Sensitive, WriteOnly) Same as `radkey`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `radkey_wo_version`. If both `radkey` and `radkey_wo` are set, `radkey_wo` takes precedence.
* `radkey_wo_version` - (Optional) An integer version tracker for `radkey_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `radnasid` - (Optional) Send the Network Access Server ID (NASID) for your Citrix ADC to the RADIUS server as the nasid part of the Radius protocol.
* `radnasip` - (Optional) Send the Citrix ADC IP (NSIP) address to the RADIUS server as the Network Access Server IP (NASIP) part of the Radius protocol. Possible values: [ ENABLED, DISABLED ]
* `radvendorid` - (Optional) Vendor ID for RADIUS group extraction.
* `serverip` - (Optional) IP address of your RADIUS server.
* `serverport` - (Optional) Port number on which the RADIUS server listens for connections. Defaults to `1812`.
* `tunnelendpointclientip` - (Optional) Send Tunnel Endpoint Client IP address to the RADIUS server. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaaradiusparams. It is a unique string prefixed with "aaaradiusparams-config".