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

package ns

/**
* Configuration for ip6 resource.
*/
type Nsip6 struct {
	/**
	* IPv6 address to create on the Citrix ADC.
	*/
	Ipv6address string `json:"ipv6address,omitempty"`
	/**
	* Scope of the IPv6 address to be created. Cannot be changed after the IP address is created.
	*/
	Scope string `json:"scope,omitempty"`
	/**
	* Type of IP address to be created on the Citrix ADC. Cannot be changed after the IP address is created.
	*/
	Type string `json:"type,omitempty"`
	/**
	* The VLAN number.
	*/
	Vlan uint32 `json:"vlan,omitempty"`
	/**
	* Respond to Neighbor Discovery (ND) requests for this IP address.
	*/
	Nd string `json:"nd,omitempty"`
	/**
	* Respond to ICMP requests for this IP address.
	*/
	Icmp string `json:"icmp,omitempty"`
	/**
	* Enable or disable the state of all the virtual servers associated with this VIP6 address.
	*/
	Vserver string `json:"vserver,omitempty"`
	/**
	* Allow Telnet access to this IP address.
	*/
	Telnet string `json:"telnet,omitempty"`
	/**
	* Allow File Transfer Protocol (FTP) access to this IP address.
	*/
	Ftp string `json:"ftp,omitempty"`
	/**
	* Allow graphical user interface (GUI) access to this IP address.
	*/
	Gui string `json:"gui,omitempty"`
	/**
	* Allow secure Shell (SSH) access to this IP address.
	*/
	Ssh string `json:"ssh,omitempty"`
	/**
	* Allow Simple Network Management Protocol (SNMP) access to this IP address.
	*/
	Snmp string `json:"snmp,omitempty"`
	/**
	* Allow access to management applications on this IP address.
	*/
	Mgmtaccess string `json:"mgmtaccess,omitempty"`
	/**
	* Block access to nonmanagement applications on this IP address. This option is applicable forMIP6s, SNIP6s, and NSIP6s, and is disabled by default. Nonmanagement applications can run on the underlying Citrix ADC Free BSD operating system.
	*/
	Restrictaccess string `json:"restrictaccess,omitempty"`
	/**
	* Allow dynamic routing on this IP address. Specific to Subnet IPv6 (SNIP6) address.
	*/
	Dynamicrouting string `json:"dynamicrouting,omitempty"`
	/**
	* Decrement Hop Limit by 1 when ENABLED.This setting is applicable only for UDP traffic.
	*/
	Decrementhoplimit string `json:"decrementhoplimit,omitempty"`
	/**
	* Option to push the VIP6 to ZebOS routing table for Kernel route redistribution through dynamic routing protocols.
	*/
	Hostroute string `json:"hostroute,omitempty"`
	/**
	* Advertise VIPs from Shared VLAN on Default Partition
	*/
	Advertiseondefaultpartition string `json:"advertiseondefaultpartition,omitempty"`
	/**
	* Option to push the SNIP6 subnet to ZebOS routing table for Kernel route redistribution through dynamic routing protocol.
	*/
	Networkroute string `json:"networkroute,omitempty"`
	/**
	* Tag value for the network/host route associated with this IP.
	*/
	Tag uint32 `json:"tag,omitempty"`
	/**
	* IPv6 address of the gateway for the route. If Gateway is not set, VIP uses :: as the gateway.
	*/
	Ip6hostrtgw string `json:"ip6hostrtgw,omitempty"`
	/**
	* Integer value to add to or subtract from the cost of the route advertised for the VIP6 address.
	*/
	Metric int32 `json:"metric,omitempty"`
	/**
	* Advertise or do not advertise the route for the Virtual IP (VIP6) address on the basis of the state of the virtual servers associated with that VIP6.
		* NONE - Advertise the route for the VIP6 address, irrespective of the state of the virtual servers associated with the address.
		* ONE VSERVER - Advertise the route for the VIP6 address if at least one of the associated virtual servers is in UP state.
		* ALL VSERVER - Advertise the route for the VIP6 address if all of the associated virtual servers are in UP state.
		* VSVR_CNTRLD.   Advertise the route for the VIP address according to the  RHIstate (RHI STATE) parameter setting on all the associated virtual servers of the VIP address along with their states.
		When Vserver RHI Level (RHI) parameter is set to VSVR_CNTRLD, the following are different RHI behaviors for the VIP address on the basis of RHIstate (RHI STATE) settings on the virtual servers associated with the VIP address:
		* If you set RHI STATE to PASSIVE on all virtual servers, the Citrix ADC always advertises the route for the VIP address.
		* If you set RHI STATE to ACTIVE on all virtual servers, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers is in UP state.
		*If you set RHI STATE to ACTIVE on some and PASSIVE on others, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers, whose RHI STATE set to ACTIVE, is in UP state.
	*/
	Vserverrhilevel string `json:"vserverrhilevel,omitempty"`
	/**
	* Type of LSAs to be used by the IPv6 OSPF protocol, running on the Citrix ADC, for advertising the route for the VIP6 address.
	*/
	Ospf6lsatype string `json:"ospf6lsatype,omitempty"`
	/**
	* ID of the area in which the Intra-Area-Prefix LSAs are to be advertised for the VIP6 address by the IPv6 OSPF protocol running on the Citrix ADC. When ospfArea is not set, VIP6 is advertised on all areas.
	*/
	Ospfarea uint32 `json:"ospfarea,omitempty"`
	/**
	* Enable or disable the IP address.
	*/
	State string `json:"state,omitempty"`
	/**
	* Mapped IPV4 address for the IPV6 address.
	*/
	Map string `json:"map,omitempty"`
	/**
	* A positive integer that uniquely identifies a VMAC address for binding to this VIP address. This binding is used to set up Citrix ADCs in an active-active configuration using VRRP.
	*/
	Vrid6 uint32 `json:"vrid6,omitempty"`
	/**
	* ID of the cluster node for which you are adding the IP address. Must be used if you want the IP address to be active only on the specific node. Can be configured only through the cluster IP address. Cannot be changed after the IP address is created.
	*/
	Ownernode uint32 `json:"ownernode,omitempty"`
	/**
	* in cluster system, if the owner node is down, whether should it respond to icmp/arp
	*/
	Ownerdownresponse string `json:"ownerdownresponse,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td uint32 `json:"td,omitempty"`
	/**
	* NdOwner in Cluster for VIPS and Striped SNIPS
	*/
	Ndowner uint32 `json:"ndowner,omitempty"`
	/**
	* If enabled, this IP will be advertised by Citrix ADC to MPTCP enabled clients as part of ADD_ADDR option.
	*/
	Mptcpadvertise string `json:"mptcpadvertise,omitempty"`

	//------- Read only Parameter ---------;

	Iptype string `json:"iptype,omitempty"`
	Curstate string `json:"curstate,omitempty"`
	Viprtadv2bsd string `json:"viprtadv2bsd,omitempty"`
	Vipvsercount string `json:"vipvsercount,omitempty"`
	Vipvserdowncount string `json:"vipvserdowncount,omitempty"`
	Systemtype string `json:"systemtype,omitempty"`
	Operationalndowner string `json:"operationalndowner,omitempty"`

}
