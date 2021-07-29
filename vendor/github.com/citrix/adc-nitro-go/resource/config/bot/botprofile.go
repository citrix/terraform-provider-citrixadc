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
* Configuration for Bot profile resource.
*/
type Botprofile struct {
	/**
	* Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of object containing bot static signature details.
	*/
	Signature string `json:"signature,omitempty"`
	/**
	* URL that Bot protection uses as the Error URL.
	*/
	Errorurl string `json:"errorurl,omitempty"`
	/**
	* URL that Bot protection uses as the Trap URL.
	*/
	Trapurl string `json:"trapurl,omitempty"`
	/**
	* Any comments about the purpose of profile, or other useful information about the profile.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Enable white-list bot detection.
	*/
	Botenablewhitelist string `json:"bot_enable_white_list,omitempty"`
	/**
	* Enable black-list bot detection.
	*/
	Botenableblacklist string `json:"bot_enable_black_list,omitempty"`
	/**
	* Enable rate-limit bot detection.
	*/
	Botenableratelimit string `json:"bot_enable_rate_limit,omitempty"`
	/**
	* Enable device-fingerprint bot detection
	*/
	Devicefingerprint string `json:"devicefingerprint,omitempty"`
	/**
	* Action to be taken for device-fingerprint based bot detection.
	*/
	Devicefingerprintaction []string `json:"devicefingerprintaction,omitempty"`
	/**
	* Enable IP-reputation bot detection.
	*/
	Botenableipreputation string `json:"bot_enable_ip_reputation,omitempty"`
	/**
	* Enable trap bot detection.
	*/
	Trap string `json:"trap,omitempty"`
	/**
	* Action to be taken for bot trap based bot detection.
	*/
	Trapaction []string `json:"trapaction,omitempty"`
	/**
	* Actions to be taken if no User-Agent header in the request (Applicable if Signature check is enabled).
	*/
	Signaturenouseragentheaderaction []string `json:"signaturenouseragentheaderaction,omitempty"`
	/**
	* Actions to be taken if multiple User-Agent headers are seen in a request (Applicable if Signature check is enabled). Log action should be combined with other actions
	*/
	Signaturemultipleuseragentheaderaction []string `json:"signaturemultipleuseragentheaderaction,omitempty"`
	/**
	* Enable TPS.
	*/
	Botenabletps string `json:"bot_enable_tps,omitempty"`
	/**
	* Enabling bot device fingerprint protection for mobile clients
	*/
	Devicefingerprintmobile []string `json:"devicefingerprintmobile,omitempty"`
	/**
	* Expression to get the client IP.
	*/
	Clientipexpression string `json:"clientipexpression,omitempty"`
	/**
	* Name of the JavaScript file that the Bot Management feature will insert in the response for keyboard-mouse based detection.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my javascript file name" or 'my javascript file name').
	*/
	Kmjavascriptname string `json:"kmjavascriptname,omitempty"`
	/**
	* Enable keyboard-mouse based bot detection.
	*/
	Kmdetection string `json:"kmdetection,omitempty"`
	/**
	* Size of the KM data send by the browser, needs to be processed on ADC
	*/
	Kmeventspostbodylimit int `json:"kmeventspostbodylimit,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
