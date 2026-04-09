---
subcategory: "Authentication"
---

# Resource: authenticationradiusaction

The authenticationradiusaction resource is used to create authentication radiusaction resource.


## Example usage

### Basic usage

```hcl
resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
  name         = "tf_radiusaction"
  serverip     = "1.2.3.4"
  serverport   = 8080
  authtimeout  = 2
  radnasip     = "DISABLED"
  passencoding = "chap"
}
```

### Using radkey (sensitive attribute - persisted in state)

```hcl
variable "radiusaction_radkey" {
  type      = string
  sensitive = true
}

resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
  name         = "tf_radiusaction"
  radkey       = var.radiusaction_radkey
  serverip     = "1.2.3.4"
  serverport   = 8080
  authtimeout  = 2
  radnasip     = "DISABLED"
  passencoding = "chap"
}
```

### Using radkey_wo (write-only/ephemeral - NOT persisted in state)

The `radkey_wo` attribute provides an ephemeral path for the RADIUS shared key. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the key value changes, increment `radkey_wo_version`.

```hcl
variable "radiusaction_radkey" {
  type      = string
  sensitive = true
}

resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
  name              = "tf_radiusaction"
  radkey_wo         = var.radiusaction_radkey
  radkey_wo_version = 1
  serverip          = "1.2.3.4"
  serverport        = 8080
  authtimeout       = 2
  radnasip          = "DISABLED"
  passencoding      = "chap"
}
```

To rotate the key, update the variable value and bump the version:

```hcl
resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
  name              = "tf_radiusaction"
  radkey_wo         = var.radiusaction_radkey
  radkey_wo_version = 2  # Bumped to trigger update
  serverip          = "1.2.3.4"
  serverport        = 8080
  authtimeout       = 2
  radnasip          = "DISABLED"
  passencoding      = "chap"
}
```


## Argument Reference

* `name` - (Required) Name for the RADIUS action.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the RADIUS action is added.
* `radkey` - (Optional, Sensitive) Key shared between the RADIUS server and the Citrix ADC. Required to allow the Citrix ADC to communicate with the RADIUS server. The value is persisted in Terraform state (encrypted). See also `radkey_wo` for an ephemeral alternative.
* `radkey_wo` - (Optional, Sensitive, WriteOnly) Same as `radkey`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `radkey_wo_version`. If both `radkey` and `radkey_wo` are set, `radkey_wo` takes precedence.
* `radkey_wo_version` - (Optional) An integer version tracker for `radkey_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the key has changed and trigger an update. Defaults to `1`.
* `accounting` - (Optional) Whether the RADIUS server is currently accepting accounting messages.
* `authentication` - (Optional) Configure the RADIUS server state to accept or refuse authentication messages.
* `authservretry` - (Optional) Number of retry by the Citrix ADC before getting response from the RADIUS server.
* `authtimeout` - (Optional) Number of seconds the Citrix ADC waits for a response from the RADIUS server.
* `callingstationid` - (Optional) Send Calling-Station-ID of the client to the RADIUS server. IP Address of the client is sent as its Calling-Station-ID.
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `ipattributetype` - (Optional) Remote IP address attribute type in a RADIUS response.
* `ipvendorid` - (Optional) Vendor ID of the intranet IP attribute in the RADIUS response. NOTE: A value of 0 indicates that the attribute is not vendor encoded.
* `passencoding` - (Optional) Encoding type for passwords in RADIUS packets that the Citrix ADC sends to the RADIUS server.
* `pwdattributetype` - (Optional) Vendor-specific password attribute type in a RADIUS response.
* `pwdvendorid` - (Optional) Vendor ID of the attribute, in the RADIUS response, used to extract the user password.
* `radattributetype` - (Optional) RADIUS attribute type, used for RADIUS group extraction.
* `radgroupseparator` - (Optional) RADIUS group separator string The group separator delimits group names within a RADIUS attribute for RADIUS group extraction.
* `radgroupsprefix` - (Optional) RADIUS groups prefix string.  This groups prefix precedes the group names within a RADIUS attribute for RADIUS group extraction.
* `radnasid` - (Optional) If configured, this string is sent to the RADIUS server as the Network Access Server ID (NASID).
* `radnasip` - (Optional) If enabled, the Citrix ADC IP address (NSIP) is sent to the RADIUS server as the  Network Access Server IP (NASIP) address.  The RADIUS protocol defines the meaning and use of the NASIP address.
* `radvendorid` - (Optional) RADIUS vendor ID attribute, used for RADIUS group extraction.
* `serverip` - (Optional) IP address assigned to the RADIUS server.
* `servername` - (Optional) RADIUS server name as a FQDN.  Mutually exclusive with RADIUS IP address.
* `serverport` - (Optional) Port number on which the RADIUS server listens for connections.
* `tunnelendpointclientip` - (Optional) Send Tunnel Endpoint Client IP address to the RADIUS server.
* `messageauthenticator` - (Optional) Control whether the Message-Authenticator attribute is included in a RADIUS Access-Request packet.
* `targetlbvserver` - (Optional) If transport mode is TLS, specify the name of LB vserver to associate. The LB vserver needs to be of type TCP and service associated needs to be SSL_TCP
* `transport` - (Optional) Transport mode to RADIUS server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationradiusaction. It has the same value as the `name` attribute.


## Import

A authenticationradiusaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationradiusaction.tf_radiusaction tf_radiusaction
```
