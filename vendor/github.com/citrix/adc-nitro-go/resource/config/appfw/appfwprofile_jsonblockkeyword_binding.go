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
* Binding class showing the jsonblockkeyword that can be bound to appfwprofile.
*/
type Appfwprofilejsonblockkeywordbinding struct {
	/**
	* Field name of json block keyword binding
	*/
	Jsonblockkeyword string `json:"jsonblockkeyword,omitempty"`
	/**
	* The json blockkeyword URL.
	*/
	Jsonblockkeywordurl string `json:"jsonblockkeywordurl,omitempty"`
	/**
	* JSON block keyword keyname
	*/
	Keynamejsonblockkeyword string `json:"keyname_json_blockkeyword,omitempty"`
	/**
	* Is JSON blockkeyword key a regular expression?
	*/
	Iskeyregexjsonblockkeyword string `json:"iskeyregex_json_blockkeyword,omitempty"`
	/**
	* JSON block keyword type
	*/
	Jsonblockkeywordtype string `json:"jsonblockkeywordtype,omitempty"`
	/**
	* Enabled.
	*/
	State string `json:"state,omitempty"`
	/**
	* Any comments about the purpose of profile, or other useful information about the profile.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Specifies rule type of binding
	*/
	Ruletype string `json:"ruletype,omitempty"`
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