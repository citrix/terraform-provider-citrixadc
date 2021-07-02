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

package ns

/**
* Configuration for Surge Protection parameter resource.
*/
type Nsspparams struct {
	/**
	* Maximum number of server connections that can be opened before surge protection is activated.
	*/
	Basethreshold int32 `json:"basethreshold,omitempty"`
	/**
	* Rate at which the system opens connections to the server.
	*/
	Throttle string `json:"throttle,omitempty"`

	//------- Read only Parameter ---------;

	Table0 string `json:"table0,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
