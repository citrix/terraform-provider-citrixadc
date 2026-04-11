---
subcategory: "Bot"
---

# Data Source: botprofile

Use this data source to retrieve information about an existing Bot Profile.

The `citrixadc_botprofile` data source allows you to retrieve details of a bot profile by its name. This is useful for referencing existing bot profiles in your Terraform configurations without managing them directly.

## Example usage

```hcl
# Retrieve an existing bot profile
data "citrixadc_botprofile" "example" {
  name = "my_bot_profile"
}

# Use the retrieved profile data in a bot policy
resource "citrixadc_botpolicy" "example_policy" {
  name        = "example_bot_policy"
  profilename = data.citrixadc_botprofile.example.name
  rule        = "true"
}

# Reference profile attributes
output "profile_errorurl" {
  value = data.citrixadc_botprofile.example.errorurl
}

output "profile_sessiontimeout" {
  value = data.citrixadc_botprofile.example.sessiontimeout
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the bot profile (same as name).

* `addcookieflags` - Add the specified flags to bot session cookies. Available settings function as follows:
  * None - Do not add flags to cookies.
  * HTTP Only - Add the HTTP Only flag to cookies, which prevents scripts from accessing cookies.
  * Secure - Add Secure flag to cookies.
  * All - Add both HTTPOnly and Secure flags to cookies.

* `bot_enable_black_list` - Enable black-list bot detection.

* `bot_enable_ip_reputation` - Enable IP-reputation bot detection.

* `bot_enable_rate_limit` - Enable rate-limit bot detection.

* `bot_enable_tps` - Enable TPS.

* `bot_enable_white_list` - Enable white-list bot detection.

* `clientipexpression` - Expression to get the client IP.

* `comment` - Any comments about the purpose of profile, or other useful information about the profile.

* `devicefingerprint` - Enable device-fingerprint bot detection.

* `devicefingerprintaction` - Action to be taken for device-fingerprint based bot detection.

* `devicefingerprintmobile` - Enabling bot device fingerprint protection for mobile clients.

* `dfprequestlimit` - Number of requests to allow without bot session cookie if device fingerprint is enabled.

* `errorurl` - URL that Bot protection uses as the Error URL.

* `headlessbrowserdetection` - Enable Headless Browser detection.

* `kmdetection` - Enable keyboard-mouse based bot detection.

* `kmeventspostbodylimit` - Size of the KM data send by the browser, needs to be processed on ADC.

* `kmjavascriptname` - Name of the JavaScript file that the Bot Management feature will insert in the response for keyboard-mouse based detection. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

* `sessioncookiename` - Name of the SessionCookie that the Bot Management feature uses for tracking. Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.

* `sessiontimeout` - Timeout, in seconds, after which a user session is terminated.

* `signature` - Name of object containing bot static signature details.

* `signaturemultipleuseragentheaderaction` - Actions to be taken if multiple User-Agent headers are seen in a request (Applicable if Signature check is enabled). Log action should be combined with other actions.

* `signaturenouseragentheaderaction` - Actions to be taken if no User-Agent header in the request (Applicable if Signature check is enabled).

* `spoofedreqaction` - Actions to be taken on a spoofed request (A request spoofing good bot user agent string).

* `trap` - Enable trap bot detection.

* `trapaction` - Action to be taken for bot trap based bot detection.

* `trapurl` - URL that Bot protection uses as the Trap URL.

* `verboseloglevel` - Bot verbose Logging. Based on the log level, ADC will log additional information whenever client is detected as a bot.

