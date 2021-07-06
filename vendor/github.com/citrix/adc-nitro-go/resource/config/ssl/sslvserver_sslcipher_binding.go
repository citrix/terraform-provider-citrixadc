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

package ssl

/**
* Binding class showing the sslcipher that can be bound to sslvserver.
*/
type Sslvserversslcipherbinding struct {
	/**
	* The name of the cipher group/alias/individual cipheri bindings.
	*/
	Cipheraliasname string `json:"cipheraliasname,omitempty"`
	/**
	* The cipher suite description.
	*/
	Description string `json:"description,omitempty"`
	/**
	* Name of the SSL virtual server.
	*/
	Vservername string `json:"vservername,omitempty"`
	/**
	* Name of the individual cipher, user-defined cipher group, or predefined (built-in) cipher alias.
	*/
	Ciphername string `json:"ciphername,omitempty"`


}