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
* Binding class showing the xmlxss that can be bound to appfwprofile.
*/
type Appfwprofilexmlxssbinding struct {
	/**
	* Exempt the specified URL from the XML cross-site scripting (XSS) check.
		An XML cross-site scripting exemption (relaxation) consists of the following items:
		* URL. URL to exempt, as a string or a PCRE-format regular expression.
		* ISREGEX flag. REGEX if URL is a regular expression, NOTREGEX if URL is a fixed string.
		* Location. ELEMENT if the attachment is located in an XML element, ATTRIBUTE if located in an XML attribute.
	*/
	Xmlxss string `json:"xmlxss,omitempty"`
	/**
	* Is the XML XSS exempted field name a regular expression?
	*/
	Isregexxmlxss string `json:"isregex_xmlxss,omitempty"`
	/**
	* Location of XSS injection exception - XML Element or Attribute.
	*/
	Asscanlocationxmlxss string `json:"as_scan_location_xmlxss,omitempty"`
	/**
	* Enabled.
	*/
	State string `json:"state,omitempty"`
	/**
	* Any comments about the purpose of profile, or other useful information about the profile.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Is the rule auto deployed by dynamic profile ?
	*/
	Isautodeployed string `json:"isautodeployed,omitempty"`
	/**
	* Send SNMP alert?
	*/
	Alertonly string `json:"alertonly,omitempty"`
	/**
	* Name of the profile to which to bind an exemption or rule.
	*/
	Name string `json:"name,omitempty"`
	/**
	* A "id" that identifies the rule.
	*/
	Resourceid string `json:"resourceid,omitempty"`
	/**
	* Specifies rule type of binding
	*/
	Ruletype string `json:"ruletype,omitempty"`


}