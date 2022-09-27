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

package smpp

/**
* Configuration for SMPP user resource.
*/
type Smppuser struct {
	/**
	* Name of the SMPP user. Must be the same as the user name specified in the SMPP server.
	*/
	Username string `json:"username,omitempty"`
	/**
	* Password for binding to the SMPP server. Must be the same as the password specified in the SMPP server.
	*/
	Password string `json:"password,omitempty"`

}
