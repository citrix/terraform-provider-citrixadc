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

package contentinspection

/**
* Configuration for WASM profile resource.
*/
type Contentinspectionwasmprofile struct {
	/**
	* Name of CI WASM profile
	*/
	Name string `json:"name,omitempty"`
	/**
	* Timeout (in milliseconds) for the connection with the CI WASM agent
	*/
	Timeout *int `json:"timeout,omitempty"`
	/**
	* Timeout action for the connection with the CI agent. Either the original request can be bypassed i.e. request/response is forwarded to the endpoint or the transaction is dropped/reset.
	*/
	Timeoutaction string `json:"timeoutaction,omitempty"`
	/**
	* Max data size (in KB) that will be sent to the CI Agent. Default is 16KB. Maximum value that can be configured is 32KB.
	*/
	Maxbodylen *int `json:"maxbodylen,omitempty"`
	/**
	* Transaction data size (in KB) greater than which a transaction is considered as anomalous. Default is 512KB.
	*/
	Anomalousdatasize *int `json:"anomalousdatasize,omitempty"`
	/**
	* Transaction time (in milliseconds) above which a transaction is considered as anomalous. Default is 1 seconds.
	*/
	Anomalousttfbtime *int `json:"anomalousttfbtime,omitempty"`
	/**
	* Name of the WASM Module
	*/
	Wasmmodule string `json:"wasmmodule,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
