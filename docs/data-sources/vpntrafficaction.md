---
subcategory: "VPN"
---

# Data Source: vpntrafficaction

The vpntrafficaction data source allows you to retrieve information about a VPN traffic action configuration.

## Example usage

```terraform
data "citrixadc_vpntrafficaction" "tf_action" {
  name = "Testing"
}

output "apptimeout" {
  value = data.citrixadc_vpntrafficaction.tf_action.apptimeout
}

output "sso" {
  value = data.citrixadc_vpntrafficaction.tf_action.sso
}
```

## Argument Reference

* `name` - (Required) Name for the traffic action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpntrafficaction. It is the same as the `name` attribute.
* `apptimeout` - Maximum amount of time, in minutes, a user can stay logged on to the web application.
* `formssoaction` - Name of the form-based single sign-on profile. Form-based single sign-on allows users to log on one time to all protected applications in your network, instead of requiring them to log on separately to access each one.
* `fta` - Specify file type association, which is a list of file extensions that users are allowed to open.
* `hdx` - Provide hdx proxy to the ICA traffic
* `kcdaccount` - Kerberos constrained delegation account name
* `passwdexpression` - expression that will be evaluated to obtain password for SingleSignOn
* `proxy` - IP address and Port of the proxy server to be used for HTTP access for this request.
* `qual` - Protocol, either HTTP or TCP, to be used with the action. Possible values: [ http, tcp ]
* `samlssoprofile` - Profile to be used for doing SAML SSO to remote relying party
* `sso` - Use single sign-on for the resource. Possible values: [ ON, OFF ]
* `userexpression` - expression that will be evaluated to obtain username for SingleSignOn
* `wanscaler` - Use the Repeater Plug-in to optimize traffic. Possible values: [ ON, OFF ]
