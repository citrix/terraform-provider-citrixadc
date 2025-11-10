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
* Binding class showing the gslbservice that can be bound to gslbvserver.
*/
type Gslbvservergslbservicebinding struct {
	/**
	* Name of the GSLB service for which to change the weight.
	*/
	Servicename string `json:"servicename,omitempty"`
	/**
	* Weight for the service.
	*/
	Weight *int `json:"weight,omitempty"`
	/**
	* The cname of the gslb service.
	*/
	Cnameentry string `json:"cnameentry,omitempty"`
	/**
	* IP address.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Port number.
	*/
	Port *int `json:"port,omitempty"`
	/**
	* Protocol used by services bound to the GSLBvirtual server.
	*/
	Gslbboundsvctype string `json:"gslbboundsvctype,omitempty"`
	/**
	* State of the gslb vserver.
	*/
	Curstate string `json:"curstate,omitempty"`
	/**
	* Weight obtained by the virtue of bound service count or weight
	*/
	Dynamicconfwt *int `json:"dynamicconfwt,omitempty"`
	/**
	* Cumulative weight is the weight of GSLB service considering both its configured weight and dynamic weight. It is equal to product of dynamic weight and configured weight of the gslb service 
	*/
	Cumulativeweight *int `json:"cumulativeweight,omitempty"`
	/**
	* Effective state of the gslb svc
	*/
	Svreffgslbstate string `json:"svreffgslbstate,omitempty"`
	/**
	* Indicates if gslb svc has reached threshold
	*/
	Gslbthreshold *int `json:"gslbthreshold,omitempty"`
	/**
	* The target site to be returned in the DNS response when a policy is successfully evaluated against the incoming DNS request. Target site is specified in dotted notation with up to 6 qualifiers. Wildcard `*' is accepted as a valid qualifier token.
	*/
	Preferredlocation string `json:"preferredlocation,omitempty"`
	/**
	* Tells whether threshold exceeded for this service participating in CUSTOMLB
	*/
	Thresholdvalue *int `json:"thresholdvalue,omitempty"`
	/**
	* is cname feature set on vserver
	*/
	Iscname string `json:"iscname,omitempty"`
	/**
	* Domain name for which to change the time to live (TTL) and/or backup service IP address.
	*/
	Domainname string `json:"domainname,omitempty"`
	/**
	* This field is introduced for displaying the cookie in cluster setup.
	*/
	Sitepersistcookie string `json:"sitepersistcookie,omitempty"`
	/**
	* Type of Site Persistence set on the bound service
	*/
	Svcsitepersistence string `json:"svcsitepersistence,omitempty"`
	/**
	* Order number to be assigned to the service when it is bound to the lb vserver.
	*/
	Order *int `json:"order,omitempty"`
	/**
	* Order number in string form assigned to the service when it is bound to the lb vserver.
	*/
	Orderstr string `json:"orderstr,omitempty"`
	/**
	* Name of the virtual server on which to perform the binding operation.
	*/
	Name string `json:"name,omitempty"`


}