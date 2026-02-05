package appfwprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AppfwprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"addcookieflags": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Add the specified flags to cookies. Available settings function as follows:\n* None - Do not add flags to cookies.\n* HTTP Only - Add the HTTP Only flag to cookies, which prevents scripts from accessing cookies.\n* Secure - Add Secure flag to cookies.\n* All - Add both HTTPOnly and Secure flags to cookies.",
			},
			"apispec": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the API Specification.",
			},
			"archivename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source for tar archive.",
			},
			"as_prof_bypass_list_enable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable bypass list for the profile.",
			},
			"as_prof_deny_list_enable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable deny list for the profile.",
			},
			"augment": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Augment Relaxation Rules during import",
			},
			"blockkeywordaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Block Keyword action. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -blockKeywordAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -blockKeywordAction none\".",
			},
			"bufferoverflowaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Buffer Overflow actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -bufferOverflowAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -bufferOverflowAction none\".",
			},
			"bufferoverflowmaxcookielength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum length, in characters, for cookies sent to your protected web sites. Requests with longer cookies are blocked.",
			},
			"bufferoverflowmaxheaderlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum length, in characters, for HTTP headers in requests sent to your protected web sites. Requests with longer headers are blocked.",
			},
			"bufferoverflowmaxquerylength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum length, in bytes, for query string sent to your protected web sites. Requests with longer query strings are blocked.",
			},
			"bufferoverflowmaxtotalheaderlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum length, in bytes, for the total HTTP header length in requests sent to your protected web sites. The minimum value of this and maxHeaderLen in httpProfile will be used. Requests with longer length are blocked.",
			},
			"bufferoverflowmaxurllength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum length, in characters, for URLs on your protected web sites. Requests with longer URLs are blocked.",
			},
			"canonicalizehtmlresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Perform HTML entity encoding for any special characters in responses sent by your protected web sites.",
			},
			"ceflogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable CEF format logs for the profile.",
			},
			"checkrequestheaders": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check request headers as well as web forms for injected SQL and cross-site scripts.",
			},
			"clientipexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression to get the client IP.",
			},
			"cmdinjectionaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Command injection action. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -cmdInjectionAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -cmdInjectionAction none\".",
			},
			"cmdinjectiongrammar": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check for CMD injection using CMD grammar",
			},
			"cmdinjectiontype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Available CMD injection types.\n-CMDSplChar              : Checks for CMD Special Chars\n-CMDKeyword              : Checks for CMD Keywords\n-CMDSplCharANDKeyword    : Checks for both and blocks if both are found\n-CMDSplCharORKeyword     : Checks for both and blocks if anyone is found,\n-None                    : Disables checking using both CMD Special Char and Keyword",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"contenttypeaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Content-type actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -contentTypeaction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -contentTypeaction none\".",
			},
			"cookieconsistencyaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Cookie Consistency actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -cookieConsistencyAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -cookieConsistencyAction none\".",
			},
			"cookieencryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of cookie encryption. Available settings function as follows:\n* None - Do not encrypt cookies.\n* Decrypt Only - Decrypt encrypted cookies, but do not encrypt cookies.\n* Encrypt Session Only - Encrypt session cookies, but not permanent cookies.\n* Encrypt All - Encrypt all cookies.",
			},
			"cookiehijackingaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more actions to prevent cookie hijacking. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\nNOTE: Cookie Hijacking feature is not supported for TLSv1.3\n\nCLI users: To enable one or more actions, type \"set appfw profile -cookieHijackingAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -cookieHijackingAction none\".",
			},
			"cookieproxying": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cookie proxy setting. Available settings function as follows:\n* None - Do not proxy cookies.\n* Session Only - Proxy session cookies by using the Citrix ADC session ID, but do not proxy permanent cookies.",
			},
			"cookiesamesiteattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cookie Samesite attribute added to support adding cookie SameSite attribute for all set-cookies including appfw session cookies. Default value will be \"SameSite=Lax\".",
			},
			"cookietransforms": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Perform the specified type of cookie transformation.\nAvailable settings function as follows:\n* Encryption - Encrypt cookies.\n* Proxying - Mask contents of server cookies by sending proxy cookie to users.\n* Cookie flags - Flag cookies as HTTP only to prevent scripts on user's browser from accessing and possibly modifying them.\nCAUTION: Make sure that this parameter is set to ON if you are configuring any cookie transformations. If it is set to OFF, no cookie transformations are performed regardless of any other settings.",
			},
			"creditcard": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Credit card types that the application firewall should protect.",
			},
			"creditcardaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Credit Card actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -creditCardAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -creditCardAction none\".",
			},
			"creditcardmaxallowed": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This parameter value is used by the block action. It represents the maximum number of credit card numbers that can appear on a web page served by your protected web sites. Pages that contain more credit card numbers are blocked.",
			},
			"creditcardxout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mask any credit card number detected in a response by replacing each digit, except the digits in the final group, with the letter \"X.\"",
			},
			"crosssitescriptingaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Cross-Site Scripting (XSS) actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -crossSiteScriptingAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -crossSiteScriptingAction none\".",
			},
			"crosssitescriptingcheckcompleteurls": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check complete URLs for cross-site scripts, instead of just the query portions of URLs.",
			},
			"crosssitescriptingtransformunsafehtml": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Transform cross-site scripts. This setting configures the application firewall to disable dangerous HTML instead of blocking the request.\nCAUTION: Make sure that this parameter is set to ON if you are configuring any cross-site scripting transformations. If it is set to OFF, no cross-site scripting transformations are performed regardless of any other settings.",
			},
			"csrftagaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Cross-Site Request Forgery (CSRF) Tagging actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -CSRFTagAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -CSRFTagAction none\".",
			},
			"customsettings": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Object name for custom settings.\nThis check is applicable to Profile Type: HTML, XML.",
			},
			"defaultcharset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default character set for protected web pages. Web pages sent by your protected web sites in response to user requests are assigned this character set if the page does not already specify a character set. The character sets supported by the application firewall are:\n* iso-8859-1 (English US)\n* big5 (Chinese Traditional)\n* gb2312 (Chinese Simplified)\n* sjis (Japanese Shift-JIS)\n* euc-jp (Japanese EUC-JP)\n* iso-8859-9 (Turkish)\n* utf-8 (Unicode)\n* euc-kr (Korean)",
			},
			"defaultfieldformatmaxlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum length, in characters, for data entered into a field that is assigned the default field type.",
			},
			"defaultfieldformatmaxoccurrences": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maxiumum allowed occurrences of the form field name in a request.",
			},
			"defaultfieldformatminlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum length, in characters, for data entered into a field that is assigned the default field type.\nTo disable the minimum and maximum length settings and allow data of any length to be entered into the field, set this parameter to zero (0).",
			},
			"defaultfieldformattype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Designate a default field type to be applied to web form fields that do not have a field type explicitly assigned to them.",
			},
			"defaults": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default configuration to apply to the profile. Basic defaults are intended for standard content that requires little further configuration, such as static web site content. Advanced defaults are intended for specialized content that requires significant specialized configuration, such as heavily scripted or dynamic content.\n\nCLI users: When adding an application firewall profile, you can set either the defaults or the type, but not both. To set both options, create the profile by using the add appfw profile command, and then use the set appfw profile command to configure the other option.",
			},
			"denyurlaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Deny URL actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nNOTE: The Deny URL check takes precedence over the Start URL check. If you enable blocking for the Deny URL check, the application firewall blocks any URL that is explicitly blocked by a Deny URL, even if the same URL would otherwise be allowed by the Start URL check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -denyURLaction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -denyURLaction none\".",
			},
			"dosecurecreditcardlogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Setting this option logs credit card numbers in the response when the match is found.",
			},
			"dynamiclearning": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more security checks. Available options are as follows:\n* SQLInjection - Enable dynamic learning for SQLInjection security check.\n* CrossSiteScripting - Enable dynamic learning for CrossSiteScripting security check.\n* fieldFormat - Enable dynamic learning for  fieldFormat security check.\n* None - Disable security checks for all security checks.\n\nCLI users: To enable dynamic learning on one or more security checks, type \"set appfw profile -dynamicLearning\" followed by the security checks to be enabled. To turn off dynamic learning on all security checks, type \"set appfw profile -dynamicLearning none\".",
			},
			"enableformtagging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable tagging of web form fields for use by the Form Field Consistency and CSRF Form Tagging checks.",
			},
			"errorurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL that application firewall uses as the Error URL.",
			},
			"excludefileuploadfromchecks": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exclude uploaded files from Form checks.",
			},
			"exemptclosureurlsfromsecuritychecks": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exempt URLs that pass the Start URL closure check from SQL injection, cross-site script, field format and field consistency security checks at locations other than headers.",
			},
			"fakeaccountdetection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fake account detection flag : ON/OFF. If set to ON fake account detection in enabled on ADC, if set to OFF fake account detection is disabled.",
			},
			"fieldconsistencyaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Form Field Consistency actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -fieldConsistencyaction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -fieldConsistencyAction none\".",
			},
			"fieldformataction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Field Format actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of suggested web form fields and field format assignments.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -fieldFormatAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -fieldFormatAction none\".",
			},
			"fieldscan": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check if formfield limit scan is ON or OFF.",
			},
			"fieldscanlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Field scan limit value for HTML",
			},
			"fileuploadmaxnum": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum allowed number of file uploads per form-submission request. The maximum setting (65535) allows an unlimited number of uploads.",
			},
			"fileuploadtypesaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more file upload types actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -fileUploadTypeAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -fileUploadTypeAction none\".",
			},
			"geolocationlogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable Geo-Location Logging in CEF format logs for the profile.",
			},
			"grpcaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "gRPC validation",
			},
			"htmlerrorobject": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name to assign to the HTML Error Object.\nMust begin with a letter, number, or the underscore character \\(_\\), and must contain only letters, numbers, and the hyphen \\(-\\), period \\(.\\) pound \\(\\#\\), space \\( \\), at (@), equals \\(=\\), colon \\(:\\), and underscore characters. Cannot be changed after the HTML error object is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks \\(for example, \"my HTML error object\" or 'my HTML error object'\\).",
			},
			"htmlerrorstatuscode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Response status code associated with HTML error page. Non-empty HTML error object must be imported to the application firewall profile for the status code.",
			},
			"htmlerrorstatusmessage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Response status message associated with HTML error page",
			},
			"importprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the profile which will be created/updated to associate the relaxation rules",
			},
			"infercontenttypexmlpayloadaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more infer content type payload actions. Available settings function as follows:\n* Block - Block connections that have mismatch in content-type header and payload.\n* Log - Log connections that have mismatch in content-type header and payload. The mismatched content-type in HTTP request header will be logged for the request.\n* Stats - Generate statistics when there is mismatch in content-type header and payload.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -inferContentTypeXMLPayloadAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -inferContentTypeXMLPayloadAction none\". Please note \"none\" action cannot be used with any other action type.",
			},
			"insertcookiesamesiteattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure whether application firewall should add samesite attribute for set-cookies",
			},
			"inspectcontenttypes": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more InspectContentType lists.\n* application/x-www-form-urlencoded\n* multipart/form-data\n* text/x-gwt-rpc\n\nCLI users: To enable, type \"set appfw profile -InspectContentTypes\" followed by the content types to be inspected.",
			},
			"inspectquerycontenttypes": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Inspect request query as well as web forms for injected SQL and cross-site scripts for following content types.",
			},
			"invalidpercenthandling": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure the method that the application firewall uses to handle percent-encoded names and values. Available settings function as follows:\n* asp_mode - Microsoft ASP format.\n* secure_mode - Secure format.",
			},
			"jsonblockkeywordaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "JSON Block Keyword action. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -JSONBlockKeywordAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -JSONBlockKeywordAction none\".",
			},
			"jsoncmdinjectionaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more JSON CMD Injection actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -JSONCMDInjectionAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -JSONCMDInjectionAction none\".",
			},
			"jsoncmdinjectiongrammar": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check for CMD injection using CMD grammar in JSON",
			},
			"jsoncmdinjectiontype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Available CMD injection types.\n-CMDSplChar              : Checks for CMD Special Chars\n-CMDKeyword              : Checks for CMD Keywords\n-CMDSplCharANDKeyword    : Checks for both and blocks if both are found\n-CMDSplCharORKeyword     : Checks for both and blocks if anyone is found,\n-None                    : Disables checking using both SQL Special Char and Keyword",
			},
			"jsondosaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more JSON Denial-of-Service (JsonDoS) actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -JSONDoSAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -JSONDoSAction none\".",
			},
			"jsonerrorobject": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name to the imported JSON Error Object to be set on application firewall profile.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks \\(for example, \"my JSON error object\" or 'my JSON error object'\\).",
			},
			"jsonerrorstatuscode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Response status code associated with JSON error page. Non-empty JSON error object must be imported to the application firewall profile for the status code.",
			},
			"jsonerrorstatusmessage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Response status message associated with JSON error page",
			},
			"jsonfieldscan": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check if JSON field limit scan is ON or OFF.",
			},
			"jsonfieldscanlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Field scan limit value for JSON",
			},
			"jsonmessagescan": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check if JSON message limit scan is ON or OFF",
			},
			"jsonmessagescanlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Message scan limit value for JSON",
			},
			"jsonsqlinjectionaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more JSON SQL Injection actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -JSONSQLInjectionAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -JSONSQLInjectionAction none\".",
			},
			"jsonsqlinjectiongrammar": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check for SQL injection using SQL grammar in JSON",
			},
			"jsonsqlinjectiontype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Available SQL injection types.\n-SQLSplChar              : Checks for SQL Special Chars\n-SQLKeyword              : Checks for SQL Keywords\n-SQLSplCharANDKeyword    : Checks for both and blocks if both are found\n-SQLSplCharORKeyword     : Checks for both and blocks if anyone is found,\n-None                    : Disables checking using both SQL Special Char and Keyword",
			},
			"jsonxssaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more JSON Cross-Site Scripting actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -JSONXssAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -JSONXssAction none\".",
			},
			"logeverypolicyhit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log every profile match, regardless of security checks results.",
			},
			"matchurlstring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Match this action url in archived Relaxation Rules to replace.",
			},
			"messagescan": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check if HTML message limit scan is ON or OFF",
			},
			"messagescanlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Message scan limit value for HTML",
			},
			"messagescanlimitcontenttypes": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Enable Message Scan Limit for following content types.",
			},
			"multipleheaderaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more multiple header actions. Available settings function as follows:\n* Block - Block connections that have multiple headers.\n* Log - Log connections that have multiple headers.\n* KeepLast - Keep only last header when multiple headers are present.\n\nRequest headers inspected:\n* Accept-Encoding\n* Content-Encoding\n* Content-Range\n* Content-Type\n* Host\n* Range\n* Referer\n\nCLI users: To enable one or more actions, type \"set appfw profile -multipleHeaderAction\" followed by the actions to be enabled.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my profile\" or 'my profile').",
			},
			"optimizepartialreqs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Optimize handle of HTTP partial requests i.e. those with range headers.\nAvailable settings are as follows:\n* ON  - Partial requests by the client result in partial requests to the backend server in most cases.\n* OFF - Partial requests by the client are changed to full requests to the backend server",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Purge existing Relaxation Rules and replace during import",
			},
			"percentdecoderecursively": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure whether the application firewall should use percentage recursive decoding",
			},
			"postbodylimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum allowed HTTP post body size, in bytes. Maximum supported value is 10GB. Citrix recommends enabling streaming option for large values of post body limit (>20MB).",
			},
			"postbodylimitaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Post Body Limit actions. Available settings function as follows:\n* Block - Block connections that violate this security check. Must always be set.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -PostBodyLimitAction block\" followed by the other actions to be enabled.",
			},
			"postbodylimitsignature": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum allowed HTTP post body size for signature inspection for location HTTP_POST_BODY in the signatures, in bytes. Note that the changes in value could impact CPU and latency profile.",
			},
			"protofileobject": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the imported proto file.",
			},
			"refererheadercheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable validation of Referer headers.\nReferer validation ensures that a web form that a user sends to your web site originally came from your web site, not an outside attacker.\nAlthough this parameter is part of the Start URL check, referer validation protects against cross-site request forgery (CSRF) attacks, not Start URL attacks.",
			},
			"relaxationrules": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Import all appfw relaxation rules",
			},
			"replaceurlstring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Replace matched url string with this action url string while restoring Relaxation Rules",
			},
			"requestcontenttype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default Content-Type header for requests.\nA Content-Type header can contain 0-255 letters, numbers, and the hyphen (-) and underscore (_) characters.",
			},
			"responsecontenttype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default Content-Type header for responses.\nA Content-Type header can contain 0-255 letters, numbers, and the hyphen (-) and underscore (_) characters.",
			},
			"restaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "rest validation",
			},
			"rfcprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Object name of the rfc profile.",
			},
			"semicolonfieldseparator": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow ';' as a form field separator in URL queries and POST form bodies.",
			},
			"sessioncookiename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the session cookie that the application firewall uses to track user sessions.\nMust begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cookie name\" or 'my cookie name').",
			},
			"sessionlessfieldconsistency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Perform sessionless Field Consistency Checks.",
			},
			"sessionlessurlclosure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable session less URL Closure Checks.\nThis check is applicable to Profile Type: HTML.",
			},
			"signatures": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Object name for signatures.\nThis check is applicable to Profile Type: HTML, XML.",
			},
			"sqlinjectionaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more HTML SQL Injection actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -SQLInjectionAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -SQLInjectionAction none\".",
			},
			"sqlinjectionchecksqlwildchars": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check for form fields that contain SQL wild chars .",
			},
			"sqlinjectiongrammar": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check for SQL injection using SQL grammar",
			},
			"sqlinjectiononlycheckfieldswithsqlchars": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check only form fields that contain SQL special strings (characters) for injected SQL code.\nMost SQL servers require a special string to activate an SQL request, so SQL code without a special string is harmless to most SQL servers.",
			},
			"sqlinjectionparsecomments": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parse HTML comments and exempt them from the HTML SQL Injection check. You must specify the type of comments that the application firewall is to detect and exempt from this security check. Available settings function as follows:\n* Check all - Check all content.\n* ANSI - Exempt content that is part of an ANSI (Mozilla-style) comment.\n* Nested - Exempt content that is part of a nested (Microsoft-style) comment.\n* ANSI Nested - Exempt content that is part of any type of comment.",
			},
			"sqlinjectionruletype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies SQL Injection rule type: ALLOW/DENY. If ALLOW rule type is configured then allow list rules are used, if DENY rule type is configured then deny rules are used.",
			},
			"sqlinjectiontransformspecialchars": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Transform injected SQL code. This setting configures the application firewall to disable SQL special strings instead of blocking the request. Since most SQL servers require a special string to activate an SQL keyword, in most cases a request that contains injected SQL code is safe if special strings are disabled.\nCAUTION: Make sure that this parameter is set to ON if you are configuring any SQL injection transformations. If it is set to OFF, no SQL injection transformations are performed regardless of any other settings.",
			},
			"sqlinjectiontype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Available SQL injection types.\n-SQLSplChar              : Checks for SQL Special Chars\n-SQLKeyword		 : Checks for SQL Keywords\n-SQLSplCharANDKeyword    : Checks for both and blocks if both are found\n-SQLSplCharORKeyword     : Checks for both and blocks if anyone is found\n-None                    : Disables checking using both SQL Special Char and Keyword",
			},
			"starturlaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Start URL actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -startURLaction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -startURLaction none\".",
			},
			"starturlclosure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Toggle  the state of Start URL Closure.",
			},
			"streaming": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Setting this option converts content-length form submission requests (requests with content-type \"application/x-www-form-urlencoded\" or \"multipart/form-data\") to chunked requests when atleast one of the following protections : Signatures, SQL injection protection, XSS protection, form field consistency protection, starturl closure, CSRF tagging, JSON SQL, JSON XSS, JSON DOS is enabled. Please make sure that the backend server accepts chunked requests before enabling this option. Citrix recommends enabling this option for large request sizes(>20MB).",
			},
			"stripcomments": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Strip HTML comments.\nThis check is applicable to Profile Type: HTML.",
			},
			"striphtmlcomments": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Strip HTML comments before forwarding a web page sent by a protected web site in response to a user request.",
			},
			"stripxmlcomments": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Strip XML comments before forwarding a web page sent by a protected web site in response to a user request.",
			},
			"trace": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Toggle  the state of trace",
			},
			"type": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Application firewall profile type, which controls which security checks and settings are applied to content that is filtered with the profile. Available settings function as follows:\n* HTML - HTML-based web sites.\n* XML -  XML-based web sites and services.\n* JSON - JSON-based web sites and services.\n* HTML XML (Web 2.0) - Sites that contain both HTML and XML content, such as ATOM feeds, blogs, and RSS feeds.\n* HTML JSON  - Sites that contain both HTML and JSON content.\n* XML JSON   - Sites that contain both XML and JSON content.\n* HTML XML JSON   - Sites that contain HTML, XML and JSON content.",
			},
			"urldecoderequestcookies": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL Decode request cookies before subjecting them to SQL and cross-site scripting checks.",
			},
			"usehtmlerrorobject": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send an imported HTML Error object to a user when a request is blocked, instead of redirecting the user to the designated Error URL.",
			},
			"verboseloglevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Detailed Logging Verbose Log Level.",
			},
			"xmlattachmentaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more XML Attachment actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLAttachmentAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLAttachmentAction none\".",
			},
			"xmldosaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more XML Denial-of-Service (XDoS) actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLDoSAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLDoSAction none\".",
			},
			"xmlerrorobject": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name to assign to the XML Error Object, which the application firewall displays when a user request is blocked.\nMust begin with a letter, number, or the underscore character \\(_\\), and must contain only letters, numbers, and the hyphen \\(-\\), period \\(.\\) pound \\(\\#\\), space \\( \\), at (@), equals \\(=\\), colon \\(:\\), and underscore characters. Cannot be changed after the XML error object is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks \\(for example, \"my XML error object\" or 'my XML error object'\\).",
			},
			"xmlerrorstatuscode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Response status code associated with XML error page. Non-empty XML error object must be imported to the application firewall profile for the status code.",
			},
			"xmlerrorstatusmessage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Response status message associated with XML error page",
			},
			"xmlformataction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more XML Format actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLFormatAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLFormatAction none\".",
			},
			"xmlsoapfaultaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more XML SOAP Fault Filtering actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n* Remove - Remove all violations for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLSOAPFaultAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLSOAPFaultAction none\".",
			},
			"xmlsqlinjectionaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more XML SQL Injection actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLSQLInjectionAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLSQLInjectionAction none\".",
			},
			"xmlsqlinjectionchecksqlwildchars": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check for form fields that contain SQL wild chars .",
			},
			"xmlsqlinjectiononlycheckfieldswithsqlchars": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check only form fields that contain SQL special characters, which most SQL servers require before accepting an SQL command, for injected SQL.",
			},
			"xmlsqlinjectionparsecomments": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parse comments in XML Data and exempt those sections of the request that are from the XML SQL Injection check. You must configure the type of comments that the application firewall is to detect and exempt from this security check. Available settings function as follows:\n* Check all - Check all content.\n* ANSI - Exempt content that is part of an ANSI (Mozilla-style) comment.\n* Nested - Exempt content that is part of a nested (Microsoft-style) comment.\n* ANSI Nested - Exempt content that is part of any type of comment.",
			},
			"xmlsqlinjectiontype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Available SQL injection types.\n-SQLSplChar              : Checks for SQL Special Chars\n-SQLKeyword              : Checks for SQL Keywords\n-SQLSplCharANDKeyword    : Checks for both and blocks if both are found\n-SQLSplCharORKeyword     : Checks for both and blocks if anyone is found",
			},
			"xmlvalidationaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more XML Validation actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLValidationAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLValidationAction none\".",
			},
			"xmlwsiaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Web Services Interoperability (WSI) actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLWSIAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLWSIAction none\".",
			},
			"xmlxssaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more XML Cross-Site Scripting actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLXSSAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLXSSAction none\".",
			},
		},
	}
}
