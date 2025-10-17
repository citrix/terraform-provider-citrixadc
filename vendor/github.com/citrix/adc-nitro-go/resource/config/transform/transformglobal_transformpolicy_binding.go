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

package transform

/**
* Binding class showing the transformpolicy that can be bound to transformglobal.
*/
type Transformglobaltransformpolicybinding struct {
	/**
	* Name of the transform policy.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Specifies the bind point to which to bind the policy. Available settings function as follows:
		* REQ_OVERRIDE. Request override. Binds the policy to the priority request queue.
		* REQ_DEFAULT. Binds the policy to the default request queue.
		* HTTPQUIC_REQ_OVERRIDE - Binds the policy to the HTTP_QUIC override request queue.
		* HTTPQUIC_REQ_DEFAULT - Binds the policy to the HTTP_QUIC default request queue.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Specifies the priority of the policy.
	*/
	Priority *int `json:"priority,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forwards the request or response to the specified virtual server or evaluates the specified policy label.
	*/
	Invoke bool `json:"invoke,omitempty"`
	/**
	* Type of invocation. Available settings function as follows:
		* reqvserver - Send the request to the specified request virtual server.
		* resvserver - Send the response to the specified response virtual server.
		* policylabel - Invoke the specified policy label.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and the label type is Policy Label.
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* The number of policies bound to the bindpoint.
	*/
	Numpol *int `json:"numpol,omitempty"`
	/**
	* flowtype of the bound transform policy.
	*/
	Flowtype *int `json:"flowtype,omitempty"`
	Globalbindtype string `json:"globalbindtype,omitempty"`


}