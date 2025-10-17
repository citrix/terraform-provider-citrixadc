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
* Binding class showing the channel that can be bound to fis.
*/
type Fischannelbinding struct {
	/**
	* Interface to be bound to the FIS, specified in slot/port notation (for example, 1/3)
	*/
	Ifnum string `json:"ifnum,omitempty"`
	/**
	* ID of the cluster node for which you are creating the FIS. Can be configured only through the cluster IP address.
	*/
	Ownernode *int `json:"ownernode,omitempty"`
	/**
	* The name of the FIS to which you want to bind interfaces.
	*/
	Name string `json:"name,omitempty"`


}