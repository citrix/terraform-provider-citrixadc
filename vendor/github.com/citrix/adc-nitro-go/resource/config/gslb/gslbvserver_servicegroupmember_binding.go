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
* Binding class showing the servicegroupmember that can be bound to gslbvserver.
*/
type Gslbvserverservicegroupmemberbinding struct {
	/**
	* The GSLB service group name bound to the selected GSLB virtual server.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* IP address.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Port number.
	*/
	Port int32 `json:"port,omitempty"`
	/**
	* Protocol used by services bound to the virtual server.
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* State of the gslb vserver.
	*/
	Curstate string `json:"curstate,omitempty"`
	/**
	* Weight to assign to the GSLB service.
	*/
	Weight uint32 `json:"weight,omitempty"`
	/**
	* Specify if the appliance should consider the service count, service weights, or ignore both when using weight-based load balancing methods. The state of the number of services bound to the virtual server help the appliance to select the service.
	*/
	Dynamicweight string `json:"dynamicweight,omitempty"`
	/**
	* The target site to be returned in the DNS response when a policy is successfully evaluated against the incoming DNS request. Target site is specified in dotted notation with up to 6 qualifiers. Wildcard `*' is accepted as a valid qualifier token.
	*/
	Preferredlocation string `json:"preferredlocation,omitempty"`
	/**
	* Effective state of the gslb svc
	*/
	Svreffgslbstate string `json:"svreffgslbstate,omitempty"`
	/**
	* Tells whether threshold exceeded for this service participating in CUSTOMLB
	*/
	Thresholdvalue int32 `json:"thresholdvalue,omitempty"`
	/**
	* Indicates if gslb svc has reached threshold
	*/
	Gslbthreshold int32 `json:"gslbthreshold,omitempty"`
	/**
	* This field is introduced for displaying the cookie in cluster setup.
	*/
	Sitepersistcookie string `json:"sitepersistcookie,omitempty"`
	/**
	* Type of Site Persistence set on the bound service
	*/
	Svcsitepersistence string `json:"svcsitepersistence,omitempty"`
	/**
	* Name of the virtual server on which to perform the binding operation.
	*/
	Name string `json:"name,omitempty"`


}