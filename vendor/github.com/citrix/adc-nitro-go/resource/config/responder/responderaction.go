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

package responder

/**
* Configuration for responder action resource.
*/
type Responderaction struct {
	/**
	* Name for the responder action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the responder policy is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my responder action" or 'my responder action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Type of responder action. Available settings function as follows:
		* respondwith <target> - Respond to the request with the expression specified as the target.
		* respondwithhtmlpage - Respond to the request with the uploaded HTML page object specified as the target.
		* redirect - Redirect the request to the URL specified as the target.
		* sqlresponse_ok - Send an SQL OK response.
		* sqlresponse_error - Send an SQL ERROR response.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Expression specifying what to respond with. Typically a URL for redirect policies or a default-syntax expression.  In addition to Citrix ADC default-syntax expressions that refer to information in the request, a stringbuilder expression can contain text and HTML, and simple escape codes that define new lines and paragraphs. Enclose each stringbuilder expression element (either a Citrix ADC default-syntax expression or a string) in double quotation marks. Use the plus (+) character to join the elements.
		Examples:
		1) Respondwith expression that sends an HTTP 1.1 200 OK response:
		"HTTP/1.1 200 OK\r\n\r\n"
		2) Redirect expression that redirects user to the specified web host and appends the request URL to the redirect.
		"http://backupsite2.com" + HTTP.REQ.URL
		3) Respondwith expression that sends an HTTP 1.1 404 Not Found response with the request URL included in the response:
		"HTTP/1.1 404 Not Found\r\n\r\n"+ "HTTP.REQ.URL.HTTP_URL_SAFE" + "does not exist on the web server."
		The following requirement applies only to the Citrix ADC CLI:
		Enclose the entire expression in single quotation marks. (Citrix ADC expression elements should be included inside the single quotation marks for the entire expression, but do not need to be enclosed in double quotation marks.)
	*/
	Target string `json:"target,omitempty"`
	/**
	* For respondwithhtmlpage policies, name of the HTML page object to use as the response. You must first import the page object.
	*/
	Htmlpage string `json:"htmlpage,omitempty"`
	/**
	* Bypass the safety check, allowing potentially unsafe expressions. An unsafe expression in a response is one that contains references to request elements that might not be present in all requests. If a response refers to a missing request element, an empty string is used instead.
	*/
	Bypasssafetycheck string `json:"bypasssafetycheck,omitempty"`
	/**
	* Comment. Any type of information about this responder action.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* HTTP response status code, for example 200, 302, 404, etc. The default value for the redirect action type is 302 and for respondwithhtmlpage is 200
	*/
	Responsestatuscode int `json:"responsestatuscode,omitempty"`
	/**
	* Expression specifying the reason phrase of the HTTP response. The reason phrase may be a string literal with quotes or a PI expression. For example: "Invalid URL: " + HTTP.REQ.URL
	*/
	Reasonphrase string `json:"reasonphrase,omitempty"`
	/**
	* One or more headers to insert into the HTTP response. Each header is specified as "name(expr)", where expr is an expression that is evaluated at runtime to provide the value for the named header. You can configure a maximum of eight headers for a responder action.
	*/
	Headers []string `json:"headers,omitempty"`
	/**
	* New name for the responder action.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my responder action" or my responder action').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Referencecount string `json:"referencecount,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
