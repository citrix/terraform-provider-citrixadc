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
* Binding class showing the vxlan that can be bound to nstrafficdomain.
*/
type Nstrafficdomainvxlanbinding struct {
	/**
	* ID of the VXLAN to bind to this traffic domain. More than one VXLAN can be bound to a traffic domain, but the same VXLAN cannot be a part of multiple traffic domains.
	*/
	Vxlan uint32 `json:"vxlan,omitempty"`
	/**
	* Integer value that uniquely identifies a traffic domain.
	*/
	Td uint32 `json:"td,omitempty"`


}