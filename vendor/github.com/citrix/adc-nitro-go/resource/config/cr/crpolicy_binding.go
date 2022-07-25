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

package cr

/**
* Binding object which returns the resources bound to crpolicy_binding. 
*/
type Crpolicybinding struct {
	/**
	* Name of the cache redirection policy to display. If this parameter is omitted, details of all the policies are displayed.<br/>Minimum value =  
	*/
	Policyname string `json:"policyname,omitempty"`


}