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
* Configuration for limit sessions resource.
*/
type Nslimitsessions struct {
	/**
	* Name of the rate limit identifier for which to display the sessions.
	*/
	Limitidentifier string `json:"limitidentifier,omitempty"`
	/**
	* Show the individual hash values.
	*/
	Detail bool `json:"detail,omitempty"`

	//------- Read only Parameter ---------;

	Timeout string `json:"timeout,omitempty"`
	Hits string `json:"hits,omitempty"`
	Drop string `json:"drop,omitempty"`
	Number string `json:"number,omitempty"`
	Name string `json:"name,omitempty"`
	Unit string `json:"unit,omitempty"`
	Flags string `json:"flags,omitempty"`
	Referencecount string `json:"referencecount,omitempty"`
	Maxbandwidth string `json:"maxbandwidth,omitempty"`
	Selectoripv61 string `json:"selectoripv61,omitempty"`
	Selectoripv62 string `json:"selectoripv62,omitempty"`
	Flag string `json:"flag,omitempty"`

}
