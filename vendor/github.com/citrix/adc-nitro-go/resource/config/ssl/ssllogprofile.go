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

package ssl

/**
* Configuration for SSL logging Profile resource.
*/
type Ssllogprofile struct {
	/**
	* The name of the ssllogprofile.
	*/
	Name string `json:"name,omitempty"`
	/**
	* log all SSL ClAuth events.
	*/
	Ssllogclauth string `json:"ssllogclauth,omitempty"`
	/**
	* log all SSL ClAuth error events.
	*/
	Ssllogclauthfailures string `json:"ssllogclauthfailures,omitempty"`
	/**
	* log all SSL HS events.
	*/
	Sslloghs string `json:"sslloghs,omitempty"`
	/**
	* log all SSL HS error events.
	*/
	Sslloghsfailures string `json:"sslloghsfailures,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
