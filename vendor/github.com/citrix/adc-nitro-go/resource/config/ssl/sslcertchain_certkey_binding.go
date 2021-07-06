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
* Binding class showing the certkey that can be bound to sslcertchain.
*/
type Sslcertchaincertkeybinding struct {
	/**
	* Name of the Linked Certificate
	*/
	Linkcertkeyname string `json:"linkcertkeyname,omitempty"`
	/**
	* Used to find if certificate is linked
	*/
	Islinked bool `json:"islinked,omitempty"`
	/**
	* Used to find if certificate is a CA
	*/
	Isca bool `json:"isca,omitempty"`
	/**
	* Used to find if certificate is linked
	*/
	Addsubject bool `json:"addsubject,omitempty"`
	/**
	* Name of the Certificate
	*/
	Certkeyname string `json:"certkeyname,omitempty"`


}