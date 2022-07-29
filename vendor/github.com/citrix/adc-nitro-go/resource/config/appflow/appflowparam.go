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

package appflow

/**
* Configuration for AppFlow parameter resource.
*/
type Appflowparam struct {
	/**
	* Refresh interval, in seconds, at which to export the template data. Because data transmission is in UDP, the templates must be resent at regular intervals.
	*/
	Templaterefresh int `json:"templaterefresh,omitempty"`
	/**
	* Interval, in seconds, at which to send Appnames to the configured collectors. Appname refers to the name of an entity (virtual server, service, or service group) in the Citrix ADC.
	*/
	Appnamerefresh int `json:"appnamerefresh,omitempty"`
	/**
	* Interval, in seconds, at which to send flow records to the configured collectors.
	*/
	Flowrecordinterval int `json:"flowrecordinterval,omitempty"`
	/**
	* Interval, in seconds, at which to send security insight flow records to the configured collectors.
	*/
	Securityinsightrecordinterval int `json:"securityinsightrecordinterval,omitempty"`
	/**
	* MTU, in bytes, for IPFIX UDP packets.
	*/
	Udppmtu int `json:"udppmtu,omitempty"`
	/**
	* Include the http URL that the Citrix ADC received from the client.
	*/
	Httpurl string `json:"httpurl,omitempty"`
	/**
	* Enable AppFlow AAA Username logging.
	*/
	Aaausername string `json:"aaausername,omitempty"`
	/**
	* Include the cookie that was in the HTTP request the appliance received from the client.
	*/
	Httpcookie string `json:"httpcookie,omitempty"`
	/**
	* Include the web page that was last visited by the client.
	*/
	Httpreferer string `json:"httpreferer,omitempty"`
	/**
	* Include the method that was specified in the HTTP request that the appliance received from the client.
	*/
	Httpmethod string `json:"httpmethod,omitempty"`
	/**
	* Include the host identified in the HTTP request that the appliance received from the client.
	*/
	Httphost string `json:"httphost,omitempty"`
	/**
	* Include the client application through which the HTTP request was received by the Citrix ADC.
	*/
	Httpuseragent string `json:"httpuseragent,omitempty"`
	/**
	* Generate AppFlow records for only the traffic from the client.
	*/
	Clienttrafficonly string `json:"clienttrafficonly,omitempty"`
	/**
	* Include the HTTP Content-Type header sent from the server to the client to determine the type of the content sent.
	*/
	Httpcontenttype string `json:"httpcontenttype,omitempty"`
	/**
	* Include the HTTP Authorization header information.
	*/
	Httpauthorization string `json:"httpauthorization,omitempty"`
	/**
	* Include the httpVia header which contains the IP address of proxy server through which the client accessed the server.
	*/
	Httpvia string `json:"httpvia,omitempty"`
	/**
	* Include the httpXForwardedFor header, which contains the original IP Address of the client using a proxy server to access the server.
	*/
	Httpxforwardedfor string `json:"httpxforwardedfor,omitempty"`
	/**
	* Include the HTTP location headers returned from the HTTP responses.
	*/
	Httplocation string `json:"httplocation,omitempty"`
	/**
	* Include the Set-cookie header sent from the server to the client in response to a HTTP request.
	*/
	Httpsetcookie string `json:"httpsetcookie,omitempty"`
	/**
	* Include the Set-cookie header sent from the server to the client in response to a HTTP request.
	*/
	Httpsetcookie2 string `json:"httpsetcookie2,omitempty"`
	/**
	* Enable connection chaining so that the client server flows of a connection are linked. Also the connection chain ID is propagated across Citrix ADCs, so that in a multi-hop environment the flows belonging to the same logical connection are linked. This id is also logged as part of appflow record
	*/
	Connectionchaining string `json:"connectionchaining,omitempty"`
	/**
	* Include the http domain request to be exported.
	*/
	Httpdomain string `json:"httpdomain,omitempty"`
	/**
	* Skip Cache http transaction. This HTTP transaction is specific to Cache Redirection module. In Case of Cache Miss there will be another HTTP transaction initiated by the cache server.
	*/
	Skipcacheredirectionhttptransaction string `json:"skipcacheredirectionhttptransaction,omitempty"`
	/**
	* Include the stream identifier name to be exported.
	*/
	Identifiername string `json:"identifiername,omitempty"`
	/**
	* Include the stream identifier session name to be exported.
	*/
	Identifiersessionname string `json:"identifiersessionname,omitempty"`
	/**
	* An observation domain groups a set of Citrix ADCs based on deployment: cluster, HA etc. A unique Observation Domain ID is required to be assigned to each such group.
	*/
	Observationdomainid int `json:"observationdomainid,omitempty"`
	/**
	* Name of the Observation Domain defined by the observation domain ID.
	*/
	Observationdomainname string `json:"observationdomainname,omitempty"`
	/**
	* Enable this option for logging end user MSISDN in L4/L7 appflow records
	*/
	Subscriberawareness string `json:"subscriberawareness,omitempty"`
	/**
	* Enable this option for obfuscating MSISDN in L4/L7 appflow records
	*/
	Subscriberidobfuscation string `json:"subscriberidobfuscation,omitempty"`
	/**
	* Algorithm(MD5 or SHA256) to be used for obfuscating MSISDN
	*/
	Subscriberidobfuscationalgo string `json:"subscriberidobfuscationalgo,omitempty"`
	/**
	* Enable this option for Gx session reporting
	*/
	Gxsessionreporting string `json:"gxsessionreporting,omitempty"`
	/**
	* Enable/disable the feature individually on appflow action.
	*/
	Securityinsighttraffic string `json:"securityinsighttraffic,omitempty"`
	/**
	* Flag to determine whether cache records need to be exported or not. If this flag is true and IC is enabled, cache records are exported instead of L7 HTTP records
	*/
	Cacheinsight string `json:"cacheinsight,omitempty"`
	/**
	* Enable/disable the feature individually on appflow action.
	*/
	Videoinsight string `json:"videoinsight,omitempty"`
	/**
	* Include the HTTP query segment along with the URL that the Citrix ADC received from the client.
	*/
	Httpquerywithurl string `json:"httpquerywithurl,omitempty"`
	/**
	* Include the URL category record.
	*/
	Urlcategory string `json:"urlcategory,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will send the Large Scale Nat(LSN) records to the configured collectors.
	*/
	Lsnlogging string `json:"lsnlogging,omitempty"`
	/**
	* TCP CQA reporting enable/disable knob.
	*/
	Cqareporting string `json:"cqareporting,omitempty"`
	/**
	* Enable AppFlow user email-id logging.
	*/
	Emailaddress string `json:"emailaddress,omitempty"`
	/**
	* On enabling this option, the NGS will send bandwidth usage record to configured collectors.
	*/
	Usagerecordinterval int `json:"usagerecordinterval,omitempty"`
	/**
	* On enabling this option, NGS will send data used by Web/saas app at the end of every HTTP transaction to configured collectors.
	*/
	Websaasappusagereporting string `json:"websaasappusagereporting,omitempty"`
	/**
	* Enable Citrix ADC Stats to be sent to the Telemetry Agent
	*/
	Metrics string `json:"metrics,omitempty"`
	/**
	* Enable Events to be sent to the Telemetry Agent
	*/
	Events string `json:"events,omitempty"`
	/**
	* Enable Auditlogs to be sent to the Telemetry Agent
	*/
	Auditlogs string `json:"auditlogs,omitempty"`
	/**
	* An observation point ID is identifier for the NetScaler from which appflow records are being exported. By default, the NetScaler IP is the observation point ID.
	*/
	Observationpointid int `json:"observationpointid,omitempty"`
	/**
	* Enable generation of the distributed tracing templates in the Appflow records
	*/
	Distributedtracing string `json:"distributedtracing,omitempty"`
	/**
	* Sampling rate for Distributed Tracing
	*/
	Disttracingsamplingrate int `json:"disttracingsamplingrate,omitempty"`
	/**
	* Interval, in seconds, at which to send tcp attack counters to the configured collectors. If 0 is configured, the record is not sent.
	*/
	Tcpattackcounterinterval int `json:"tcpattackcounterinterval,omitempty"`
	/**
	* To use the Citrix ADC IP to send Logstream records instead of the SNIP
	*/
	Logstreamovernsip string `json:"logstreamovernsip,omitempty"`
	/**
	* Authentication token to be set by the agent.
	*/
	Analyticsauthtoken string `json:"analyticsauthtoken,omitempty"`
	/**
	* To use the Citrix ADC IP to send Time series data such as metrics and events, instead of the SNIP
	*/
	Timeseriesovernsip string `json:"timeseriesovernsip,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Tcpburstreporting string `json:"tcpburstreporting,omitempty"`
	Tcpburstreportingthreshold string `json:"tcpburstreportingthreshold,omitempty"`

}
