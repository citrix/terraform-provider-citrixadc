---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile

The appfwprofile data source allows you to retrieve information about an Application Firewall profile.


## Example usage

```hcl
# Retrieve an existing appfwprofile
data "citrixadc_appfwprofile" "example" {
  name = "my_appfw_profile"
}

# Use the retrieved profile data in a policy
resource "citrixadc_appfwpolicy" "example_policy" {
  name        = "example_policy"
  profilename = data.citrixadc_appfwprofile.example.name
  rule        = "true"
  comment     = "Policy using existing profile"
}

# Reference profile attributes
output "profile_type" {
  value = data.citrixadc_appfwprofile.example.type
}

output "sql_injection_action" {
  value = data.citrixadc_appfwprofile.example.sqlinjectionaction
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the Application Firewall profile to retrieve. Must match an existing profile name.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the Application Firewall profile (same as name).

### Security Check Actions

* `sqlinjectionaction` - SQL Injection actions (Block, Log, Stats, None).
* `crosssitescriptingaction` - Cross-Site Scripting (XSS) actions (Block, Learn, Log, Stats, None).
* `bufferoverflowaction` - Buffer Overflow actions (Block, Log, Stats, None).
* `cookieconsistencyaction` - Cookie Consistency actions (Block, Learn, Log, Stats, None).
* `cookiehijackingaction` - Cookie Hijacking prevention actions (Block, Log, Stats, None).
* `fieldconsistencyaction` - Form Field Consistency actions (Block, Learn, Log, Stats, None).
* `fieldformataction` - Field Format actions (Block, Learn, Log, Stats, None).
* `csrftagaction` - Cross-Site Request Forgery (CSRF) Tagging actions (Block, Learn, Log, Stats, None).
* `creditcardaction` - Credit Card protection actions (Block, Log, Stats, None).
* `contenttypeaction` - Content-Type actions (Block, Learn, Log, Stats, None).
* `starturlaction` - Start URL actions (Block, Learn, Log, Stats, None).
* `denyurlaction` - Deny URL actions (Block, Log, Stats, None).
* `cmdinjectionaction` - Command Injection actions (Block, Log, Stats, None).
* `fileuploadtypesaction` - File Upload Types actions (Block, Learn, Log, Stats, None).
* `blockkeywordaction` - Block Keyword actions (Block, Log, Stats, None).

### XML Security Actions

* `xmlformataction` - XML Format actions (Block, Log, Stats, None).
* `xmlsqlinjectionaction` - XML SQL Injection actions (Block, Log, Stats, None).
* `xmlxssaction` - XML Cross-Site Scripting actions (Block, Log, Stats, None).
* `xmldosaction` - XML Denial of Service actions (Block, Learn, Log, Stats, None).
* `xmlattachmentaction` - XML Attachment actions (Block, Learn, Log, Stats, None).
* `xmlvalidationaction` - XML Validation actions (Block, Log, Stats, None).
* `xmlwsiaction` - Web Services Interoperability (WSI) actions (Block, Learn, Log, Stats, None).
* `xmlsoapfaultaction` - XML SOAP Fault Filtering actions (Block, Log, Stats, Remove, None).

### JSON Security Actions

* `jsondosaction` - JSON Denial of Service actions (Block, Log, Stats, None).
* `jsonsqlinjectionaction` - JSON SQL Injection actions (Block, Log, Stats, None).
* `jsonxssaction` - JSON Cross-Site Scripting actions (Block, Log, Stats, None).
* `jsoncmdinjectionaction` - JSON Command Injection actions (Block, Log, Stats, None).

### Profile Configuration

* `type` - Application Firewall profile types (HTML, XML, JSON, etc.).
* `comment` - Comments about the profile purpose or usage.
* `signatures` - Name of the signature object associated with the profile.
* `apispec` - Name of the API Specification associated with the profile.

### Cookie Settings

* `cookieencryption` - Type of cookie encryption (none, decryptOnly, encryptSessionOnly, encryptAll).
* `cookietransforms` - Enable/disable cookie transformations (ON, OFF).
* `cookieproxying` - Cookie proxy setting (none, sessionOnly).
* `addcookieflags` - Add flags to cookies (none, httpOnly, secure, all).
* `cookiesamesiteattribute` - Cookie SameSite attribute setting.

### Buffer Overflow Settings

* `bufferoverflowmaxurllength` - Maximum length for URLs in characters.
* `bufferoverflowmaxheaderlength` - Maximum length for HTTP headers in characters.
* `bufferoverflowmaxquerylength` - Maximum length for query strings in bytes.
* `bufferoverflowmaxcookielength` - Maximum length for cookies in characters.
* `bufferoverflowmaxtotalheaderlength` - Maximum total HTTP header length in bytes.

### Credit Card Protection

* `creditcard` - Credit card types to protect.
* `creditcardmaxallowed` - Maximum number of credit card numbers allowed on a page.
* `creditcardxout` - Mask credit card numbers in responses (ON, OFF).

### SQL Injection Settings

* `sqlinjectiononlycheckfieldswithsqlchars` - Check only fields with SQL special characters (ON, OFF).
* `sqlinjectiontype` - SQL injection check types (SQLSplChar, SQLKeyword, SQLSplCharANDKeyword, SQLSplCharORKeyword).
* `sqlinjectionchecksqlwildchars` - Check for SQL wildcard characters (ON, OFF).
* `sqlinjectionparsecomments` - Parse and exempt comments from SQL injection checks.

### XSS Settings

* `crosssitescriptingcheckcompleteurls` - Check complete URLs for XSS (ON, OFF).
* `crosssitescriptingtransformunsafehtml` - Transform cross-site scripts instead of blocking (ON, OFF).

### Command Injection Settings

* `cmdinjectiontype` - Command injection check types (CMDSplChar, CMDKeyword, CMDSplCharANDKeyword, CMDSplCharORKeyword, None).
* `cmdinjectiongrammar` - Check for CMD injection using CMD grammar.

### Additional Settings

* `defaults` - Default configuration applied (basic, advanced).
* `refererheadercheck` - Referer header validation setting (OFF, if_present, AlwaysExceptStartURLs, AlwaysExceptFirstRequest).
* `checkrequestheaders` - Check request headers for injected SQL and scripts (ON, OFF).
* `optimizepartialreqs` - Optimize handling of HTTP partial requests (ON, OFF).
* `urldecoderequestcookies` - URL decode request cookies (ON, OFF).
* `canonicalizehtmlresponse` - Perform HTML entity encoding for special characters (ON, OFF).
* `inspectcontenttypes` - Content types to inspect (application/x-www-form-urlencoded, multipart/form-data, text/x-gwt-rpc, none).
* `starturlclosure` - Enable/disable Start URL Closure (ON, OFF).
* `dynamiclearning` - Dynamic learning settings for various security checks.
* `responsecontenttype` - Response content type to be used for enforcement.
* `ceflogging` - Enable CEF format logs (ON, OFF).
* `clientipexpression` - Expression to extract client IP address.
* `multipleheaderaction` - Actions for multiple headers (Block, Log, KeepLast).
* `archivename` - Source for tar archive.
* `overwrite` - Overwrite existing configurations during import.
* `augment` - Augment Relaxation Rules during import.

### Advanced Configuration

* `xmlerrorobject` - Name of the XML Error Object to display when requests are blocked.
* `xmlerrorstatuscode` - Response status code associated with XML error page.
* `xmlerrorstatusmessage` - Response status message associated with XML error page.
* `as_prof_bypass_list_enable` - Enable bypass list for the profile.
* `as_prof_deny_list_enable` - Enable deny list for the profile.

### Read-only appfwprofile metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_appfwprofile` resource). They are GET-only/Computed and are `null` when the appliance does not return them.

