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
* Configuration for ip resource.
*/
type Nsip struct {
	/**
	* IPv4 address to create on the Citrix ADC. Cannot be changed after the IP address is created.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Subnet mask associated with the IP address.
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* Type of the IP address to create on the Citrix ADC. Cannot be changed after the IP address is created. The following are the different types of Citrix ADC owned IP addresses:
		* A Subnet IP (SNIP) address is used by the Citrix ADC to communicate with the servers. The Citrix ADC also uses the subnet IP address when generating its own packets, such as packets related to dynamic routing protocols, or to send monitor probes to check the health of the servers.
		* A Virtual IP (VIP) address is the IP address associated with a virtual server. It is the IP address to which clients connect. An appliance managing a wide range of traffic may have many VIPs configured. Some of the attributes of the VIP address are customized to meet the requirements of the virtual server.
		* A GSLB site IP (GSLBIP) address is associated with a GSLB site. It is not mandatory to specify a GSLBIP address when you initially configure the Citrix ADC. A GSLBIP address is used only when you create a GSLB site.
		* A Cluster IP (CLIP) address is the management address of the cluster. All cluster configurations must be performed by accessing the cluster through this IP address.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Respond to ARP requests for this IP address.
	*/
	Arp string `json:"arp,omitempty"`
	/**
	* Respond to ICMP requests for this IP address.
	*/
	Icmp string `json:"icmp,omitempty"`
	/**
	* Use this option to set (enable or disable) the virtual server attribute for this IP address.
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
	* Allow secure shell (SSH) access to this IP address.
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
	* Block access to nonmanagement applications on this IP. This option is applicable for MIPs, SNIPs, and NSIP, and is disabled by default. Nonmanagement applications can run on the underlying Citrix ADC Free BSD operating system.
	*/
	Restrictaccess string `json:"restrictaccess,omitempty"`
	/**
	* Allow dynamic routing on this IP address. Specific to Subnet IP (SNIP) address.
	*/
	Dynamicrouting string `json:"dynamicrouting,omitempty"`
	/**
	* Decrement TTL by 1 when ENABLED.This setting is applicable only for UDP traffic.
	*/
	Decrementttl string `json:"decrementttl,omitempty"`
	/**
	* Use this option to enable or disable OSPF on this IP address for the entity.
	*/
	Ospf string `json:"ospf,omitempty"`
	/**
	* Use this option to enable or disable BGP on this IP address for the entity.
	*/
	Bgp string `json:"bgp,omitempty"`
	/**
	* Use this option to enable or disable RIP on this IP address for the entity.
	*/
	Rip string `json:"rip,omitempty"`
	/**
	* Option to push the VIP to ZebOS routing table for Kernel route redistribution through dynamic routing protocols
	*/
	Hostroute string `json:"hostroute,omitempty"`
	/**
	* Advertise VIPs from Shared VLAN on Default Partition.
	*/
	Advertiseondefaultpartition string `json:"advertiseondefaultpartition,omitempty"`
	/**
	* Option to push the SNIP subnet to ZebOS routing table for Kernel route redistribution through dynamic routing protocol.
	*/
	Networkroute string `json:"networkroute,omitempty"`
	/**
	* Tag value for the network/host route associated with this IP.
	*/
	Tag int `json:"tag,omitempty"`
	/**
	* IP address of the gateway of the route for this VIP address.
	*/
	Hostrtgw string `json:"hostrtgw,omitempty"`
	/**
	* Integer value to add to or subtract from the cost of the route advertised for the VIP address.
	*/
	Metric int `json:"metric,omitempty"`
	/**
	* Advertise the route for the Virtual IP (VIP) address on the basis of the state of the virtual servers associated with that VIP.
		* NONE - Advertise the route for the VIP address, regardless of the state of the virtual servers associated with the address.
		* ONE VSERVER - Advertise the route for the VIP address if at least one of the associated virtual servers is in UP state.
		* ALL VSERVER - Advertise the route for the VIP address if all of the associated virtual servers are in UP state.
		* VSVR_CNTRLD - Advertise the route for the VIP address according to the  RHIstate (RHI STATE) parameter setting on all the associated virtual servers of the VIP address along with their states.
		When Vserver RHI Level (RHI) parameter is set to VSVR_CNTRLD, the following are different RHI behaviors for the VIP address on the basis of RHIstate (RHI STATE) settings on the virtual servers associated with the VIP address:
		* If you set RHI STATE to PASSIVE on all virtual servers, the Citrix ADC always advertises the route for the VIP address.
		* If you set RHI STATE to ACTIVE on all virtual servers, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers is in UP state.
		*If you set RHI STATE to ACTIVE on some and PASSIVE on others, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers, whose RHI STATE set to ACTIVE, is in UP state.
	*/
	Vserverrhilevel string `json:"vserverrhilevel,omitempty"`
	/**
	* Type of LSAs to be used by the OSPF protocol, running on the Citrix ADC, for advertising the route for this VIP address.
	*/
	Ospflsatype string `json:"ospflsatype,omitempty"`
	/**
	* ID of the area in which the type1 link-state advertisements (LSAs) are to be advertised for this virtual IP (VIP)  address by the OSPF protocol running on the Citrix ADC.  When this parameter is not set, the VIP is advertised on all areas.
	*/
	Ospfarea int `json:"ospfarea,omitempty"`
	/**
	* Enable or disable the IP address.
	*/
	State string `json:"state,omitempty"`
	/**
	* A positive integer that uniquely identifies a VMAC address for binding to this VIP address. This binding is used to set up Citrix ADCs in an active-active configuration using VRRP.
	*/
	Vrid int `json:"vrid,omitempty"`
	/**
	* Respond to ICMP requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP. Available settings function as follows:
		* NONE - The Citrix ADC responds to any ICMP request for the VIP address, irrespective of the states of the virtual servers associated with the address.
		* ONE VSERVER - The Citrix ADC responds to any ICMP request for the VIP address if at least one of the associated virtual servers is in UP state.
		* ALL VSERVER - The Citrix ADC responds to any ICMP request for the VIP address if all of the associated virtual servers are in UP state.
		* VSVR_CNTRLD - The behavior depends on the ICMP VSERVER RESPONSE setting on all the associated virtual servers.
		The following settings can be made for the ICMP VSERVER RESPONSE parameter on a virtual server:
		* If you set ICMP VSERVER RESPONSE to PASSIVE on all virtual servers, Citrix ADC always responds.
		* If you set ICMP VSERVER RESPONSE to ACTIVE on all virtual servers, Citrix ADC responds if even one virtual server is UP.
		* When you set ICMP VSERVER RESPONSE to ACTIVE on some and PASSIVE on others, Citrix ADC responds if even one virtual server set to ACTIVE is UP.
	*/
	Icmpresponse string `json:"icmpresponse,omitempty"`
	/**
	* The owner node in a Cluster for this IP address. Owner node can vary from 0 to 31. If ownernode is not specified then the IP is treated as Striped IP.
	*/
	Ownernode int `json:"ownernode,omitempty"`
	/**
	* Respond to ARP requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP. Available settings function as follows:
		* NONE - The Citrix ADC responds to any ARP request for the VIP address, irrespective of the states of the virtual servers associated with the address.
		* ONE VSERVER - The Citrix ADC responds to any ARP request for the VIP address if at least one of the associated virtual servers is in UP state.
		* ALL VSERVER - The Citrix ADC responds to any ARP request for the VIP address if all of the associated virtual servers are in UP state.
	*/
	Arpresponse string `json:"arpresponse,omitempty"`
	/**
	* in cluster system, if the owner node is down, whether should it respond to icmp/arp
	*/
	Ownerdownresponse string `json:"ownerdownresponse,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. TD id 4095 is used reserved for  LSN use  
	*/
	Td int `json:"td,omitempty"`
	/**
	* The arp owner in a Cluster for this IP address. It can vary from 0 to 31.
	*/
	Arpowner int `json:"arpowner,omitempty"`
	/**
	* If enabled, this IP will be advertised by Citrix ADC to MPTCP enabled clients as part of ADD_ADDR option.
	*/
	Mptcpadvertise string `json:"mptcpadvertise,omitempty"`

	//------- Read only Parameter ---------;

	Flags string `json:"flags,omitempty"`
	Hostrtgwact string `json:"hostrtgwact,omitempty"`
	Ospfareaval string `json:"ospfareaval,omitempty"`
	Viprtadv2bsd string `json:"viprtadv2bsd,omitempty"`
	Vipvsercount string `json:"vipvsercount,omitempty"`
	Vipvserdowncount string `json:"vipvserdowncount,omitempty"`
	Vipvsrvrrhiactivecount string `json:"vipvsrvrrhiactivecount,omitempty"`
	Vipvsrvrrhiactiveupcount string `json:"vipvsrvrrhiactiveupcount,omitempty"`
	Freeports string `json:"freeports,omitempty"`
	Iptype string `json:"iptype,omitempty"`
	Operationalarpowner string `json:"operationalarpowner,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
