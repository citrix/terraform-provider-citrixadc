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

package gslb

/**
* Configuration for GSLB parameter resource.
*/
type Gslbparameter struct {
	/**
	* Time, in seconds, after which an inactive LDNS entry is removed.
	*/
	Ldnsentrytimeout int `json:"ldnsentrytimeout,omitempty"`
	/**
	* Tolerance, in milliseconds, for newly learned round-trip time (RTT) values. If the difference between the old RTT value and the newly computed RTT value is less than or equal to the specified tolerance value, the LDNS entry in the network metric table is not updated with the new RTT value. Prevents the exchange of metrics when variations in RTT values are negligible.
	*/
	Rtttolerance int `json:"rtttolerance,omitempty"`
	/**
	* The IPv4 network mask with which to create LDNS entries.
	*/
	Ldnsmask string `json:"ldnsmask,omitempty"`
	/**
	* Mask for creating LDNS entries for IPv6 source addresses. The mask is defined as the number of leading bits to consider, in the source IP address, when creating an LDNS entry.
	*/
	V6ldnsmasklen int `json:"v6ldnsmasklen,omitempty"`
	/**
	* Order in which monitors should be initiated to calculate RTT.
	*/
	Ldnsprobeorder []string `json:"ldnsprobeorder,omitempty"`
	/**
	* Drop LDNS requests if round-trip time (RTT) information is not available.
	*/
	Dropldnsreq string `json:"dropldnsreq,omitempty"`
	/**
	* Amount of delay in updating the state of GSLB service to DOWN when MEP goes down.
		This parameter is applicable only if monitors are not bound to GSLB services
	*/
	Gslbsvcstatedelaytime int `json:"gslbsvcstatedelaytime,omitempty"`
	/**
	* Time (in seconds) within which local or child site services remain in learning phase. GSLB site will enter the learning phase after reboot, HA failover, Cluster GSLB owner node changes or MEP being enabled on local node.  Backup parent (if configured) will selectively move the adopted children's GSLB services to learning phase when primary parent goes down. While a service is in learning period, remote site will not honour the state and stats got through MEP for that service. State can be learnt from health monitor if bound explicitly.
	*/
	Svcstatelearningtime int `json:"svcstatelearningtime,omitempty"`
	/**
	* GSLB configuration will be synced automatically to remote gslb sites if enabled.
	*/
	Automaticconfigsync string `json:"automaticconfigsync,omitempty"`
	/**
	* Time duartion (in seconds) during which if no new packets received by Local gslb site from Remote gslb site then mark the MEP connection DOWN
	*/
	Mepkeepalivetimeout int `json:"mepkeepalivetimeout,omitempty"`
	/**
	* Time duartion (in seconds) for which the gslb sync process will wait before checking for config changes.
	*/
	Gslbsyncinterval int `json:"gslbsyncinterval,omitempty"`
	/**
	* Mode in which configuration will be synced from master site to remote sites.
	*/
	Gslbsyncmode string `json:"gslbsyncmode,omitempty"`
	/**
	* If disabled, Location files will not be synced to the remote sites as part of automatic sync.
	*/
	Gslbsynclocfiles string `json:"gslbsynclocfiles,omitempty"`
	/**
	* If enabled, remote gslb site's rsync port will be monitored and site is considered for configuration sync only when the monitor is successful.
	*/
	Gslbconfigsyncmonitor string `json:"gslbconfigsyncmonitor,omitempty"`

	//------- Read only Parameter ---------;

	Flags string `json:"flags,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Incarnation string `json:"incarnation,omitempty"`

}
