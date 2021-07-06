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
* Binding object which returns the resources bound to vlan_binding. 
*/
type Vlanbinding struct {
	/**
	* Integer that uniquely identifies the VLAN for which the details are to be displayed.<br/>Minimum value =  1<br/>Maximum value =  4094
	*/
	Id int `json:"id,omitempty"`


}