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

package dns


type Dnsnsecrec struct {
	/**
	* Name of the domain.
	*/
	Hostname string `json:"hostname,omitempty"`
	/**
	* Type of records to display. Available settings function as follows:
		* ADNS - Display all authoritative address records.
		* PROXY - Display all proxy address records.
		* ALL - Display all address records.
	*/
	Type string `json:"type,omitempty"`

	//------- Read only Parameter ---------;

	Nextnsec string `json:"nextnsec,omitempty"`
	Nextrecs string `json:"nextrecs,omitempty"`
	Ttl string `json:"ttl,omitempty"`
	Ecssubnet string `json:"ecssubnet,omitempty"`

}
