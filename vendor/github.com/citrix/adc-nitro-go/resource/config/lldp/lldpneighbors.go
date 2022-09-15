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

package lldp

/**
* Configuration for lldp neighbors resource.
*/
type Lldpneighbors struct {
	/**
	* Interface Name
	*/
	Ifnum string `json:"ifnum,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Chassisidsubtype string `json:"chassisidsubtype,omitempty"`
	Chassisid string `json:"chassisid,omitempty"`
	Portidsubtype string `json:"portidsubtype,omitempty"`
	Portid string `json:"portid,omitempty"`
	Ttl string `json:"ttl,omitempty"`
	Portdescription string `json:"portdescription,omitempty"`
	Sys string `json:"sys,omitempty"`
	Sysdesc string `json:"sysdesc,omitempty"`
	Mgmtaddresssubtype string `json:"mgmtaddresssubtype,omitempty"`
	Mgmtaddress string `json:"mgmtaddress,omitempty"`
	Iftype string `json:"iftype,omitempty"`
	Ifnumber string `json:"ifnumber,omitempty"`
	Vlan string `json:"vlan,omitempty"`
	Vlanid string `json:"vlanid,omitempty"`
	Portprotosupported string `json:"portprotosupported,omitempty"`
	Portprotoenabled string `json:"portprotoenabled,omitempty"`
	Portprotoid string `json:"portprotoid,omitempty"`
	Portvlanid string `json:"portvlanid,omitempty"`
	Protocolid string `json:"protocolid,omitempty"`
	Linkaggrcapable string `json:"linkaggrcapable,omitempty"`
	Linkaggrenabled string `json:"linkaggrenabled,omitempty"`
	Linkaggrid string `json:"linkaggrid,omitempty"`
	Flag string `json:"flag,omitempty"`
	Syscapabilities string `json:"syscapabilities,omitempty"`
	Syscapenabled string `json:"syscapenabled,omitempty"`
	Autonegsupport string `json:"autonegsupport,omitempty"`
	Autonegenabled string `json:"autonegenabled,omitempty"`
	Autonegadvertised string `json:"autonegadvertised,omitempty"`
	Autonegmautype string `json:"autonegmautype,omitempty"`
	Mtu string `json:"mtu,omitempty"`

}
