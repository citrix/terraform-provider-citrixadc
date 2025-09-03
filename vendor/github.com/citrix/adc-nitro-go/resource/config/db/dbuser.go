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
* Configuration for DB user resource.
*/
type Dbuser struct {
	/**
	* Name of the database user. Must be the same as the user name specified in the database.
	*/
	Username string `json:"username,omitempty"`
	/**
	* Password for logging on to the database. Must be the same as the password specified in the database.
	*/
	Password string `json:"password,omitempty"`
	/**
	* Display the names of all database users currently logged on to the Citrix ADC.
	*/
	Loggedin bool `json:"loggedin,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
