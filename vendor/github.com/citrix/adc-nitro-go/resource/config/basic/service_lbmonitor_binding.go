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

package basic

/**
* Binding class showing the lbmonitor that can be bound to service.
*/
type Servicelbmonitorbinding struct {
	/**
	* The monitor Names.
	*/
	Monitorname string `json:"monitor_name,omitempty"`
	/**
	* The configured state (enable/disable) of the monitor on this server.
	*/
	Monstate string `json:"monstate,omitempty"`
	/**
	* The running state of the monitor on this service.
	*/
	Monitorstate string `json:"monitor_state,omitempty"`
	/**
	* Added this field for getting state value from table.
	*/
	Dupstate string `json:"dup_state,omitempty"`
	/**
	* Weight to assign to the monitor-service binding. When a monitor is UP, the weight assigned to its binding with the service determines how much the monitor contributes toward keeping the health of the service above the value configured for the Monitor Threshold parameter.
	*/
	Weight int `json:"weight,omitempty"`
	/**
	* The weight of the monitor.
	*/
	Dupweight int `json:"dup_weight,omitempty"`
	/**
	* The total number of probs sent.
	*/
	Totalprobes int `json:"totalprobes,omitempty"`
	/**
	* The total number of failed probs.
	*/
	Totalfailedprobes int `json:"totalfailedprobes,omitempty"`
	/**
	* Number of the current failed monitoring probes.
	*/
	Failedprobes int `json:"failedprobes,omitempty"`
	/**
	* The code indicating the monitor response.
	*/
	Monstatcode int `json:"monstatcode,omitempty"`
	/**
	* The string form of monstatcode.
	*/
	Lastresponse string `json:"lastresponse,omitempty"`
	/**
	* First parameter for use with message code.
	*/
	Monstatparam1 int `json:"monstatparam1,omitempty"`
	/**
	* Second parameter for use with message code.
	*/
	Monstatparam2 int `json:"monstatparam2,omitempty"`
	/**
	* Third parameter for use with message code.
	*/
	Monstatparam3 int `json:"monstatparam3,omitempty"`
	/**
	* Response time of this monitor.
	*/
	Responsetime int `json:"responsetime,omitempty"`
	/**
	* Total number of probes sent to monitor this service.
	*/
	Monitortotalprobes int `json:"monitortotalprobes,omitempty"`
	/**
	* Total number of failed probes
	*/
	Monitortotalfailedprobes int `json:"monitortotalfailedprobes,omitempty"`
	/**
	* Total number of currently failed probes
	*/
	Monitorcurrentfailedprobes int `json:"monitorcurrentfailedprobes,omitempty"`
	/**
	* Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.
	*/
	Passive bool `json:"passive,omitempty"`
	/**
	* Name of the service to which to bind a policy or monitor.
	*/
	Name string `json:"name,omitempty"`


}