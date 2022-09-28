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
* Configuration for Subscriber Params resource.
*/
type Subscriberparam struct {
	/**
	* Type of subscriber key type IP or IPANDVLAN. IPANDVLAN option can be used only when the interfaceType is set to gxOnly.
		Changing the lookup method should result to the subscriber session database being flushed.
	*/
	Keytype string `json:"keytype,omitempty"`
	/**
	* Subscriber Interface refers to Citrix ADC interaction with control plane protocols, RADIUS and GX.
		Types of subscriber interface: NONE, RadiusOnly, RadiusAndGx, GxOnly.
		NONE: Only static subscribers can be configured.
		RadiusOnly: GX interface is absent. Subscriber information is obtained through RADIUS Accounting messages.
		RadiusAndGx: Subscriber ID obtained through RADIUS Accounting is used to query PCRF. Subscriber information is obtained from both RADIUS and PCRF.
		GxOnly: RADIUS interface is absent. Subscriber information is queried using Subscriber IP or IP+VLAN.
	*/
	Interfacetype string `json:"interfacetype,omitempty"`
	/**
	* q!Idle Timeout, in seconds, after which Citrix ADC will take an idleAction on a subscriber session (refer to 'idleAction' arguement in 'set subscriber param' for more details on idleAction). Any data-plane or control plane activity updates the idleTimeout on subscriber session. idleAction could be to 'just delete the session' or 'delete and CCR-T' (if PCRF is configured) or 'do not delete but send a CCR-U'. 
		Zero value disables the idle timeout. !
	*/
	Idlettl int `json:"idlettl"`
	/**
	* q!Once idleTTL exprires on a subscriber session, Citrix ADC will take an idle action on that session. idleAction could be chosen from one of these ==>
		1. ccrTerminate: (default) send CCR-T to inform PCRF about session termination and delete the session.  
		2. delete: Just delete the subscriber session without informing PCRF.
		3. ccrUpdate: Do not delete the session and instead send a CCR-U to PCRF requesting for an updated session. !
	*/
	Idleaction string `json:"idleaction,omitempty"`
	/**
	*  The ipv6PrefixLookupList should consist of all the ipv6 prefix lengths assigned to the UE's'
	*/
	Ipv6prefixlookuplist []int `json:"ipv6prefixlookuplist,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
