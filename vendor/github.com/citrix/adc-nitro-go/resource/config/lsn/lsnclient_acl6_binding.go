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

package lsn

/**
* Binding class showing the acl6 that can be bound to lsnclient.
*/
type Lsnclientacl6binding struct {
	/**
	* Name of any configured extended ACL6 whose action is ALLOW. The condition specified in the extended ACL6 rule is used as the condition for the LSN client.
	*/
	Acl6name string `json:"acl6name,omitempty"`
	/**
	* ID of the traffic domain on which this subscriber or the subscriber network (as specified by the network parameter) belongs. 
		If you do not specify an ID, the subscriber or the subscriber network becomes part of the default traffic domain.
	*/
	Td uint32 `json:"td,omitempty"`
	/**
	* Name for the LSN client entity. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN client is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn client1" or 'lsn client1'). 
	*/
	Clientname string `json:"clientname,omitempty"`


}