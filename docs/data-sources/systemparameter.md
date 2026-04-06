---
subcategory: "System"
---

# Data Source: systemparameter

The systemparameter data source allows you to retrieve information about system parameters configured on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_systemparameter" "tf_systemparameter" {
}

output "rbaonresponse" {
  value = data.citrixadc_systemparameter.tf_systemparameter.rbaonresponse
}

output "timeout" {
  value = data.citrixadc_systemparameter.tf_systemparameter.timeout
}
```

## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `basicauth` - Enable or disable basic authentication for Nitro API.
* `cliloglevel` - Audit log level, which specifies the types of events to log for cli executed commands. Possible values: [ EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG ]
* `daystoexpire` - Password expiry days for all the system users. The daystoexpire value ranges from 30 to 255.
* `doppler` - Enable or disable Doppler. Possible values: [ ENABLED, DISABLED ]
* `fipsusermode` - Use this option to set the FIPS mode for key user-land processes. When enabled, these user-land processes will operate in FIPS mode.
* `forcepasswordchange` - Enable or disable force password change for nsroot user. Possible values: [ ENABLED, DISABLED ]
* `googleanalytics` - Enable or disable Google analytics. Possible values: [ ENABLED, DISABLED ]
* `id` - The id of the systemparameter. It is a system-generated identifier.
* `localauth` - When enabled, local users can access Citrix ADC even when external authentication is configured. Possible values: [ ENABLED, DISABLED ]
* `maxsessionperuser` - Maximum number of client connection allowed per user. The maxsessionperuser value ranges from 1 to 40.
* `minpasswordlen` - Minimum length of system user password. When strong password is enabled default minimum length is 8.
* `natpcbforceflushlimit` - Flush the system if the number of Network Address Translation Protocol Control Blocks (NATPCBs) exceeds this value.
* `natpcbrstontimeout` - Send a reset signal to client and server connections when their NATPCBs time out. Possible values: [ ENABLED, DISABLED ]
* `passwordhistorycontrol` - Enables or disable password expiry feature for system users. Possible values: [ ENABLED, DISABLED ]
* `promptstring` - String to display at the command-line prompt.
* `pwdhistorycount` - Number of passwords to be maintained as history for system users. The pwdhistorycount value ranges from 1 to 10.
* `rbaonresponse` - Enable or disable Role-Based Authentication (RBA) on responses. Possible values: [ ENABLED, DISABLED ]
* `reauthonauthparamchange` - Enable or disable External user reauthentication when authentication parameter changes. Possible values: [ ENABLED, DISABLED ]
* `removesensitivefiles` - Use this option to remove the sensitive files from the system like authorise keys, public keys etc. Possible values: [ ENABLED, DISABLED ]
* `restrictedtimeout` - Enable/Disable the restricted timeout behaviour. Possible values: [ ENABLED, DISABLED ]
* `strongpassword` - After enabling strong password, all the passwords / sensitive information must have required complexity.
* `timeout` - CLI session inactivity timeout, in seconds.
* `totalauthtimeout` - Total time a request can take for authentication/authorization.
* `wafprotection` - Configure WAF protection for endpoints used by NetScaler management interfaces. Possible values: [ DEFAULT, GUI, DISABLED ]
* `warnpriorndays` - Number of days before which password expiration warning would be thrown with respect to daystoexpire. The warnpriorndays value ranges from 5 to 40.
