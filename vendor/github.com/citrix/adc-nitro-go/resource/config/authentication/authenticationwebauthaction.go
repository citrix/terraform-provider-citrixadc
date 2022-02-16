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
* Configuration for Web authentication action resource.
*/
type Authenticationwebauthaction struct {
	/**
	* Name for the Web Authentication action. 
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication action" or 'my authentication action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* IP address of the web server to be used for authentication.
	*/
	Serverip string `json:"serverip,omitempty"`
	/**
	* Port on which the web server accepts connections.
	*/
	Serverport int `json:"serverport,omitempty"`
	/**
	* Exact HTTP request, in the form of an expression, which the Citrix ADC sends to the authentication server.
		The Citrix ADC does not check the validity of this request. One must manually validate the request.
	*/
	Fullreqexpr string `json:"fullreqexpr,omitempty"`
	/**
	* Type of scheme for the web server.
	*/
	Scheme string `json:"scheme,omitempty"`
	/**
	* Expression, that checks to see if authentication is successful.
	*/
	Successrule string `json:"successrule,omitempty"`
	/**
	* This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
	*/
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute1 from the webauth response
	*/
	Attribute1 string `json:"attribute1,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute2 from the webauth response
	*/
	Attribute2 string `json:"attribute2,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute3 from the webauth response
	*/
	Attribute3 string `json:"attribute3,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute4 from the webauth response
	*/
	Attribute4 string `json:"attribute4,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute5 from the webauth response
	*/
	Attribute5 string `json:"attribute5,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute6 from the webauth response
	*/
	Attribute6 string `json:"attribute6,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute7 from the webauth response
	*/
	Attribute7 string `json:"attribute7,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute8 from the webauth response
	*/
	Attribute8 string `json:"attribute8,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute9 from the webauth response
	*/
	Attribute9 string `json:"attribute9,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute10 from the webauth response
	*/
	Attribute10 string `json:"attribute10,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute11 from the webauth response
	*/
	Attribute11 string `json:"attribute11,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute12 from the webauth response
	*/
	Attribute12 string `json:"attribute12,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute13 from the webauth response
	*/
	Attribute13 string `json:"attribute13,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute14 from the webauth response
	*/
	Attribute14 string `json:"attribute14,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute15 from the webauth response
	*/
	Attribute15 string `json:"attribute15,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute16 from the webauth response
	*/
	Attribute16 string `json:"attribute16,omitempty"`

}
