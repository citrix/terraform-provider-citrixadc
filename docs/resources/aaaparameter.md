---
subcategory: "AAA"
---

# Resource: aaaparameter

The aaaparameter resource is used to create aaaparameter.


## Example usage

```hcl
resource "citrixadc_aaaparameter" "tf_aaaparameter" {
  enablestaticpagecaching    = "NO"
  enableenhancedauthfeedback = "YES"
  defaultauthtype            = "LDAP"
  maxaaausers                = 3
  maxloginattempts           = 5
  failedlogintimeout         = 15
}
```


## Argument Reference

* `enablestaticpagecaching` - (Optional) The default state of VPN Static Page caching. If nothing is specified, the default value is set to YES. Possible values: [ YES, NO ]
* `enableenhancedauthfeedback` - (Optional) Enhanced auth feedback provides more information to the end user about the reason for an authentication failure.  The default value is set to NO. Possible values: [ YES, NO ]
* `defaultauthtype` - (Optional) The default authentication server type. Possible values: [ LOCAL, LDAP, RADIUS, TACACS, CERT ]
* `maxaaausers` - (Optional) Maximum number of concurrent users allowed to log on to VPN simultaneously. Minimum value =  1
* `maxloginattempts` - (Optional) Maximum Number of login Attempts. Minimum value =  1
* `failedlogintimeout` - (Optional) Number of minutes an account will be locked if user exceeds maximum permissible attempts. Minimum value =  1 Maximum value =  525600
* `aaadnatip` - (Optional) Source IP address to use for traffic that is sent to the authentication server.
* `enablesessionstickiness` - (Optional) Enables/Disables stickiness to authentication servers. Possible values: [ YES, NO ]
* `aaasessionloglevel` - (Optional) Audit log level, which specifies the types of events to log for cli executed commands. Available values function as follows: * EMERGENCY - Events that indicate an immediate crisis on the server. * ALERT - Events that might require action. * CRITICAL - Events that indicate an imminent server crisis. * ERROR - Events that indicate some type of error. * WARNING - Events that require action in the near future. * NOTICE - Events that the administrator should know about. * INFORMATIONAL - All but low-level events. * DEBUG - All events, in extreme detail. Possible values: [ EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG ]
* `aaadloglevel` - (Optional) AAAD log level, which specifies the types of AAAD events to log in nsvpn.log. Available values function as follows: * EMERGENCY - Events that indicate an immediate crisis on the server. * ALERT - Events that might require action. * CRITICAL - Events that indicate an imminent server crisis. * ERROR - Events that indicate some type of error. * WARNING - Events that require action in the near future. * NOTICE - Events that the administrator should know about. * INFORMATIONAL - All but low-level events. * DEBUG - All events, in extreme detail. Possible values: [ EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG ]
* `dynaddr` - (Optional) Set by the DHCP client when the IP address was fetched dynamically. Possible values: [ on, off ]
* `ftmode` - (Optional) First time user mode determines which configuration options are shown by default when logging in to the GUI. This setting is controlled by the GUI. Possible values: [ ON, HA, OFF ]
* `maxsamldeflatesize` - (Optional) This will set the maximum deflate size in case of SAML Redirect binding.
* `persistentloginattempts` - (Optional) Persistent storage of unsuccessful user login attempts. Possible values: [ ENABLED, DISABLED ]
* `pwdexpirynotificationdays` - (Optional) This will set the threshold time in days for password expiry notification. Default value is 0, which means no notification is sent.
* `maxkbquestions` - (Optional) This will set maximum number of Questions to be asked for KB Validation. Default value is 2, Max Value is 6. Minimum value =  2 Maximum value =  6
* `loginencryption` - (Optional) Parameter to encrypt login information for nFactor flow. Possible values: [ ENABLED, DISABLED ]
* `samesite` - (Optional) SameSite attribute value for Cookies generated in AAATM context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite. Possible values: [ None, LAX, STRICT ]
* `apitokencache` - (Optional) Option to enable/disable API cache feature. Possible values: [ ENABLED, DISABLED ]
* `tokenintrospectioninterval` - (Optional) Frequency at which a token must be verified at the Authorization Server (AS) despite being found in cache.
* `defaultcspheader` - (Optional) Parameter to enable/disable default CSP header. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaaparameter. It is a unique string prefixed with `tf-aaaparameter-`.