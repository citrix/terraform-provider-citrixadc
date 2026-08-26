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

package policy

/**
* Configuration for policy parameter resource.
*/
type Policyparam struct {
	/**
	* Maximum time in milliseconds to allow for processing expressions and policies without interruption. If the timeout is reached then the evaluation causes an UNDEF to be raised and no further processing is performed.
	*/
	Timeout *int `json:"timeout,omitempty"`
	/**
	* Maximum event size in kilobytes that the policy engine will process. When event data exceeds this limit, the action specified by maxEventSizeExceedAction is taken. This parameter helps prevent resource exhaustion from processing extremely large events.
	*/
	Maxeventsize *int `json:"maxeventsize,omitempty"`
	/**
	* Action to take when event data exceeds maxEventSize:
		* RESET - Terminate the connection immediately with TCP RST (most secure).
		* BYPASS - When the limit is exceeded, forward the entire event to the client without policy evaluation or processing. All event data (both parsed and remaining) is sent as-is until the event boundary is reached.
	*/
	Maxeventsizeexceedaction string `json:"maxeventsizeexceedaction,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
