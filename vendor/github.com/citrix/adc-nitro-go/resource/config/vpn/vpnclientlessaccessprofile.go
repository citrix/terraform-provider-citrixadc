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

package vpn

/**
* Configuration for Clientless VPN rewrite profile resource.
*/
type Vpnclientlessaccessprofile struct {
	/**
	* Name for the Citrix Gateway clientless access profile. Must begin with an ASCII alphabetic or underscore (_) character, and must consist only of ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the profile is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
	*/
	Profilename string `json:"profilename,omitempty"`
	/**
	* Name of the configured URL rewrite policy label. If you do not specify a policy label name, then URLs are not rewritten.
	*/
	Urlrewritepolicylabel string `json:"urlrewritepolicylabel,omitempty"`
	/**
	* Name of the configured JavaScript rewrite policy label.  If you do not specify a policy label name, then JAVA scripts are not rewritten.
	*/
	Javascriptrewritepolicylabel string `json:"javascriptrewritepolicylabel,omitempty"`
	/**
	* Name of the configured Request rewrite policy label.  If you do not specify a policy label name, then requests are not rewritten.
	*/
	Reqhdrrewritepolicylabel string `json:"reqhdrrewritepolicylabel,omitempty"`
	/**
	* Name of the configured Response rewrite policy label.
	*/
	Reshdrrewritepolicylabel string `json:"reshdrrewritepolicylabel,omitempty"`
	/**
	* Name of the pattern set that contains the regular expressions, which match the URL in Java script.
	*/
	Regexforfindingurlinjavascript string `json:"regexforfindingurlinjavascript,omitempty"`
	/**
	* Name of the pattern set that contains the regular expressions, which match the URL in the CSS.
	*/
	Regexforfindingurlincss string `json:"regexforfindingurlincss,omitempty"`
	/**
	* Name of the pattern set that contains the regular expressions, which match the URL in X Component.
	*/
	Regexforfindingurlinxcomponent string `json:"regexforfindingurlinxcomponent,omitempty"`
	/**
	* Name of the pattern set that contains the regular expressions, which match the URL in XML.
	*/
	Regexforfindingurlinxml string `json:"regexforfindingurlinxml,omitempty"`
	/**
	* Name of the pattern set that contains the regular expressions, which match the URLs in the custom content type other than HTML, CSS, XML, XCOMP, and JavaScript. The custom content type should be included in the patset ns_cvpn_custom_content_types.
	*/
	Regexforfindingcustomurls string `json:"regexforfindingcustomurls,omitempty"`
	/**
	* Specify the name of the pattern set containing the names of the cookies, which are allowed between the client and the server. If a pattern set is not specified, Citrix Gateway does not allow any cookies between the client and the server. A cookie that is not specified in the pattern set is handled by Citrix Gateway on behalf of the client.
	*/
	Clientconsumedcookies string `json:"clientconsumedcookies,omitempty"`
	/**
	* Specify whether a persistent session cookie is set and accepted for clientless access. If this parameter is set to ON, COM objects, such as MSOffice, which are invoked by the browser can access the files using clientless access. Use caution because the persistent cookie is stored on the disk.
	*/
	Requirepersistentcookie string `json:"requirepersistentcookie,omitempty"`

	//------- Read only Parameter ---------;

	Cssrewritepolicylabel string `json:"cssrewritepolicylabel,omitempty"`
	Xmlrewritepolicylabel string `json:"xmlrewritepolicylabel,omitempty"`
	Xcomponentrewritepolicylabel string `json:"xcomponentrewritepolicylabel,omitempty"`
	Isdefault string `json:"isdefault,omitempty"`
	Description string `json:"description,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
