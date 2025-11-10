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
* Configuration for Global Server Load Balancing Virtual Server resource.
*/
type Gslbvserver struct {
	/**
	* Name for the GSLB virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the virtual server is created.
		CLI Users:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Protocol used by services bound to the virtual server.
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* The IP type for this GSLB vserver.
	*/
	Iptype string `json:"iptype,omitempty"`
	/**
	* DNS record type to associate with the GSLB virtual server's domain name.
	*/
	Dnsrecordtype string `json:"dnsrecordtype,omitempty"`
	/**
	* Load balancing method for the GSLB virtual server.
	*/
	Lbmethod string `json:"lbmethod,omitempty"`
	/**
	* A non zero value enables the feature whose minimum value is 2 minutes. The feature can be disabled by setting the value to zero. The created session is in effect for a specific client per domain.
	*/
	Backupsessiontimeout *int `json:"backupsessiontimeout,omitempty"`
	/**
	* Backup load balancing method. Becomes operational if the primary load balancing method fails or cannot be used. Valid only if the primary method is based on either round-trip time (RTT) or static proximity.
	*/
	Backuplbmethod string `json:"backuplbmethod,omitempty"`
	/**
	* IPv4 network mask for use in the SOURCEIPHASH load balancing method.
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* Number of bits to consider, in an IPv6 source IP address, for creating the hash that is required by the SOURCEIPHASH load balancing method.
	*/
	V6netmasklen *int `json:"v6netmasklen,omitempty"`
	/**
	* Expression, or name of a named expression, against which traffic is evaluated.
		This field is applicable only if gslb method or gslb backup method are set to API.
		The following requirements apply only to the Citrix ADC CLI:
		* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.
		* If the expression itself includes double quotation marks, escape the quotations by using the \ character.
		* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Tolerance in milliseconds. Tolerance value is used in deciding which sites in a GSLB configuration must be considered for implementing the RTT load balancing method. The sites having the RTT value less than or equal to the sum of the lowest RTT and tolerance value are considered. NetScaler implements the round robin method of global server load balancing among these considered sites. The sites that have RTT value greater than this value are not considered. The logic is applied for each LDNS and based on the LDNS, the sites that are considered might change. For example, a site that is considered for requests coming from LDNS1 might not be considered for requests coming from LDNS2.
	*/
	Tolerance *int `json:"tolerance,omitempty"`
	/**
	* Use source IP address based persistence for the virtual server.
		After the load balancing method selects a service for the first packet, the IP address received in response to the DNS query is used for subsequent requests from the same client.
	*/
	Persistencetype string `json:"persistencetype,omitempty"`
	/**
	* The persistence ID for the GSLB virtual server. The ID is a positive integer that enables GSLB sites to identify the GSLB virtual server, and is required if source IP address based or spill over based persistence is enabled on the virtual server.
	*/
	Persistenceid *int `json:"persistenceid,omitempty"`
	/**
	* The optional IPv4 network mask applied to IPv4 addresses to establish source IP address based persistence.
	*/
	Persistmask string `json:"persistmask,omitempty"`
	/**
	* Number of bits to consider in an IPv6 source IP address when creating source IP address based persistence sessions.
	*/
	V6persistmasklen *int `json:"v6persistmasklen,omitempty"`
	/**
	* Idle time, in minutes, after which a persistence entry is cleared.
	*/
	Timeout *int `json:"timeout,omitempty"`
	/**
	* Send clients an empty DNS response when the GSLB virtual server is DOWN.
	*/
	Edr string `json:"edr,omitempty"`
	/**
	* If enabled, respond with EDNS Client Subnet (ECS) option in the response for a DNS query with ECS. The ECS address will be used for persistence and spillover persistence (if enabled) instead of the LDNS address. Persistence mask is ignored if ECS is enabled.
	*/
	Ecs string `json:"ecs,omitempty"`
	/**
	* Validate if ECS address is a private or unroutable address and in such cases, use the LDNS IP.
	*/
	Ecsaddrvalidation string `json:"ecsaddrvalidation,omitempty"`
	/**
	* Include multiple IP addresses in the DNS responses sent to clients.
	*/
	Mir string `json:"mir,omitempty"`
	/**
	* Continue to direct traffic to the backup chain even after the primary GSLB virtual server returns to the UP state. Used when spillover is configured for the virtual server.
	*/
	Disableprimaryondown string `json:"disableprimaryondown,omitempty"`
	/**
	* Specify if the appliance should consider the service count, service weights, or ignore both when using weight-based load balancing methods. The state of the number of services bound to the virtual server help the appliance to select the service.
	*/
	Dynamicweight string `json:"dynamicweight,omitempty"`
	/**
	* State of the GSLB virtual server.
	*/
	State string `json:"state,omitempty"`
	/**
	* If the primary state of all bound GSLB services is DOWN, consider the effective states of all the GSLB services, obtained through the Metrics Exchange Protocol (MEP), when determining the state of the GSLB virtual server. To consider the effective state, set the parameter to STATE_ONLY. To disregard the effective state, set the parameter to NONE.
		The effective state of a GSLB service is the ability of the corresponding virtual server to serve traffic. The effective state of the load balancing virtual server, which is transferred to the GSLB service, is UP even if only one virtual server in the backup chain of virtual servers is in the UP state.
	*/
	Considereffectivestate string `json:"considereffectivestate,omitempty"`
	/**
	* Any comments that you might want to associate with the GSLB virtual server.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Type of threshold that, when exceeded, triggers spillover. Available settings function as follows:
		* CONNECTION - Spillover occurs when the number of client connections exceeds the threshold.
		* DYNAMICCONNECTION - Spillover occurs when the number of client connections at the GSLB virtual server exceeds the sum of the maximum client (Max Clients) settings for bound GSLB services. Do not specify a spillover threshold for this setting, because the threshold is implied by the Max Clients settings of the bound GSLB services.
		* BANDWIDTH - Spillover occurs when the bandwidth consumed by the GSLB virtual server's incoming and outgoing traffic exceeds the threshold.
		* HEALTH - Spillover occurs when the percentage of weights of the GSLB services that are UP drops below the threshold. For example, if services gslbSvc1, gslbSvc2, and gslbSvc3 are bound to a virtual server, with weights 1, 2, and 3, and the spillover threshold is 50%, spillover occurs if gslbSvc1 and gslbSvc3 or gslbSvc2 and gslbSvc3 transition to DOWN.
		* NONE - Spillover does not occur.
	*/
	Somethod string `json:"somethod,omitempty"`
	/**
	* If spillover occurs, maintain source IP address based persistence for both primary and backup GSLB virtual servers.
	*/
	Sopersistence string `json:"sopersistence,omitempty"`
	/**
	* Timeout for spillover persistence, in minutes.
	*/
	Sopersistencetimeout *int `json:"sopersistencetimeout,omitempty"`
	/**
	* Threshold at which spillover occurs. Specify an integer for the CONNECTION spillover method, a bandwidth value in kilobits per second for the BANDWIDTH method (do not enter the units), or a percentage for the HEALTH method (do not enter the percentage symbol).
	*/
	Sothreshold *int `json:"sothreshold,omitempty"`
	/**
	* Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists
	*/
	Sobackupaction string `json:"sobackupaction,omitempty"`
	/**
	* Enable logging appflow flow information
	*/
	Appflowlog string `json:"appflowlog,omitempty"`
	/**
	* Configure this option to toggle order preference
	*/
	Toggleorder string `json:"toggleorder,omitempty"`
	/**
	* This option is used to to specify the threshold of minimum number of services to be UP in an order, for it to be considered in Lb decision.
	*/
	Orderthreshold *int `json:"orderthreshold,omitempty"`
	/**
	* Name of the backup GSLB virtual server to which the appliance should to forward requests if the status of the primary GSLB virtual server is down or exceeds its spillover threshold.
	*/
	Backupvserver string `json:"backupvserver,omitempty"`
	/**
	* Name of the GSLB service for which to change the weight.
	*/
	Servicename string `json:"servicename,omitempty"`
	/**
	* Weight for the service.
	*/
	Weight *int `json:"weight,omitempty"`
	/**
	* The GSLB service group name bound to the selected GSLB virtual server.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* Domain name for which to change the time to live (TTL) and/or backup service IP address.
	*/
	Domainname string `json:"domainname,omitempty"`
	/**
	* Time to live (TTL) for the domain.
	*/
	Ttl *int `json:"ttl,omitempty"`
	/**
	* The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.
	*/
	Backupip string `json:"backupip,omitempty"`
	/**
	* The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.
	*/
	Cookiedomain string `json:"cookie_domain,omitempty"`
	/**
	* Timeout, in minutes, for the GSLB site cookie.
	*/
	Cookietimeout *int `json:"cookietimeout,omitempty"`
	/**
	* TTL, in seconds, for all internally created site domains (created when a site prefix is configured on a GSLB service) that are associated with this virtual server.
	*/
	Sitedomainttl *int `json:"sitedomainttl,omitempty"`
	/**
	* Order number to be assigned to the service when it is bound to the lb vserver.
	*/
	Order *int `json:"order,omitempty"`
	/**
	* New name for the GSLB virtual server.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Curstate string `json:"curstate,omitempty"`
	Status string `json:"status,omitempty"`
	Lbrrreason string `json:"lbrrreason,omitempty"`
	Iscname string `json:"iscname,omitempty"`
	Sitepersistence string `json:"sitepersistence,omitempty"`
	Totalservices string `json:"totalservices,omitempty"`
	Activeservices string `json:"activeservices,omitempty"`
	Statechangetimesec string `json:"statechangetimesec,omitempty"`
	Statechangetimemsec string `json:"statechangetimemsec,omitempty"`
	Tickssincelaststatechange string `json:"tickssincelaststatechange,omitempty"`
	Health string `json:"health,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Priority string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Type string `json:"type,omitempty"`
	Vsvrbindsvcip string `json:"vsvrbindsvcip,omitempty"`
	Vsvrbindsvcport string `json:"vsvrbindsvcport,omitempty"`
	Servername string `json:"servername,omitempty"`
	Nodefaultbindings string `json:"nodefaultbindings,omitempty"`
	Currentactiveorder string `json:"currentactiveorder,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
