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
* Configuration for limit Indetifier resource.
*/
type Nslimitidentifier struct {
	/**
	* Name for a rate limit identifier. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Reserved words must not be used.
	*/
	Limitidentifier string `json:"limitidentifier,omitempty"`
	/**
	* Maximum number of requests that are allowed in the given timeslice when requests (mode is set as REQUEST_RATE) are tracked per timeslice.
		When connections (mode is set as CONNECTION) are tracked, it is the total number of connections that would be let through.
	*/
	Threshold int `json:"threshold,omitempty"`
	/**
	* Time interval, in milliseconds, specified in multiples of 10, during which requests are tracked to check if they cross the threshold. This argument is needed only when the mode is set to REQUEST_RATE.
	*/
	Timeslice int `json:"timeslice,omitempty"`
	/**
	* Defines the type of traffic to be tracked.
		* REQUEST_RATE - Tracks requests/timeslice.
		* CONNECTION - Tracks active transactions.
		Examples
		1. To permit 20 requests in 10 ms and 2 traps in 10 ms:
		add limitidentifier limit_req -mode request_rate -limitType smooth -timeslice 1000 -Threshold 2000 -trapsInTimeSlice 200
		2. To permit 50 requests in 10 ms:
		set  limitidentifier limit_req -mode request_rate -timeslice 1000 -Threshold 5000 -limitType smooth
		3. To permit 1 request in 40 ms:
		set limitidentifier limit_req -mode request_rate -timeslice 2000 -Threshold 50 -limitType smooth
		4. To permit 1 request in 200 ms and 1 trap in 130 ms:
		set limitidentifier limit_req -mode request_rate -timeslice 1000 -Threshold 5 -limitType smooth -trapsInTimeSlice 8
		5. To permit 5000 requests in 1000 ms and 200 traps in 1000 ms:
		set limitidentifier limit_req  -mode request_rate -timeslice 1000 -Threshold 5000 -limitType BURSTY
	*/
	Mode string `json:"mode,omitempty"`
	/**
	* Smooth or bursty request type.
		* SMOOTH - When you want the permitted number of requests in a given interval of time to be spread evenly across the timeslice
		* BURSTY - When you want the permitted number of requests to exhaust the quota anytime within the timeslice.
		This argument is needed only when the mode is set to REQUEST_RATE.
	*/
	Limittype string `json:"limittype,omitempty"`
	/**
	* Name of the rate limit selector. If this argument is NULL, rate limiting will be applied on all traffic received by the virtual server or the Citrix ADC (depending on whether the limit identifier is bound to a virtual server or globally) without any filtering.
	*/
	Selectorname string `json:"selectorname,omitempty"`
	/**
	* Maximum bandwidth permitted, in kbps.
	*/
	Maxbandwidth int `json:"maxbandwidth,omitempty"`
	/**
	* Number of traps to be sent in the timeslice configured. A value of 0 indicates that traps are disabled.
	*/
	Trapsintimeslice int `json:"trapsintimeslice,omitempty"`

	//------- Read only Parameter ---------;

	Ngname string `json:"ngname,omitempty"`
	Hits string `json:"hits,omitempty"`
	Drop string `json:"drop,omitempty"`
	Rule string `json:"rule,omitempty"`
	Time string `json:"time,omitempty"`
	Total string `json:"total,omitempty"`
	Trapscomputedintimeslice string `json:"trapscomputedintimeslice,omitempty"`
	Computedtraptimeslice string `json:"computedtraptimeslice,omitempty"`
	Referencecount string `json:"referencecount,omitempty"`

}
