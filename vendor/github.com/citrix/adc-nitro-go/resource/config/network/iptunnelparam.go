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
* Configuration for ip tunnel parameter resource.
*/
type Iptunnelparam struct {
	/**
	* Common source-IP address for all tunnels. For a specific tunnel, this global setting is overridden if you have specified another source IP address. Must be a MIP or SNIP address.
	*/
	Srcip string `json:"srcip,omitempty"`
	/**
	* Drop any IP packet that requires fragmentation before it is sent through the tunnel.
	*/
	Dropfrag string `json:"dropfrag,omitempty"`
	/**
	* Threshold value, as a percentage of CPU usage, at which to drop packets that require fragmentation to use the IP tunnel. Applies only if dropFragparameter is set to NO. The default value, 0, specifies that this parameter is not set.
	*/
	Dropfragcputhreshold int `json:"dropfragcputhreshold,omitempty"`
	/**
	* Use a different source IP address for each new session through a particular IP tunnel, as determined by round robin selection of one of the SNIP addresses. This setting is ignored if a common global source IP address has been specified for all the IP tunnels. This setting does not apply to a tunnel for which a source IP address has been specified.
	*/
	Srciproundrobin string `json:"srciproundrobin,omitempty"`
	/**
	* Strict PBR check for IPSec packets received through tunnel
	*/
	Enablestrictrx string `json:"enablestrictrx,omitempty"`
	/**
	* Strict PBR check for packets to be sent IPSec protected
	*/
	Enablestricttx string `json:"enablestricttx,omitempty"`
	/**
	* The shared MAC used for shared IP between cluster nodes/HA peers
	*/
	Mac string `json:"mac,omitempty"`
	/**
	* Use client source IP as source IP for outer tunnel IP header
	*/
	Useclientsourceip string `json:"useclientsourceip,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
