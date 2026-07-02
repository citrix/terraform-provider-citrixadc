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

package protocol

/**
* Configuration for HTTP band resource.
*/
type Protocolhttpband struct {
	/**
	* Band size, in bytes, for HTTP request band statistics.
	*/
	Reqbandsize *int `json:"reqbandsize,omitempty"`
	/**
	* Band size, in bytes, for HTTP response band statistics.
	*/
	Respbandsize *int `json:"respbandsize,omitempty"`
	/**
	* Type of statistics to display (show/clear filter only).
	*/
	Type string `json:"type,omitempty"`
	/**
	* Unique number that identifies the cluster node (GET filter only).
	*/
	Nodeid *int `json:"nodeid,omitempty"`
}
