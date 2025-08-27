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

package ipsecalg

/**
* Configuration for IPSEC ALG session resource.
*/
type Ipsecalgsession struct {
	/**
	* Original Source IP address
	*/
	Sourceipalg string `json:"sourceip_alg,omitempty"`
	/**
	* Natted Source IP address
	*/
	Natipalg string `json:"natip_alg,omitempty"`
	/**
	* Destination IP address
	*/
	Destipalg string `json:"destip_alg,omitempty"`
	/**
	* Original Source IP address
	*/
	Sourceip string `json:"sourceip,omitempty"`
	/**
	* Natted Source IP address
	*/
	Natip string `json:"natip,omitempty"`
	/**
	* Destination IP address
	*/
	Destip string `json:"destip,omitempty"`

	//------- Read only Parameter ---------;

	Spiin string `json:"spiin,omitempty"`
	Spiout string `json:"spiout,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
