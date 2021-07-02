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
* Binding class showing the crosssitescripting that can be bound to appfwprofile.
*/
type Appfwprofilecrosssitescriptingbinding struct {
	/**
	* The web form field name.
	*/
	Crosssitescripting string `json:"crosssitescripting,omitempty"`
	/**
	* Is the web form field name a regular expression?
	*/
	Isregexxss string `json:"isregex_xss,omitempty"`
	/**
	* The web form action URL.
	*/
	Formactionurlxss string `json:"formactionurl_xss,omitempty"`
	/**
	* Location of cross-site scripting exception - form field, header, cookie or URL.
	*/
	Asscanlocationxss string `json:"as_scan_location_xss,omitempty"`
	/**
	* The web form value type.
	*/
	Asvaluetypexss string `json:"as_value_type_xss,omitempty"`
	/**
	* The web form value expression.
	*/
	Asvalueexprxss string `json:"as_value_expr_xss,omitempty"`
	/**
	* Is the web form field value a regular expression?
	*/
	Isvalueregexxss string `json:"isvalueregex_xss,omitempty"`
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
	* A "id" that identifies the rule.
	*/
	Resourceid string `json:"resourceid,omitempty"`
	/**
	* Name of the profile to which to bind an exemption or rule.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Specifies rule type of binding
	*/
	Ruletype string `json:"ruletype,omitempty"`


}