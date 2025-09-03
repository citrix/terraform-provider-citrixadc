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

package snmp

/**
* Configuration for SNMP mib resource.
*/
type Snmpmib struct {
	/**
	* Name of the administrator for this Citrix ADC. Along with the name, you can include information on how to contact this person, such as a phone number or an email address. Can consist of 1 to 127 characters that include uppercase and  lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the information includes one or more spaces, enclose it in double or single quotation marks (for example, "my contact" or 'my contact').
	*/
	Contact string `json:"contact,omitempty"`
	/**
	* Name for this Citrix ADC. Can consist of 1 to 127 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  You should choose a name that helps identify the Citrix ADC appliance.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose it in double or single quotation marks (for example, "my name" or 'my name').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Physical location of the Citrix ADC. For example, you can specify building name, lab number, and rack number. Can consist of 1 to 127 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the location includes one or more spaces, enclose it in double or single quotation marks (for example, "my location" or 'my location').
	*/
	Location string `json:"location,omitempty"`
	/**
	* Custom identification number for the Citrix ADC. Can consist of 1 to 127 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters. You should choose a custom identification that helps identify the Citrix ADC appliance.
		The following requirement applies only to the Citrix ADC CLI:
		If the ID includes one or more spaces, enclose it in double or single quotation marks (for example, "my ID" or 'my ID').
	*/
	Customid string `json:"customid,omitempty"`
	/**
	* ID of the cluster node for which we are setting the mib. This is a mandatory argument to set snmp mib on CLIP.
	*/
	Ownernode int `json:"ownernode,omitempty"`

	//------- Read only Parameter ---------;

	Sysdesc string `json:"sysdesc,omitempty"`
	Sysuptime string `json:"sysuptime,omitempty"`
	Sysservices string `json:"sysservices,omitempty"`
	Sysoid string `json:"sysoid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
