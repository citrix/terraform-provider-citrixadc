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
* Binding class showing the sqlinjection that can be bound to appfwprofile.
*/
type Appfwprofilesqlinjectionbinding struct {
	/**
	* The web form field name.
	*/
	Sqlinjection string `json:"sqlinjection,omitempty"`
	/**
	* Is the web form field name a regular expression?
	*/
	Isregexsql string `json:"isregex_sql,omitempty"`
	/**
	* The web form action URL.
	*/
	Formactionurlsql string `json:"formactionurl_sql,omitempty"`
	/**
	* Location of SQL injection exception - form field, header or cookie.
	*/
	Asscanlocationsql string `json:"as_scan_location_sql,omitempty"`
	/**
	* The web form value type.
	*/
	Asvaluetypesql string `json:"as_value_type_sql,omitempty"`
	/**
	* The web form value expression.
	*/
	Asvalueexprsql string `json:"as_value_expr_sql,omitempty"`
	/**
	* Is the web form field value a regular expression?
	*/
	Isvalueregexsql string `json:"isvalueregex_sql,omitempty"`
	/**
	* Specifies rule type of binding
	*/
	Ruletype string `json:"ruletype,omitempty"`
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


}