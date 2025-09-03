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

package subscriber

/**
* Configuration for Subscriber Profile resource.
*/
type Subscriberprofile struct {
	/**
	* Subscriber ip address
	*/
	Ip string `json:"ip,omitempty"`
	/**
	* The vlan number on which the subscriber is located.
	*/
	Vlan int `json:"vlan"` // Zero is a valid value
	/**
	* Rules configured for this subscriber. This is similar to rules received from PCRF for dynamic subscriber sessions.
	*/
	Subscriberrules []string `json:"subscriberrules,omitempty"`
	/**
	* Subscription-Id type
	*/
	Subscriptionidtype string `json:"subscriptionidtype,omitempty"`
	/**
	* Subscription-Id value
	*/
	Subscriptionidvalue string `json:"subscriptionidvalue,omitempty"`
	/**
	*  Name of the servicepath to be taken for this subscriber.
	*/
	Servicepath string `json:"servicepath,omitempty"`

	//------- Read only Parameter ---------;

	Flags string `json:"flags,omitempty"`
	Ttl string `json:"ttl,omitempty"`
	Avpdisplaybuffer string `json:"avpdisplaybuffer,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
