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

package policy


type Policytracing struct {
	/**
	* Policy tracing filter expression. For example: http.req.url.startswith("/this").
	*/
	Filterexpr string `json:"filterexpr,omitempty"`
	/**
	* protocol type for which policy records needs to be collected
	*/
	Protocoltype string `json:"protocoltype,omitempty"`
	/**
	* Set it to yes if need to capture the SSL handshake policies
	*/
	Capturesslhandshakepolicies string `json:"capturesslhandshakepolicies,omitempty"`
	/**
	* Unique ID to identify the current transaction
	*/
	Transactionid string `json:"transactionid,omitempty"`
	/**
	* Show detailed information of the captured records
	*/
	Detail string `json:"detail,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Packetengineid string `json:"packetengineid,omitempty"`
	Clientip string `json:"clientip,omitempty"`
	Destip string `json:"destip,omitempty"`
	Srcport string `json:"srcport,omitempty"`
	Destport string `json:"destport,omitempty"`
	Transactiontime string `json:"transactiontime,omitempty"`
	Policytracingmodule string `json:"policytracingmodule,omitempty"`
	Url string `json:"url,omitempty"`
	Policynames string `json:"policynames,omitempty"`
	Isresponse string `json:"isresponse,omitempty"`
	Isundefpolicy string `json:"isundefpolicy,omitempty"`
	Policytracingrecordcount string `json:"policytracingrecordcount,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
