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

package aaa

/**
* Configuration for pre authentication parameter resource.
*/
type Aaapreauthenticationparameter struct {
	/**
	* Deny or allow login on the basis of end point analysis results.
	*/
	Preauthenticationaction string `json:"preauthenticationaction,omitempty"`
	/**
	* Name of the Citrix ADC named rule, or an expression, to be evaluated by the EPA tool.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* String specifying the name of a process to be terminated by the EPA tool.
	*/
	Killprocess string `json:"killprocess,omitempty"`
	/**
	* String specifying the path(s) to and name(s) of the files to be deleted by the EPA tool, as a string of between 1 and 1023 characters.
	*/
	Deletefiles string `json:"deletefiles,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
