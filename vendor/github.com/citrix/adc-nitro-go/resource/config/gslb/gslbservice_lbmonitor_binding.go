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
* Binding class showing the lbmonitor that can be bound to gslbservice.
*/
type Gslbservicelbmonitorbinding struct {
	/**
	* Monitor name.
	*/
	Monitorname string `json:"monitor_name,omitempty"`
	/**
	* State of the monitor bound to gslb service.
	*/
	Monstate string `json:"monstate,omitempty"`
	/**
	* The running state of the monitor on this service.
	*/
	Monitorstate string `json:"monitor_state,omitempty"`
	/**
	* Weight to assign to the monitor-service binding. A larger number specifies a greater weight. Contributes to the monitoring threshold, which determines the state of the service.
	*/
	Weight *int `json:"weight,omitempty"`
	/**
	* The total number of failed probs.
	*/
	Totalfailedprobes *int `json:"totalfailedprobes,omitempty"`
	/**
	* Number of the current failed monitoring probes.
	*/
	Failedprobes *int `json:"failedprobes,omitempty"`
	/**
	* The code indicating the monitor response.
	*/
	Monstatcode *int `json:"monstatcode,omitempty"`
	/**
	* First parameter for use with message code.
	*/
	Monstatparam1 *int `json:"monstatparam1,omitempty"`
	/**
	* Second parameter for use with message code.
	*/
	Monstatparam2 *int `json:"monstatparam2,omitempty"`
	/**
	* Third parameter for use with message code.
	*/
	Monstatparam3 *int `json:"monstatparam3,omitempty"`
	/**
	* Response time of this monitor.
	*/
	Responsetime *int `json:"responsetime,omitempty"`
	/**
	* Total number of probes sent to monitor this service.
	*/
	Monitortotalprobes *int `json:"monitortotalprobes,omitempty"`
	/**
	* Total number of failed probes
	*/
	Monitortotalfailedprobes *int `json:"monitortotalfailedprobes,omitempty"`
	/**
	* Total number of currently failed probes
	*/
	Monitorcurrentfailedprobes *int `json:"monitorcurrentfailedprobes,omitempty"`
	/**
	* Displays the gslb monitor status in string format.
	*/
	Lastresponse string `json:"lastresponse,omitempty"`
	/**
	* Name of the GSLB service.
	*/
	Servicename string `json:"servicename,omitempty"`


}