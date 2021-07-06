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

package ns

/**
* Binding class showing the autoscalepolicy that can be bound to nstimer.
*/
type Nstimerautoscalepolicybinding struct {
	/**
	* The timer policy associated with the timer.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Specifies the priority of the timer policy.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Name of the vserver which provides the context for the rule in timer policy. When not specified it is treated as a Global Default context.
	*/
	Vserver string `json:"vserver,omitempty"`
	/**
	* Denotes the sample size. Sample size value of 'x' means that previous '(x - 1)' policy's rule evaluation results and the current evaluation result are present with the binding. For example, sample size of 10 means that there is a state of previous 9 policy evaluation results and also the current policy evaluation result.
	*/
	Samplesize int `json:"samplesize,omitempty"`
	/**
	* Denotes the threshold. If the rule of the policy in the binding relation evaluates 'threshold size' number of times in 'sample size' to true, then the corresponding action is taken. Its value needs to be less than or equal to the sample size value.
	*/
	Threshold int `json:"threshold,omitempty"`
	/**
	* Timer name.
	*/
	Name string `json:"name,omitempty"`


}