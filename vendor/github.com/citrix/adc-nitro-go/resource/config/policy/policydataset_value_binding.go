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

/**
* Binding class showing the value that can be bound to policydataset.
*/
type Policydatasetvaluebinding struct {
	/**
	* Value of the specified type that is associated with the dataset. For ipv4 and ipv6, value can be a subnet using the slash notation address/n, where address is the beginning of the subnet and n is the number of left-most bits set in the subnet mask, defining the end of the subnet. The start address will be masked by the subnet mask if necessary, for example for 192.128.128.0/10, the start address will be 192.128.0.0.
	*/
	Value string `json:"value,omitempty"`
	/**
	* The index of the value (ipv4, ipv6, number) associated with the set.
	*/
	Index int `json:"index,omitempty"`
	/**
	* Any comments to preserve information about this dataset or a data bound to this dataset.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* The dataset entry is a range from <value> through <end_range>, inclusive. endRange cannot be used if value is an ipv4 or ipv6 subnet and endRange cannot itself be a subnet.
	*/
	Endrange string `json:"endrange,omitempty"`
	/**
	* Name of the dataset to which to bind the value.
	*/
	Name string `json:"name,omitempty"`


}