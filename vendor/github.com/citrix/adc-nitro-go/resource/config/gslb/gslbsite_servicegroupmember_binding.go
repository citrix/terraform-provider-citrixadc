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
* Binding class showing the servicegroupmember that can be bound to gslbsite.
*/
type Gslbsiteservicegroupmemberbinding struct {
	/**
	* Servicegroup name.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* IP Address of the gslb service.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Port number of the gslb service.
	*/
	Port int32 `json:"port,omitempty"`
	/**
	* Service type.
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* State of the gslb service.
	*/
	State string `json:"state,omitempty"`
	/**
	* Name of the GSLB site. If you specify a site name, details of all the site's constituent services are also displayed.
	*/
	Sitename string `json:"sitename,omitempty"`


}