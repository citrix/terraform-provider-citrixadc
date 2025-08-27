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

package gslb

/**
* Configuration for GSLB site resource.
*/
type Gslbsite struct {
	/**
	* Name for the GSLB site. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the virtual server is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my gslbsite" or 'my gslbsite').
	*/
	Sitename string `json:"sitename,omitempty"`
	/**
	* Type of site to create. If the type is not specified, the appliance automatically detects and sets the type on the basis of the IP address being assigned to the site. If the specified site IP address is owned by the appliance (for example, a MIP address or SNIP address), the site is a local site. Otherwise, it is a remote site.
	*/
	Sitetype string `json:"sitetype,omitempty"`
	/**
	* IP address for the GSLB site. The GSLB site uses this IP address to communicate with other GSLB sites. For a local site, use any IP address that is owned by the appliance (for example, a SNIP or MIP address, or the IP address of the ADNS service).
	*/
	Siteipaddress string `json:"siteipaddress,omitempty"`
	/**
	* Public IP address for the local site. Required only if the appliance is deployed in a private address space and the site has a public IP address hosted on an external firewall or a NAT device.
	*/
	Publicip string `json:"publicip,omitempty"`
	/**
	* Exchange metrics with other sites. Metrics are exchanged by using Metric Exchange Protocol (MEP). The appliances in the GSLB setup exchange health information once every second.
		If you disable metrics exchange, you can use only static load balancing methods (such as round robin, static proximity, or the hash-based methods), and if you disable metrics exchange when a dynamic load balancing method (such as least connection) is in operation, the appliance falls back to round robin. Also, if you disable metrics exchange, you must use a monitor to determine the state of GSLB services. Otherwise, the service is marked as DOWN.
	*/
	Metricexchange string `json:"metricexchange,omitempty"`
	/**
	* Exchange, with other GSLB sites, network metrics such as round-trip time (RTT), learned from communications with various local DNS (LDNS) servers used by clients. RTT information is used in the dynamic RTT load balancing method, and is exchanged every 5 seconds.
	*/
	Nwmetricexchange string `json:"nwmetricexchange,omitempty"`
	/**
	* Exchange persistent session entries with other GSLB sites every five seconds.
	*/
	Sessionexchange string `json:"sessionexchange,omitempty"`
	/**
	* Specify the conditions under which the GSLB service must be monitored by a monitor, if one is bound. Available settings function as follows:
		* ALWAYS - Monitor the GSLB service at all times.
		* MEPDOWN - Monitor the GSLB service only when the exchange of metrics through the Metrics Exchange Protocol (MEP) is disabled.
		MEPDOWN_SVCDOWN - Monitor the service in either of the following situations:
		* The exchange of metrics through MEP is disabled.
		* The exchange of metrics through MEP is enabled but the status of the service, learned through metrics exchange, is DOWN.
	*/
	Triggermonitor string `json:"triggermonitor,omitempty"`
	/**
	* Parent site of the GSLB site, in a parent-child topology.
	*/
	Parentsite string `json:"parentsite,omitempty"`
	/**
	* Cluster IP address. Specify this parameter to connect to the remote cluster site for GSLB auto-sync. Note: The cluster IP address is defined when creating the cluster.
	*/
	Clip string `json:"clip,omitempty"`
	/**
	* IP address to be used to globally access the remote cluster when it is deployed behind a NAT. It can be same as the normal cluster IP address.
	*/
	Publicclip string `json:"publicclip,omitempty"`
	/**
	* The naptr replacement suffix configured here will be used to construct the naptr replacement field in NAPTR record.
	*/
	Naptrreplacementsuffix string `json:"naptrreplacementsuffix,omitempty"`
	/**
	* The list of backup gslb sites configured in preferred order. Need to be parent gsb sites.
	*/
	Backupparentlist []string `json:"backupparentlist,omitempty"`
	/**
	* Password to be used for mep communication between gslb site nodes.
	*/
	Sitepassword string `json:"sitepassword,omitempty"`
	/**
	* New name for the GSLB site.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Status string `json:"status,omitempty"`
	Persistencemepstatus string `json:"persistencemepstatus,omitempty"`
	Version string `json:"version,omitempty"`
	Curbackupparentip string `json:"curbackupparentip,omitempty"`
	Sitestate string `json:"sitestate,omitempty"`
	Oldname string `json:"oldname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
