---
subcategory: "Application Firewall"
---

# Resource: appfwprofile

The `appfwprofile` resource is used to create Applicatin Firewall Profile.

## Example usage

``` hcl
resource "citrixadc_appfwprofile" "tf_appfwprofile1" {
  name = "tf_appfwprofile1"
  sqlinjectionaction = [
    "block",
    "log",
    "stats",
  ]
  type = [
    "HTML",
    "JSON",
    "XML",
  ]
}

```

## Argument Reference

* `name` - Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `postbodylimitaction` One or more Post Body Limit actions. Available settings function as follows: Block - Block connections that violate this security check. Must always be set. Log - Log violations of this security check. Stats - Generate statistics for this security check. Possible values: [block, log, stats]
* `bufferoverflowmaxquerylength` Maximum length, in bytes, for query string sent to your protected web sites. Requests with longer query strings are blocked.Default value: 65535 Minimum value = 0 Maximum value = 65535
* `cookiehijackingaction` One or more actions to prevent cookie hijacking. Available settings function as follows:Block - Block connections that violate this security check. Log - Log violations of this security check. Stats - Generate statistics for this security check. None - Disable all actions for this security check. NOTE: Cookie Hijacking feature is not supported for TLSv1.3. Possible values: [none, block, log, stats]
* `infercontenttypexmlpayloadaction` One or more infer content type payload actions. Available settings function as follows: Block - Block connections that have mismatch in content-type header and payload. Log - Log connections that have mismatch in content-type header and payload. The mismatched content-type in HTTP request header will be logged for the request. Stats - Generate statistics when there is mismatch in content-type header and payload. None - Disable all actions for this security check. Possible values: [ block, log, stats, none ]
* `cmdinjectionaction` Command injection action. Available settings function as follows: Block - Block connections that violate this security check. Log - Log violations of this security check. Stats - Generate statistics for this security check. None - Disable all actions for this security check. Possible values: [none, block, log, stats]
* `defaults` - (Optional) Default configuration to apply to the profile. Basic defaults are intended for standard content that requires little further configuration, such as static web site content. Advanced defaults are intended for specialized content that requires significant specialized configuration, such as heavily scripted or dynamic content. CLI users: When adding an application firewall profile, you can set either the defaults or the type, but not both. To set both options, create the profile by using the add appfw profile command, and then use the set appfw profile command to configure the other option. Possible values: [ basic, advanced ]
* `starturlaction` - (Optional) One or more Start URL actions. Available settings function as follows: * Block - Block connections that violate this security check. * Learn - Use the learning engine to generate a list of exceptions to this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -startURLaction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -startURLaction none". Possible values: [ none, block, learn, log, stats ]
* `contenttypeaction` - (Optional) One or more Content-type actions. Available settings function as follows: * Block - Block connections that violate this security check. * Learn - Use the learning engine to generate a list of exceptions to this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -contentTypeaction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -contentTypeaction none". Possible values: [ none, block, learn, log, stats ]
* `inspectcontenttypes` - (Optional) One or more InspectContentType lists. * application/x-www-form-urlencoded * multipart/form-data * text/x-gwt-rpc CLI users: To enable, type "set appfw profile -InspectContentTypes" followed by the content types to be inspected. Possible values: [ none, application/x-www-form-urlencoded, multipart/form-data, text/x-gwt-rpc ]
* `starturlclosure` - (Optional) Toggle  the state of Start URL Closure. Possible values: [ on, off ]
* `denyurlaction` - (Optional) One or more Deny URL actions. Available settings function as follows: * Block - Block connections that violate this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. NOTE: The Deny URL check takes precedence over the Start URL check. If you enable blocking for the Deny URL check, the application firewall blocks any URL that is explicitly blocked by a Deny URL, even if the same URL would otherwise be allowed by the Start URL check. CLI users: To enable one or more actions, type "set appfw profile -denyURLaction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -denyURLaction none". Possible values: [ none, block, log, stats ]
* `refererheadercheck` - (Optional) Enable validation of Referer headers. Referer validation ensures that a web form that a user sends to your web site originally came from your web site, not an outside attacker. Although this parameter is part of the Start URL check, referer validation protects against cross-site request forgery (CSRF) attacks, not Start URL attacks. Possible values: [ OFF, if_present, AlwaysExceptStartURLs, AlwaysExceptFirstRequest ]
* `cookieconsistencyaction` - (Optional) One or more Cookie Consistency actions. Available settings function as follows: * Block - Block connections that violate this security check. * Learn - Use the learning engine to generate a list of exceptions to this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -cookieConsistencyAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -cookieConsistencyAction none". Possible values: [ none, block, learn, log, stats ]
* `cookietransforms` - (Optional) Perform the specified type of cookie transformation. Available settings function as follows: * Encryption - Encrypt cookies. * Proxying - Mask contents of server cookies by sending proxy cookie to users. * Cookie flags - Flag cookies as HTTP only to prevent scripts on user's browser from accessing and possibly modifying them. CAUTION: Make sure that this parameter is set to ON if you are configuring any cookie transformations. If it is set to OFF, no cookie transformations are performed regardless of any other settings. Possible values: [ on, off ]
* `cookieencryption` - (Optional) Type of cookie encryption. Available settings function as follows: * None - Do not encrypt cookies. * Decrypt Only - Decrypt encrypted cookies, but do not encrypt cookies. * Encrypt Session Only - Encrypt session cookies, but not permanent cookies. * Encrypt All - Encrypt all cookies. Possible values: [ none, decryptOnly, encryptSessionOnly, encryptAll ]
* `cookieproxying` - (Optional) Cookie proxy setting. Available settings function as follows: * None - Do not proxy cookies. * Session Only - Proxy session cookies by using the Citrix ADC session ID, but do not proxy permanent cookies. Possible values: [ none, sessionOnly ]
* `addcookieflags` - (Optional) Add the specified flags to cookies. Available settings function as follows: * None - Do not add flags to cookies. * HTTP Only - Add the HTTP Only flag to cookies, which prevents scripts from accessing cookies. * Secure - Add Secure flag to cookies. * All - Add both HTTPOnly and Secure flags to cookies. Possible values: [ none, httpOnly, secure, all ]
* `fieldconsistencyaction` - (Optional) One or more Form Field Consistency actions. Available settings function as follows: * Block - Block connections that violate this security check. * Learn - Use the learning engine to generate a list of exceptions to this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -fieldConsistencyaction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -fieldConsistencyAction none". Possible values: [ none, block, learn, log, stats ]
* `csrftagaction` - (Optional) One or more Cross-Site Request Forgery (CSRF) Tagging actions. Available settings function as follows: * Block - Block connections that violate this security check. * Learn - Use the learning engine to generate a list of exceptions to this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -CSRFTagAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -CSRFTagAction none". Possible values: [ none, block, learn, log, stats ]
* `crosssitescriptingaction` - (Optional) One or more Cross-Site Scripting (XSS) actions. Available settings function as follows: * Block - Block connections that violate this security check. * Learn - Use the learning engine to generate a list of exceptions to this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -crossSiteScriptingAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -crossSiteScriptingAction none". Possible values: [ none, block, learn, log, stats ]
* `crosssitescriptingtransformunsafehtml` - (Optional) Transform cross-site scripts. This setting configures the application firewall to disable dangerous HTML instead of blocking the request. CAUTION: Make sure that this parameter is set to ON if you are configuring any cross-site scripting transformations. If it is set to OFF, no cross-site scripting transformations are performed regardless of any other settings. Possible values: [ on, off ]
* `crosssitescriptingcheckcompleteurls` - (Optional) Check complete URLs for cross-site scripts, instead of just the query portions of URLs. Possible values: [ on, off ]
* `sqlinjectionaction` - (Optional) One or more HTML SQL Injection actions. Available settings function as follows: * Block - Block connections that violate this security check. * Learn - Use the learning engine to generate a list of exceptions to this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -SQLInjectionAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -SQLInjectionAction none". Possible values: [ none, block, learn, log, stats ]
* `sqlinjectiontransformspecialchars` - (Optional) Transform injected SQL code. This setting configures the application firewall to disable SQL special strings instead of blocking the request. Since most SQL servers require a special string to activate an SQL keyword, in most cases a request that contains injected SQL code is safe if special strings are disabled. CAUTION: Make sure that this parameter is set to ON if you are configuring any SQL injection transformations. If it is set to OFF, no SQL injection transformations are performed regardless of any other settings. Possible values: [ on, off ]
* `sqlinjectiononlycheckfieldswithsqlchars` - (Optional) Check only form fields that contain SQL special strings (characters) for injected SQL code. Most SQL servers require a special string to activate an SQL request, so SQL code without a special string is harmless to most SQL servers. Possible values: [ on, off ]
* `sqlinjectiontype` - (Optional) Available SQL injection types. -SQLSplChar              : Checks for SQL Special Chars -SQLKeyword		 : Checks for SQL Keywords -SQLSplCharANDKeyword    : Checks for both and blocks if both are found -SQLSplCharORKeyword     : Checks for both and blocks if anyone is found. Possible values: [ SQLSplChar, SQLKeyword, SQLSplCharORKeyword, SQLSplCharANDKeyword ]
* `sqlinjectionchecksqlwildchars` - (Optional) Check for form fields that contain SQL wild chars . Possible values: [ on, off ]
* `fieldformataction` - (Optional) One or more Field Format actions. Available settings function as follows: * Block - Block connections that violate this security check. * Learn - Use the learning engine to generate a list of suggested web form fields and field format assignments. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -fieldFormatAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -fieldFormatAction none". Possible values: [ none, block, learn, log, stats ]
* `defaultfieldformattype` - (Optional) Designate a default field type to be applied to web form fields that do not have a field type explicitly assigned to them.
* `defaultfieldformatminlength` - (Optional) To disable the minimum and maximum length settings and allow data of any length to be entered into the field, set this parameter to zero (0).
* `defaultfieldformatmaxlength` - (Optional)
* `bufferoverflowaction` - (Optional) One or more Buffer Overflow actions. Available settings function as follows: * Block - Block connections that violate this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -bufferOverflowAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -bufferOverflowAction none". Possible values: [ none, block, log, stats ]
* `bufferoverflowmaxurllength` - (Optional)
* `bufferoverflowmaxheaderlength` - (Optional)
* `bufferoverflowmaxcookielength` - (Optional)
* `creditcardaction` - (Optional) One or more Credit Card actions. Available settings function as follows: * Block - Block connections that violate this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -creditCardAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -creditCardAction none". Possible values: [ none, block, learn, log, stats ]
* `creditcard` - (Optional) Credit card types that the application firewall should protect. Possible values: [ none, visa, mastercard, discover, amex, jcb, dinersclub ]
* `creditcardmaxallowed` - (Optional) This parameter value is used by the block action. It represents the maximum number of credit card numbers that can appear on a web page served by your protected web sites. Pages that contain more credit card numbers are blocked.
* `creditcardxout` - (Optional) Mask any credit card number detected in a response by replacing each digit, except the digits in the final group, with the letter "X.". Possible values: [ on, off ]
* `dosecurecreditcardlogging` - (Optional) Setting this option logs credit card numbers in the response when the match is found. Possible values: [ on, off ]
* `streaming` - (Optional) Setting this option converts content-length form submission requests (requests with content-type "application/x-www-form-urlencoded" or "multipart/form-data") to chunked requests when atleast one of the following protections : SQL injection protection, XSS protection, form field consistency protection, starturl closure, CSRF tagging is enabled. Please make sure that the backend server accepts chunked requests before enabling this option. Possible values: [ on, off ]
* `trace` - (Optional) Toggle  the state of trace. Possible values: [ on, off ]
* `requestcontenttype` - (Optional) Default Content-Type header for requests. A Content-Type header can contain 0-255 letters, numbers, and the hyphen (-) and underscore (_) characters.
* `responsecontenttype` - (Optional) Default Content-Type header for responses. A Content-Type header can contain 0-255 letters, numbers, and the hyphen (-) and underscore (_) characters.
* `jsonerrorobject` - (Optional) Name to the imported JSON Error Object to be set on application firewall profile. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my JSON error object" or 'my JSON error object'\).
* `jsondosaction` - (Optional) One or more JSON Denial-of-Service (JsonDoS) actions. Available settings function as follows: * Block - Block connections that violate this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -JSONDoSAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -JSONDoSAction none". Possible values: [ none, block, log, stats ]
* `jsonsqlinjectionaction` - (Optional) One or more JSON SQL Injection actions. Available settings function as follows: * Block - Block connections that violate this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -JSONSQLInjectionAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -JSONSQLInjectionAction none". Possible values: [ none, block, log, stats ]
* `jsonsqlinjectiontype` - (Optional) Available SQL injection types. -SQLSplChar              : Checks for SQL Special Chars -SQLKeyword              : Checks for SQL Keywords -SQLSplCharANDKeyword    : Checks for both and blocks if both are found -SQLSplCharORKeyword     : Checks for both and blocks if anyone is found. Possible values: [ SQLSplChar, SQLKeyword, SQLSplCharORKeyword, SQLSplCharANDKeyword ]
* `jsonxssaction` - (Optional) One or more JSON Cross-Site Scripting actions. Available settings function as follows: * Block - Block connections that violate this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -JSONXssAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -JSONXssAction none". Possible values: [ none, block, log, stats ]
* `xmldosaction` - (Optional) One or more XML Denial-of-Service (XDoS) actions. Available settings function as follows: * Block - Block connections that violate this security check. * Learn - Use the learning engine to generate a list of exceptions to this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -XMLDoSAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLDoSAction none". Possible values: [ none, block, learn, log, stats ]
* `xmlformataction` - (Optional) One or more XML Format actions. Available settings function as follows: * Block - Block connections that violate this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -XMLFormatAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLFormatAction none". Possible values: [ none, block, log, stats ]
* `xmlsqlinjectionaction` - (Optional) One or more XML SQL Injection actions. Available settings function as follows: * Block - Block connections that violate this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -XMLSQLInjectionAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLSQLInjectionAction none". Possible values: [ none, block, log, stats ]
* `xmlsqlinjectiononlycheckfieldswithsqlchars` - (Optional) Check only form fields that contain SQL special characters, which most SQL servers require before accepting an SQL command, for injected SQL. Possible values: [ on, off ]
* `xmlsqlinjectiontype` - (Optional) Available SQL injection types. -SQLSplChar              : Checks for SQL Special Chars -SQLKeyword              : Checks for SQL Keywords -SQLSplCharANDKeyword    : Checks for both and blocks if both are found -SQLSplCharORKeyword     : Checks for both and blocks if anyone is found. Possible values: [ SQLSplChar, SQLKeyword, SQLSplCharORKeyword, SQLSplCharANDKeyword ]
* `xmlsqlinjectionchecksqlwildchars` - (Optional) Check for form fields that contain SQL wild chars . Possible values: [ on, off ]
* `xmlsqlinjectionparsecomments` - (Optional) Parse comments in XML Data and exempt those sections of the request that are from the XML SQL Injection check. You must configure the type of comments that the application firewall is to detect and exempt from this security check. Available settings function as follows: * Check all - Check all content. * ANSI - Exempt content that is part of an ANSI (Mozilla-style) comment. * Nested - Exempt content that is part of a nested (Microsoft-style) comment. * ANSI Nested - Exempt content that is part of any type of comment. Possible values: [ checkall, ansi, nested, ansinested ]
* `xmlxssaction` - (Optional) One or more XML Cross-Site Scripting actions. Available settings function as follows: * Block - Block connections that violate this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -XMLXSSAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLXSSAction none". Possible values: [ none, block, learn, log, stats ]
* `xmlwsiaction` - (Optional) One or more Web Services Interoperability (WSI) actions. Available settings function as follows: * Block - Block connections that violate this security check. * Learn - Use the learning engine to generate a list of exceptions to this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -XMLWSIAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLWSIAction none". Possible values: [ none, block, learn, log, stats ]
* `xmlattachmentaction` - (Optional) One or more XML Attachment actions. Available settings function as follows: * Block - Block connections that violate this security check. * Learn - Use the learning engine to generate a list of exceptions to this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -XMLAttachmentAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLAttachmentAction none". Possible values: [ none, block, learn, log, stats ]
* `xmlvalidationaction` - (Optional) One or more XML Validation actions. Available settings function as follows: * Block - Block connections that violate this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -XMLValidationAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLValidationAction none". Possible values: [ none, block, log, stats ]
* `xmlerrorobject` - (Optional) Name to assign to the XML Error Object, which the application firewall displays when a user request is blocked. Must begin with a letter, number, or the underscore character \(_\), and must contain only letters, numbers, and the hyphen \(-\), period \(.\) pound \(\#\), space \( \), at (@), equals \(=\), colon \(:\), and underscore characters. Cannot be changed after the XML error object is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my XML error object" or 'my XML error object'\).
* `customsettings` - (Optional) Object name for custom settings. This check is applicable to Profile Type: HTML, XML. .
* `signatures` - (Optional) Object name for signatures. This check is applicable to Profile Type: HTML, XML. .
* `xmlsoapfaultaction` - (Optional) One or more XML SOAP Fault Filtering actions. Available settings function as follows: * Block - Block connections that violate this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. * Remove - Remove all violations for this security check. CLI users: To enable one or more actions, type "set appfw profile -XMLSOAPFaultAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLSOAPFaultAction none". Possible values: [ none, block, log, remove, stats ]
* `usehtmlerrorobject` - (Optional) Send an imported HTML Error object to a user when a request is blocked, instead of redirecting the user to the designated Error URL. Possible values: [ on, off ]
* `errorurl` - (Optional) URL that application firewall uses as the Error URL.
* `htmlerrorobject` - (Optional) Name to assign to the HTML Error Object. Must begin with a letter, number, or the underscore character \(_\), and must contain only letters, numbers, and the hyphen \(-\), period \(.\) pound \(\#\), space \( \), at (@), equals \(=\), colon \(:\), and underscore characters. Cannot be changed after the HTML error object is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my HTML error object" or 'my HTML error object'\).
* `logeverypolicyhit` - (Optional) Log every profile match, regardless of security checks results. Possible values: [ on, off ]
* `stripcomments` - (Optional) Strip HTML comments. This check is applicable to Profile Type: HTML. . Possible values: [ on, off ]
* `striphtmlcomments` - (Optional) Strip HTML comments before forwarding a web page sent by a protected web site in response to a user request. Possible values: [ none, all, exclude_script_tag ]
* `stripxmlcomments` - (Optional) Strip XML comments before forwarding a web page sent by a protected web site in response to a user request. Possible values: [ none, all ]
* `exemptclosureurlsfromsecuritychecks` - (Optional) Exempt URLs that pass the Start URL closure check from SQL injection, cross-site script, field format and field consistency security checks at locations other than headers. Possible values: [ on, off ]
* `defaultcharset` - (Optional) Default character set for protected web pages. Web pages sent by your protected web sites in response to user requests are assigned this character set if the page does not already specify a character set. The character sets supported by the application firewall are: * iso-8859-1 (English US) * big5 (Chinese Traditional) * gb2312 (Chinese Simplified) * sjis (Japanese Shift-JIS) * euc-jp (Japanese EUC-JP) * iso-8859-9 (Turkish) * utf-8 (Unicode) * euc-kr (Korean).
* `dynamiclearning` - (Optional) One or more security checks. Available options are as follows: * SQLInjection - Enable dynamic learning for SQLInjection security check. * CrossSiteScripting - Enable dynamic learning for CrossSiteScripting security check. * fieldFormat - Enable dynamic learning for  fieldFormat security check. * None - Disable security checks for all security checks. CLI users: To enable dynamic learning on one or more security checks, type "set appfw profile -dynamicLearning" followed by the security checks to be enabled. To turn off dynamic learning on all security checks, type "set appfw profile -dynamicLearning none". Possible values: [ none, SQLInjection, CrossSiteScripting, fieldFormat, startURL, cookieConsistency, fieldConsistency, CSRFtag, ContentType ]
* `postbodylimit` - (Optional)
* `postbodylimitsignature` - (Optional)
* `fileuploadmaxnum` - (Optional)
* `canonicalizehtmlresponse` - (Optional) Perform HTML entity encoding for any special characters in responses sent by your protected web sites. Possible values: [ on, off ]
* `enableformtagging` - (Optional) Enable tagging of web form fields for use by the Form Field Consistency and CSRF Form Tagging checks. Possible values: [ on, off ]
* `sessionlessfieldconsistency` - (Optional) Perform sessionless Field Consistency Checks. Possible values: [ OFF, ON, postOnly ]
* `sessionlessurlclosure` - (Optional) Enable session less URL Closure Checks. This check is applicable to Profile Type: HTML. . Possible values: [ on, off ]
* `semicolonfieldseparator` - (Optional) Allow '; ' as a form field separator in URL queries and POST form bodies. . Possible values: [ on, off ]
* `excludefileuploadfromchecks` - (Optional) Exclude uploaded files from Form checks. Possible values: [ on, off ]
* `sqlinjectionparsecomments` - (Optional) Parse HTML comments and exempt them from the HTML SQL Injection check. You must specify the type of comments that the application firewall is to detect and exempt from this security check. Available settings function as follows: * Check all - Check all content. * ANSI - Exempt content that is part of an ANSI (Mozilla-style) comment. * Nested - Exempt content that is part of a nested (Microsoft-style) comment. * ANSI Nested - Exempt content that is part of any type of comment. Possible values: [ checkall, ansi, nested, ansinested ]
* `invalidpercenthandling` - (Optional) Configure the method that the application firewall uses to handle percent-encoded names and values. Available settings function as follows: * apache_mode - Apache format. * asp_mode - Microsoft ASP format. * secure_mode - Secure format. Possible values: [ apache_mode, asp_mode, secure_mode ]
* `type` - (Optional) Application firewall profile type, which controls which security checks and settings are applied to content that is filtered with the profile. Available settings function as follows: * HTML - HTML-based web sites. * XML -  XML-based web sites and services. * JSON - JSON-based web sites and services. * HTML XML (Web 2.0) - Sites that contain both HTML and XML content, such as ATOM feeds, blogs, and RSS feeds. * HTML JSON  - Sites that contain both HTML and JSON content. * XML JSON   - Sites that contain both XML and JSON content. * HTML XML JSON   - Sites that contain HTML, XML and JSON content. Possible values: [ HTML, XML, JSON ]
* `checkrequestheaders` - (Optional) Check request headers as well as web forms for injected SQL and cross-site scripts. Possible values: [ on, off ]
* `optimizepartialreqs` - (Optional) Optimize handle of HTTP partial requests i.e. those with range headers. Available settings are as follows: * ON  - Partial requests by the client result in partial requests to the backend server in most cases. * OFF - Partial requests by the client are changed to full requests to the backend server. Possible values: [ on, off ]
* `urldecoderequestcookies` - (Optional) URL Decode request cookies before subjecting them to SQL and cross-site scripting checks. Possible values: [ on, off ]
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `percentdecoderecursively` - (Optional) Configure whether the application firewall should use percentage recursive decoding. Possible values: [ on, off ]
* `multipleheaderaction` - (Optional) One or more multiple header actions. Available settings function as follows: * Block - Block connections that have multiple headers. * Log - Log connections that have multiple headers. * KeepLast - Keep only last header when multiple headers are present. CLI users: To enable one or more actions, type "set appfw profile -multipleHeaderAction" followed by the actions to be enabled. Possible values: [ block, keepLast, log, none ]
* `rfcprofile` - (Optional) Object name of the rfc profile.
* `fileuploadtypesaction` - (Optional) One or more file upload types actions. Available settings function as follows: * Block - Block connections that violate this security check. * Log - Log violations of this security check. * Stats - Generate statistics for this security check. * None - Disable all actions for this security check. CLI users: To enable one or more actions, type "set appfw profile -fileUploadTypeAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -fileUploadTypeAction none". Possible values: [ none, block, log, stats ]
* `verboseloglevel` - (Optional) Detailed Logging Verbose Log Level. Possible values: [ pattern, patternPayload, patternPayloadHeader ]
* `archivename` - (Optional) Source for tar archive.
* `htmlerrorstatuscode` - (Optional) Response status code associated with HTML error page. Non-empty HTML error object must be imported to the application firewall profile for the status code.
    * Default value: 200
    * Minimum value = 1
    * Maximum value = 999
* `htmlerrorstatusmessage` - (Optional) Response status message associated with HTML error page.
* `bufferoverflowmaxtotalheaderlength` - (Optional) Maximum length, in bytes, for the total HTTP header length in requests sent to your protected web sites. The minimum value of this and maxHeaderLen in httpProfile will be used. Requests with longer length are blocked.
    * Default value: 65535
    * Minimum value = 0
    * Maximum value = 65535
* `sqlinjectiongrammar` - (Optional) Check for SQL injection using SQL grammar. Default value: OFF  Possible values = ON, OFF
* `cmdinjectiontype` - (Optional) Available CMD injection types.
    * Default value: CMDSplCharANDKeyword
    * Possible values = CMDSplChar, CMDKeyword, CMDSplCharORKeyword, CMDSplCharANDKeyword, None
    * CMDSplChar : Checks for CMD Special Chars
    * CMDKeyword : Checks for CMD Keywords
    * CMDSplCharANDKeyword : Checks for both and blocks if both are found
    * CMDSplCharORKeyword : Checks for both and blocks if anyone is found,
    * None : Disables checking using both CMD Special Char and Keyword.
* `apispec` - (Optional) Name of the API Specification.
* `as_prof_bypass_list_enable` - (Optional) Enable bypass list for the profile.
* `as_prof_deny_list_enable` - (Optional) Enable deny list for the profile.
* `augment` - (Optional) Augment Relaxation Rules during import.
* `blockkeywordaction` - (Optional) Block Keyword action. Available settings function as follows:
    * Block - Block connections that violate this security check.
    * Log - Log violations of this security check.
    * Stats - Generate statistics for this security check.
    * None - Disable all actions for this security check.
    Possible values: [ none, block, log, stats ]
* `ceflogging` - (Optional) Enable CEF format logs for the profile.
* `clientipexpression` - (Optional) Expression to get the client IP.
* `cmdinjectiongrammar` - (Optional) Check for CMD injection using CMD grammar.
* `cookiesamesiteattribute` - (Optional) Cookie Samesite attribute added to support adding cookie SameSite attribute for all set-cookies including appfw session cookies. Default value will be "SameSite=Lax".
* `defaultfieldformatmaxoccurrences` - (Optional) Maximum allowed occurrences of the form field name in a request.
* `fakeaccountdetection` - (Optional) Fake account detection flag : ON/OFF. If set to ON fake account detection is enabled on ADC, if set to OFF fake account detection is disabled.
* `fieldscan` - (Optional) Check if formfield limit scan is ON or OFF.
* `fieldscanlimit` - (Optional) Field scan limit value for HTML.
* `geolocationlogging` - (Optional) Enable Geo-Location Logging in CEF format logs for the profile.
* `grpcaction` - (Optional) gRPC validation. Possible values: [ none, block, log, stats ]
* `importprofilename` - (Optional) Name of the profile which will be created/updated to associate the relaxation rules.
* `insertcookiesamesiteattribute` - (Optional) Configure whether application firewall should add samesite attribute for set-cookies.
* `inspectquerycontenttypes` - (Optional) Inspect request query as well as web forms for injected SQL and cross-site scripts for following content types.
* `jsonblockkeywordaction` - (Optional) JSON Block Keyword action. Available settings function as follows:
    * Block - Block connections that violate this security check.
    * Log - Log violations of this security check.
    * Stats - Generate statistics for this security check.
    * None - Disable all actions for this security check.
    Possible values: [ none, block, log, stats ]
* `jsoncmdinjectionaction` - (Optional) One or more JSON CMD Injection actions. Available settings function as follows:
    * Block - Block connections that violate this security check.
    * Log - Log violations of this security check.
    * Stats - Generate statistics for this security check.
    * None - Disable all actions for this security check.
    Possible values: [ none, block, log, stats ]
* `jsoncmdinjectiongrammar` - (Optional) Check for CMD injection using CMD grammar in JSON.
* `jsoncmdinjectiontype` - (Optional) Available CMD injection types.
    * CMDSplChar : Checks for CMD Special Chars
    * CMDKeyword : Checks for CMD Keywords
    * CMDSplCharANDKeyword : Checks for both and blocks if both are found
    * CMDSplCharORKeyword : Checks for both and blocks if anyone is found,
    * None : Disables checking using both SQL Special Char and Keyword
    Possible values: [ CMDSplChar, CMDKeyword, CMDSplCharORKeyword, CMDSplCharANDKeyword, None ]
* `jsonerrorstatuscode` - (Optional) Response status code associated with JSON error page. Non-empty JSON error object must be imported to the application firewall profile for the status code.
* `jsonerrorstatusmessage` - (Optional) Response status message associated with JSON error page.
* `jsonfieldscan` - (Optional) Check if JSON field limit scan is ON or OFF.
* `jsonfieldscanlimit` - (Optional) Field scan limit value for JSON.
* `jsonmessagescan` - (Optional) Check if JSON message limit scan is ON or OFF.
* `jsonmessagescanlimit` - (Optional) Message scan limit value for JSON.
* `jsonsqlinjectiongrammar` - (Optional) Check for SQL injection using SQL grammar in JSON.
* `matchurlstring` - (Optional) Match this action url in archived Relaxation Rules to replace.
* `messagescan` - (Optional) Check if HTML message limit scan is ON or OFF.
* `messagescanlimit` - (Optional) Message scan limit value for HTML.
* `messagescanlimitcontenttypes` - (Optional) Enable Message Scan Limit for following content types.
* `overwrite` - (Optional) Purge existing Relaxation Rules and replace during import.
* `protofileobject` - (Optional) Name of the imported proto file.
* `relaxationrules` - (Optional) Import all appfw relaxation rules.
* `replaceurlstring` - (Optional) Replace matched url string with this action url string while restoring Relaxation Rules.
* `restaction` - (Optional) rest validation. Possible values: [ none, block, log, stats ]
* `sessioncookiename` - (Optional) Name of the session cookie that the application firewall uses to track user sessions.
    Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.
    The following requirement applies only to the Citrix ADC CLI:
    If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cookie name" or 'my cookie name').
* `sqlinjectionruletype` - (Optional) Specifies SQL Injection rule type: ALLOW/DENY. If ALLOW rule type is configured then allow list rules are used, if DENY rule type is configured then deny rules are used.
* `xmlerrorstatuscode` - (Optional) Response status code associated with XML error page. Non-empty XML error object must be imported to the application firewall profile for the status code.
* `xmlerrorstatusmessage` - (Optional) Response status message associated with XML error page.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwprofile`. It has the same value as the `name` attribute.