* `state` - Enabled state of the profile (for example `ENABLED`, `DISABLED`).
* `learning` - Profile level learning option that overrides the protection level learning (for example `ON`, `OFF`).
* `csrftag` - The web form originating URL.
* `builtin` - Indicates that a profile is a built-in entity.
* `customsettings` - Object name for custom settings. This check is applicable to Profile Type: HTML, XML.
* `defaultcharset` - Default character set for protected web pages. Web pages sent by your protected web sites in response to user requests are assigned this character set if the page does not already specify a character set. The character sets supported by the application firewall are: * iso-8859-1 (English US) * big5 (Chinese Traditional) * gb2312 (Chinese Simplified) * sjis (Japanese Shift-JIS) * euc-jp (Japanese EUC-JP) * iso-8859-9 (Turkish) * utf-8 (Unicode) * euc-kr (Korean)
* `defaultfieldformatmaxlength` - Maximum length, in characters, for data entered into a field that is assigned the default field type.
* `defaultfieldformatmaxoccurrences` - Maxiumum allowed occurrences of the form field name in a request.
* `defaultfieldformatminlength` - Minimum length, in characters, for data entered into a field that is assigned the default field type. To disable the minimum and maximum length settings and allow data of any length to be entered into the field, set this parameter to zero (0).
* `defaultfieldformattype` - Designate a default field type to be applied to web form fields that do not have a field type explicitly assigned to them.
* `dosecurecreditcardlogging` - Setting this option logs credit card numbers in the response when the match is found.
* `enableformtagging` - Enable tagging of web form fields for use by the Form Field Consistency and CSRF Form Tagging checks.
* `errorurl` - URL that application firewall uses as the Error URL.
* `excludefileuploadfromchecks` - Exclude uploaded files from Form checks.
* `exemptclosureurlsfromsecuritychecks` - Exempt URLs that pass the Start URL closure check from SQL injection, cross-site script, field format and field consistency security checks at locations other than headers.
* `fakeaccountdetection` - Fake account detection flag : ON/OFF. If set to ON fake account detection in enabled on ADC, if set to OFF fake account detection is disabled.
* `fieldscan` - Check if formfield limit scan is ON or OFF.
* `fieldscanlimit` - Field scan limit value for HTML
* `fileuploadmaxnum` - Maximum allowed number of file uploads per form-submission request. The maximum setting (65535) allows an unlimited number of uploads.
* `geolocationlogging` - Enable Geo-Location Logging in CEF format logs for the profile.
* `grpcaction` - gRPC validation
* `htmlerrorobject` - Name to assign to the HTML Error Object. Must begin with a letter, number, or the underscore character \(_\), and must contain only letters, numbers, and the hyphen \(-\), period \(.\) pound \(\#\), space \( \), at (@), equals \(=\), colon \(:\), and underscore characters. Cannot be changed after the HTML error object is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my HTML error object" or 'my HTML error object'\).
* `htmlerrorstatuscode` - Response status code associated with HTML error page. Non-empty HTML error object must be imported to the application firewall profile for the status code.
* `htmlerrorstatusmessage` - Response status message associated with HTML error page
* `importprofilename` - Name of the profile which will be created/updated to associate the relaxation rules
* `infercontenttypexmlpayloadaction` - One or more infer content type payload actions. Available settings function as follows: * Block - Block connections that have mismatch in content-type header and payload. * Log - Log connections that have mismatch in content-type header and payload. The mismatched content-type in HTTP request header will be logged for the request. * Stats - Generate statistics when there is mismatch in content-type header and payload. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -inferContentTypeXMLPayloadAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -inferContentTypeXMLPayloadAction none". Please note "none" action cannot be used with any other action type.
* `insertcookiesamesiteattribute` - Configure whether application firewall should add samesite attribute for set-cookies
* `inspectquerycontenttypes` - Inspect request query as well as web forms for injected SQL and cross-site scripts for following content types.
* `invalidpercenthandling` - Configure the method that the application firewall uses to handle percent-encoded names and values. Available settings function as follows: * asp_mode - Microsoft ASP format. * secure_mode - Secure format.
* `jsonblockkeywordaction` - JSON Block Keyword action. Available settings function as follows: * Block - Block connections that violate this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -JSONBlockKeywordAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -JSONBlockKeywordAction none".
* `jsoncmdinjectiongrammar` - Check for CMD injection using CMD grammar in JSON
* `jsoncmdinjectiontype` - Available CMD injection types. -CMDSplChar : Checks for CMD Special Chars -CMDKeyword : Checks for CMD Keywords -CMDSplCharANDKeyword : Checks for both and blocks if both are found -CMDSplCharORKeyword : Checks for both and blocks if anyone is found, -None : Disables checking using both SQL Special Char and Keyword
* `jsonerrorobject` - Name to the imported JSON Error Object to be set on application firewall profile. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my JSON error object" or 'my JSON error object'\).
* `jsonerrorstatuscode` - Response status code associated with JSON error page. Non-empty JSON error object must be imported to the application firewall profile for the status code.
* `jsonerrorstatusmessage` - Response status message associated with JSON error page
* `jsonfieldscan` - Check if JSON field limit scan is ON or OFF.
* `jsonfieldscanlimit` - Field scan limit value for JSON
* `jsonmessagescan` - Check if JSON message limit scan is ON or OFF
* `jsonmessagescanlimit` - Message scan limit value for JSON
* `jsonsqlinjectiongrammar` - Check for SQL injection using SQL grammar in JSON
* `jsonsqlinjectiontype` - Available SQL injection types. -SQLSplChar : Checks for SQL Special Chars -SQLKeyword : Checks for SQL Keywords -SQLSplCharANDKeyword : Checks for both and blocks if both are found -SQLSplCharORKeyword : Checks for both and blocks if anyone is found, -None : Disables checking using both SQL Special Char and Keyword
* `logeverypolicyhit` - Log every profile match, regardless of security checks results.
* `matchurlstring` - Match this action url in archived Relaxation Rules to replace.
* `messagescan` - Check if HTML message limit scan is ON or OFF
* `messagescanlimit` - Message scan limit value for HTML
* `messagescanlimitcontenttypes` - Enable Message Scan Limit for following content types.
* `percentdecoderecursively` - Configure whether the application firewall should use percentage recursive decoding
* `postbodylimit` - Maximum allowed HTTP post body size, in bytes. Maximum supported value is 10GB. Citrix recommends enabling streaming option for large values of post body limit (>20MB).
* `postbodylimitaction` - One or more Post Body Limit actions. Available settings function as follows: * Block - Block connections that violate this security check. Must always be set. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. CLI users: To enable one or more actions, type "set appfw profile -PostBodyLimitAction block" followed by the other actions to be enabled.
* `postbodylimitsignature` - Maximum allowed HTTP post body size for signature inspection for location HTTP_POST_BODY in the signatures, in bytes. Note that the changes in value could impact CPU and latency profile.
* `protofileobject` - Name of the imported proto file.
* `relaxationrules` - Import all appfw relaxation rules
* `replaceurlstring` - Replace matched url string with this action url string while restoring Relaxation Rules
* `requestcontenttype` - Default Content-Type header for requests. A Content-Type header can contain 0-255 letters, numbers, and the hyphen (-) and underscore (_) characters.
* `restaction` - rest validation
* `rfcprofile` - Object name of the rfc profile.
* `semicolonfieldseparator` - Allow ';' as a form field separator in URL queries and POST form bodies.
* `sessioncookiename` - Name of the session cookie that the application firewall uses to track user sessions. Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cookie name" or 'my cookie name').
* `sessionlessfieldconsistency` - Perform sessionless Field Consistency Checks.
* `sessionlessurlclosure` - Enable session less URL Closure Checks. This check is applicable to Profile Type: HTML.
* `sqlinjectiongrammar` - Check for SQL injection using SQL grammar
* `sqlinjectionruletype` - Specifies SQL Injection rule type: ALLOW/DENY. If ALLOW rule type is configured then allow list rules are used, if DENY rule type is configured then deny rules are used.
* `sqlinjectiontransformspecialchars` - Transform injected SQL code. This setting configures the application firewall to disable SQL special strings instead of blocking the request. Since most SQL servers require a special string to activate an SQL keyword, in most cases a request that contains injected SQL code is safe if special strings are disabled. CAUTION: Make sure that this parameter is set to ON if you are configuring any SQL injection transformations. If it is set to OFF, no SQL injection transformations are performed regardless of any other settings.
* `streaming` - Setting this option converts content-length form submission requests (requests with content-type "application/x-www-form-urlencoded" or "multipart/form-data") to chunked requests when atleast one of the following protections : Signatures, SQL injection protection, XSS protection, form field consistency protection, starturl closure, CSRF tagging, JSON SQL, JSON XSS, JSON DOS is enabled. Please make sure that the backend server accepts chunked requests before enabling this option. Citrix recommends enabling this option for large request sizes(>20MB).
* `stripcomments` - Strip HTML comments. This check is applicable to Profile Type: HTML.
* `striphtmlcomments` - Strip HTML comments before forwarding a web page sent by a protected web site in response to a user request.
* `stripxmlcomments` - Strip XML comments before forwarding a web page sent by a protected web site in response to a user request.
* `trace` - Toggle the state of trace
* `usehtmlerrorobject` - Send an imported HTML Error object to a user when a request is blocked, instead of redirecting the user to the designated Error URL.
* `verboseloglevel` - Detailed Logging Verbose Log Level.
* `xmlsqlinjectionchecksqlwildchars` - Check for form fields that contain SQL wild chars .
* `xmlsqlinjectiononlycheckfieldswithsqlchars` - Check only form fields that contain SQL special characters, which most SQL servers require before accepting an SQL command, for injected SQL.
* `xmlsqlinjectionparsecomments` - Parse comments in XML Data and exempt those sections of the request that are from the XML SQL Injection check. You must configure the type of comments that the application firewall is to detect and exempt from this security check. Available settings function as follows: * Check all - Check all content. * ANSI - Exempt content that is part of an ANSI (Mozilla-style) comment. * Nested - Exempt content that is part of a nested (Microsoft-style) comment. * ANSI Nested - Exempt content that is part of any type of comment.
* `xmlsqlinjectiontype` - Available SQL injection types. -SQLSplChar : Checks for SQL Special Chars -SQLKeyword : Checks for SQL Keywords -SQLSplCharANDKeyword : Checks for both and blocks if both are found -SQLSplCharORKeyword : Checks for both and blocks if anyone is found
