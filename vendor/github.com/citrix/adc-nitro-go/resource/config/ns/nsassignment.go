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
* Configuration for assignment resource.
*/
type Nsassignment struct {
	/**
	* Name for the assignment. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the assignment is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my assignment" or my assignment).
	*/
	Name string `json:"name,omitempty"`
	/**
	* Left hand side of the assigment, of the form $variable-name (for a singleton variabled) or $variable-name[key-expression], where key-expression is an expression that evaluates to a text string and provides the key to select a map entry
	*/
	Variable string `json:"variable,omitempty"`
	/**
	* Right hand side of the assignment. The expression is evaluated and assigned to the left hand variable.
	*/
	Set string `json:"set,omitempty"`
	/**
	* Right hand side of the assignment. The expression is evaluated and added to the left hand variable.
	*/
	Add string `json:"Add,omitempty"`
	/**
	* Right hand side of the assignment. The expression is evaluated and subtracted from the left hand variable.
	*/
	Sub string `json:"sub,omitempty"`
	/**
	* Right hand side of the assignment. The expression is evaluated and appended to the left hand variable.
	*/
	Append string `json:"append,omitempty"`
	/**
	* Clear the variable value. Deallocates a text value, and for a map, the text key.
	*/
	Clear bool `json:"clear,omitempty"`
	/**
	* Comment. Can be used to preserve information about this rewrite action.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* New name for the assignment.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the rewrite policy is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my assignment" or my assignment).
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Referencecount string `json:"referencecount,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
