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
* Binding class showing the gslbservicegroupmember that can be bound to gslbdomain.
*/
type Gslbdomaingslbservicegroupmemberbinding struct {
	/**
	* The GSLB service group name bound to the selected GSLB virtual server.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* The Ip address of the service
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Port Number
	*/
	Port int `json:"port,omitempty"`
	/**
	* The type GSLB service
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* weight assigned
	*/
	Weight int `json:"weight,omitempty"`
	/**
	* GSLB server state
	*/
	Svreffgslbstate string `json:"svreffgslbstate,omitempty"`
	/**
	* The threshold value of the service
	*/
	Gslbthreshold int `json:"gslbthreshold,omitempty"`
	/**
	* Order number assigned to the service when it is bound to the gslb vserver.
	*/
	Order int `json:"order,omitempty"`
	/**
	* Name of the Domain
	*/
	Name string `json:"name,omitempty"`


}