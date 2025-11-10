---
subcategory: "Bot"
---

# Resource: botprofile

The Botprofile resource is used to create a collection of profile settings to configure bot management on the appliance.


## Example usage

```hcl
resource "citrixadc_botprofile" "tf_botprofile_name" {
  name                   = "botprofile_name"
  comment                = "My botprofile"
  bot_enable_white_list  = "ON"
  devicefingerprint      = "ON"
}

```


## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile'). Minimum length =  1 Maximum length =  31
* `signature` - (Optional) Name of object containing bot static signature details. Minimum length =  1
* `errorurl` - (Optional) URL that Bot protection uses as the Error URL. Minimum length =  1
* `trapurl` - (Optional) URL that Bot protection uses as the Trap URL. Minimum length =  1 Maximum length =  127
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile. Minimum length =  1
* `bot_enable_white_list` - (Optional) Enable white-list bot detection. Possible values: [ on, off ]
* `bot_enable_black_list` - (Optional) Enable black-list bot detection. Possible values: [ on, off ]
* `bot_enable_rate_limit` - (Optional) Enable rate-limit bot detection. Possible values: [ on, off ]
* `devicefingerprint` - (Optional) Enable device-fingerprint bot detection. Possible values: [ on, off ]
* `devicefingerprintaction` - (Optional) Action to be taken for device-fingerprint based bot detection. Possible values: [ NONE, LOG, DROP, REDIRECT, RESET, MITIGATION ]
* `bot_enable_ip_reputation` - (Optional) Enable IP-reputation bot detection. Possible values: [ on, off ]
* `trap` - (Optional) Enable trap bot detection. Possible values: [ on, off ]
* `trapaction` - (Optional) Action to be taken for bot trap based bot detection. Possible values: [ NONE, LOG, DROP, REDIRECT, RESET ]
* `signaturenouseragentheaderaction` - (Optional) Actions to be taken if no User-Agent header in the request (Applicable if Signature check is enabled). Possible values: [ NONE, LOG, DROP, REDIRECT, RESET ]
* `signaturemultipleuseragentheaderaction` - (Optional) Actions to be taken if multiple User-Agent headers are seen in a request (Applicable if Signature check is enabled). Log action should be combined with other actions. Possible values: [ CHECKLAST, LOG, DROP, REDIRECT, RESET ]
* `bot_enable_tps` - (Optional) Enable TPS. Possible values: [ on, off ]
* `devicefingerprintmobile` - (Optional) Enabling bot device fingerprint protection for mobile clients. Possible values: [ NONE, Android ]
* `clientipexpression` - (Optional) Expression to get the client IP.
* `kmjavascriptname` - (Optional) Name of the JavaScript file that the Bot Management feature will insert in the response for keyboard-mouse based detection. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my javascript file name" or 'my javascript file name').
* `kmdetection` - (Optional) Enable keyboard-mouse based bot detection. Possible values: [ on, off ]
* `kmeventspostbodylimit` - (Optional) Size of the KM data send by the browser, needs to be processed on ADC. Minimum value =  1 Maximum value =  204800
* `addcookieflags` - (Optional) Add the specified flags to bot session cookies. Available settings function as follows: * None - Do not add flags to cookies. * HTTP Only - Add the HTTP Only flag to cookies, which prevents scripts from accessing cookies. * Secure - Add Secure flag to cookies. * All - Add both HTTPOnly and Secure flags to cookies.
* `dfprequestlimit` - (Optional) Number of requests to allow without bot session cookie if device fingerprint is enabled
* `headlessbrowserdetection` - (Optional) Enable Headless Browser detection.
* `sessioncookiename` - (Optional) Name of the SessionCookie that the Bot Management feature uses for tracking. Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cookie name" or 'my cookie name').
* `sessiontimeout` - (Optional) Timeout, in seconds, after which a user session is terminated.
* `spoofedreqaction` - (Optional) Actions to be taken on a spoofed request (A request spoofing good bot user agent string).
* `verboseloglevel` - (Optional) Bot verbose Logging. Based on the log level, ADC will log additional information whenever client is detected as a bot.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile. It has the same value as the `name` attribute.

## Import

A botprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_botprofile.tf_botprofile botprofile_name
```
