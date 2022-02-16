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

package authentication

/**
* Configuration for authentication policy label resource.
*/
type Authenticationpolicylabel struct {
	/**
	* Name for the new authentication policy label.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy label" or 'authentication policy label').
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* Type of feature (aaatm or rba) against which to match the policies bound to this policy label.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Any comments to preserve information about this authentication policy label.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Login schema associated with authentication policy label. Login schema defines the UI rendering by providing customization option of the fields. If user intervention is not needed for a given factor such as group extraction, a loginSchema whose authentication schema is "noschema" should be used.
	*/
	Loginschema string `json:"loginschema,omitempty"`
	/**
	* The new name of the auth policy label
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Numpol string `json:"numpol,omitempty"`
	Hits string `json:"hits,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Priority string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Flowtype string `json:"flowtype,omitempty"`
	Description string `json:"description,omitempty"`

}
