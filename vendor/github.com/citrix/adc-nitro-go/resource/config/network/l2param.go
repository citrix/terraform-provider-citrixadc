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
* Configuration for Layer 2 related parameter resource.
*/
type L2param struct {
	/**
	* When mbf_instant_learning is enabled, learn any changes in peer's MAC after this time interval, which is in 10ms ticks.
	*/
	Mbfpeermacupdate int `json:"mbfpeermacupdate,omitempty"`
	/**
	* Maximum bridge collision for loop detection 
	*/
	Maxbridgecollision int `json:"maxbridgecollision,omitempty"`
	/**
	* Set/reset proxy ARP in bridge group deployment
	*/
	Bdggrpproxyarp string `json:"bdggrpproxyarp,omitempty"`
	/**
	* Bridging settings for C2C behavior. If enabled, each PE will learn MAC entries independently. Otherwise, when L2 mode is ON, learned MAC entries on a PE will be broadcasted to all other PEs.
	*/
	Bdgsetting string `json:"bdgsetting,omitempty"`
	/**
	* Send GARP messagess on VRID-configured interfaces upon failover 
	*/
	Garponvridintf string `json:"garponvridintf,omitempty"`
	/**
	* Allows MAC mode vserver to pick and forward the packets even if it is destined to Citrix ADC owned VIP.
	*/
	Macmodefwdmypkt string `json:"macmodefwdmypkt,omitempty"`
	/**
	* Use Citrix ADC MAC for all outgoing packets.
	*/
	Usemymac string `json:"usemymac,omitempty"`
	/**
	* Proxies the ARP as Citrix ADC MAC for FreeBSD.
	*/
	Proxyarp string `json:"proxyarp,omitempty"`
	/**
	* Set/reset REPLY form of GARP 
	*/
	Garpreply string `json:"garpreply,omitempty"`
	/**
	* Enable instant learning of MAC changes in MBF mode.
	*/
	Mbfinstlearning string `json:"mbfinstlearning,omitempty"`
	/**
	* Enable the reset interface upon HA failover.
	*/
	Rstintfonhafo string `json:"rstintfonhafo,omitempty"`
	/**
	* Control source parameters (IP and Port) for FreeBSD initiated traffic. If Enabled, source parameters are retained. Else proxy the source parameters based on next hop.
	*/
	Skipproxyingbsdtraffic string `json:"skipproxyingbsdtraffic,omitempty"`
	/**
	*  Return to ethernet sender.
	*/
	Returntoethernetsender string `json:"returntoethernetsender,omitempty"`
	/**
	* Stop Update of server mac change to NAT sessions.
	*/
	Stopmacmoveupdate string `json:"stopmacmoveupdate,omitempty"`
	/**
	* Time-out value for the bridge table entries, in seconds. The new value applies only to the entries that are dynamically learned after the new value is set. Previously existing bridge table entries expire after the previously configured time-out value.
	*/
	Bridgeagetimeout int `json:"bridgeagetimeout,omitempty"`
	/**
	* Control source parameters (IP and Port) for FreeBSD initiated traffic. If enabled proxy the source parameters based on netprofile source ip. If netprofile does not have ip configured, then it will continue to use NSIP as earlier.
	*/
	Usenetprofilebsdtraffic string `json:"usenetprofilebsdtraffic,omitempty"`

}
