---
subcategory: "AAA"
---

# Data Source `aaaparameter`

The aaaparameter data source allows you to retrieve information about AAA parameters configuration.


## Example usage

```terraform
data "citrixadc_aaaparameter" "tf_aaaparameter" {
}

output "defaultauthtype" {
  value = data.citrixadc_aaaparameter.tf_aaaparameter.defaultauthtype
}

output "maxloginattempts" {
  value = data.citrixadc_aaaparameter.tf_aaaparameter.maxloginattempts
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `enablestaticpagecaching` - Enable or disable static page caching. Possible values: [ YES, NO ]
* `enableenhancedauthfeedback` - Enable or disable enhanced authentication feedback. Possible values: [ YES, NO ]
* `defaultauthtype` - Default authentication type for the AAA users. Possible values: [ LOCAL, LDAP, RADIUS, TACACS, CERT ]
* `maxloginattempts` - Maximum number of login attempts before lockout.
* `failedlogintimeout` - Number of minutes an account will be locked after exceeding maximum login attempts.
* `aaadloglevel` - AAAD log level. Possible values: [ EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG ]
* `aaadnatip` - Source IP address to use for traffic that is sent to the authentication server by the authentication proxy.
* `enablesessionstickiness` - Enable or disable stickiness for AAA authenticated users. Possible values: [ YES, NO ]
* `aaasessionloglevel` - Audit log level, which specifies the types of events to log for cli executed commands. Possible values: [ EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG ]
* `aaadloglevel` - AAAD log level. Possible values: [ EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG ]
* `dynaddr` - Enable or disable dynamic address allocation. Possible values: [ ON, OFF ]
* `ftmode` - Enable or disable fault tolerance for AAA. Possible values: [ ON, OFF, HA ]

## Attribute Reference

* `id` - The id of the aaaparameter. It is a system-generated identifier.
