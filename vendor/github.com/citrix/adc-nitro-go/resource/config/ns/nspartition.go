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
* Configuration for admin partition resource.
*/
type Nspartition struct {
	/**
	* Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
	*/
	Partitionname string `json:"partitionname,omitempty"`
	/**
	* Maximum bandwidth, in Kbps, that the partition can consume. A zero value indicates the bandwidth is unrestricted on the partition and it can consume up to the system limits.
	*/
	Maxbandwidth int `json:"maxbandwidth"` // Zero is a valid value
	/**
	* Minimum bandwidth, in Kbps, that the partition can consume. A zero value indicates the bandwidth is unrestricted on the partition and it can consume up to the system limits
	*/
	Minbandwidth int `json:"minbandwidth"` // Zero is a valid value
	/**
	* Maximum number of concurrent connections that can be open in the partition. A zero value indicates no limit on number of open connections.
	*/
	Maxconn int `json:"maxconn"` // Zero is a valid value
	/**
	* Maximum memory, in megabytes, allocated to the partition.  A zero value indicates the memory is unlimited on the partition and it can consume up to the system limits.
	*/
	Maxmemlimit int `json:"maxmemlimit"` // Zero is a valid value
	/**
	* Special MAC address for the partition which is used for communication over shared vlans in this partition. If not specified, the MAC address is auto-generated.
	*/
	Partitionmac string `json:"partitionmac,omitempty"`
	/**
	* Switches to new admin partition without prompt for saving configuration. Configuration will not be saved
	*/
	Force bool `json:"force,omitempty"`
	/**
	* Switches to new admin partition without prompt for saving configuration. Configuration will be saved
	*/
	Save bool `json:"save,omitempty"`

	//------- Read only Parameter ---------;

	Partitionid string `json:"partitionid,omitempty"`
	Partitiontype string `json:"partitiontype,omitempty"`
	Pmacinternal string `json:"pmacinternal,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
