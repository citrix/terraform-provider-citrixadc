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
* Configuration for snmp trap resource.
*/
type Snmptrap struct {
	/**
	* Type of trap messages that the Citrix ADC sends to the trap listener: Generic or the enterprise-specific messages defined in the MIB file.
	*/
	Trapclass string `json:"trapclass,omitempty"`
	/**
	* IPv4 or the IPv6 address of the trap listener to which the Citrix ADC is to send SNMP trap messages.
	*/
	Trapdestination string `json:"trapdestination,omitempty"`
	/**
	* SNMP version, which determines the format of trap messages sent to the trap listener. 
		This setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.
	*/
	Version string `json:"version,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td *int `json:"td,omitempty"`
	/**
	* UDP port at which the trap listener listens for trap messages. This setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.
	*/
	Destport *int `json:"destport,omitempty"`
	/**
	* Password (string) sent with the trap messages, so that the trap listener can authenticate them. Can include 1 to 31 uppercase or lowercase letters, numbers, and hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters.  
		You must specify the same community string on the trap listener device. Otherwise, the trap listener drops the trap messages.
		The following requirement applies only to the Citrix ADC CLI:
		If the string includes one or more spaces, enclose the name in double or single quotation marks (for example, "my string" or 'my string').
	*/
	Communityname string `json:"communityname,omitempty"`
	/**
	* IPv4 or IPv6 address that the Citrix ADC inserts as the source IP address in all SNMP trap messages that it sends to this trap listener. By default this is the appliance's NSIP or NSIP6 address, but you can specify an IPv4 MIP or SNIP/SNIP6 address. In cluster setup, the default value is the individual node's NSIP, but it can be set to CLIP or Striped SNIP address. In non default partition, this parameter must be set to the SNIP/SNIP6 address.
	*/
	Srcip string `json:"srcip,omitempty"`
	/**
	* Severity level at or above which the Citrix ADC sends trap messages to this trap listener. The severity levels, in increasing order of severity, are Informational, Warning, Minor, Major, Critical. This parameter can be set for trap listeners of type SPECIFIC only. The default is to send all levels of trap messages. 
		Important: Trap messages are not assigned severity levels unless you specify severity levels when configuring SNMP alarms.
	*/
	Severity string `json:"severity,omitempty"`
	/**
	* Send traps of all partitions to this destination.
	*/
	Allpartitions string `json:"allpartitions,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
