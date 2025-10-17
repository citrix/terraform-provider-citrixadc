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
* Binding class showing the snmpuser that can be bound to snmptrap.
*/
type Snmptrapsnmpuserbinding struct {
	/**
	* Name of the SNMP user that will send the SNMPv3 traps.
	*/
	Username string `json:"username,omitempty"`
	/**
	* Security level of the SNMPv3 trap.
	*/
	Securitylevel string `json:"securitylevel,omitempty"`
	/**
	* Type of trap messages that the Citrix ADC sends to the trap listener: Generic or the enterprise-specific messages defined in the MIB file.
	*/
	Trapclass string `json:"trapclass,omitempty"`
	/**
	* IPv4 or the IPv6 address of the trap listener to which the Citrix ADC is to send SNMP trap messages.
	*/
	Trapdestination string `json:"trapdestination,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td *int `json:"td,omitempty"`
	/**
	* SNMP version, which determines the format of trap messages sent to the trap listener. 
		This setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.
	*/
	Version string `json:"version,omitempty"`


}