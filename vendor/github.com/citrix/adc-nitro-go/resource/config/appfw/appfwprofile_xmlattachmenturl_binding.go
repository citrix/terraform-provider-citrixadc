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
* Binding class showing the xmlattachmenturl that can be bound to appfwprofile.
*/
type Appfwprofilexmlattachmenturlbinding struct {
	/**
	* XML attachment URL regular expression length.
	*/
	Xmlattachmenturl string `json:"xmlattachmenturl,omitempty"`
	/**
	* State if XML Max attachment size Check is ON or OFF. Protects against XML requests with large attachment data.
	*/
	Xmlmaxattachmentsizecheck string `json:"xmlmaxattachmentsizecheck,omitempty"`
	/**
	* Specify maximum attachment size.
	*/
	Xmlmaxattachmentsize *int `json:"xmlmaxattachmentsize,omitempty"`
	/**
	* State if XML attachment content-type check is ON or OFF. Protects against XML requests with illegal attachments.
	*/
	Xmlattachmentcontenttypecheck string `json:"xmlattachmentcontenttypecheck,omitempty"`
	/**
	* Specify content-type regular expression.
	*/
	Xmlattachmentcontenttype string `json:"xmlattachmentcontenttype,omitempty"`
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