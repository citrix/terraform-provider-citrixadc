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
* Binding class showing the gslbservice that can be bound to gslbdomain.
*/
type Gslbdomaingslbservicebinding struct {
	/**
	* The service name.
	*/
	Servicename string `json:"servicename,omitempty"`
	/**
	* The type GSLB service
	*/
	Servicetype string `json:"servicetype,omitempty"`
	Vservername string `json:"vservername,omitempty"`
	/**
	* The Ip address of the service
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Port Number
	*/
	Port int `json:"port,omitempty"`
	/**
	* The state of the vserver
	*/
	State string `json:"state,omitempty"`
	/**
	* weight assigned
	*/
	Weight int `json:"weight,omitempty"`
	/**
	* dynamic weight
	*/
	Dynamicconfwt int `json:"dynamicconfwt,omitempty"`
	/**
	* cumlative weight
	*/
	Cumulativeweight int `json:"cumulativeweight,omitempty"`
	/**
	* GSLB server state
	*/
	Svreffgslbstate string `json:"svreffgslbstate,omitempty"`
	/**
	* The threshold value of the service
	*/
	Gslbthreshold int `json:"gslbthreshold,omitempty"`
	/**
	* The cname of the gslb service
	*/
	Cnameentry string `json:"cnameentry,omitempty"`
	/**
	* Order number assigned to the service when it is bound to the gslb vserver.
	*/
	Order int `json:"order,omitempty"`
	/**
	* Name of the Domain
	*/
	Name string `json:"name,omitempty"`


}