---
subcategory: "Traffic Management"
---

# Resource: tmtrafficaction

The tmtrafficaction resource is used to create tmtrafficaction.


## Example usage

```hcl
resource "citrixadc_tmtrafficaction" "tf_tmtrafficaction" {
  name             = "my_traffic_action"
  apptimeout       = 5
  sso              = "OFF"
  persistentcookie = "ON"
}
```


## Argument Reference

* `name` - (Required) Name for the traffic action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a traffic action is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action'). Minimum length =  1
* `apptimeout` - (Optional) Time interval, in minutes, of user inactivity after which the connection is closed. Minimum value =  1 Maximum value =  715827
* `sso` - (Optional) Use single sign-on for the resource that the user is accessing now. Possible values: [ on, off ]
* `formssoaction` - (Optional) Name of the configured form-based single sign-on profile.
* `persistentcookie` - (Optional) Use persistent cookies for the traffic session. A persistent cookie remains on the user device and is sent with each HTTP request. The cookie becomes stale if the session ends. Possible values: [ on, off ]
* `initiatelogout` - (Optional) Initiate logout for the traffic management (TM) session if the policy evaluates to true. The session is then terminated after two minutes. Possible values: [ on, off ]
* `kcdaccount` - (Optional) Kerberos constrained delegation account name. Minimum length =  1 Maximum length =  32
* `samlssoprofile` - (Optional) Profile to be used for doing SAML SSO to remote relying party. Minimum length =  1
* `forcedtimeout` - (Optional) Setting to start, stop or reset TM session force timer. Possible values: [ START, STOP, RESET ]
* `forcedtimeoutval` - (Optional) Time interval, in minutes, for which force timer should be set.
* `userexpression` - (Optional) expression that will be evaluated to obtain username for SingleSignOn. Maximum length =  256
* `passwdexpression` - (Optional) expression that will be evaluated to obtain password for SingleSignOn. Maximum length =  256


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the tmtrafficaction. It has the same value as the `name` attribute.


## Import

A tmtrafficaction can be imported using its name, e.g.

```shell
terraform import citrixadc_tmtrafficaction.tf_tmtrafficaction my_traffic_action
```
