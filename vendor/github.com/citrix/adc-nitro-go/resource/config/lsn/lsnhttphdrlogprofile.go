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

package lsn

/**
* Configuration for LSN HTTP header logging Profile resource.
*/
type Lsnhttphdrlogprofile struct {
	/**
	* The name of the HTTP header logging Profile.
	*/
	Httphdrlogprofilename string `json:"httphdrlogprofilename,omitempty"`
	/**
	* URL information is logged if option is enabled.
	*/
	Logurl string `json:"logurl,omitempty"`
	/**
	* HTTP method information is logged if option is enabled.
	*/
	Logmethod string `json:"logmethod,omitempty"`
	/**
	* Version information is logged if option is enabled.
	*/
	Logversion string `json:"logversion,omitempty"`
	/**
	* Host information is logged if option is enabled.
	*/
	Loghost string `json:"loghost,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
