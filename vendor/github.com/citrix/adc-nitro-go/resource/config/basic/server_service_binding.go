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

package basic

/**
* Binding class showing the service that can be bound to server.
*/
type Serverservicebinding struct {
	/**
	* The services attatched to the server.
	*/
	Servicename string `json:"servicename,omitempty"`
	/**
	* The type of bound service
	*/
	Svctype string `json:"svctype,omitempty"`
	/**
	* The IP address of the bound service
	*/
	Serviceipaddress string `json:"serviceipaddress,omitempty"`
	/**
	* The port number to be used for the bound service.
	*/
	Port *int `json:"port,omitempty"`
	/**
	* The state of the bound service
	*/
	Svrstate string `json:"svrstate,omitempty"`
	/**
	* This field has been intorduced to show the dbs services ip
	*/
	Serviceipstr string `json:"serviceipstr,omitempty"`
	/**
	* Name of the server for which to display parameters.
	*/
	Name string `json:"name,omitempty"`


}