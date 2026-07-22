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
* Configuration for HTTP request/response band resource.
*/
type Protocolhttpband struct {
	/**
	* Band size, in bytes, for HTTP request band statistics. For example, if you specify a band size of 100 bytes, statistics will be maintained and displayed for the following size ranges:
		0 - 99 bytes
		100 - 199 bytes
		200 - 299 bytes and so on.
	*/
	Reqbandsize *int `json:"reqbandsize,omitempty"`
	/**
	* Band size, in bytes, for HTTP response band statistics. For example, if you specify a band size of 100 bytes, statistics will be maintained and displayed for the following size ranges:
		0 - 99 bytes
		100 - 199 bytes
		200 - 299 bytes and so on.
	*/
	Respbandsize *int `json:"respbandsize,omitempty"`
	/**
	* Type of statistics to display.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid *int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Bandrange string `json:"bandrange,omitempty"`
	Numberofbands string `json:"numberofbands,omitempty"`
	Totalbandsize string `json:"totalbandsize,omitempty"`
	Avgbandsize string `json:"avgbandsize,omitempty"`
	Avgbandsizenew string `json:"avgbandsizenew,omitempty"`
	Banddata string `json:"banddata,omitempty"`
	Banddatanew string `json:"banddatanew,omitempty"`
	Accesscount string `json:"accesscount,omitempty"`
	Accessratio string `json:"accessratio,omitempty"`
	Accessrationew string `json:"accessrationew,omitempty"`
	Totals string `json:"totals,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
