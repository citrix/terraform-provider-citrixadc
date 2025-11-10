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

package contentinspection

/**
* Binding class showing the contentinspectionpolicy that can be bound to contentinspectionglobal.
*/
type Contentinspectionglobalcontentinspectionpolicybinding struct {
	/**
	* Name of the contentInspection policy.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* The bindpoint to which to policy is bound.
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
	* Terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
	*/
	Invoke bool `json:"invoke,omitempty"`
	/**
	* Type of invocation. Available settings function as follows:
		* reqvserver - Forward the request to the specified request virtual server.
		* resvserver - Forward the response to the specified response virtual server.
		* policylabel - Invoke the specified policy label.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* * If labelType is policylabel, name of the policy label to invoke.
		* If labelType is reqvserver or resvserver, name of the virtual server to which to forward the request of response.
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* The number of policies bound to the bindpoint.
	*/
	Numpol *int `json:"numpol,omitempty"`
	/**
	* flowtype of the bound contentInspection policy.
	*/
	Flowtype *int `json:"flowtype,omitempty"`
	Globalbindtype string `json:"globalbindtype,omitempty"`


}