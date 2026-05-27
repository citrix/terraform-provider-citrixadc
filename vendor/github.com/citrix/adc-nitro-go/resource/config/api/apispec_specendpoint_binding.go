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

package api

/**
* Binding class showing the specendpoint that can be bound to apispec.
*/
type Apispecspecendpointbinding struct {
	/**
	* API method used for display purposes
	*/
	Apiname string `json:"apiname,omitempty"`
	/**
	* API endpoint (could be GRPC service, REST path) used for display purposes
	*/
	Apiservice string `json:"apiservice,omitempty"`
	/**
	* GRPC API option method: REST way to reach GRPC Service. One per service/method combination
	*/
	Httpmethod string `json:"httpmethod,omitempty"`
	/**
	* GRPC API option endpoint: REST way to reach GRPC Service. One per service/method combination
	*/
	Httpurlpath string `json:"httpurlpath,omitempty"`
	/**
	* Name of the spec for which to show detailed information.
	*/
	Name string `json:"name,omitempty"`


}