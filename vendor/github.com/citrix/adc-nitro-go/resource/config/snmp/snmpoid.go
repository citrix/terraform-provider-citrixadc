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
* Configuration for SNMP Object Identifier resource.
*/
type Snmpoid struct {
	/**
	* The type of entity whose SNMP OIDs you want to displayType of entity whose SNMP OIDs you want the Citrix ADC to display.
	*/
	Entitytype string `json:"entitytype,omitempty"`
	/**
	* Name of the entity whose SNMP OID you want the Citrix ADC to display.
	*/
	Name string `json:"name,omitempty"`

	//------- Read only Parameter ---------;

	Snmpoid string `json:"Snmpoid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
