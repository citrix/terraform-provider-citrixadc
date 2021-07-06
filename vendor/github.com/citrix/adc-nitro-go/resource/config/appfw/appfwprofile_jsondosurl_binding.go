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
* Binding class showing the jsondosurl that can be bound to appfwprofile.
*/
type Appfwprofilejsondosurlbinding struct {
	/**
	* The URL on which we need to enforce the specified JSON denial-of-service (JSONDoS) attack protections.
		An JSON DoS configuration consists of the following items:
		* URL. PCRE-format regular expression for the URL.
		* Maximum-document-length-check toggle.  ON to enable this check, OFF to disable it.
		* Maximum document length. Positive integer representing the maximum length of the JSON document.
		* Maximum-container-depth-check toggle. ON to enable, OFF to disable.
		* Maximum container depth. Positive integer representing the maximum container depth of the JSON document.
		* Maximum-object-key-count-check toggle. ON to enable, OFF to disable.
		* Maximum object key count. Positive integer representing the maximum allowed number of keys in any of the  JSON object.
		* Maximum-object-key-length-check toggle. ON to enable, OFF to disable.
		* Maximum object key length. Positive integer representing the maximum allowed length of key in any of the  JSON object.
		* Maximum-array-value-count-check toggle. ON to enable, OFF to disable.
		* Maximum array value count. Positive integer representing the maximum allowed number of values in any of the JSON array.
		* Maximum-string-length-check toggle. ON to enable, OFF to disable.
		* Maximum string length. Positive integer representing the maximum length of string in JSON.
	*/
	Jsondosurl string `json:"jsondosurl,omitempty"`
	/**
	* State if JSON Max document length check is ON or OFF.
	*/
	Jsonmaxdocumentlengthcheck string `json:"jsonmaxdocumentlengthcheck,omitempty"`
	/**
	* Maximum document length of JSON document, in bytes.
	*/
	Jsonmaxdocumentlength int `json:"jsonmaxdocumentlength,omitempty"`
	/**
	* State if JSON Max depth check is ON or OFF.
	*/
	Jsonmaxcontainerdepthcheck string `json:"jsonmaxcontainerdepthcheck,omitempty"`
	/**
	* Maximum allowed nesting depth  of JSON document. JSON allows one to nest the containers (object and array) in any order to any depth. This check protects against documents that have excessive depth of hierarchy.
	*/
	Jsonmaxcontainerdepth int `json:"jsonmaxcontainerdepth,omitempty"`
	/**
	* State if JSON Max object key count check is ON or OFF.
	*/
	Jsonmaxobjectkeycountcheck string `json:"jsonmaxobjectkeycountcheck,omitempty"`
	/**
	* Maximum key count in the any of JSON object. This check protects against objects that have large number of keys.
	*/
	Jsonmaxobjectkeycount int `json:"jsonmaxobjectkeycount,omitempty"`
	/**
	* State if JSON Max object key length check is ON or OFF.
	*/
	Jsonmaxobjectkeylengthcheck string `json:"jsonmaxobjectkeylengthcheck,omitempty"`
	/**
	* Maximum key length in the any of JSON object. This check protects against objects that have large keys.
	*/
	Jsonmaxobjectkeylength int `json:"jsonmaxobjectkeylength,omitempty"`
	/**
	* State if JSON Max array value count check is ON or OFF.
	*/
	Jsonmaxarraylengthcheck string `json:"jsonmaxarraylengthcheck,omitempty"`
	/**
	* Maximum array length in the any of JSON object. This check protects against arrays having large lengths.
	*/
	Jsonmaxarraylength int `json:"jsonmaxarraylength,omitempty"`
	/**
	* State if JSON Max string value count check is ON or OFF.
	*/
	Jsonmaxstringlengthcheck string `json:"jsonmaxstringlengthcheck,omitempty"`
	/**
	* Maximum string length in the JSON. This check protects against strings that have large length.
	*/
	Jsonmaxstringlength int `json:"jsonmaxstringlength,omitempty"`
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