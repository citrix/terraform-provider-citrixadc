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
* Configuration for aptlicense resource.
*/
type Nsaptlicense struct {
	/**
	* Hardware Serial Number/License Activation Code(LAC)
	*/
	Serialno string `json:"serialno,omitempty"`
	/**
	* Specifies whether to use the licenseproxyserver to reach the internet. Make sure to configure licenseproxyserver to use this option.
	*/
	Useproxy string `json:"useproxy,omitempty"`
	/**
	* License ID
	*/
	Id string `json:"id,omitempty"`
	/**
	* Session ID
	*/
	Sessionid string `json:"sessionid,omitempty"`
	/**
	* Bind type
	*/
	Bindtype string `json:"bindtype,omitempty"`
	/**
	* The user can allocate one or more licenses. Ensure the value is less than (for partial allocation) or equal to the total number of available licenses
	*/
	Countavailable string `json:"countavailable,omitempty"`
	/**
	* License Directory
	*/
	Licensedir string `json:"licensedir,omitempty"`

	//------- Read only Parameter ---------;

	Response string `json:"response,omitempty"`
	Counttotal string `json:"counttotal,omitempty"`
	Name string `json:"name,omitempty"`
	Relevance string `json:"relevance,omitempty"`
	Datepurchased string `json:"datepurchased,omitempty"`
	Datesa string `json:"datesa,omitempty"`
	Dateexp string `json:"dateexp,omitempty"`
	Features string `json:"features,omitempty"`

}
