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

package network

/**
* Configuration for Link aggregation control protocol resource.
*/
type Lacp struct {
	/**
	* Priority number that determines which peer device of an LACP LA channel can have control over the LA channel. This parameter is globally applied to all LACP channels on the Citrix ADC. The lower the number, the higher the priority.
	*/
	Syspriority int `json:"syspriority,omitempty"`
	/**
	* The owner node in a cluster for which we want to set the lacp priority. Owner node can vary from 0 to 31. Ownernode value of 254 is used for Cluster.
	*/
	Ownernode int `json:"ownernode"` // Zero is a valid value

	//------- Read only Parameter ---------;

	Devicename string `json:"devicename,omitempty"`
	Mac string `json:"mac,omitempty"`
	Flags string `json:"flags,omitempty"`
	Lacpkey string `json:"lacpkey,omitempty"`
	Clustersyspriority string `json:"clustersyspriority,omitempty"`
	Clustermac string `json:"clustermac,omitempty"`

}
