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
* Binding class showing the cmdinjection that can be bound to appfwprofile.
*/
type Appfwprofilecmdinjectionbinding struct {
	/**
	* Name of the relaxed web form field/header/cookie
	*/
	Cmdinjection string `json:"cmdinjection,omitempty"`
	/**
	* Is the relaxed web form field name/header/cookie a regular expression?
	*/
	Isregexcmd string `json:"isregex_cmd,omitempty"`
	/**
	* The web form action URL.
	*/
	Formactionurlcmd string `json:"formactionurl_cmd,omitempty"`
	/**
	* Location of command injection exception - form field, header or cookie.
	*/
	Asscanlocationcmd string `json:"as_scan_location_cmd,omitempty"`
	/**
	* Type of the relaxed web form value
	*/
	Asvaluetypecmd string `json:"as_value_type_cmd,omitempty"`
	/**
	* The web form/header/cookie value expression.
	*/
	Asvalueexprcmd string `json:"as_value_expr_cmd,omitempty"`
	/**
	* Is the web form field/header/cookie value a regular expression?
	*/
	Isvalueregexcmd string `json:"isvalueregex_cmd,omitempty"`
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