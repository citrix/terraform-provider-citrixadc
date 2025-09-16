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
* Configuration for SNMP option resource.
*/
type Snmpoption struct {
	/**
	* Accept SNMP SET requests sent to the Citrix ADC, and allow SNMP managers to write values to MIB objects that are configured for write access.
	*/
	Snmpset string `json:"snmpset,omitempty"`
	/**
	* Log any SNMP trap events (for SNMP alarms in which logging is enabled) even if no trap listeners are configured. With the default setting, SNMP trap events are logged if at least one trap listener is configured on the appliance.
	*/
	Snmptraplogging string `json:"snmptraplogging,omitempty"`
	/**
	* Send partition name as a varbind in traps. By default the partition names are not sent as a varbind.
	*/
	Partitionnameintrap string `json:"partitionnameintrap,omitempty"`
	/**
	* Audit log level of SNMP trap logs. The default value is INFORMATIONAL.
	*/
	Snmptraplogginglevel string `json:"snmptraplogginglevel,omitempty"`
	/**
	* By default, the severity level info of the trap is not mentioned in the trap message. Enable this option to send severity level of trap as one of the varbind in the trap message.
	*/
	Severityinfointrap string `json:"severityinfointrap,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
