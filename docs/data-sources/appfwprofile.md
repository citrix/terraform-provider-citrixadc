---
subcategory: "Application Firewall"
---

# Data Source: citrixadc_appfwprofile

Use this data source to retrieve information about an existing Application Firewall Profile.

The `citrixadc_appfwprofile` data source allows you to retrieve details of an Application Firewall profile by its name. This is useful for referencing existing profiles in your Terraform configurations without managing them directly.

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

## Usage Notes

1. The data source requires the profile to exist on the Citrix ADC. If the profile doesn't exist, Terraform will return an error.

2. When using this data source with a resource in the same configuration, use `depends_on` to ensure the resource is created before the data source attempts to read it.

3. All attributes returned by the data source are read-only and reflect the current state of the profile on the Citrix ADC.

4. The `name` attribute is case-sensitive and must exactly match the profile name on the Citrix ADC.

5. Many attributes are lists representing multiple actions that can be enabled simultaneously (e.g., `["block", "log", "stats"]`).
