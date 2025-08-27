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
* Configuration for application firewall form field type resource.
*/
type Appfwfieldtype struct {
	/**
	* Name for the field type.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the field type is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my field type" or 'my field type').
	*/
	Name string `json:"name,omitempty"`
	/**
	* PCRE - format regular expression defining the characters and length allowed for this field type.
	*/
	Regex string `json:"regex,omitempty"`
	/**
	* Positive integer specifying the priority of the field type. A lower number specifies a higher priority. Field types are checked in the order of their priority numbers.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Comment describing the type of field that this field type is intended to match.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* will not show internal field types added as part of FieldFormat learn rules deployment
	*/
	Nocharmaps bool `json:"nocharmaps,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
