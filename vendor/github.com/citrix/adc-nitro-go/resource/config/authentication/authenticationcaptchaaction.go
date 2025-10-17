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

package authentication

/**
* Configuration for Captcha Action resource.
*/
type Authenticationcaptchaaction struct {
	/**
	* Name for the new captcha action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.
		The following requirement applies only to the NetScaler CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* This is the endpoint at which captcha response is validated.
	*/
	Serverurl string `json:"serverurl,omitempty"`
	/**
	* Secret of gateway as established at the captcha source.
	*/
	Secretkey string `json:"secretkey,omitempty"`
	/**
	* Sitekey to identify gateway fqdn while loading captcha.
	*/
	Sitekey string `json:"sitekey,omitempty"`
	/**
	* This is the group that is added to user sessions that match current policy.
	*/
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	/**
	* This is the score threshold value for recaptcha v3.
	*/
	Scorethreshold *int `json:"scorethreshold,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
