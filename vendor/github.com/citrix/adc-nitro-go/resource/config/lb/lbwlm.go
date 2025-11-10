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
* Configuration for web log manager resource.
*/
type Lbwlm struct {
	/**
	* The name of the Work Load Manager.
	*/
	Wlmname string `json:"wlmname,omitempty"`
	/**
	* The IP address of the WLM.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* The port of the WLM.
	*/
	Port *int `json:"port,omitempty"`
	/**
	* The LBUID for the Load Balancer to communicate to the Work Load Manager.
	*/
	Lbuid string `json:"lbuid,omitempty"`
	/**
	* The idle time period after which Citrix ADC would probe the WLM. The value ranges from 1 to 1440 minutes.
	*/
	Katimeout *int `json:"katimeout,omitempty"`

	//------- Read only Parameter ---------;

	Secure string `json:"secure,omitempty"`
	State string `json:"state,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
