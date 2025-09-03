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

package ica

/**
* Configuration for Config Parameters for NS ICA resource.
*/
type Icaparameter struct {
	/**
	* Enable/Disable Session Reliability on HA failover. The default value is No
	*/
	Enablesronhafailover string `json:"enablesronhafailover,omitempty"`
	/**
	* Enable/Disable HDXInsight for Non NSAP ICA Sessions. The default value is Yes
	*/
	Hdxinsightnonnsap string `json:"hdxinsightnonnsap,omitempty"`
	/**
	* Enable/Disable DF enforcement for EDT PMTUD Control Blocks
	*/
	Edtpmtuddf string `json:"edtpmtuddf,omitempty"`
	/**
	* DF enforcement timeout for EDTPMTUDDF
	*/
	Edtpmtuddftimeout int `json:"edtpmtuddftimeout,omitempty"`
	/**
	* Specify the time interval/period for which L7 Client Latency value is to be calculated. By default, L7 Client Latency is calculated for every packet. The default value is 0
	*/
	L7latencyfrequency int `json:"l7latencyfrequency,omitempty"`
	/**
	* Enable/Disable EDT Loss Tolerant feature
	*/
	Edtlosstolerant string `json:"edtlosstolerant,omitempty"`
	/**
	* Enable/Disable EDT PMTUD Rediscovery
	*/
	Edtpmtudrediscovery string `json:"edtpmtudrediscovery,omitempty"`
	/**
	* Enable/Disable DF Persistence
	*/
	Dfpersistence string `json:"dfpersistence,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
