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

package dns

/**
* Configuration for Active DNS resolution entries resource.
*/
type Dnsresolverentry struct {
	/**
	* To get a detailed view of ongoing DNS resolution entries
	*/
	Detail bool `json:"detail,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid *int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Hostname string `json:"hostname,omitempty"`
	Nsrectype string `json:"nsrectype,omitempty"`
	Cnameentry string `json:"cnameentry,omitempty"`
	Dnsqueryid string `json:"dnsqueryid,omitempty"`
	Passiverips string `json:"passiverips,omitempty"`
	Resolutiondepth string `json:"resolutiondepth,omitempty"`
	Cnameresolutiondepth string `json:"cnameresolutiondepth,omitempty"`
	Sourceip string `json:"sourceip,omitempty"`
	Srcport string `json:"srcport,omitempty"`
	Destip string `json:"destip,omitempty"`
	Destport string `json:"destport,omitempty"`
	Issecondaryrip string `json:"issecondaryrip,omitempty"`
	Isresolutioninprogress string `json:"isresolutioninprogress,omitempty"`
	Ispassiverip string `json:"ispassiverip,omitempty"`
	Isloopcase string `json:"isloopcase,omitempty"`
	Isedns string `json:"isedns,omitempty"`
	Dnsprotocol string `json:"dnsprotocol,omitempty"`
	Activeripdevno string `json:"activeripdevno,omitempty"`
	Peid string `json:"peid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
