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
* Configuration for responser parameter resource.
*/
type Responderparam struct {
	/**
	* Action to perform when policy evaluation creates an UNDEF condition. Available settings function as follows:
		* NOOP - Send the request to the protected server.
		* RESET - Reset the request and notify the user's browser, so that the user can resend the request.
		* DROP - Drop the request without sending a response to the user.
	*/
	Undefaction string `json:"undefaction,omitempty"`
	/**
	* Maximum time in milliseconds to allow for processing all the policies and their selected actions without interruption. If the timeout is reached then the evaluation causes an UNDEF to be raised and no further processing is performed.
	*/
	Timeout int `json:"timeout,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
