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
* Configuration for Parameter for extended memory used by LSN and Subscriber Store resource.
*/
type Extendedmemoryparam struct {
	/**
	* Amount of Citrix ADC memory to reserve for the memory used by LSN and Subscriber Session Store feature, in multiples of 2MB.
		Note: If you later reduce the value of this parameter, the amount of active memory is not reduced. Changing the configured memory limit can only increase the amount of active memory.
	*/
	Memlimit int `json:"memlimit,omitempty"`

	//------- Read only Parameter ---------;

	Memlimitactive string `json:"memlimitactive,omitempty"`
	Maxmemlimit string `json:"maxmemlimit,omitempty"`
	Minrequiredmemory string `json:"minrequiredmemory,omitempty"`

}
