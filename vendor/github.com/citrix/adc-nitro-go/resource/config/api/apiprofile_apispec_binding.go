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

package api

/**
* Binding class showing the apispec that can be bound to apiprofile.
*/
type Apiprofileapispecbinding struct {
	/**
	* Name for the API spec which will be binded to the profile.
	*/
	Apispec string `json:"apispec,omitempty"`
	/**
	* Name of the API profile in which to bind the API apispec(s).
	*/
	Name string `json:"name,omitempty"`


}