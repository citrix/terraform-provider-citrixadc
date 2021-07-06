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
* Configuration for link set resource.
*/
type Linkset struct {
	/**
	* Unique identifier for the linkset. Must be of the form LS/x, where x can be an integer from 1 to 32.
	*/
	Id string `json:"id,omitempty"`

	//------- Read only Parameter ---------;

	Ifnum string `json:"ifnum,omitempty"`

}
