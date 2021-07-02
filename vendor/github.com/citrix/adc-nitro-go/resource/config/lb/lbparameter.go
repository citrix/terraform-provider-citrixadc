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

package lb

/**
* Configuration for LB parameter resource.
*/
type Lbparameter struct {
	/**
	* Include the HttpOnly attribute in persistence cookies. The HttpOnly attribute limits the scope of a cookie to HTTP requests and helps mitigate the risk of cross-site scripting attacks.
	*/
	Httponlycookieflag string `json:"httponlycookieflag,omitempty"`
	/**
	* Encode persistence cookie values using SHA2 hash.
	*/
	Usesecuredpersistencecookie string `json:"usesecuredpersistencecookie,omitempty"`
	/**
	* Encode persistence cookie values using SHA2 hash.
	*/
	Useencryptedpersistencecookie string `json:"useencryptedpersistencecookie,omitempty"`
	/**
	* Use this parameter to specify the passphrase used to generate secured persistence cookie value. It specifies the passphrase with a maximum of 31 characters.
	*/
	Cookiepassphrase string `json:"cookiepassphrase,omitempty"`
	/**
	* To find the service with the fewest connections, the virtual server uses the consolidated connection statistics from all the packet engines. The NO setting allows consideration of only the number of connections on the packet engine that received the new connection.
	*/
	Consolidatedlconn string `json:"consolidatedlconn,omitempty"`
	/**
	* Include the port number of the service when creating a hash for hash based load balancing methods. With the NO setting, only the IP address of the service is considered when creating a hash.
	*/
	Useportforhashlb string `json:"useportforhashlb,omitempty"`
	/**
	* Perform route lookup for traffic received by the Citrix ADC, and forward the traffic according to configured routes. Do not set this parameter if you want a wildcard virtual server to direct packets received by the appliance to an intermediary device, such as a firewall, even if their destination is directly connected to the appliance. Route lookup is performed after the packets have been processed and returned by the intermediary device.
	*/
	Preferdirectroute string `json:"preferdirectroute,omitempty"`
	/**
	* Number of requests, per service, for which to apply the round robin load balancing method before switching to the configured load balancing method, thus allowing services to ramp up gradually to full load. Until the specified number of requests is distributed, the Citrix ADC is said to be implementing the slow start mode (or startup round robin). Implemented for a virtual server when one of the following is true:
		* The virtual server is newly created.
		* One or more services are newly bound to the virtual server. 
		* One or more services bound to the virtual server are enabled.
		* The load balancing method is changed.
		This parameter applies to all the load balancing virtual servers configured on the Citrix ADC, except for those virtual servers for which the virtual server-level slow start parameters (New Service Startup Request Rate and Increment Interval) are configured. If the global slow start parameter and the slow start parameters for a given virtual server are not set, the appliance implements a default slow start for the virtual server, as follows:
		* For a newly configured virtual server, the appliance implements slow start for the first 100 requests received by the virtual server.
		* For an existing virtual server, if one or more services are newly bound or newly enabled, or if the load balancing method is changed, the appliance dynamically computes the number of requests for which to implement startup round robin. It obtains this number by multiplying the request rate by the number of bound services (it includes services that are marked as DOWN). For example, if the current request rate is 20 requests/s and ten services are bound to the virtual server, the appliance performs startup round robin for 200 requests.
		Not applicable to a virtual server for which a hash based load balancing method is configured.
	*/
	Startuprrfactor uint32 `json:"startuprrfactor,omitempty"`
	/**
	* When a monitor initiates a connection to a service, do not check to determine whether the number of connections to the service has reached the limit specified by the service's Max Clients setting. Enables monitoring to continue even if the service has reached its connection limit.
	*/
	Monitorskipmaxclient string `json:"monitorskipmaxclient,omitempty"`
	/**
	* Close monitoring connections by sending the service a connection termination message with the specified bit set.
	*/
	Monitorconnectionclose string `json:"monitorconnectionclose,omitempty"`
	/**
	* Allow a MAC-mode virtual server to accept traffic returned by an intermediary device, such as a firewall, to which the traffic was previously forwarded by another MAC-mode virtual server. The second virtual server can then distribute that traffic across the destination server farm. Also useful when load balancing Branch Repeater appliances.
		Note: The second virtual server can also send the traffic to another set of intermediary devices, such as another set of firewalls. If necessary, you can configure multiple MAC-mode virtual servers to pass traffic successively through multiple sets of intermediary devices.
	*/
	Vserverspecificmac string `json:"vserverspecificmac,omitempty"`
	/**
	* This is used, to enable/disable the option of svc/svcgroup removal, if it is bound to one or more vserver. If it is enabled, the svc/svcgroup can be removed, even if it bound to vservers. If disabled, an error will be thrown, when the user tries to remove a svc/svcgroup without unbinding from its vservers.
	*/
	Allowboundsvcremoval string `json:"allowboundsvcremoval,omitempty"`
	/**
	* This option is used to retain the original state of service or servicegroup member when an enable server command is issued.
	*/
	Retainservicestate string `json:"retainservicestate,omitempty"`
	/**
	* Specify the TTL for DNS record for domain based service. The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors
	*/
	Dbsttl uint64 `json:"dbsttl,omitempty"`
	/**
	* Maximum number of concurrent requests to allow on a single client connection, which is identified by the <clientip:port>-<vserver ip:port> tuple. This parameter is applicable to ANY service type and all UDP service types (except DNS) and only when "svrTimeout" is set to zero. A value of 0 (zero) applies no limit to the number of concurrent requests allowed on a single client connection
	*/
	Maxpipelinenat uint32 `json:"maxpipelinenat,omitempty"`
	/**
	* String configured as LiteralADCCookieAttribute will be appended as attribute for Citrix ADC cookie (for example: LB cookie persistence , GSLB site persistence, CS cookie persistence, LB group cookie persistence).
		Sample usage -
		set lb parameter -LiteralADCCookieAttribute ";SameSite=None"
	*/
	Literaladccookieattribute string `json:"literaladccookieattribute,omitempty"`
	/**
	* ComputedADCCookieAttribute accepts ns variable as input in form of string starting with $ (to understand how to configure ns variable, please check man add ns variable). policies can be configured to modify this variable for every transaction and the final value of the variable after policy evaluation will be appended as attribute to Citrix ADC cookie (for example: LB cookie persistence , GSLB sitepersistence, CS cookie persistence, LB group cookie persistence). Only one of ComputedADCCookieAttribute, LiteralADCCookieAttribute can be set.
		Sample usage -
		add ns variable lbvar -type TEXT(100) -scope Transaction
		add ns assignment lbassign -variable $lbvar -set "\\";SameSite=Strict\\""
		add rewrite policy lbpol <valid policy expression> lbassign
		bind rewrite global lbpol 100 next -type RES_OVERRIDE
		set lb param -ComputedADCCookieAttribute "$lbvar"
		For incoming client request, if above policy evaluates TRUE, then SameSite=Strict will be appended to ADC generated cookie
	*/
	Computedadccookieattribute string `json:"computedadccookieattribute,omitempty"`
	/**
	* This option allows to store the MQTT clientid and username in transactional logs
	*/
	Storemqttclientidandusername string `json:"storemqttclientidandusername,omitempty"`
	/**
	* When this option is enabled, MQTT messages of length greater than 64k will be dropped and the client/server connections will be reset.
	*/
	Dropmqttjumbomessage string `json:"dropmqttjumbomessage,omitempty"`
	/**
	* This option dictates the hashing algorithm used for hash based LB methods (URLHASH, DOMAINHASH, SOURCEIPHASH, DESTINATIONIPHASH, SRCIPDESTIPHASH, SRCIPSRCPORTHASH, TOKEN, USER_TOKEN, CALLIDHASH).
	*/
	Lbhashalgorithm string `json:"lbhashalgorithm,omitempty"`
	/**
	* This option is used to specify the number of fingers to be used in PRAC and JARH algorithms for hash based LB methods. Increasing the number of fingers might give better distribution of traffic at the expense of additional memory
	*/
	Lbhashfingers uint32 `json:"lbhashfingers,omitempty"`

	//------- Read only Parameter ---------;

	Sessionsthreshold string `json:"sessionsthreshold,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Adccookieattributewarningmsg string `json:"adccookieattributewarningmsg,omitempty"`
	Lbhashalgowinsize string `json:"lbhashalgowinsize,omitempty"`

}
