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
* Binding class showing the servicegroupmemberlist that can be bound to servicegroup.
*/
type Servicegroupservicegroupmemberlistbinding struct {
	/**
	* Desired servicegroupmember binding set. Any existing servicegroupmember which is not part of the input will be deleted or disabled based on graceful setting on servicegroup.
	*/
	Members []Members `json:"members,omitempty"`
	/**
	* List of servicegroupmembers which couldn't be bound.
	*/
	Failedmembers []Failedmembers `json:"failedmembers,omitempty"`
	/**
	* Name of the service group.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`


}
type Failedmembers struct {
	/**
	* IP Address.
	*/
	Ip string `json:"ip,omitempty"`
	/**
	* The port number of the service to be enabled.
	*/
	Port int `json:"port,omitempty"`
}

type Members struct {
	/**
	* IP Address.
	*/
	Ip string `json:"ip,omitempty"`
	/**
	* The port number of the service to be enabled.
	*/
	Port int `json:"port,omitempty"`
	Weight int `json:"weight,omitempty"`
	/**
	* Initial state of the service group.
	*/
	State string `json:"state,omitempty"`
	/**
	* Order number to be assigned to the servicegroup member
	*/
	Order int `json:"order,omitempty"`
}
