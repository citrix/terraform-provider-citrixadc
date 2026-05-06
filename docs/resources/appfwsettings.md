---
subcategory: "Application Firewall"
---

# Resource: appfwsettings

The `appfwsettings` resource is used to configure Application Firewall global settings.


## Example usage

### Using proxypassword (sensitive attribute - persisted in state)

```hcl
variable "appfwsettings_proxypassword" {
  type      = string
  sensitive = true
}

resource "citrixadc_appfwsettings" "example" {
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
  malformedreqaction       = ["block", "log", "stats"]
  centralizedlearning      = "OFF"
  proxyserver              = "192.0.2.10"
  proxyport                = 8080
  proxyusername            = "proxyuser"
  proxypassword            = var.appfwsettings_proxypassword
}
```

### Using proxypassword_wo (write-only/ephemeral - NOT persisted in state)

The `proxypassword_wo` attribute provides an ephemeral path for the proxy password. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `proxypassword_wo_version`.

```hcl
variable "appfwsettings_proxypassword" {
  type      = string
  sensitive = true
}

resource "citrixadc_appfwsettings" "example" {
  defaultprofile           = "APPFW_BYPASS"
  undefaction              = "APPFW_BLOCK"
  sessiontimeout           = 900
  learnratelimit           = 400
  proxyserver              = "192.0.2.10"
  proxyport                = 8080
  proxyusername            = "proxyuser"
  proxypassword_wo         = var.appfwsettings_proxypassword
  proxypassword_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_appfwsettings" "example" {
  defaultprofile           = "APPFW_BYPASS"
  undefaction              = "APPFW_BLOCK"
  sessiontimeout           = 900
  learnratelimit           = 400
  proxyserver              = "192.0.2.10"
  proxyport                = 8080
  proxyusername            = "proxyuser"
  proxypassword_wo         = var.appfwsettings_proxypassword
  proxypassword_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `ceflogging` - (Optional) Enable CEF format logs.
* `centralizedlearning` - (Optional) Flag used to enable/disable ADM centralized learning.
* `clientiploggingheader` - (Optional) Name of an HTTP header that contains the IP address that the client used to connect to the protected web site or service.
* `cookieflags` - (Optional) Add the specified flags to AppFW cookies. Available settings function as follows: None - Do not add flags to AppFW cookies. HTTP Only - Add the HTTP Only flag to AppFW cookies, which prevent scripts from accessing them. Secure - Add Secure flag to AppFW cookies. All - Add both HTTPOnly and Secure flag to AppFW cookies.
* `cookiepostencryptprefix` - (Optional) String that is prepended to all encrypted cookie values.
* `defaultprofile` - (Optional) Profile to use when a connection does not match any policy. Default setting is APPFW_BYPASS, which sends unmatched connections back to the Citrix ADC without attempting to filter them further.
* `entitydecoding` - (Optional) Transform multibyte (double- or half-width) characters to single width characters.
* `geolocationlogging` - (Optional) Enable Geo-Location Logging in CEF format logs.
* `importsizelimit` - (Optional) Maximum cumulative size in bytes of all objects imported to Netscaler. The user is not allowed to import an object if the operation exceeds the currently configured limit.
* `learnratelimit` - (Optional) Maximum number of connections per second that the application firewall learning engine examines to generate new relaxations for learning-enabled security checks. The application firewall drops any connections above this limit from the list of connections used by the learning engine.
* `logmalformedreq` - (Optional) Log requests that are so malformed that application firewall parsing doesn't occur. Note: Changing this attribute forces resource replacement.
* `malformedreqaction` - (Optional) Flag to define action on malformed requests that application firewall cannot parse.
* `proxypassword` - (Optional, Sensitive) Password with which proxy user logs on. The value is persisted in Terraform state (encrypted). See also `proxypassword_wo` for an ephemeral alternative.
* `proxypassword_wo` - (Optional, Sensitive, WriteOnly) Same as `proxypassword`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `proxypassword_wo_version`. If both `proxypassword` and `proxypassword_wo` are set, `proxypassword_wo` takes precedence.
* `proxypassword_wo_version` - (Optional) An integer version tracker for `proxypassword_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `proxyport` - (Optional) Proxy Server Port to get updated signatures from AWS.
* `proxyserver` - (Optional) Proxy Server IP to get updated signatures from AWS.
* `proxyusername` - (Optional) Proxy Username.
* `sessioncookiename` - (Optional) Name of the session cookie that the application firewall uses to track user sessions. Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (\_) symbols. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cookie name" or 'my cookie name').
* `sessionlifetime` - (Optional) Maximum amount of time (in seconds) that the application firewall allows a user session to remain active, regardless of user activity. After this time, the user session is terminated. Before continuing to use the protected web site, the user must establish a new session by opening a designated start URL. A value of 0 represents infinite time.
* `sessionlimit` - (Optional) Maximum number of sessions that the application firewall allows to be active, regardless of user activity. After the max_limit reaches, No more user session will be created.
* `sessiontimeout` - (Optional) Timeout, in seconds, after which a user session is terminated. Before continuing to use the protected web site, the user must establish a new session by opening a designated start URL.
* `signatureautoupdate` - (Optional) Flag used to enable/disable auto update signatures.
* `signatureurl` - (Optional) URL to download the mapping file from server.
* `undefaction` - (Optional) Profile to use when an application firewall policy evaluates to undefined (UNDEF). An UNDEF event indicates an internal error condition. The APPFW_BLOCK built-in profile is the default setting. You can specify a different built-in or user-created profile as the UNDEF profile.
* `useconfigurablesecretkey` - (Optional) Use configurable secret key in AppFw operations.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the appfwsettings resource. It is a unique string prefixed with `appfwsettings-config`.
