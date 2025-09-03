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

package ns

/**
* Configuration for Variable runtime data information resource.
*/
type Nsvariablevalues struct {
	Name string `json:"name,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Variablekey string `json:"variablekey,omitempty"`
	Variablevalue string `json:"variablevalue,omitempty"`
	Variabledata string `json:"variabledata,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
