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
* Binding class showing the fileuploadtype that can be bound to appfwprofile.
*/
type Appfwprofilefileuploadtypebinding struct {
	/**
	* FileUploadTypes to allow/deny.
	*/
	Fileuploadtype string `json:"fileuploadtype,omitempty"`
	/**
	* FileUploadTypes action URL.
	*/
	Asfileuploadtypesurl string `json:"as_fileuploadtypes_url,omitempty"`
	/**
	* Is field name a regular expression?
	*/
	Isnameregex string `json:"isnameregex,omitempty"`
	/**
	* Enabled.
	*/
	State string `json:"state,omitempty"`
	/**
	* Any comments about the purpose of profile, or other useful information about the profile.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* FileUploadTypes file types.
	*/
	Filetype []string `json:"filetype,omitempty"`
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
	* Is a regular expression?
	*/
	Isregexfileuploadtypesurl string `json:"isregex_fileuploadtypes_url,omitempty"`
	/**
	* Specifies rule type of binding
	*/
	Ruletype string `json:"ruletype,omitempty"`


}