---
subcategory: "Traffic Management"
---

# Resource: tmsessionaction

The tmsessionaction resource is used to create tmsessionaction.


## Example usage

```hcl
resource "citrixadc_tmsessionaction" "tf_tmsessionaction" {
  name                       = "my_tmsession_action"
  sesstimeout                = 10
  defaultauthorizationaction = "ALLOW"
  sso                        = "OFF"
}
```


## Argument Reference

* `name` - (Required) Name for the session action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a session action is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action'). Minimum length =  1
* `sesstimeout` - (Optional) Session timeout, in minutes. If there is no traffic during the timeout period, the user is disconnected and must reauthenticate to access intranet resources. Minimum value =  1
* `defaultauthorizationaction` - (Optional) Allow or deny access to content for which there is no specific authorization policy. Possible values: [ ALLOW, DENY ]
* `sso` - (Optional) Use single sign-on (SSO) to log users on to all web applications automatically after they authenticate, or pass users to the web application logon page to authenticate to each application individually. Note that this configuration does not honor the following authentication types for security reason. BASIC, DIGEST, and NTLM (without Negotiate NTLM2 Key or Negotiate Sign Flag). Use TM TrafficAction to configure SSO for these authentication types. Possible values: [ on, off ]
* `ssocredential` - (Optional) Use the primary or secondary authentication credentials for single sign-on (SSO). Possible values: [ PRIMARY, SECONDARY ]
* `ssodomain` - (Optional) Domain to use for single sign-on (SSO). Minimum length =  1 Maximum length =  32
* `httponlycookie` - (Optional) Allow only an HTTP session cookie, in which case the cookie cannot be accessed by scripts. Possible values: [ YES, NO ]
* `kcdaccount` - (Optional) Kerberos constrained delegation account name. Minimum length =  1 Maximum length =  32
* `persistentcookie` - (Optional) Enable or disable persistent SSO cookies for the traffic management (TM) session. A persistent cookie remains on the user device and is sent with each HTTP request. The cookie becomes stale if the session ends. This setting is overwritten if a traffic action sets persistent cookie to OFF. Note: If persistent cookie is enabled, make sure you set the persistent cookie validity. Possible values: [ on, off ]
* `persistentcookievalidity` - (Optional) Integer specifying the number of minutes for which the persistent cookie remains valid. Can be set only if the persistent cookie setting is enabled. Minimum value =  1
* `homepage` - (Optional) Web address of the home page that a user is displayed when authentication vserver is bookmarked and used to login.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the tmsessionaction. It has the same value as the `name` attribute.


## Import

A tmsessionaction can be imported using its name, e.g.

```shell
terraform import citrixadc_tmsessionaction.tf_tmsessionaction my_tmsession_action
```
