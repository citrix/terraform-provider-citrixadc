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
* Configuration for TYPE set resource.
*/
type Policydataset struct {
	/**
	* Name of the dataset. Must not exceed 127 characters.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Type of value to bind to the dataset.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Any comments to preserve information about this dataset or a data bound to this dataset.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* File which contains list of patterns that needs to be bound to the dataset. A patsetfile cannot be associated with multiple datasets.
	*/
	Patsetfile string `json:"patsetfile,omitempty"`
	/**
	* This is used to populate internal dataset information so that the dataset can also be used dynamically in an expression. Here dynamically means the dataset name can also be derived using an expression. For example for a given dataset name "allow_test" it can be used dynamically as client.ip.src.equals_any("allow_" + http.req.url.path.get(1)). This cannot be used with default datasets.
	*/
	Dynamic string `json:"dynamic,omitempty"`
	/**
	* Shows only dynamic datasets when set true.
	*/
	Dynamiconly bool `json:"dynamiconly,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
