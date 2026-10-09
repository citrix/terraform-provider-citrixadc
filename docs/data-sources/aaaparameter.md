---
subcategory: "AAA"
---

# Data Source: aaaparameter

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
* `id` - The id of the aaaparameter. It is a system-generated identifier.
* `apitokencache` - Option to enable/disable API cache feature.
* `classicendpoints` - Parameter to enable/disable classic endpoints.
* `defaultcspheader` - Parameter to enable/disable default CSP header
* `enhancedepa` - Parameter to enable/disable EPA v2 functionality
* `httponlycookie` - Parameter to set/reset HttpOnly Flag for NSC_AAAC/NSC_TMAS cookies in nfactor
* `loginencryption` - Parameter to encrypt login information for nFactor flow
* `maxaaausers` - Maximum number of concurrent users allowed to log on to VPN simultaneously.
* `maxkbquestions` - This will set maximum number of Questions to be asked for KB Validation. Default value is 2, Max Value is 6
* `maxsamldeflatesize` - This will set the maximum deflate size in case of SAML Redirect binding.
* `persistentloginattempts` - Persistent storage of unsuccessful user login attempts
* `pwdexpirynotificationdays` - This will set the threshold time in days for password expiry notification. Default value is 0, which means no notification is sent
* `samesite` - SameSite attribute value for Cookies generated in AAATM context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite
* `securityinsights` - On enabling this option, the Citrix ADC will send the security insight records to the configured collectors when request comes to Authentication endpoint.
  * If cs vserver is frontend with Authentication vserver as target for cs action, then record is sent using Authentication vserver name.
  * If vpn/lb/cs vserver are configured with Authentication ON, then then record is sent using vpn/lb/cs vserver name accordingly.
  * If authentication vserver is frontend, then record is sent using Authentication vserver name.
* `tokenintrospectioninterval` - Frequency at which a token must be verified at the Authorization Server (AS) despite being found in cache.
* `wafprotection` - Entities for which WAF Protection need to be applied.
  Available settings function as follows:
  * DEFAULT - No Endpoint WAF protection.
  * AUTH - Endpoints used for Authentication applicable for both AAATM, IDP, GATEWAY use cases.
  * VPN - Endpoints used for Gateway use cases.
  * PORTAL - Endpoints related to web portal.
  * DISABLED - No Endpoint WAF protection.
  Currently supported only in default partition
* `webviewendpoints` - Parameter to enable/disable webview endpoints.

### Read-only aaaparameter metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_aaaparameter` resource) and are Computed/GET-only. Any attribute the appliance does not return is `null`.

* `builtin` - Flag to determine if aaa param is built-in or not. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ]. A list of strings.
* `feature` - The feature to be checked while applying this config.
