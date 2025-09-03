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

package vpn

/**
* Binding class showing the sslcertkey that can be bound to vpnglobal.
*/
type Vpnglobalsslcertkeybinding struct {
	/**
	* SSL certkey to use in signing tokens. Only RSA cert key is allowed
	*/
	Certkeyname string `json:"certkeyname,omitempty"`
	/**
	* The state of the CRL check parameter (Mandatory/Optional).
	*/
	Crlcheck string `json:"crlcheck,omitempty"`
	/**
	* The state of the OCSP check parameter (Mandatory/Optional).
	*/
	Ocspcheck string `json:"ocspcheck,omitempty"`
	/**
	* Certificate to be used for encrypting user data like KB Question and Answers, Alternate Email Address, etc.
	*/
	Userdataencryptionkey string `json:"userdataencryptionkey,omitempty"`
	/**
	* The name of the CA certificate binding.
	*/
	Cacert string `json:"cacert,omitempty"`
	/**
	* Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`


}