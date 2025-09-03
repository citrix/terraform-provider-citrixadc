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


type Nsconfigview struct {
	/**
	* State is a session-level setting that controls the configuration shown to the user. Based on the selected option user will be able to see only the configuration created through Classic Interface (CLI/API) or Only Configurations created through Next-Gen API, Or both Classic & Next-Gen API interfaces.
		Configurations created by Nextgenapi are read-only and can only be modified via Next-Gen API REST endpoints. Configurations created by classic interfaces (CLI, NITRO) are editable for CLASSIC or ALL views. Possible values:
		- CLASSIC: Config view is limited to the configuration entities created through classic interfaces
		(CLI, NITRO). This is the default behaviour.
		- NEXTGENAPI: Config view is limited to the configuration entities created through Next-Gen API Interface.
		- ALL: Configurations created by both Classic and Next-Gen API interfaces are visible.
	*/
	State string `json:"state,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
