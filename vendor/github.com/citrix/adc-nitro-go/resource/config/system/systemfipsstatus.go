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

package system

/**
* Configuration for FIPS INIT Status resource.
*/
type Systemfipsstatus struct {

	//------- Read only Parameter ---------;

	Fipsstatus string `json:"fipsstatus,omitempty"`
	Netscalercontrolplanecryptographiclibraryversion string `json:"netscalercontrolplanecryptographiclibraryversion,omitempty"`
	Netscalerdataplanecryptographiclibraryversion string `json:"netscalerdataplanecryptographiclibraryversion,omitempty"`
	Netscalercrytographicmoduleversion string `json:"netscalercrytographicmoduleversion,omitempty"`
	Intelhwcryptographicacceleratorversion string `json:"intelhwcryptographicacceleratorversion,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
