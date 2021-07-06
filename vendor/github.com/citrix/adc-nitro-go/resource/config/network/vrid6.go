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

package network

/**
* Configuration for Virtual Router ID for IPv6 resource.
*/
type Vrid6 struct {
	/**
	* Integer value that uniquely identifies a VMAC6 address.
	*/
	Id int `json:"id,omitempty"`
	/**
	* Base priority (BP), in an active-active mode configuration, which ordinarily determines the master VIP address.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* In an active-active mode configuration, make a backup VIP address the master if its priority becomes higher than that of a master VIP address bound to this VMAC address.
		If you disable pre-emption while a backup VIP address is the master, the backup VIP address remains master until the original master VIP's priority becomes higher than that of the current master.
	*/
	Preemption string `json:"preemption,omitempty"`
	/**
	* In an active-active mode configuration, enable the backup VIP address to process any traffic instead of dropping it.
	*/
	Sharing string `json:"sharing,omitempty"`
	/**
	* The effective priority (EP) value, relative to the base priority (BP) value in an active-active mode configuration. When EP is set to a value other than None, it is EP, not BP, which determines the master VIP address.
		Available settings function as follows:
		* NONE - No tracking. EP = BP
		* ALL -  If the status of all virtual servers is UP, EP = BP. Otherwise, EP = 0.
		* ONE - If the status of at least one virtual server is UP, EP = BP. Otherwise, EP = 0.
		* PROGRESSIVE - If the status of all virtual servers is UP, EP = BP. If the status of all virtual servers is DOWN, EP = 0. Otherwise EP = BP (1 - K/N), where N is the total number of virtual servers associated with the VIP address and K is the number of virtual servers for which the status is DOWN.
		Default: NONE.
	*/
	Tracking string `json:"tracking,omitempty"`
	/**
	* Preemption delay time in seconds, in an active-active configuration. If any high priority node will come in network, it will wait for these many seconds before becoming master.
	*/
	Preemptiondelaytimer int `json:"preemptiondelaytimer,omitempty"`
	/**
	* Priority by which the Effective priority will be reduced if any of the tracked interfaces goes down in an active-active configuration.
	*/
	Trackifnumpriority int `json:"trackifnumpriority,omitempty"`
	/**
	* In a cluster setup, assign a cluster node as the owner of this VMAC address for IP based VRRP configuration. If no owner is configured, ow ner node is displayed as ALL and one node is dynamically elected as the owner.
	*/
	Ownernode int `json:"ownernode,omitempty"`
	/**
	* Remove all configured VMAC6 addresses from the Citrix ADC.
	*/
	All bool `json:"all,omitempty"`

	//------- Read only Parameter ---------;

	Ifaces string `json:"ifaces,omitempty"`
	Ifnum string `json:"ifnum,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Flags string `json:"flags,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Effectivepriority string `json:"effectivepriority,omitempty"`
	Operationalownernode string `json:"operationalownernode,omitempty"`

}
