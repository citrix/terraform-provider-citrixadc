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

package contentinspection

/**
* Configuration for ContentInspection policy label resource.
*/
type Contentinspectionpolicylabel struct {
	/**
	* Name for the contentInspection policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the contentInspection policy label is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my contentInspection policy label" or 'my contentInspection policy label').
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* Type of packets (request or response packets) against which to match the policies bound to this policy label.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Any comments to preserve information about this contentInspection policy label.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* New name for the contentInspection policy label.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy label" or 'my policy label').
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
	Isdefault string `json:"isdefault,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
