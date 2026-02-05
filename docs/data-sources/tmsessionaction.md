---
subcategory: "Traffic Management"
---

# Data Source: tmsessionaction

The tmsessionaction data source allows you to retrieve information about a TM session action.

## Example usage

```terraform
data "citrixadc_tmsessionaction" "tf_tmsessionaction" {
  name = "my_tmsession_action"
}

output "sesstimeout" {
  value = data.citrixadc_tmsessionaction.tf_tmsessionaction.sesstimeout
}

output "defaultauthorizationaction" {
  value = data.citrixadc_tmsessionaction.tf_tmsessionaction.defaultauthorizationaction
}
```

## Argument Reference

* `name` - (Required) Name for the session action.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `defaultauthorizationaction` - Allow or deny access to content for which there is no specific authorization policy.
* `homepage` - Web address of the home page that a user is displayed when authentication vserver is bookmarked and used to login.
* `httponlycookie` - Allow only an HTTP session cookie, in which case the cookie cannot be accessed by scripts.
* `kcdaccount` - Kerberos constrained delegation account name.
* `persistentcookie` - Enable or disable persistent SSO cookies for the traffic management (TM) session. A persistent cookie remains on the user device and is sent with each HTTP request. The cookie becomes stale if the session ends. This setting is overwritten if a traffic action sets persistent cookie to OFF. Note: If persistent cookie is enabled, make sure you set the persistent cookie validity.
* `persistentcookievalidity` - Integer specifying the number of minutes for which the persistent cookie remains valid. Can be set only if the persistent cookie setting is enabled.
* `sesstimeout` - Session timeout, in minutes. If there is no traffic during the timeout period, the user is disconnected and must reauthenticate to access intranet resources.
* `sso` - Use single sign-on (SSO) to log users on to all web applications automatically after they authenticate, or pass users to the web application logon page to authenticate to each application individually. Note that this configuration does not honor the following authentication types for security reason. BASIC, DIGEST, and NTLM (without Negotiate NTLM2 Key or Negotiate Sign Flag). Use TM TrafficAction to configure SSO for these authentication types.
* `ssocredential` - Use the primary or secondary authentication credentials for single sign-on (SSO).
* `ssodomain` - Domain to use for single sign-on (SSO).

## Attribute Reference

* `id` - The id of the tmsessionaction. It has the same value as the `name` attribute.

## Import

A tmsessionaction can be imported using its name, e.g.

```shell
terraform import citrixadc_tmsessionaction.tf_tmsessionaction my_tmsession_action
```
