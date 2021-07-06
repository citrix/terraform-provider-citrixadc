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
* Configuration for GSLB service resource.
*/
type Gslbservice struct {
	/**
	* Name for the GSLB service. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the GSLB service is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my gslbsvc" or 'my gslbsvc').
	*/
	Servicename string `json:"servicename,omitempty"`
	/**
	* Canonical name of the GSLB service. Used in CNAME-based GSLB.
	*/
	Cnameentry string `json:"cnameentry,omitempty"`
	/**
	* IP address for the GSLB service. Should represent a load balancing, content switching, or VPN virtual server on the Citrix ADC, or the IP address of another load balancing device.
	*/
	Ip string `json:"ip,omitempty"`
	/**
	* Name of the server hosting the GSLB service.
	*/
	Servername string `json:"servername,omitempty"`
	/**
	* Type of service to create.
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* Port on which the load balancing entity represented by this GSLB service listens.
	*/
	Port int `json:"port,omitempty"`
	/**
	* The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.
	*/
	Publicip string `json:"publicip,omitempty"`
	/**
	* The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.
	*/
	Publicport int `json:"publicport,omitempty"`
	/**
	* The maximum number of open connections that the service can support at any given time. A GSLB service whose connection count reaches the maximum is not considered when a GSLB decision is made, until the connection count drops below the maximum.
	*/
	Maxclient int `json:"maxclient,omitempty"`
	/**
	* Monitor the health of the GSLB service.
	*/
	Healthmonitor string `json:"healthmonitor,omitempty"`
	/**
	* Name of the GSLB site to which the service belongs.
	*/
	Sitename string `json:"sitename,omitempty"`
	/**
	* Enable or disable the service.
	*/
	State string `json:"state,omitempty"`
	/**
	* In the request that is forwarded to the GSLB service, insert a header that stores the client's IP address. Client IP header insertion is used in connection-proxy based site persistence.
	*/
	Cip string `json:"cip,omitempty"`
	/**
	* Name for the HTTP header that stores the client's IP address. Used with the Client IP option. If client IP header insertion is enabled on the service and a name is not specified for the header, the Citrix ADC uses the name specified by the cipHeader parameter in the set ns param command or, in the GUI, the Client IP Header parameter in the Configure HTTP Parameters dialog box.
	*/
	Cipheader string `json:"cipheader,omitempty"`
	/**
	* Use cookie-based site persistence. Applicable only to HTTP and SSL GSLB services.
	*/
	Sitepersistence string `json:"sitepersistence,omitempty"`
	/**
	* Timeout value, in minutes, for the cookie, when cookie based site persistence is enabled.
	*/
	Cookietimeout int `json:"cookietimeout,omitempty"`
	/**
	* The site's prefix string. When the service is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound service-domain pair by concatenating the site prefix of the service and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.
	*/
	Siteprefix string `json:"siteprefix,omitempty"`
	/**
	* Idle time, in seconds, after which a client connection is terminated. Applicable if connection proxy based site persistence is used.
	*/
	Clttimeout int `json:"clttimeout,omitempty"`
	/**
	* Idle time, in seconds, after which a server connection is terminated. Applicable if connection proxy based site persistence is used.
	*/
	Svrtimeout int `json:"svrtimeout,omitempty"`
	/**
	* Integer specifying the maximum bandwidth allowed for the service. A GSLB service whose bandwidth reaches the maximum is not considered when a GSLB decision is made, until its bandwidth consumption drops below the maximum.
	*/
	Maxbandwidth int `json:"maxbandwidth,omitempty"`
	/**
	* Flush all active transactions associated with the GSLB service when its state transitions from UP to DOWN. Do not enable this option for services that must complete their transactions. Applicable if connection proxy based site persistence is used.
	*/
	Downstateflush string `json:"downstateflush,omitempty"`
	/**
	* Maximum number of SSL VPN users that can be logged on concurrently to the VPN virtual server that is represented by this GSLB service. A GSLB service whose user count reaches the maximum is not considered when a GSLB decision is made, until the count drops below the maximum.
	*/
	Maxaaausers int `json:"maxaaausers,omitempty"`
	/**
	* Monitoring threshold value for the GSLB service. If the sum of the weights of the monitors that are bound to this GSLB service and are in the UP state is not equal to or greater than this threshold value, the service is marked as DOWN.
	*/
	Monthreshold int `json:"monthreshold,omitempty"`
	/**
	* Unique hash identifier for the GSLB service, used by hash based load balancing methods.
	*/
	Hashid int `json:"hashid,omitempty"`
	/**
	* Any comments that you might want to associate with the GSLB service.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Enable logging appflow flow information
	*/
	Appflowlog string `json:"appflowlog,omitempty"`
	/**
	* The replacement domain name for this NAPTR.
	*/
	Naptrreplacement string `json:"naptrreplacement,omitempty"`
	/**
	* An integer specifying the order in which the NAPTR records MUST be processed in order to accurately represent the ordered list of Rules. The ordering is from lowest to highest
	*/
	Naptrorder int `json:"naptrorder,omitempty"`
	/**
	* Service Parameters applicable to this delegation path.
	*/
	Naptrservices string `json:"naptrservices,omitempty"`
	/**
	* Modify the TTL of the internally created naptr domain
	*/
	Naptrdomainttl int `json:"naptrdomainttl,omitempty"`
	/**
	* An integer specifying the preference of this NAPTR among NAPTR records having same order. lower the number, higher the preference.
	*/
	Naptrpreference int `json:"naptrpreference,omitempty"`
	/**
	* The new IP address of the service.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Name of the DNS view of the service. A DNS view is used in global server load balancing (GSLB) to return a predetermined IP address to a specific group of clients, which are identified by using a DNS policy.
	*/
	Viewname string `json:"viewname,omitempty"`
	/**
	* IP address to be used for the given view
	*/
	Viewip string `json:"viewip,omitempty"`
	/**
	* Weight to assign to the monitor-service binding. A larger number specifies a greater weight. Contributes to the monitoring threshold, which determines the state of the service.
	*/
	Weight int `json:"weight,omitempty"`
	/**
	* Name of the monitor to bind to the service.
	*/
	Monitornamesvc string `json:"monitor_name_svc,omitempty"`
	/**
	* New name for the GSLB service.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Gslb string `json:"gslb,omitempty"`
	Svrstate string `json:"svrstate,omitempty"`
	Svreffgslbstate string `json:"svreffgslbstate,omitempty"`
	Gslbthreshold string `json:"gslbthreshold,omitempty"`
	Gslbsvcstats string `json:"gslbsvcstats,omitempty"`
	Monstate string `json:"monstate,omitempty"`
	Preferredlocation string `json:"preferredlocation,omitempty"`
	Monitorstate string `json:"monitor_state,omitempty"`
	Statechangetimesec string `json:"statechangetimesec,omitempty"`
	Tickssincelaststatechange string `json:"tickssincelaststatechange,omitempty"`
	Threshold string `json:"threshold,omitempty"`
	Clmonowner string `json:"clmonowner,omitempty"`
	Clmonview string `json:"clmonview,omitempty"`
	Gslbsvchealth string `json:"gslbsvchealth,omitempty"`
	Glsbsvchealthdescr string `json:"glsbsvchealthdescr,omitempty"`
	Nodefaultbindings string `json:"nodefaultbindings,omitempty"`

}
