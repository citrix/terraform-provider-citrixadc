---
subcategory: "Authentication"
---

# Data Source `authenticationvserver`

The authenticationvserver data source allows you to retrieve information about an existing authentication virtual server.


## Example usage

```terraform
data "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name = "my_authenticationvserver"
}

output "servicetype" {
  value = data.citrixadc_authenticationvserver.tf_authenticationvserver.servicetype
}

output "comment" {
  value = data.citrixadc_authenticationvserver.tf_authenticationvserver.comment
}
```


## Argument Reference

* `name` - (Required) Name for the new authentication virtual server. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the authentication virtual server is added by using the rename authentication vserver command.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationvserver. It has the same value as the `name` attribute.
* `appflowlog` - Log AppFlow flow information.
* `authentication` - Require users to be authenticated before sending traffic through this virtual server.
* `authenticationdomain` - The domain of the authentication cookie set by Authentication vserver.
* `certkeynames` - Name of the certificate key that was bound to the corresponding SSL virtual server as the Certificate Authority for the device certificate.
* `comment` - Any comments associated with this virtual server.
* `failedlogintimeout` - Number of minutes an account will be locked if user exceeds maximum permissible attempts.
* `ipv46` - IP address of the authentication virtual server, if a single IP address is assigned to the virtual server.
* `maxloginattempts` - Maximum Number of login Attempts.
* `newname` - New name of the authentication virtual server. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `port` - TCP port on which the virtual server accepts connections.
* `range` - If you are creating a series of virtual servers with a range of IP addresses assigned to them, the length of the range. The new range of authentication virtual servers will have IP addresses consecutively numbered, starting with the primary address specified with the IP Address parameter.
* `samesite` - SameSite attribute value for Cookies generated in AAATM context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite.
* `servicetype` - Protocol type of the authentication virtual server. Always SSL.
* `state` - Initial state of the new virtual server.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.


## Import

A authenticationvserver can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationvserver.tf_authenticationvserver my_authenticationvserver
```
