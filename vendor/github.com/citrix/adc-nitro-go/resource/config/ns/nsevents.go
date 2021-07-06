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
* Configuration for events resource.
*/
type Nsevents struct {
	/**
	* Event number starting from which events must be shown.
	*/
	Eventno int `json:"eventno,omitempty"`

	//------- Read only Parameter ---------;

	Time string `json:"time,omitempty"`
	Eventcode string `json:"eventcode,omitempty"`
	Devid string `json:"devid,omitempty"`
	Devname string `json:"devname,omitempty"`
	Text string `json:"text,omitempty"`
	Data0 string `json:"data0,omitempty"`
	Data1 string `json:"data1,omitempty"`
	Data2 string `json:"data2,omitempty"`
	Data3 string `json:"data3,omitempty"`

}
