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

package pcp

/**
* Configuration for server resource.
*/
type Pcpmap struct {
	/**
	* Type of sessions to be displayed.
	*/
	Nattype string `json:"nattype,omitempty"`

	//------- Read only Parameter ---------;

	Pcpsrcip string `json:"pcpsrcip,omitempty"`
	Subscrip string `json:"subscrip,omitempty"`
	Pcpsrcport string `json:"pcpsrcport,omitempty"`
	Pcpdstip string `json:"pcpdstip,omitempty"`
	Pcpdstport string `json:"pcpdstport,omitempty"`
	Pcpnatip string `json:"pcpnatip,omitempty"`
	Pcpnatport string `json:"pcpnatport,omitempty"`
	Pcpprotocol string `json:"pcpprotocol,omitempty"`
	Pcpaddr string `json:"pcpaddr,omitempty"`
	Pcpnounce string `json:"pcpnounce,omitempty"`
	Pcprefcnt string `json:"pcprefcnt,omitempty"`
	Pcplifetime string `json:"pcplifetime,omitempty"`

}
