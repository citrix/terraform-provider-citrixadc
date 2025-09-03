/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package appfw

/**
* Configuration for application firewall profile resource.
*/
type Appfwprofile struct {
	/**
	* Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Default configuration to apply to the profile. Basic defaults are intended for standard content that requires little further configuration, such as static web site content. Advanced defaults are intended for specialized content that requires significant specialized configuration, such as heavily scripted or dynamic content.
		CLI users: When adding an application firewall profile, you can set either the defaults or the type, but not both. To set both options, create the profile by using the add appfw profile command, and then use the set appfw profile command to configure the other option.
	*/
	Defaults string `json:"defaults,omitempty"`
	/**
	* One or more Start URL actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Learn - Use the learning engine to generate a list of exceptions to this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -startURLaction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -startURLaction none".
	*/
	Starturlaction []string `json:"starturlaction,omitempty"`
	/**
	* One or more infer content type payload actions. Available settings function as follows:
		* Block - Block connections that have mismatch in content-type header and payload.
		* Log - Log connections that have mismatch in content-type header and payload. The mismatched content-type in HTTP request header will be logged for the request.
		* Stats - Generate statistics when there is mismatch in content-type header and payload.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -inferContentTypeXMLPayloadAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -inferContentTypeXMLPayloadAction none". Please note "none" action cannot be used with any other action type.
	*/
	Infercontenttypexmlpayloadaction []string `json:"infercontenttypexmlpayloadaction,omitempty"`
	/**
	* One or more Content-type actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Learn - Use the learning engine to generate a list of exceptions to this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -contentTypeaction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -contentTypeaction none".
	*/
	Contenttypeaction []string `json:"contenttypeaction,omitempty"`
	/**
	* One or more InspectContentType lists.
		* application/x-www-form-urlencoded
		* multipart/form-data
		* text/x-gwt-rpc
		CLI users: To enable, type "set appfw profile -InspectContentTypes" followed by the content types to be inspected.
	*/
	Inspectcontenttypes []string `json:"inspectcontenttypes,omitempty"`
	/**
	* Toggle  the state of Start URL Closure.
	*/
	Starturlclosure string `json:"starturlclosure,omitempty"`
	/**
	* One or more Deny URL actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		NOTE: The Deny URL check takes precedence over the Start URL check. If you enable blocking for the Deny URL check, the application firewall blocks any URL that is explicitly blocked by a Deny URL, even if the same URL would otherwise be allowed by the Start URL check.
		CLI users: To enable one or more actions, type "set appfw profile -denyURLaction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -denyURLaction none".
	*/
	Denyurlaction []string `json:"denyurlaction,omitempty"`
	/**
	* Enable validation of Referer headers.
		Referer validation ensures that a web form that a user sends to your web site originally came from your web site, not an outside attacker.
		Although this parameter is part of the Start URL check, referer validation protects against cross-site request forgery (CSRF) attacks, not Start URL attacks.
	*/
	Refererheadercheck string `json:"refererheadercheck,omitempty"`
	/**
	* One or more Cookie Consistency actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Learn - Use the learning engine to generate a list of exceptions to this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -cookieConsistencyAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -cookieConsistencyAction none".
	*/
	Cookieconsistencyaction []string `json:"cookieconsistencyaction,omitempty"`
	/**
	* One or more actions to prevent cookie hijacking. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		NOTE: Cookie Hijacking feature is not supported for TLSv1.3
		CLI users: To enable one or more actions, type "set appfw profile -cookieHijackingAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -cookieHijackingAction none".
	*/
	Cookiehijackingaction []string `json:"cookiehijackingaction,omitempty"`
	/**
	* Perform the specified type of cookie transformation.
		Available settings function as follows:
		* Encryption - Encrypt cookies.
		* Proxying - Mask contents of server cookies by sending proxy cookie to users.
		* Cookie flags - Flag cookies as HTTP only to prevent scripts on user's browser from accessing and possibly modifying them.
		CAUTION: Make sure that this parameter is set to ON if you are configuring any cookie transformations. If it is set to OFF, no cookie transformations are performed regardless of any other settings.
	*/
	Cookietransforms string `json:"cookietransforms,omitempty"`
	/**
	* Type of cookie encryption. Available settings function as follows:
		* None - Do not encrypt cookies.
		* Decrypt Only - Decrypt encrypted cookies, but do not encrypt cookies.
		* Encrypt Session Only - Encrypt session cookies, but not permanent cookies.
		* Encrypt All - Encrypt all cookies.
	*/
	Cookieencryption string `json:"cookieencryption,omitempty"`
	/**
	* Cookie proxy setting. Available settings function as follows:
		* None - Do not proxy cookies.
		* Session Only - Proxy session cookies by using the Citrix ADC session ID, but do not proxy permanent cookies.
	*/
	Cookieproxying string `json:"cookieproxying,omitempty"`
	/**
	* Add the specified flags to cookies. Available settings function as follows:
		* None - Do not add flags to cookies.
		* HTTP Only - Add the HTTP Only flag to cookies, which prevents scripts from accessing cookies.
		* Secure - Add Secure flag to cookies.
		* All - Add both HTTPOnly and Secure flags to cookies.
	*/
	Addcookieflags string `json:"addcookieflags,omitempty"`
	/**
	* One or more Form Field Consistency actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Learn - Use the learning engine to generate a list of exceptions to this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -fieldConsistencyaction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -fieldConsistencyAction none".
	*/
	Fieldconsistencyaction []string `json:"fieldconsistencyaction,omitempty"`
	/**
	* One or more Cross-Site Request Forgery (CSRF) Tagging actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Learn - Use the learning engine to generate a list of exceptions to this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -CSRFTagAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -CSRFTagAction none".
	*/
	Csrftagaction []string `json:"csrftagaction,omitempty"`
	/**
	* One or more Cross-Site Scripting (XSS) actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Learn - Use the learning engine to generate a list of exceptions to this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -crossSiteScriptingAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -crossSiteScriptingAction none".
	*/
	Crosssitescriptingaction []string `json:"crosssitescriptingaction,omitempty"`
	/**
	* Transform cross-site scripts. This setting configures the application firewall to disable dangerous HTML instead of blocking the request.
		CAUTION: Make sure that this parameter is set to ON if you are configuring any cross-site scripting transformations. If it is set to OFF, no cross-site scripting transformations are performed regardless of any other settings.
	*/
	Crosssitescriptingtransformunsafehtml string `json:"crosssitescriptingtransformunsafehtml,omitempty"`
	/**
	* Check complete URLs for cross-site scripts, instead of just the query portions of URLs.
	*/
	Crosssitescriptingcheckcompleteurls string `json:"crosssitescriptingcheckcompleteurls,omitempty"`
	/**
	* One or more HTML SQL Injection actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Learn - Use the learning engine to generate a list of exceptions to this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -SQLInjectionAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -SQLInjectionAction none".
	*/
	Sqlinjectionaction []string `json:"sqlinjectionaction,omitempty"`
	/**
	* Command injection action. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -cmdInjectionAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -cmdInjectionAction none".
	*/
	Cmdinjectionaction []string `json:"cmdinjectionaction,omitempty"`
	/**
	* Available CMD injection types.
		-CMDSplChar              : Checks for CMD Special Chars
		-CMDKeyword              : Checks for CMD Keywords
		-CMDSplCharANDKeyword    : Checks for both and blocks if both are found
		-CMDSplCharORKeyword     : Checks for both and blocks if anyone is found,
		-None                    : Disables checking using both CMD Special Char and Keyword
	*/
	Cmdinjectiontype string `json:"cmdinjectiontype,omitempty"`
	/**
	* Check for SQL injection using SQL grammar
	*/
	Sqlinjectiongrammar string `json:"sqlinjectiongrammar,omitempty"`
	/**
	* Check for CMD injection using CMD grammar
	*/
	Cmdinjectiongrammar string `json:"cmdinjectiongrammar,omitempty"`
	/**
	* Check if formfield limit scan is ON or OFF.
	*/
	Fieldscan string `json:"fieldscan,omitempty"`
	/**
	*  Field scan limit value for HTML
	*/
	Fieldscanlimit int `json:"fieldscanlimit,omitempty"`
	/**
	* Check if JSON field limit scan is ON or OFF.
	*/
	Jsonfieldscan string `json:"jsonfieldscan,omitempty"`
	/**
	*  Field scan limit value for JSON
	*/
	Jsonfieldscanlimit int `json:"jsonfieldscanlimit,omitempty"`
	/**
	* Check if HTML message limit scan is ON or OFF
	*/
	Messagescan string `json:"messagescan,omitempty"`
	/**
	* Message scan limit value for HTML
	*/
	Messagescanlimit int `json:"messagescanlimit,omitempty"`
	/**
	* Check if JSON message limit scan is ON or OFF
	*/
	Jsonmessagescan string `json:"jsonmessagescan,omitempty"`
	/**
	* Message scan limit value for JSON
	*/
	Jsonmessagescanlimit int `json:"jsonmessagescanlimit,omitempty"`
	/**
	* Enable Message Scan Limit for following content types.
	*/
	Messagescanlimitcontenttypes []string `json:"messagescanlimitcontenttypes,omitempty"`
	/**
	* Transform injected SQL code. This setting configures the application firewall to disable SQL special strings instead of blocking the request. Since most SQL servers require a special string to activate an SQL keyword, in most cases a request that contains injected SQL code is safe if special strings are disabled.
		CAUTION: Make sure that this parameter is set to ON if you are configuring any SQL injection transformations. If it is set to OFF, no SQL injection transformations are performed regardless of any other settings.
	*/
	Sqlinjectiontransformspecialchars string `json:"sqlinjectiontransformspecialchars,omitempty"`
	/**
	* Check only form fields that contain SQL special strings (characters) for injected SQL code.
		Most SQL servers require a special string to activate an SQL request, so SQL code without a special string is harmless to most SQL servers.
	*/
	Sqlinjectiononlycheckfieldswithsqlchars string `json:"sqlinjectiononlycheckfieldswithsqlchars,omitempty"`
	/**
	* Available SQL injection types.
		-SQLSplChar              : Checks for SQL Special Chars
		-SQLKeyword		 : Checks for SQL Keywords
		-SQLSplCharANDKeyword    : Checks for both and blocks if both are found
		-SQLSplCharORKeyword     : Checks for both and blocks if anyone is found
		-None                    : Disables checking using both SQL Special Char and Keyword
	*/
	Sqlinjectiontype string `json:"sqlinjectiontype,omitempty"`
	/**
	* Check for form fields that contain SQL wild chars .
	*/
	Sqlinjectionchecksqlwildchars string `json:"sqlinjectionchecksqlwildchars,omitempty"`
	/**
	* One or more Field Format actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Learn - Use the learning engine to generate a list of suggested web form fields and field format assignments.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -fieldFormatAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -fieldFormatAction none".
	*/
	Fieldformataction []string `json:"fieldformataction,omitempty"`
	/**
	* Designate a default field type to be applied to web form fields that do not have a field type explicitly assigned to them.
	*/
	Defaultfieldformattype string `json:"defaultfieldformattype,omitempty"`
	/**
	* Minimum length, in characters, for data entered into a field that is assigned the default field type.
		To disable the minimum and maximum length settings and allow data of any length to be entered into the field, set this parameter to zero (0).
	*/
	Defaultfieldformatminlength int `json:"defaultfieldformatminlength,omitempty"`
	/**
	* Maximum length, in characters, for data entered into a field that is assigned the default field type.
	*/
	Defaultfieldformatmaxlength int `json:"defaultfieldformatmaxlength,omitempty"`
	/**
	* Maxiumum allowed occurrences of the form field name in a request.
	*/
	Defaultfieldformatmaxoccurrences int `json:"defaultfieldformatmaxoccurrences,omitempty"`
	/**
	* One or more Buffer Overflow actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -bufferOverflowAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -bufferOverflowAction none".
	*/
	Bufferoverflowaction []string `json:"bufferoverflowaction,omitempty"`
	/**
	* gRPC validation
	*/
	Grpcaction []string `json:"grpcaction,omitempty"`
	/**
	* rest validation
	*/
	Restaction []string `json:"restaction,omitempty"`
	/**
	* Maximum length, in characters, for URLs on your protected web sites. Requests with longer URLs are blocked.
	*/
	Bufferoverflowmaxurllength int `json:"bufferoverflowmaxurllength,omitempty"`
	/**
	* Maximum length, in characters, for HTTP headers in requests sent to your protected web sites. Requests with longer headers are blocked.
	*/
	Bufferoverflowmaxheaderlength int `json:"bufferoverflowmaxheaderlength,omitempty"`
	/**
	* Maximum length, in characters, for cookies sent to your protected web sites. Requests with longer cookies are blocked.
	*/
	Bufferoverflowmaxcookielength int `json:"bufferoverflowmaxcookielength,omitempty"`
	/**
	* Maximum length, in bytes, for query string sent to your protected web sites. Requests with longer query strings are blocked.
	*/
	Bufferoverflowmaxquerylength int `json:"bufferoverflowmaxquerylength,omitempty"`
	/**
	* Maximum length, in bytes, for the total HTTP header length in requests sent to your protected web sites. The minimum value of this and maxHeaderLen in httpProfile will be used. Requests with longer length are blocked.
	*/
	Bufferoverflowmaxtotalheaderlength int `json:"bufferoverflowmaxtotalheaderlength,omitempty"`
	/**
	* One or more Credit Card actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -creditCardAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -creditCardAction none".
	*/
	Creditcardaction []string `json:"creditcardaction,omitempty"`
	/**
	* Credit card types that the application firewall should protect.
	*/
	Creditcard []string `json:"creditcard,omitempty"`
	/**
	* This parameter value is used by the block action. It represents the maximum number of credit card numbers that can appear on a web page served by your protected web sites. Pages that contain more credit card numbers are blocked.
	*/
	Creditcardmaxallowed int `json:"creditcardmaxallowed,omitempty"`
	/**
	* Mask any credit card number detected in a response by replacing each digit, except the digits in the final group, with the letter "X."
	*/
	Creditcardxout string `json:"creditcardxout,omitempty"`
	/**
	* Setting this option logs credit card numbers in the response when the match is found.
	*/
	Dosecurecreditcardlogging string `json:"dosecurecreditcardlogging,omitempty"`
	/**
	* Setting this option converts content-length form submission requests (requests with content-type "application/x-www-form-urlencoded" or "multipart/form-data") to chunked requests when atleast one of the following protections : Signatures, SQL injection protection, XSS protection, form field consistency protection, starturl closure, CSRF tagging, JSON SQL, JSON XSS, JSON DOS is enabled. Please make sure that the backend server accepts chunked requests before enabling this option. Citrix recommends enabling this option for large request sizes(>20MB).
	*/
	Streaming string `json:"streaming,omitempty"`
	/**
	* Toggle  the state of trace
	*/
	Trace string `json:"trace,omitempty"`
	/**
	* Default Content-Type header for requests.
		A Content-Type header can contain 0-255 letters, numbers, and the hyphen (-) and underscore (_) characters.
	*/
	Requestcontenttype string `json:"requestcontenttype,omitempty"`
	/**
	* Default Content-Type header for responses.
		A Content-Type header can contain 0-255 letters, numbers, and the hyphen (-) and underscore (_) characters.
	*/
	Responsecontenttype string `json:"responsecontenttype,omitempty"`
	/**
	* Name to the imported JSON Error Object to be set on application firewall profile.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my JSON error object" or 'my JSON error object'\).
	*/
	Jsonerrorobject string `json:"jsonerrorobject,omitempty"`
	/**
	* Name of the API Specification.
	*/
	Apispec string `json:"apispec,omitempty"`
	/**
	* Name of the imported proto file.
	*/
	Protofileobject string `json:"protofileobject,omitempty"`
	/**
	* Response status code associated with JSON error page. Non-empty JSON error object must be imported to the application firewall profile for the status code.
	*/
	Jsonerrorstatuscode int `json:"jsonerrorstatuscode,omitempty"`
	/**
	* Response status message associated with JSON error page
	*/
	Jsonerrorstatusmessage string `json:"jsonerrorstatusmessage,omitempty"`
	/**
	* One or more JSON Denial-of-Service (JsonDoS) actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -JSONDoSAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -JSONDoSAction none".
	*/
	Jsondosaction []string `json:"jsondosaction,omitempty"`
	/**
	* One or more JSON SQL Injection actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -JSONSQLInjectionAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -JSONSQLInjectionAction none".
	*/
	Jsonsqlinjectionaction []string `json:"jsonsqlinjectionaction,omitempty"`
	/**
	* Available SQL injection types.
		-SQLSplChar              : Checks for SQL Special Chars
		-SQLKeyword              : Checks for SQL Keywords
		-SQLSplCharANDKeyword    : Checks for both and blocks if both are found
		-SQLSplCharORKeyword     : Checks for both and blocks if anyone is found,
		-None                    : Disables checking using both SQL Special Char and Keyword
	*/
	Jsonsqlinjectiontype string `json:"jsonsqlinjectiontype,omitempty"`
	/**
	* Check for SQL injection using SQL grammar in JSON
	*/
	Jsonsqlinjectiongrammar string `json:"jsonsqlinjectiongrammar,omitempty"`
	/**
	* One or more JSON CMD Injection actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -JSONCMDInjectionAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -JSONCMDInjectionAction none".
	*/
	Jsoncmdinjectionaction []string `json:"jsoncmdinjectionaction,omitempty"`
	/**
	* Available CMD injection types.
		-CMDSplChar              : Checks for CMD Special Chars
		-CMDKeyword              : Checks for CMD Keywords
		-CMDSplCharANDKeyword    : Checks for both and blocks if both are found
		-CMDSplCharORKeyword     : Checks for both and blocks if anyone is found,
		-None                    : Disables checking using both SQL Special Char and Keyword
	*/
	Jsoncmdinjectiontype string `json:"jsoncmdinjectiontype,omitempty"`
	/**
	* Check for CMD injection using CMD grammar in JSON
	*/
	Jsoncmdinjectiongrammar string `json:"jsoncmdinjectiongrammar,omitempty"`
	/**
	* One or more JSON Cross-Site Scripting actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -JSONXssAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -JSONXssAction none".
	*/
	Jsonxssaction []string `json:"jsonxssaction,omitempty"`
	/**
	* One or more XML Denial-of-Service (XDoS) actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Learn - Use the learning engine to generate a list of exceptions to this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -XMLDoSAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLDoSAction none".
	*/
	Xmldosaction []string `json:"xmldosaction,omitempty"`
	/**
	* One or more XML Format actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -XMLFormatAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLFormatAction none".
	*/
	Xmlformataction []string `json:"xmlformataction,omitempty"`
	/**
	* One or more XML SQL Injection actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -XMLSQLInjectionAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLSQLInjectionAction none".
	*/
	Xmlsqlinjectionaction []string `json:"xmlsqlinjectionaction,omitempty"`
	/**
	* Check only form fields that contain SQL special characters, which most SQL servers require before accepting an SQL command, for injected SQL.
	*/
	Xmlsqlinjectiononlycheckfieldswithsqlchars string `json:"xmlsqlinjectiononlycheckfieldswithsqlchars,omitempty"`
	/**
	* Available SQL injection types.
		-SQLSplChar              : Checks for SQL Special Chars
		-SQLKeyword              : Checks for SQL Keywords
		-SQLSplCharANDKeyword    : Checks for both and blocks if both are found
		-SQLSplCharORKeyword     : Checks for both and blocks if anyone is found
	*/
	Xmlsqlinjectiontype string `json:"xmlsqlinjectiontype,omitempty"`
	/**
	* Check for form fields that contain SQL wild chars .
	*/
	Xmlsqlinjectionchecksqlwildchars string `json:"xmlsqlinjectionchecksqlwildchars,omitempty"`
	/**
	* Parse comments in XML Data and exempt those sections of the request that are from the XML SQL Injection check. You must configure the type of comments that the application firewall is to detect and exempt from this security check. Available settings function as follows:
		* Check all - Check all content.
		* ANSI - Exempt content that is part of an ANSI (Mozilla-style) comment.
		* Nested - Exempt content that is part of a nested (Microsoft-style) comment.
		* ANSI Nested - Exempt content that is part of any type of comment.
	*/
	Xmlsqlinjectionparsecomments string `json:"xmlsqlinjectionparsecomments,omitempty"`
	/**
	* One or more XML Cross-Site Scripting actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -XMLXSSAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLXSSAction none".
	*/
	Xmlxssaction []string `json:"xmlxssaction,omitempty"`
	/**
	* One or more Web Services Interoperability (WSI) actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Learn - Use the learning engine to generate a list of exceptions to this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -XMLWSIAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLWSIAction none".
	*/
	Xmlwsiaction []string `json:"xmlwsiaction,omitempty"`
	/**
	* One or more XML Attachment actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Learn - Use the learning engine to generate a list of exceptions to this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -XMLAttachmentAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLAttachmentAction none".
	*/
	Xmlattachmentaction []string `json:"xmlattachmentaction,omitempty"`
	/**
	* One or more XML Validation actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -XMLValidationAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLValidationAction none".
	*/
	Xmlvalidationaction []string `json:"xmlvalidationaction,omitempty"`
	/**
	* Name to assign to the XML Error Object, which the application firewall displays when a user request is blocked.
		Must begin with a letter, number, or the underscore character \(_\), and must contain only letters, numbers, and the hyphen \(-\), period \(.\) pound \(\#\), space \( \), at (@), equals \(=\), colon \(:\), and underscore characters. Cannot be changed after the XML error object is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my XML error object" or 'my XML error object'\).
	*/
	Xmlerrorobject string `json:"xmlerrorobject,omitempty"`
	/**
	* Response status code associated with XML error page. Non-empty XML error object must be imported to the application firewall profile for the status code.
	*/
	Xmlerrorstatuscode int `json:"xmlerrorstatuscode,omitempty"`
	/**
	* Response status message associated with XML error page
	*/
	Xmlerrorstatusmessage string `json:"xmlerrorstatusmessage,omitempty"`
	/**
	* Object name for custom settings.
		This check is applicable to Profile Type: HTML, XML. 
	*/
	Customsettings string `json:"customsettings,omitempty"`
	/**
	* Object name for signatures.
		This check is applicable to Profile Type: HTML, XML. 
	*/
	Signatures string `json:"signatures,omitempty"`
	/**
	* One or more XML SOAP Fault Filtering actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		* Remove - Remove all violations for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -XMLSOAPFaultAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -XMLSOAPFaultAction none".
	*/
	Xmlsoapfaultaction []string `json:"xmlsoapfaultaction,omitempty"`
	/**
	* Send an imported HTML Error object to a user when a request is blocked, instead of redirecting the user to the designated Error URL.
	*/
	Usehtmlerrorobject string `json:"usehtmlerrorobject,omitempty"`
	/**
	* URL that application firewall uses as the Error URL.
	*/
	Errorurl string `json:"errorurl,omitempty"`
	/**
	* Name to assign to the HTML Error Object.
		Must begin with a letter, number, or the underscore character \(_\), and must contain only letters, numbers, and the hyphen \(-\), period \(.\) pound \(\#\), space \( \), at (@), equals \(=\), colon \(:\), and underscore characters. Cannot be changed after the HTML error object is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my HTML error object" or 'my HTML error object'\).
	*/
	Htmlerrorobject string `json:"htmlerrorobject,omitempty"`
	/**
	* Response status code associated with HTML error page. Non-empty HTML error object must be imported to the application firewall profile for the status code.
	*/
	Htmlerrorstatuscode int `json:"htmlerrorstatuscode,omitempty"`
	/**
	* Response status message associated with HTML error page
	*/
	Htmlerrorstatusmessage string `json:"htmlerrorstatusmessage,omitempty"`
	/**
	* Log every profile match, regardless of security checks results.
	*/
	Logeverypolicyhit string `json:"logeverypolicyhit,omitempty"`
	/**
	* Strip HTML comments.
		This check is applicable to Profile Type: HTML. 
	*/
	Stripcomments string `json:"stripcomments,omitempty"`
	/**
	* Strip HTML comments before forwarding a web page sent by a protected web site in response to a user request.
	*/
	Striphtmlcomments string `json:"striphtmlcomments,omitempty"`
	/**
	* Strip XML comments before forwarding a web page sent by a protected web site in response to a user request.
	*/
	Stripxmlcomments string `json:"stripxmlcomments,omitempty"`
	/**
	* Exempt URLs that pass the Start URL closure check from SQL injection, cross-site script, field format and field consistency security checks at locations other than headers.
	*/
	Exemptclosureurlsfromsecuritychecks string `json:"exemptclosureurlsfromsecuritychecks,omitempty"`
	/**
	* Default character set for protected web pages. Web pages sent by your protected web sites in response to user requests are assigned this character set if the page does not already specify a character set. The character sets supported by the application firewall are:
		* iso-8859-1 (English US)
		* big5 (Chinese Traditional)
		* gb2312 (Chinese Simplified)
		* sjis (Japanese Shift-JIS)
		* euc-jp (Japanese EUC-JP)
		* iso-8859-9 (Turkish)
		* utf-8 (Unicode)
		* euc-kr (Korean)
	*/
	Defaultcharset string `json:"defaultcharset,omitempty"`
	/**
	* Expression to get the client IP.
	*/
	Clientipexpression string `json:"clientipexpression,omitempty"`
	/**
	* One or more security checks. Available options are as follows:
		* SQLInjection - Enable dynamic learning for SQLInjection security check.
		* CrossSiteScripting - Enable dynamic learning for CrossSiteScripting security check.
		* fieldFormat - Enable dynamic learning for  fieldFormat security check.
		* None - Disable security checks for all security checks.
		CLI users: To enable dynamic learning on one or more security checks, type "set appfw profile -dynamicLearning" followed by the security checks to be enabled. To turn off dynamic learning on all security checks, type "set appfw profile -dynamicLearning none".
	*/
	Dynamiclearning []string `json:"dynamiclearning,omitempty"`
	/**
	* Maximum allowed HTTP post body size, in bytes. Maximum supported value is 10GB. Citrix recommends enabling streaming option for large values of post body limit (>20MB).
	*/
	Postbodylimit int `json:"postbodylimit,omitempty"`
	/**
	* One or more Post Body Limit actions. Available settings function as follows:
		* Block - Block connections that violate this security check. Must always be set.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -PostBodyLimitAction block" followed by the other actions to be enabled.
	*/
	Postbodylimitaction []string `json:"postbodylimitaction,omitempty"`
	/**
	* Maximum allowed HTTP post body size for signature inspection for location HTTP_POST_BODY in the signatures, in bytes. Note that the changes in value could impact CPU and latency profile.
	*/
	Postbodylimitsignature int `json:"postbodylimitsignature,omitempty"`
	/**
	* Maximum allowed number of file uploads per form-submission request. The maximum setting (65535) allows an unlimited number of uploads.
	*/
	Fileuploadmaxnum int `json:"fileuploadmaxnum,omitempty"`
	/**
	* Perform HTML entity encoding for any special characters in responses sent by your protected web sites.
	*/
	Canonicalizehtmlresponse string `json:"canonicalizehtmlresponse,omitempty"`
	/**
	* Enable tagging of web form fields for use by the Form Field Consistency and CSRF Form Tagging checks.
	*/
	Enableformtagging string `json:"enableformtagging,omitempty"`
	/**
	* Perform sessionless Field Consistency Checks.
	*/
	Sessionlessfieldconsistency string `json:"sessionlessfieldconsistency,omitempty"`
	/**
	* Enable session less URL Closure Checks.
		This check is applicable to Profile Type: HTML. 
	*/
	Sessionlessurlclosure string `json:"sessionlessurlclosure,omitempty"`
	/**
	* Allow ';' as a form field separator in URL queries and POST form bodies. 
	*/
	Semicolonfieldseparator string `json:"semicolonfieldseparator,omitempty"`
	/**
	* Exclude uploaded files from Form checks.
	*/
	Excludefileuploadfromchecks string `json:"excludefileuploadfromchecks,omitempty"`
	/**
	* Parse HTML comments and exempt them from the HTML SQL Injection check. You must specify the type of comments that the application firewall is to detect and exempt from this security check. Available settings function as follows:
		* Check all - Check all content.
		* ANSI - Exempt content that is part of an ANSI (Mozilla-style) comment.
		* Nested - Exempt content that is part of a nested (Microsoft-style) comment.
		* ANSI Nested - Exempt content that is part of any type of comment.
	*/
	Sqlinjectionparsecomments string `json:"sqlinjectionparsecomments,omitempty"`
	/**
	* Configure the method that the application firewall uses to handle percent-encoded names and values. Available settings function as follows:
		* asp_mode - Microsoft ASP format.
		* secure_mode - Secure format.
	*/
	Invalidpercenthandling string `json:"invalidpercenthandling,omitempty"`
	/**
	* Application firewall profile type, which controls which security checks and settings are applied to content that is filtered with the profile. Available settings function as follows:
		* HTML - HTML-based web sites.
		* XML -  XML-based web sites and services.
		* JSON - JSON-based web sites and services.
		* HTML XML (Web 2.0) - Sites that contain both HTML and XML content, such as ATOM feeds, blogs, and RSS feeds.
		* HTML JSON  - Sites that contain both HTML and JSON content.
		* XML JSON   - Sites that contain both XML and JSON content.
		* HTML XML JSON   - Sites that contain HTML, XML and JSON content.
	*/
	Type []string `json:"type,omitempty"`
	/**
	* Check request headers as well as web forms for injected SQL and cross-site scripts.
	*/
	Checkrequestheaders string `json:"checkrequestheaders,omitempty"`
	/**
	* Inspect request query as well as web forms for injected SQL and cross-site scripts for following content types.
	*/
	Inspectquerycontenttypes []string `json:"inspectquerycontenttypes,omitempty"`
	/**
	* Optimize handle of HTTP partial requests i.e. those with range headers.
		Available settings are as follows:
		* ON  - Partial requests by the client result in partial requests to the backend server in most cases.
		* OFF - Partial requests by the client are changed to full requests to the backend server
	*/
	Optimizepartialreqs string `json:"optimizepartialreqs,omitempty"`
	/**
	* URL Decode request cookies before subjecting them to SQL and cross-site scripting checks.
	*/
	Urldecoderequestcookies string `json:"urldecoderequestcookies,omitempty"`
	/**
	* Any comments about the purpose of profile, or other useful information about the profile.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Configure whether the application firewall should use percentage recursive decoding
	*/
	Percentdecoderecursively string `json:"percentdecoderecursively,omitempty"`
	/**
	* One or more multiple header actions. Available settings function as follows:
		* Block - Block connections that have multiple headers.
		* Log - Log connections that have multiple headers.
		* KeepLast - Keep only last header when multiple headers are present.
		Request headers inspected:
		* Accept-Encoding
		* Content-Encoding
		* Content-Range
		* Content-Type
		* Host
		* Range
		* Referer
		CLI users: To enable one or more actions, type "set appfw profile -multipleHeaderAction" followed by the actions to be enabled.
	*/
	Multipleheaderaction []string `json:"multipleheaderaction,omitempty"`
	/**
	* Object name of the rfc profile.
	*/
	Rfcprofile string `json:"rfcprofile,omitempty"`
	/**
	* One or more file upload types actions. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -fileUploadTypeAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -fileUploadTypeAction none".
	*/
	Fileuploadtypesaction []string `json:"fileuploadtypesaction,omitempty"`
	/**
	* Detailed Logging Verbose Log Level.
	*/
	Verboseloglevel string `json:"verboseloglevel,omitempty"`
	/**
	* Configure whether application firewall should add samesite attribute for set-cookies
	*/
	Insertcookiesamesiteattribute string `json:"insertcookiesamesiteattribute,omitempty"`
	/**
	* Cookie Samesite attribute added to support adding cookie SameSite attribute for all set-cookies including appfw session cookies. Default value will be "SameSite=Lax".
	*/
	Cookiesamesiteattribute string `json:"cookiesamesiteattribute,omitempty"`
	/**
	* Specifies SQL Injection rule type: ALLOW/DENY. If ALLOW rule type is configured then allow list rules are used, if DENY rule type is configured then deny rules are used.
	*/
	Sqlinjectionruletype string `json:"sqlinjectionruletype,omitempty"`
	/**
	* Fake account detection flag : ON/OFF. If set to ON fake account detection in enabled on ADC, if set to OFF fake account detection is disabled.
	*/
	Fakeaccountdetection string `json:"fakeaccountdetection,omitempty"`
	/**
	* Enable Geo-Location Logging in CEF format logs for the profile.
	*/
	Geolocationlogging string `json:"geolocationlogging,omitempty"`
	/**
	* Enable CEF format logs for the profile.
	*/
	Ceflogging string `json:"ceflogging,omitempty"`
	/**
	* Block Keyword action. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -blockKeywordAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -blockKeywordAction none".
	*/
	Blockkeywordaction []string `json:"blockkeywordaction,omitempty"`
	/**
	* JSON Block Keyword action. Available settings function as follows:
		* Block - Block connections that violate this security check.
		* Log - Log violations of this security check.
		* Stats - Generate statistics for this security check.
		* None - Disable all actions for this security check.
		CLI users: To enable one or more actions, type "set appfw profile -JSONBlockKeywordAction" followed by the actions to be enabled. To turn off all actions, type "set appfw profile -JSONBlockKeywordAction none".
	*/
	Jsonblockkeywordaction []string `json:"jsonblockkeywordaction,omitempty"`
	/**
	* Enable bypass list for the profile.
	*/
	Asprofbypasslistenable string `json:"as_prof_bypass_list_enable,omitempty"`
	/**
	* Enable deny list for the profile.
	*/
	Asprofdenylistenable string `json:"as_prof_deny_list_enable,omitempty"`
	/**
	* Name of the session cookie that the application firewall uses to track user sessions.
		Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cookie name" or 'my cookie name').
	*/
	Sessioncookiename string `json:"sessioncookiename,omitempty"`
	/**
	* Source for tar archive.
	*/
	Archivename string `json:"archivename,omitempty"`
	/**
	* Import all appfw relaxation rules
	*/
	Relaxationrules bool `json:"relaxationrules,omitempty"`
	/**
	* Name of the profile which will be created/updated to associate the relaxation rules
	*/
	Importprofilename string `json:"importprofilename,omitempty"`
	/**
	* Match this action url in archived Relaxation Rules to replace.
	*/
	Matchurlstring string `json:"matchurlstring,omitempty"`
	/**
	* Replace matched url string with this action url string while restoring Relaxation Rules
	*/
	Replaceurlstring string `json:"replaceurlstring,omitempty"`
	/**
	* Purge existing Relaxation Rules and replace during import
	*/
	Overwrite bool `json:"overwrite,omitempty"`
	/**
	* Augment Relaxation Rules during import
	*/
	Augment bool `json:"augment,omitempty"`

	//------- Read only Parameter ---------;

	State string `json:"state,omitempty"`
	Learning string `json:"learning,omitempty"`
	Csrftag string `json:"csrftag,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
