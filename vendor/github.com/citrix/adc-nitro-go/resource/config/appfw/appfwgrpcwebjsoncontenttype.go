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
* Configuration for gRRPC-Web-json content type resource.
*/
type Appfwgrpcwebjsoncontenttype struct {
	/**
	* Content type to be classified as gRPC-web-json
	*/
	Grpcwebjsoncontenttypevalue string `json:"grpcwebjsoncontenttypevalue,omitempty"`
	/**
	* Is gRPC-web-json content type a regular expression?
	*/
	Isregex string `json:"isregex,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
