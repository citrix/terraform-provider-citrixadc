---
subcategory: "Authentication"
---

# Resource: authenticationvserver

The authenticationvserver resource is used to create authentication virtual server resource.


## Example usage

```hcl
resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name        = "tf_authenticationvserver"
  servicetype = "SSL"
  comment     = "Comments"
	authentication = "ON"
	state          = "DISABLED"
}
```


## Argument Reference

* `name` - (Required) Name for the new authentication virtual server.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the authentication virtual server is added by using the rename authentication vserver command.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').
* `appflowlog` - (Optional) Log AppFlow flow information.
* `authentication` - (Optional) Require users to be authenticated before sending traffic through this virtual server.
* `authenticationdomain` - (Optional) The domain of the authentication cookie set by Authentication vserver
* `certkeynames` - (Optional) Name of the certificate key that was bound to the corresponding SSL virtual server as the Certificate Authority for the device certificate
* `comment` - (Optional) Any comments associated with this virtual server.
* `failedlogintimeout` - (Optional) Number of minutes an account will be locked if user exceeds maximum permissible attempts
* `ipv46` - (Optional) IP address of the authentication virtual server, if a single IP address is assigned to the virtual server.
* `maxloginattempts` - (Optional) Maximum Number of login Attempts
* `newname` - (Optional) New name of the authentication virtual server.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, 'my authentication policy' or "my authentication policy").
* `port` - (Optional) TCP port on which the virtual server accepts connections.
* `range` - (Optional) If you are creating a series of virtual servers with a range of IP addresses assigned to them, the length of the range.  The new range of authentication virtual servers will have IP addresses consecutively numbered, starting with the primary address specified with the IP Address parameter.
* `samesite` - (Optional) SameSite attribute value for Cookies generated in AAATM context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite
* `servicetype` - (Optional) Protocol type of the authentication virtual server. Always SSL.
* `state` - (Optional) Initial state of the new virtual server.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationvserver. It has the same value as the `name` attribute.


## Import

A authenticationvserver can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationvserver.tf_authenticationvserver tf_authenticationvserver
```
