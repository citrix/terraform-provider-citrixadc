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
* Configuration for ns mode resource.
*/
type Nsmode struct {
	/**
	* Mode to be enabled. Multiple modes can be specified by providing a blank space between each mode.
	*/
	Mode []string `json:"mode,omitempty"`

	//------- Read only Parameter ---------;

	Fr string `json:"fr,omitempty"`
	L2 string `json:"l2,omitempty"`
	Usip string `json:"usip,omitempty"`
	Cka string `json:"cka,omitempty"`
	Tcpb string `json:"tcpb,omitempty"`
	Mbf string `json:"mbf,omitempty"`
	Edge string `json:"edge,omitempty"`
	Usnip string `json:"usnip,omitempty"`
	L3 string `json:"l3,omitempty"`
	Pmtud string `json:"pmtud,omitempty"`
	Mediaclassification string `json:"mediaclassification,omitempty"`
	Sradv string `json:"sradv,omitempty"`
	Dradv string `json:"dradv,omitempty"`
	Iradv string `json:"iradv,omitempty"`
	Sradv6 string `json:"sradv6,omitempty"`
	Dradv6 string `json:"dradv6,omitempty"`
	Bridgebpdus string `json:"bridgebpdus,omitempty"`
	Ulfd string `json:"ulfd,omitempty"`

}
