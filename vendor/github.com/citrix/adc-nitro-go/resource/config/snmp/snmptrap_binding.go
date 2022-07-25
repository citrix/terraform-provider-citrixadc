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
* Binding object which returns the resources bound to snmptrap_binding. 
*/
type Snmptrapbinding struct {
	/**
	* Trap type specified in the trap listener entry.<br/>Possible values = generic, specific
	*/
	Trapclass string `json:"trapclass,omitempty"`
	/**
	* IP address specified in the trap listener entry.<br/>Minimum value =  
	*/
	Trapdestination string `json:"trapdestination,omitempty"`
	/**
	* The SNMP version of the trap specified in the trap listener entry.<br/>Possible values = V1, V2, V3
	*/
	Version string `json:"version,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.<br/>Minimum value =  0<br/>Maximum value =  4094
	*/
	Td int `json:"td,omitempty"`


}