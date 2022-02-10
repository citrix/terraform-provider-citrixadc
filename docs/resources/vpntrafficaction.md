---
subcategory: "Vpn"
---

# Resource: vpntrafficaction

The vpntrafficaction resource is used to create vpn traffic action.


## Example usage

```hcl
resource "citrixadc_vpntrafficaction" "tf_vpntrafficaction" {
  name       = "Testing"
  qual       = "tcp"
  apptimeout = 20
  fta        = "OFF"
  hdx        = "OFF"
  sso        = "ON"
}
```


## Argument Reference

* `name` - (Required) Name for the traffic action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a traffic action is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
* `qual` - (Required) Protocol, either HTTP or TCP, to be used with the action.
* `apptimeout` - (Optional) Maximum amount of time, in minutes, a user can stay logged on to the web application.
* `formssoaction` - (Optional) Name of the form-based single sign-on profile. Form-based single sign-on allows users to log on one time to all protected applications in your network, instead of requiring them to log on separately to access each one.
* `fta` - (Optional) Specify file type association, which is a list of file extensions that users are allowed to open.
* `hdx` - (Optional) Provide hdx proxy to the ICA traffic
* `kcdaccount` - (Optional) Kerberos constrained delegation account name
* `passwdexpression` - (Optional) expression that will be evaluated to obtain password for SingleSignOn
* `proxy` - (Optional) IP address and Port of the proxy server to be used for HTTP access for this request.
* `samlssoprofile` - (Optional) Profile to be used for doing SAML SSO to remote relying party
* `sso` - (Optional) Provide single sign-on to the web application. 	    NOTE : Authentication mechanisms like Basic-authentication  require the user credentials to be sent in plaintext which is not secure if the server is running on HTTP (instead of HTTPS).
* `userexpression` - (Optional) expression that will be evaluated to obtain username for SingleSignOn
* `wanscaler` - (Optional) Use the Repeater Plug-in to optimize network traffic.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpntrafficaction. It has the same value as the `name` attribute.


## Import

A vpntrafficaction can be imported using its name, e.g.

```shell
terraform import citrixadc_vpntrafficaction.tf_vpntrafficaction Testing
```
