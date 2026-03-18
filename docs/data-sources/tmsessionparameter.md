---
subcategory: "Traffic Management"
---

# Data Source: tmsessionparameter

The tmsessionparameter data source allows you to retrieve information about the global TM session parameters.

## Example usage

```terraform
data "citrixadc_tmsessionparameter" "tf_tmsessionparameter" {
}

output "sesstimeout" {
  value = data.citrixadc_tmsessionparameter.tf_tmsessionparameter.sesstimeout
}

output "defaultauthorizationaction" {
  value = data.citrixadc_tmsessionparameter.tf_tmsessionparameter.defaultauthorizationaction
}
```

## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `defaultauthorizationaction` - Allow or deny access to content for which there is no specific authorization policy.
* `homepage` - Web address of the home page that a user is displayed when authentication vserver is bookmarked and used to login.
* `httponlycookie` - Allow only an HTTP session cookie, in which case the cookie cannot be accessed by scripts.
* `kcdaccount` - Kerberos constrained delegation account name.
* `persistentcookie` - Use persistent SSO cookies for the traffic session. A persistent cookie remains on the user device and is sent with each HTTP request. The cookie becomes stale if the session ends.
* `persistentcookievalidity` - Integer specifying the number of minutes for which the persistent cookie remains valid. Can be set only if the persistence cookie setting is enabled.
* `sesstimeout` - Session timeout, in minutes. If there is no traffic during the timeout period, the user is disconnected and must reauthenticate to access the intranet resources.
* `sso` - Log users on to all web applications automatically after they authenticate, or pass users to the web application logon page to authenticate for each application. Note that this configuration does not honor the following authentication types for security reason. BASIC, DIGEST, and NTLM (without Negotiate NTLM2 Key or Negotiate Sign Flag). Use TM TrafficAction to configure SSO for these authentication types.
* `ssocredential` - Use primary or secondary authentication credentials for single sign-on.
* `ssodomain` - Domain to use for single sign-on.

## Attribute Reference

* `id` - The id of the tmsessionparameter. It is a system-generated identifier.
