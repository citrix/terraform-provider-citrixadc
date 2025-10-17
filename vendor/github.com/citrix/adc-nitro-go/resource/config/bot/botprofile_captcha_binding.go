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

package bot

/**
* Binding class showing the captcha that can be bound to botprofile.
*/
type Botprofilecaptchabinding struct {
	/**
	* Captcha action binding. For each URL, only one binding is allowed. To update the values of an existing URL binding, user has to first unbind that binding, and then needs to bind the URL again with new values. Maximum 30 bindings can be configured per profile.
	*/
	Captcharesource bool `json:"captcharesource,omitempty"`
	/**
	* URL for which the Captcha action, if configured under IP reputation, TPS or device fingerprint, need to be applied.
	*/
	Botcaptchaurl string `json:"bot_captcha_url,omitempty"`
	/**
	* Enable or disable the captcha binding.
	*/
	Botcaptchaenabled string `json:"bot_captcha_enabled,omitempty"`
	/**
	* Wait time in seconds for which ADC needs to wait for the Captcha response. This is to avoid DOS attacks.
	*/
	Waittime *int `json:"waittime,omitempty"`
	/**
	* Time (in seconds) duration for which no new captcha challenge is sent after current captcha challenge has been answered successfully.
	*/
	Graceperiod *int `json:"graceperiod,omitempty"`
	/**
	* Time (in seconds) duration for which client which failed captcha need to wait until allowed to try again. The requests from this client are silently dropped during the mute period.
	*/
	Muteperiod *int `json:"muteperiod,omitempty"`
	/**
	* Length of body request (in Bytes) up to (equal or less than) which captcha challenge will be provided to client. Above this length threshold the request will be dropped. This is to avoid DOS and DDOS attacks.
	*/
	Requestsizelimit *int `json:"requestsizelimit,omitempty"`
	/**
	* Number of times client can retry solving the captcha.
	*/
	Retryattempts *int `json:"retryattempts,omitempty"`
	/**
	* One or more actions to be taken when client fails captcha challenge. Only, log action can be configured with DROP, REDIRECT or RESET action.
	*/
	Botcaptchaaction []string `json:"bot_captcha_action,omitempty"`
	/**
	* Message to be logged for this binding.
	*/
	Logmessage string `json:"logmessage,omitempty"`
	/**
	* Any comments about this binding.
	*/
	Botbindcomment string `json:"bot_bind_comment,omitempty"`
	/**
	* Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
	*/
	Name string `json:"name,omitempty"`


}