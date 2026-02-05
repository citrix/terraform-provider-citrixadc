---
subcategory: "Authentication"
---

# Data Source `authenticationradiusaction`

The authenticationradiusaction data source allows you to retrieve information about authentication RADIUS actions.


## Example usage

```terraform
data "citrixadc_authenticationradiusaction" "tf_radiusaction" {
  name = "my_radiusaction"
}

output "serverip" {
  value = data.citrixadc_authenticationradiusaction.tf_radiusaction.serverip
}

output "serverport" {
  value = data.citrixadc_authenticationradiusaction.tf_radiusaction.serverport
}
```


## Argument Reference

* `name` - (Required) Name for the RADIUS action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the RADIUS action is added.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `accounting` - Whether the RADIUS server is currently accepting accounting messages.
* `authentication` - Configure the RADIUS server state to accept or refuse authentication messages.
* `authservretry` - Number of retry by the Citrix ADC before getting response from the RADIUS server.
* `authtimeout` - Number of seconds the Citrix ADC waits for a response from the RADIUS server.
* `callingstationid` - Send Calling-Station-ID of the client to the RADIUS server. IP Address of the client is sent as its Calling-Station-ID.
* `defaultauthenticationgroup` - This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `ipattributetype` - Remote IP address attribute type in a RADIUS response.
* `ipvendorid` - Vendor ID of the intranet IP attribute in the RADIUS response. NOTE: A value of 0 indicates that the attribute is not vendor encoded.
* `messageauthenticator` - Control whether the Message-Authenticator attribute is included in a RADIUS Access-Request packet.
* `passencoding` - Encoding type for passwords in RADIUS packets that the Citrix ADC sends to the RADIUS server.
* `pwdattributetype` - Vendor-specific password attribute type in a RADIUS response.
* `pwdvendorid` - Vendor ID of the attribute, in the RADIUS response, used to extract the user password.
* `radattributetype` - RADIUS attribute type, used for RADIUS group extraction.
* `radgroupseparator` - RADIUS group separator string. The group separator delimits group names within a RADIUS attribute for RADIUS group extraction.
* `radgroupsprefix` - RADIUS groups prefix string. This groups prefix precedes the group names within a RADIUS attribute for RADIUS group extraction.
* `radkey` - Key shared between the RADIUS server and the Citrix ADC. Required to allow the Citrix ADC to communicate with the RADIUS server.
* `radnasid` - If configured, this string is sent to the RADIUS server as the Network Access Server ID (NASID).
* `radnasip` - If enabled, the Citrix ADC IP address (NSIP) is sent to the RADIUS server as the Network Access Server IP (NASIP) address. The RADIUS protocol defines the meaning and use of the NASIP address.
* `radvendorid` - RADIUS vendor ID attribute, used for RADIUS group extraction.
* `serverip` - IP address assigned to the RADIUS server.
* `servername` - RADIUS server name as a FQDN. Mutually exclusive with RADIUS IP address.
* `serverport` - Port number on which the RADIUS server listens for connections.
* `targetlbvserver` - If transport mode is TLS, specify the name of LB vserver to associate. The LB vserver needs to be of type TCP and service associated needs to be SSL_TCP.
* `transport` - Transport mode to RADIUS server.
* `tunnelendpointclientip` - Send Tunnel Endpoint Client IP address to the RADIUS server.

## Attribute Reference

* `id` - The id of the authenticationradiusaction. It has the same value as the `name` attribute.


## Import

A authenticationradiusaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationradiusaction.tf_radiusaction my_radiusaction
```
