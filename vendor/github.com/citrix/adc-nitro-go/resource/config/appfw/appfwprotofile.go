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

package appfw

/**
* Configuration for gRPC protofile resource.
*/
type Appfwprotofile struct {
	/**
	* Name of the gRPC schema object.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Indicates source path of the gRPC schema file.
	*/
	Src string `json:"src,omitempty"`
	/**
	* Overwrite any existing gRPC schema object of the same name.
	*/
	Overwrite bool `json:"overwrite,omitempty"`
	/**
	* Comments associated with this gRPC schema file.
	*/
	Comment string `json:"comment,omitempty"`

	//------- Read only Parameter ---------;

	Response string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
