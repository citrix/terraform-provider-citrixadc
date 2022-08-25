---
subcategory: "Traffic Management"
---

# Resource: tmsessionparameter

The tmsessionparameter resource is used to create tmsessionparameter.


## Example usage

```hcl
resource "citrixadc_tmsessionparameter" "tf_tmsessionparameter" {
  sesstimeout                = 40
  defaultauthorizationaction = "ALLOW"
  sso                        = "OFF"
  ssodomain                  = 3
}

```


## Argument Reference

* `sesstimeout` - (Optional) Session timeout, in minutes. If there is no traffic during the timeout period, the user is disconnected and must reauthenticate to access the intranet resources. Minimum value =  1
* `defaultauthorizationaction` - (Optional) Allow or deny access to content for which there is no specific authorization policy. Possible values: [ ALLOW, DENY ]
* `sso` - (Optional) Log users on to all web applications automatically after they authenticate, or pass users to the web application logon page to authenticate for each application. Note that this configuration does not honor the following authentication types for security reason. BASIC, DIGEST, and NTLM (without Negotiate NTLM2 Key or Negotiate Sign Flag). Use TM TrafficAction to configure SSO for these authentication types. Possible values: [ on, off ]
* `ssocredential` - (Optional) Use primary or secondary authentication credentials for single sign-on. Possible values: [ PRIMARY, SECONDARY ]
* `ssodomain` - (Optional) Domain to use for single sign-on. Minimum length =  1 Maximum length =  32
* `kcdaccount` - (Optional) Kerberos constrained delegation account name. Minimum length =  1 Maximum length =  32
* `httponlycookie` - (Optional) Allow only an HTTP session cookie, in which case the cookie cannot be accessed by scripts. Possible values: [ YES, NO ]
* `persistentcookie` - (Optional) Use persistent SSO cookies for the traffic session. A persistent cookie remains on the user device and is sent with each HTTP request. The cookie becomes stale if the session ends. Possible values: [ on, off ]
* `persistentcookievalidity` - (Optional) Integer specifying the number of minutes for which the persistent cookie remains valid. Can be set only if the persistence cookie setting is enabled. Minimum value =  1
* `homepage` - (Optional) Web address of the home page that a user is displayed when authentication vserver is bookmarked and used to login.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the tmsessionparameter. It is a unique string prefixed with `tf-tmsessionparameter-` attribute.
