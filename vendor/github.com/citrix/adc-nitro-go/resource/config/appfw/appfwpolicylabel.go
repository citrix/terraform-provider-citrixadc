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

package appfw

/**
* Configuration for application firewall policy label resource.
*/
type Appfwpolicylabel struct {
	/**
	* Name for the policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the policy label is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy label" or 'my policy label').
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* Type of transformations allowed by the policies bound to the label. Always http_req for application firewall policy labels.
	*/
	Policylabeltype string `json:"policylabeltype,omitempty"`
	/**
	* The new name of the application firewall policylabel.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Numpol string `json:"numpol,omitempty"`
	Hits string `json:"hits,omitempty"`
	Priority string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Invokelabelname string `json:"invoke_labelname,omitempty"`
	Description string `json:"description,omitempty"`
	Policytype string `json:"policytype,omitempty"`

}
