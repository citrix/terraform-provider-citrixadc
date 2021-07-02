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
* Configuration for Source IP Mac Cache Table. resource.
*/
type Nssourceroutecachetable struct {

	//------- Read only Parameter ---------;

	Sourceip string `json:"sourceip,omitempty"`
	Sourcemac string `json:"sourcemac,omitempty"`
	Vlan string `json:"vlan,omitempty"`
	Interface string `json:"Interface,omitempty"`

}
