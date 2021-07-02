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
* Configuration for dns policy label resource.
*/
type Dnspolicylabel struct {
	/**
	* Name of the dns policy label.
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* The type of transformations allowed by the policies bound to the label.
	*/
	Transform string `json:"transform,omitempty"`
	/**
	* The new name of the dns policylabel.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Numpol string `json:"numpol,omitempty"`
	Hits string `json:"hits,omitempty"`
	Priority string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Invokelabelname string `json:"invoke_labelname,omitempty"`
	Flowtype string `json:"flowtype,omitempty"`
	Description string `json:"description,omitempty"`
	Isdefault string `json:"isdefault,omitempty"`

}
