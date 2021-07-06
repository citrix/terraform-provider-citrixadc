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

package cs

/**
* Binding class showing the gslbvserver that can be bound to csvserver.
*/
type Csvservergslbvserverbinding struct {
	/**
	* Name of the default gslb or vpn vserver bound to CS vserver of type GSLB/VPN. For Example: bind cs vserver cs1 -vserver gslb1 or bind cs vserver cs1 -vserver vpn1
	*/
	Vserver string `json:"vserver,omitempty"`
	/**
	* Number of hits.
	*/
	Hits int `json:"hits,omitempty"`
	/**
	* Name of the content switching virtual server to which the content switching policy applies.
	*/
	Name string `json:"name,omitempty"`


}