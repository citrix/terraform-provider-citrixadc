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

package cs

/**
* Binding class showing the lbvserver that can be bound to csvserver.
*/
type Csvserverlbvserverbinding struct {
	/**
	* Name of the default lb vserver bound. Use this param for Default binding only. For Example: bind cs vserver cs1 -lbvserver lb1
	*/
	Lbvserver string `json:"lbvserver,omitempty"`
	/**
	* Number of hits.
	*/
	Hits int `json:"hits,omitempty"`
	/**
	* Vserver Id of vserver
	*/
	Vserverid string `json:"vserverid,omitempty"`
	/**
	* Vserver id of the lb vserver that is inserted into the set-cookie HTTP header
	*/
	Cookieipport string `json:"cookieipport,omitempty"`
	/**
	* Name of the content switching virtual server to which the content switching policy applies.
	*/
	Name string `json:"name,omitempty"`
	/**
	* The virtual server name (created with the add lb vserver command) to which content will be switched.
	*/
	Targetvserver string `json:"targetvserver,omitempty"`


}