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
* Configuration for Contentinspection parameter resource.
*/
type Contentinspectionparameter struct {
	/**
	* Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an error condition in evaluating the expression.
		Available settings function as follows:
		* NOINSPECTION - Do not Inspect the traffic.
		* RESET - Reset the connection and notify the user's browser, so that the user can resend the request.
		* DROP - Drop the message without sending a response to the user.
	*/
	Undefaction string `json:"undefaction,omitempty"`

}
