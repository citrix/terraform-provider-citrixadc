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
* Binding class showing the xmldosurl that can be bound to appfwprofile.
*/
type Appfwprofilexmldosurlbinding struct {
	/**
	* XML DoS URL regular expression length.
	*/
	Xmldosurl string `json:"xmldosurl,omitempty"`
	/**
	* State if XML Max element depth check is ON or OFF.
	*/
	Xmlmaxelementdepthcheck string `json:"xmlmaxelementdepthcheck,omitempty"`
	/**
	* Maximum nesting (depth) of XML elements. This check protects against documents that have excessive hierarchy depths.
	*/
	Xmlmaxelementdepth *int `json:"xmlmaxelementdepth,omitempty"`
	/**
	* State if XML Max element name length check is ON or OFF.
	*/
	Xmlmaxelementnamelengthcheck string `json:"xmlmaxelementnamelengthcheck,omitempty"`
	/**
	* Specify the longest name of any element (including the expanded namespace) to protect against overflow attacks.
	*/
	Xmlmaxelementnamelength *int `json:"xmlmaxelementnamelength,omitempty"`
	/**
	* State if XML Max elements check is ON or OFF.
	*/
	Xmlmaxelementscheck string `json:"xmlmaxelementscheck,omitempty"`
	/**
	* Specify the maximum number of XML elements allowed. Protects against overflow attacks.
	*/
	Xmlmaxelements *int `json:"xmlmaxelements,omitempty"`
	/**
	* State if XML Max element children check is ON or OFF.
	*/
	Xmlmaxelementchildrencheck string `json:"xmlmaxelementchildrencheck,omitempty"`
	/**
	* Specify the maximum number of children allowed per XML element. Protects against overflow attacks.
	*/
	Xmlmaxelementchildren *int `json:"xmlmaxelementchildren,omitempty"`
	/**
	* State if XML Max nodes check is ON or OFF.
	*/
	Xmlmaxnodescheck string `json:"xmlmaxnodescheck,omitempty"`
	/**
	* Specify the maximum number of XML nodes. Protects against overflow attacks.
	*/
	Xmlmaxnodes *int `json:"xmlmaxnodes,omitempty"`
	/**
	* State if XML Max Entity Expansions Check is ON or OFF.
	*/
	Xmlmaxentityexpansionscheck string `json:"xmlmaxentityexpansionscheck,omitempty"`
	/**
	* Specify maximum allowed number of entity expansions. Protects aganist Entity Expansion Attack.
	*/
	Xmlmaxentityexpansions *int `json:"xmlmaxentityexpansions,omitempty"`
	/**
	* State if XML Max Entity Expansions Depth Check is ON or OFF.
	*/
	Xmlmaxentityexpansiondepthcheck string `json:"xmlmaxentityexpansiondepthcheck,omitempty"`
	/**
	* Specify maximum entity expansion depth. Protects aganist Entity Expansion Attack.
	*/
	Xmlmaxentityexpansiondepth *int `json:"xmlmaxentityexpansiondepth,omitempty"`
	/**
	* State if XML Max attributes check is ON or OFF.
	*/
	Xmlmaxattributescheck string `json:"xmlmaxattributescheck,omitempty"`
	/**
	* Specify maximum number of attributes per XML element. Protects against overflow attacks.
	*/
	Xmlmaxattributes *int `json:"xmlmaxattributes,omitempty"`
	/**
	* State if XML Max attribute name length check is ON or OFF.
	*/
	Xmlmaxattributenamelengthcheck string `json:"xmlmaxattributenamelengthcheck,omitempty"`
	/**
	* Specify the longest name of any XML attribute. Protects against overflow attacks.
	*/
	Xmlmaxattributenamelength *int `json:"xmlmaxattributenamelength,omitempty"`
	/**
	* State if XML Max atribute value length is ON or OFF.
	*/
	Xmlmaxattributevaluelengthcheck string `json:"xmlmaxattributevaluelengthcheck,omitempty"`
	/**
	* Specify the longest value of any XML attribute. Protects against overflow attacks.
	*/
	Xmlmaxattributevaluelength *int `json:"xmlmaxattributevaluelength,omitempty"`
	/**
	* State if XML Max namespaces check is ON or OFF.
	*/
	Xmlmaxnamespacescheck string `json:"xmlmaxnamespacescheck,omitempty"`
	/**
	* Specify maximum number of active namespaces. Protects against overflow attacks.
	*/
	Xmlmaxnamespaces *int `json:"xmlmaxnamespaces,omitempty"`
	/**
	* State if XML Max namespace URI length check is ON or OFF.
	*/
	Xmlmaxnamespaceurilengthcheck string `json:"xmlmaxnamespaceurilengthcheck,omitempty"`
	/**
	* Specify the longest URI of any XML namespace. Protects against overflow attacks.
	*/
	Xmlmaxnamespaceurilength *int `json:"xmlmaxnamespaceurilength,omitempty"`
	/**
	* State if XML Max CDATA length check is ON or OFF.
	*/
	Xmlmaxchardatalengthcheck string `json:"xmlmaxchardatalengthcheck,omitempty"`
	/**
	* Specify the maximum size of CDATA. Protects against overflow attacks and large quantities of unparsed data within XML messages.
	*/
	Xmlmaxchardatalength *int `json:"xmlmaxchardatalength,omitempty"`
	/**
	* State if XML Max file size check is ON or OFF.
	*/
	Xmlmaxfilesizecheck string `json:"xmlmaxfilesizecheck,omitempty"`
	/**
	* Specify the maximum size of XML messages. Protects against overflow attacks.
	*/
	Xmlmaxfilesize *int `json:"xmlmaxfilesize,omitempty"`
	/**
	* State if XML Min file size check is ON or OFF.
	*/
	Xmlminfilesizecheck string `json:"xmlminfilesizecheck,omitempty"`
	/**
	* Enforces minimum message size.
	*/
	Xmlminfilesize *int `json:"xmlminfilesize,omitempty"`
	/**
	* State if XML Block PI is ON or OFF. Protects resources from denial of service attacks as SOAP messages cannot have processing instructions (PI) in messages.
	*/
	Xmlblockpi string `json:"xmlblockpi,omitempty"`
	/**
	* State if XML DTD is ON or OFF. Protects against recursive Document Type Declaration (DTD) entity expansion attacks. Also, SOAP messages cannot have DTDs in messages. 
	*/
	Xmlblockdtd string `json:"xmlblockdtd,omitempty"`
	/**
	* State if XML Block External Entities Check is ON or OFF. Protects against XML External Entity (XXE) attacks that force applications to parse untrusted external entities (sources) in XML documents.
	*/
	Xmlblockexternalentities string `json:"xmlblockexternalentities,omitempty"`
	/**
	* State if XML SOAP Array check is ON or OFF.
	*/
	Xmlsoaparraycheck string `json:"xmlsoaparraycheck,omitempty"`
	/**
	* XML Max Total SOAP Array Size. Protects against SOAP Array Abuse attack.
	*/
	Xmlmaxsoaparraysize *int `json:"xmlmaxsoaparraysize,omitempty"`
	/**
	* XML Max Individual SOAP Array Rank. This is the dimension of the SOAP array.
	*/
	Xmlmaxsoaparrayrank *int `json:"xmlmaxsoaparrayrank,omitempty"`
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