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
* Configuration for ip6 tunnel parameter resource.
*/
type Ip6tunnelparam struct {
	/**
	* Common source IPv6 address for all IPv6 tunnels. Must be a SNIP6 or VIP6 address.
	*/
	Srcip string `json:"srcip,omitempty"`
	/**
	* Drop any packet that requires fragmentation.
	*/
	Dropfrag string `json:"dropfrag,omitempty"`
	/**
	* Threshold value, as a percentage of CPU usage, at which to drop packets that require fragmentation. Applies only if dropFragparameter is set to NO.
	*/
	Dropfragcputhreshold *int `json:"dropfragcputhreshold,omitempty"`
	/**
	* Use a different source IPv6 address for each new session through a particular IPv6 tunnel, as determined by round robin selection of one of the SNIP6 addresses. This setting is ignored if a common global source IPv6 address has been specified for all the IPv6 tunnels. This setting does not apply to a tunnel for which a source IPv6 address has been specified.
	*/
	Srciproundrobin string `json:"srciproundrobin,omitempty"`
	/**
	* Use client source IPv6 address as source IPv6 address for outer tunnel IPv6 header
	*/
	Useclientsourceipv6 string `json:"useclientsourceipv6,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
