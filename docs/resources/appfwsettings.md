---
subcategory: "Application Firewall"
---

# Resource: appfwsettings

The appfwsettings resource is used to create configure aplication firewall settings.


## Example usage

```hcl
resource "citrixadc_appfwsettings" "tf_appfwsettings" {
  defaultprofile           = "APPFW_BYPASS"
  undefaction              = "APPFW_BLOCK"
  sessiontimeout           = 900
  learnratelimit           = 400
  sessionlifetime          = 0
  sessioncookiename        = "citrix_ns_id"
  importsizelimit          = 134217728
  signatureautoupdate      = "OFF"
  signatureurl             = "https://s3.amazonaws.com/NSAppFwSignatures/SignaturesMapping.xml"
  cookiepostencryptprefix  = "ENC"
  geolocationlogging       = "OFF"
  ceflogging               = "OFF"
  entitydecoding           = "OFF"
  useconfigurablesecretkey = "OFF"
  sessionlimit             = 100000
  malformedreqaction = [
    "block",
    "log",
    "stats"
  ]
  centralizedlearning = "OFF"
  proxyport           = 8080
}
```


## Argument Reference

* `ceflogging` - (Optional) Enable CEF format logs.
* `centralizedlearning` - (Optional) Flag used to enable/disable ADM centralized learning
* `clientiploggingheader` - (Optional) Name of an HTTP header that contains the IP address that the client used to connect to the protected web site or service.
* `cookiepostencryptprefix` - (Optional) String that is prepended to all encrypted cookie values.
* `defaultprofile` - (Optional) Profile to use when a connection does not match any policy. Default setting is APPFW_BYPASS, which sends unmatched connections back to the Citrix ADC without attempting to filter them further.
* `entitydecoding` - (Optional) Transform multibyte (double- or half-width) characters to single width characters.
* `geolocationlogging` - (Optional) Enable Geo-Location Logging in CEF format logs.
* `importsizelimit` - (Optional) Cumulative total maximum number of bytes in web forms imported to a protected web site. If a user attempts to upload files with a total byte count higher than the specified limit, the application firewall blocks the request.
* `learnratelimit` - (Optional) Maximum number of connections per second that the application firewall learning engine examines to generate new relaxations for learning-enabled security checks. The application firewall drops any connections above this limit from the list of connections used by the learning engine.
* `logmalformedreq` - (Optional) Log requests that are so malformed that application firewall parsing doesn't occur.
* `malformedreqaction` - (Optional) flag to define action on malformed requests that application firewall cannot parse
* `proxyport` - (Optional) Proxy Server Port to get updated signatures from AWS.
* `proxyserver` - (Optional) Proxy Server IP to get updated signatures from AWS.
* `sessioncookiename` - (Optional) Name of the session cookie that the application firewall uses to track user sessions.  Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (\_) symbols.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cookie name" or 'my cookie name').
* `sessionlifetime` - (Optional) Maximum amount of time (in seconds) that the application firewall allows a user session to remain active, regardless of user activity. After this time, the user session is terminated. Before continuing to use the protected web site, the user must establish a new session by opening a designated start URL.
* `sessionlimit` - (Optional) Maximum number of sessions that the application firewall allows to be active, regardless of user activity. After the max_limit reaches, No more user session will be created .
* `sessiontimeout` - (Optional) Timeout, in seconds, after which a user session is terminated. Before continuing to use the protected web site, the user must establish a new session by opening a designated start URL.
* `signatureautoupdate` - (Optional) Flag used to enable/disable auto update signatures
* `signatureurl` - (Optional) URL to download the mapping file from server
* `undefaction` - (Optional) Profile to use when an application firewall policy evaluates to undefined (UNDEF).  An UNDEF event indicates an internal error condition. The APPFW_BLOCK built-in profile is the default setting. You can specify a different built-in or user-created profile as the UNDEF profile.
* `useconfigurablesecretkey` - (Optional) Use configurable secret key in AppFw operations

* `proxyusername` - (Optional) Proxy Username for signature updates.
* `proxypassword` - (Optional, Sensitive) Password with which proxy user logs on.
* `cookieflags` - (Optional) Add the specified flags to AppFW cookies. Available settings: None, HTTP Only, Secure, All.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwsettings. It is a unique string prefixed with "tf-appfwsettings-".
