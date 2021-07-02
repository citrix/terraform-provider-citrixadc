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
* Configuration for hardware resource.
*/
type Nshardware struct {

	//------- Read only Parameter ---------;

	Hwdescription string `json:"hwdescription,omitempty"`
	Sysid string `json:"sysid,omitempty"`
	Manufactureday string `json:"manufactureday,omitempty"`
	Manufacturemonth string `json:"manufacturemonth,omitempty"`
	Manufactureyear string `json:"manufactureyear,omitempty"`
	Cpufrequncy string `json:"cpufrequncy,omitempty"`
	Hostid string `json:"hostid,omitempty"`
	Host string `json:"host,omitempty"`
	Serialno string `json:"serialno,omitempty"`
	Encodedserialno string `json:"encodedserialno,omitempty"`
	Netscaleruuid string `json:"netscaleruuid,omitempty"`
	Bmcrevision string `json:"bmcrevision,omitempty"`

}
