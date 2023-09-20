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
* Binding class showing the jsonsqlurl that can be bound to appfwprofile.
*/
type Appfwprofilejsonsqlurlbinding struct {
	/**
	* A regular expression that designates a URL on the Json SQL URL list for which SQL violations are relaxed.
		Enclose URLs in double quotes to ensure preservation of any embedded spaces or non-alphanumeric characters.
	*/
	Jsonsqlurl string `json:"jsonsqlurl,omitempty"`
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
	/**
	* Is the JSON SQL key value a regular expression?
	*/
	Isvalueregex_json_sql string `json:"isvalueregex_json_sql,omitempty"`
	/**
	* Is the key name a regular expression?
	*/
	Iskeyregex_json_sql string `json:"iskeyregex_json_sql,omitempty"`
	/**
	* Type of the relaxed JSON SQL key value.
	*/
	As_value_type_json_sql string `json:"as_value_type_json_sql,omitempty"`
	/**
	* The JSON SQL key value expression.
	*/
	As_value_expr_json_sql string `json:"as_value_expr_json_sql,omitempty"`
	/**
	* An expression that designates a keyname on the JSON SQL URL for which SQL injection violations are relaxed.
	*/
	Keyname_json_sql string `json:"keyname_json_sql,omitempty"`

}