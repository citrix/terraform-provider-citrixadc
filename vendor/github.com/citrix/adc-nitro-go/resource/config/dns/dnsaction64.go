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

/**
* Configuration for dns64 action resource.
*/
type Dnsaction64 struct {
	/**
	* Name of the dns64 action.
	*/
	Actionname string `json:"actionname,omitempty"`
	/**
	* The dns64 prefix to be used if the after evaluating the rules
	*/
	Prefix string `json:"prefix,omitempty"`
	/**
	* The expression to select the criteria for ipv4 addresses to be used for synthesis.
		Only if the mappedrule is evaluated to true the corresponding ipv4 address is used for synthesis using respective prefix,
		otherwise the A RR is discarded
	*/
	Mappedrule string `json:"mappedrule,omitempty"`
	/**
	* The expression to select the criteria for eliminating the corresponding ipv6 addresses from the response.
	*/
	Excluderule string `json:"excluderule,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
