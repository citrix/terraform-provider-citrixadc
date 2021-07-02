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
* Binding class showing the monitor that can be bound to gslbdomain.
*/
type Gslbdomainmonitorbinding struct {
	/**
	* Monitor name
	*/
	Monitorname string `json:"monitorname,omitempty"`
	/**
	* The service name.
	*/
	Servicename string `json:"servicename,omitempty"`
	Vservername string `json:"vservername,omitempty"`
	/**
	* Monitor state
	*/
	Monstate string `json:"monstate,omitempty"`
	/**
	* HTTP request to the backend server
	*/
	Httprequest string `json:"httprequest,omitempty"`
	/**
	* The state of the monitor for tunneled devices.
	*/
	Iptunnel string `json:"iptunnel,omitempty"`
	/**
	* The string that is sent to the service. Applicable to HTTP ,HTTP-ECV and RTSP monitor types.
	*/
	Customheaders string `json:"customheaders,omitempty"`
	/**
	* The response codes.
	*/
	Respcode string `json:"respcode,omitempty"`
	/**
	* Total monitor probes
	*/
	Monitortotalprobes uint32 `json:"monitortotalprobes,omitempty"`
	/**
	* Total probes failed
	*/
	Monitortotalfailedprobes uint32 `json:"monitortotalfailedprobes,omitempty"`
	/**
	* Total number of current failed probes
	*/
	Monitorcurrentfailedprobes uint32 `json:"monitorcurrentfailedprobes,omitempty"`
	/**
	* Response time of this monitor.
	*/
	Responsetime uint64 `json:"responsetime,omitempty"`
	/**
	* The code indicating the monitor response.
	*/
	Monstatcode int32 `json:"monstatcode,omitempty"`
	/**
	* The string form of monstatcode.
	*/
	Lastresponse string `json:"lastresponse,omitempty"`
	/**
	* Name of the Domain
	*/
	Name string `json:"name,omitempty"`


}