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
* Binding object which returns the resources bound to nsservicepath_binding. 
*/
type Nsservicepathbinding struct {
	/**
	* Name for the Service path. Must begin with an ASCII alphanumeric or underscore (_) character, and must
		contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-)
		characters.<br/>Minimum value =  
	*/
	Servicepathname string `json:"servicepathname,omitempty"`


}