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

package responder

/**
* Binding class showing the responderpolicy that can be bound to responderglobal.
*/
type Responderglobalresponderpolicybinding struct {
	/**
	* Name of the responder policy.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Specifies the bind point whose policies you want to display. Available settings function as follows:
		* REQ_OVERRIDE - Request override. Binds the policy to the priority request queue.
		* REQ_DEFAULT - Binds the policy to the default request queue.
		* OTHERTCP_REQ_OVERRIDE - Binds the policy to the non-HTTP TCP priority request queue.
		* OTHERTCP_REQ_DEFAULT - Binds the policy to the non-HTTP TCP default request queue..
		* SIPUDP_REQ_OVERRIDE - Binds the policy to the SIP UDP priority response queue..
		* SIPUDP_REQ_DEFAULT - Binds the policy to the SIP UDP default response queue.
		* RADIUS_REQ_OVERRIDE - Binds the policy to the RADIUS priority response queue..
		* RADIUS_REQ_DEFAULT - Binds the policy to the RADIUS default response queue.
		* MSSQL_REQ_OVERRIDE - Binds the policy to the Microsoft SQL priority response queue..
		* MSSQL_REQ_DEFAULT - Binds the policy to the Microsoft SQL default response queue.
		* MYSQL_REQ_OVERRIDE - Binds the policy to the MySQL priority response queue.
		* MYSQL_REQ_DEFAULT - Binds the policy to the MySQL default response queue.
		* HTTPQUIC_REQ_OVERRIDE - Binds the policy to the HTTP_QUIC override response queue.
		* HTTPQUIC_REQ_DEFAULT - Binds the policy to the HTTP_QUIC default response queue.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Specifies the priority of the policy.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
	*/
	Invoke bool `json:"invoke,omitempty"`
	/**
	* Type of invocation, Available settings function as follows:
		* vserver - Forward the request to the specified virtual server.
		* policylabel - Invoke the specified policy label.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the policy label to invoke. If the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is policylabel.
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* number of polices bound to label.
	*/
	Numpol int `json:"numpol,omitempty"`
	/**
	* flowtype of the bound responder policy.
	*/
	Flowtype int `json:"flowtype,omitempty"`
	Globalbindtype string `json:"globalbindtype,omitempty"`


}