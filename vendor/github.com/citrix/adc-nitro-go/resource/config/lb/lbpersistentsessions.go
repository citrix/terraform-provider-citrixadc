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

package lb

/**
* Configuration for persistence session resource.
*/
type Lbpersistentsessions struct {
	/**
	* The name of the virtual server.
	*/
	Vserver string `json:"vserver,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid int `json:"nodeid,omitempty"`
	/**
	* The persistence parameter whose persistence sessions are to be flushed.
	*/
	Persistenceparameter string `json:"persistenceparameter,omitempty"`

	//------- Read only Parameter ---------;

	Type string `json:"type,omitempty"`
	Typestring string `json:"typestring,omitempty"`
	Srcip string `json:"srcip,omitempty"`
	Srcipv6 string `json:"srcipv6,omitempty"`
	Destip string `json:"destip,omitempty"`
	Destipv6 string `json:"destipv6,omitempty"`
	Flags string `json:"flags,omitempty"`
	Destport string `json:"destport,omitempty"`
	Vservername string `json:"vservername,omitempty"`
	Timeout string `json:"timeout,omitempty"`
	Referencecount string `json:"referencecount,omitempty"`
	Persistenceparam string `json:"persistenceparam,omitempty"`
	Cnamepersparam string `json:"cnamepersparam,omitempty"`

}
