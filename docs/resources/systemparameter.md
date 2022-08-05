---
subcategory: "System"
---

# Resource: systemparameter

The systemparameter resource is used to update systemparameter.


## Example usage

```hcl
resource "citrixadc_systemparameter" "tf_systemparameter" {
    rbaonresponse = "ENABLED"
    natpcbforceflushlimit = 3000
    natpcbrstontimeout = "DISABLED"
    timeout = 500
    doppler = "ENABLED"
}
```


## Argument Reference

* `basicauth` - (Optional) Enable or disable basic authentication for Nitro API.
* `cliloglevel` - (Optional) Audit log level, which specifies the types of events to log for cli executed commands. Available values function as follows: * EMERGENCY - Events that indicate an immediate crisis on the server. * ALERT - Events that might require action. * CRITICAL - Events that indicate an imminent server crisis. * ERROR - Events that indicate some type of error. * WARNING - Events that require action in the near future. * NOTICE - Events that the administrator should know about. * INFORMATIONAL - All but low-level events. * DEBUG - All events, in extreme detail.
* `doppler` - (Optional) Enable or disable Doppler
* `fipsusermode` - (Optional) Use this option to set the FIPS mode for key user-land processes. When enabled, these user-land processes will operate in FIPS mode. In this mode, theses processes will use FIPS 140-2 Level-1 certified crypto algorithms. Default is disabled, wherein, these user-land processes will not operate in FIPS mode.
* `forcepasswordchange` - (Optional) Enable or disable force password change for nsroot user
* `googleanalytics` - (Optional) Enable or disable Google analytics
* `localauth` - (Optional) When enabled, local users can access Citrix ADC even when external authentication is configured. When disabled, local users are not allowed to access the Citrix ADC, Local users can access the Citrix ADC only when the configured external authentication servers are unavailable. This parameter is not applicable to SSH Key-based authentication
* `minpasswordlen` - (Optional) Minimum length of system user password. When strong password is enabled default minimum length is 4. User entered value can be greater than or equal to 4. Default mininum value is 1 when strong password is disabled. Maximum value is 127 in both cases.
* `natpcbforceflushlimit` - (Optional) Flush the system if the number of Network Address Translation Protocol Control Blocks (NATPCBs) exceeds this value.
* `natpcbrstontimeout` - (Optional) Send a reset signal to client and server connections when their NATPCBs time out. Avoids the buildup of idle TCP connections on both the sides.
* `promptstring` - (Optional) String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (_), and the following variables:  * %u - Will be replaced by the user name. * %h - Will be replaced by the hostname of the Citrix ADC. * %t - Will be replaced by the current time in 12-hour format. * %T - Will be replaced by the current time in 24-hour format. * %d - Will be replaced by the current date. * %s - Will be replaced by the state of the Citrix ADC.  Note: The 63-character limit for the length of the string does not apply to the characters that replace the variables.
* `rbaonresponse` - (Optional) Enable or disable Role-Based Authentication (RBA) on responses.
* `reauthonauthparamchange` - (Optional) Enable or disable External user reauthentication when authentication parameter changes
* `removesensitivefiles` - (Optional) Use this option to remove the sensitive files from the system like authorise keys, public keys etc. The commands which will remove sensitive files when this system paramter is enabled are rm cluster instance, rm cluster node, rm ha node, clear config full, join cluster and add cluster instance.
* `restrictedtimeout` - (Optional) Enable/Disable the restricted timeout behaviour. When enabled, timeout cannot be configured beyond admin configured timeout  and also it will have the [minimum - maximum] range check. When disabled, timeout will have the old behaviour. By default the value is disabled
* `strongpassword` - (Optional) After enabling strong password (enableall / enablelocal - not included in exclude list), all the passwords / sensitive information must have - Atleast 1 Lower case character, Atleast 1 Upper case character, Atleast 1 numeric character, Atleast 1 special character ( ~, `, !, @, #, $, %, ^, &, *, -, _, =, +, {, }, [, ], |, \, :, <, >, /, ., ,, " "). Exclude list in case of enablelocal is - NS_FIPS, NS_CRL, NS_RSAKEY, NS_PKCS12, NS_PKCS8, NS_LDAP, NS_TACACS, NS_TACACSACTION, NS_RADIUS, NS_RADIUSACTION, NS_ENCRYPTION_PARAMS. So no Strong Password checks will be performed on these ObjectType commands for enablelocal case.
* `timeout` - (Optional) CLI session inactivity timeout, in seconds. If Restrictedtimeout argument is enabled, Timeout can have values in the range [300-86400] seconds. If Restrictedtimeout argument is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.
* `totalauthtimeout` - (Optional) Total time a request can take for authentication/authorization


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemparameter. It is a unique string prefixed with "tf-systemparameter-".
