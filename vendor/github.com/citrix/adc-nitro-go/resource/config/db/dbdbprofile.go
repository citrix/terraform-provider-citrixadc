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

package db

/**
* Configuration for DB profile resource.
*/
type Dbdbprofile struct {
	/**
	* Name for the database profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile'). 
	*/
	Name string `json:"name,omitempty"`
	/**
	* If ENABLED, inspect the query and update the connection information, if required. If DISABLED, forward the query to the server.
	*/
	Interpretquery string `json:"interpretquery,omitempty"`
	/**
	* If the queries are related to each other, forward to the same backend server.
	*/
	Stickiness string `json:"stickiness,omitempty"`
	/**
	* Name of the KCD account that is used for Windows authentication.
	*/
	Kcdaccount string `json:"kcdaccount,omitempty"`
	/**
	* Use the same server-side connection for multiple client-side requests. Default is enabled.
	*/
	Conmultiplex string `json:"conmultiplex,omitempty"`
	/**
	* Enable caching when connection multiplexing is OFF.
	*/
	Enablecachingconmuxoff string `json:"enablecachingconmuxoff,omitempty"`

	//------- Read only Parameter ---------;

	Refcnt string `json:"refcnt,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
