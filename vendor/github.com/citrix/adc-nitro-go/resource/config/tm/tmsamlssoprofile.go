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

package tm

/**
* Configuration for SAML sso action resource.
*/
type Tmsamlssoprofile struct {
	/**
	* Name for the new saml single sign-on profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an SSO action is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the SSL certificate that is used to Sign Assertion.
	*/
	Samlsigningcertname string `json:"samlsigningcertname,omitempty"`
	/**
	* URL to which the assertion is to be sent.
	*/
	Assertionconsumerserviceurl string `json:"assertionconsumerserviceurl,omitempty"`
	/**
	* Expression to extract relaystate to be sent along with assertion. Evaluation of this expression should return TEXT content. This is typically a targ
		et url to which user is redirected after the recipient validates SAML token
	*/
	Relaystaterule string `json:"relaystaterule,omitempty"`
	/**
	* Option to send password in assertion.
	*/
	Sendpassword string `json:"sendpassword,omitempty"`
	/**
	* The name to be used in requests sent from	Citrix ADC to IdP to uniquely identify Citrix ADC.
	*/
	Samlissuername string `json:"samlissuername,omitempty"`
	/**
	* Algorithm to be used to sign/verify SAML transactions
	*/
	Signaturealg string `json:"signaturealg,omitempty"`
	/**
	* Algorithm to be used to compute/verify digest for SAML transactions
	*/
	Digestmethod string `json:"digestmethod,omitempty"`
	/**
	* Audience for which assertion sent by IdP is applicable. This is typically entity name or url that represents ServiceProvider
	*/
	Audience string `json:"audience,omitempty"`
	/**
	* Format of Name Identifier sent in Assertion.
	*/
	Nameidformat string `json:"nameidformat,omitempty"`
	/**
	* Expression that will be evaluated to obtain NameIdentifier to be sent in assertion
	*/
	Nameidexpr string `json:"nameidexpr,omitempty"`
	/**
	* Name of attribute1 that needs to be sent in SAML Assertion
	*/
	Attribute1 string `json:"attribute1,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute1's value to be sent in Assertion
	*/
	Attribute1expr string `json:"attribute1expr,omitempty"`
	/**
	* User-Friendly Name of attribute1 that needs to be sent in SAML Assertion
	*/
	Attribute1friendlyname string `json:"attribute1friendlyname,omitempty"`
	/**
	* Format of Attribute1 to be sent in Assertion.
	*/
	Attribute1format string `json:"attribute1format,omitempty"`
	/**
	* Name of attribute2 that needs to be sent in SAML Assertion
	*/
	Attribute2 string `json:"attribute2,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute2's value to be sent in Assertion
	*/
	Attribute2expr string `json:"attribute2expr,omitempty"`
	/**
	* User-Friendly Name of attribute2 that needs to be sent in SAML Assertion
	*/
	Attribute2friendlyname string `json:"attribute2friendlyname,omitempty"`
	/**
	* Format of Attribute2 to be sent in Assertion.
	*/
	Attribute2format string `json:"attribute2format,omitempty"`
	/**
	* Name of attribute3 that needs to be sent in SAML Assertion
	*/
	Attribute3 string `json:"attribute3,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute3's value to be sent in Assertion
	*/
	Attribute3expr string `json:"attribute3expr,omitempty"`
	/**
	* User-Friendly Name of attribute3 that needs to be sent in SAML Assertion
	*/
	Attribute3friendlyname string `json:"attribute3friendlyname,omitempty"`
	/**
	* Format of Attribute3 to be sent in Assertion.
	*/
	Attribute3format string `json:"attribute3format,omitempty"`
	/**
	* Name of attribute4 that needs to be sent in SAML Assertion
	*/
	Attribute4 string `json:"attribute4,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute4's value to be sent in Assertion
	*/
	Attribute4expr string `json:"attribute4expr,omitempty"`
	/**
	* User-Friendly Name of attribute4 that needs to be sent in SAML Assertion
	*/
	Attribute4friendlyname string `json:"attribute4friendlyname,omitempty"`
	/**
	* Format of Attribute4 to be sent in Assertion.
	*/
	Attribute4format string `json:"attribute4format,omitempty"`
	/**
	* Name of attribute5 that needs to be sent in SAML Assertion
	*/
	Attribute5 string `json:"attribute5,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute5's value to be sent in Assertion
	*/
	Attribute5expr string `json:"attribute5expr,omitempty"`
	/**
	* User-Friendly Name of attribute5 that needs to be sent in SAML Assertion
	*/
	Attribute5friendlyname string `json:"attribute5friendlyname,omitempty"`
	/**
	* Format of Attribute5 to be sent in Assertion.
	*/
	Attribute5format string `json:"attribute5format,omitempty"`
	/**
	* Name of attribute6 that needs to be sent in SAML Assertion
	*/
	Attribute6 string `json:"attribute6,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute6's value to be sent in Assertion
	*/
	Attribute6expr string `json:"attribute6expr,omitempty"`
	/**
	* User-Friendly Name of attribute6 that needs to be sent in SAML Assertion
	*/
	Attribute6friendlyname string `json:"attribute6friendlyname,omitempty"`
	/**
	* Format of Attribute6 to be sent in Assertion.
	*/
	Attribute6format string `json:"attribute6format,omitempty"`
	/**
	* Name of attribute7 that needs to be sent in SAML Assertion
	*/
	Attribute7 string `json:"attribute7,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute7's value to be sent in Assertion
	*/
	Attribute7expr string `json:"attribute7expr,omitempty"`
	/**
	* User-Friendly Name of attribute7 that needs to be sent in SAML Assertion
	*/
	Attribute7friendlyname string `json:"attribute7friendlyname,omitempty"`
	/**
	* Format of Attribute7 to be sent in Assertion.
	*/
	Attribute7format string `json:"attribute7format,omitempty"`
	/**
	* Name of attribute8 that needs to be sent in SAML Assertion
	*/
	Attribute8 string `json:"attribute8,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute8's value to be sent in Assertion
	*/
	Attribute8expr string `json:"attribute8expr,omitempty"`
	/**
	* User-Friendly Name of attribute8 that needs to be sent in SAML Assertion
	*/
	Attribute8friendlyname string `json:"attribute8friendlyname,omitempty"`
	/**
	* Format of Attribute8 to be sent in Assertion.
	*/
	Attribute8format string `json:"attribute8format,omitempty"`
	/**
	* Name of attribute9 that needs to be sent in SAML Assertion
	*/
	Attribute9 string `json:"attribute9,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute9's value to be sent in Assertion
	*/
	Attribute9expr string `json:"attribute9expr,omitempty"`
	/**
	* User-Friendly Name of attribute9 that needs to be sent in SAML Assertion
	*/
	Attribute9friendlyname string `json:"attribute9friendlyname,omitempty"`
	/**
	* Format of Attribute9 to be sent in Assertion.
	*/
	Attribute9format string `json:"attribute9format,omitempty"`
	/**
	* Name of attribute10 that needs to be sent in SAML Assertion
	*/
	Attribute10 string `json:"attribute10,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute10's value to be sent in Assertion
	*/
	Attribute10expr string `json:"attribute10expr,omitempty"`
	/**
	* User-Friendly Name of attribute10 that needs to be sent in SAML Assertion
	*/
	Attribute10friendlyname string `json:"attribute10friendlyname,omitempty"`
	/**
	* Format of Attribute10 to be sent in Assertion.
	*/
	Attribute10format string `json:"attribute10format,omitempty"`
	/**
	* Name of attribute11 that needs to be sent in SAML Assertion
	*/
	Attribute11 string `json:"attribute11,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute11's value to be sent in Assertion
	*/
	Attribute11expr string `json:"attribute11expr,omitempty"`
	/**
	* User-Friendly Name of attribute11 that needs to be sent in SAML Assertion
	*/
	Attribute11friendlyname string `json:"attribute11friendlyname,omitempty"`
	/**
	* Format of Attribute11 to be sent in Assertion.
	*/
	Attribute11format string `json:"attribute11format,omitempty"`
	/**
	* Name of attribute12 that needs to be sent in SAML Assertion
	*/
	Attribute12 string `json:"attribute12,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute12's value to be sent in Assertion
	*/
	Attribute12expr string `json:"attribute12expr,omitempty"`
	/**
	* User-Friendly Name of attribute12 that needs to be sent in SAML Assertion
	*/
	Attribute12friendlyname string `json:"attribute12friendlyname,omitempty"`
	/**
	* Format of Attribute12 to be sent in Assertion.
	*/
	Attribute12format string `json:"attribute12format,omitempty"`
	/**
	* Name of attribute13 that needs to be sent in SAML Assertion
	*/
	Attribute13 string `json:"attribute13,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute13's value to be sent in Assertion
	*/
	Attribute13expr string `json:"attribute13expr,omitempty"`
	/**
	* User-Friendly Name of attribute13 that needs to be sent in SAML Assertion
	*/
	Attribute13friendlyname string `json:"attribute13friendlyname,omitempty"`
	/**
	* Format of Attribute13 to be sent in Assertion.
	*/
	Attribute13format string `json:"attribute13format,omitempty"`
	/**
	* Name of attribute14 that needs to be sent in SAML Assertion
	*/
	Attribute14 string `json:"attribute14,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute14's value to be sent in Assertion
	*/
	Attribute14expr string `json:"attribute14expr,omitempty"`
	/**
	* User-Friendly Name of attribute14 that needs to be sent in SAML Assertion
	*/
	Attribute14friendlyname string `json:"attribute14friendlyname,omitempty"`
	/**
	* Format of Attribute14 to be sent in Assertion.
	*/
	Attribute14format string `json:"attribute14format,omitempty"`
	/**
	* Name of attribute15 that needs to be sent in SAML Assertion
	*/
	Attribute15 string `json:"attribute15,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute15's value to be sent in Assertion
	*/
	Attribute15expr string `json:"attribute15expr,omitempty"`
	/**
	* User-Friendly Name of attribute15 that needs to be sent in SAML Assertion
	*/
	Attribute15friendlyname string `json:"attribute15friendlyname,omitempty"`
	/**
	* Format of Attribute15 to be sent in Assertion.
	*/
	Attribute15format string `json:"attribute15format,omitempty"`
	/**
	* Name of attribute16 that needs to be sent in SAML Assertion
	*/
	Attribute16 string `json:"attribute16,omitempty"`
	/**
	* Expression that will be evaluated to obtain attribute16's value to be sent in Assertion
	*/
	Attribute16expr string `json:"attribute16expr,omitempty"`
	/**
	* User-Friendly Name of attribute16 that needs to be sent in SAML Assertion
	*/
	Attribute16friendlyname string `json:"attribute16friendlyname,omitempty"`
	/**
	* Format of Attribute16 to be sent in Assertion.
	*/
	Attribute16format string `json:"attribute16format,omitempty"`
	/**
	* Option to encrypt assertion when Citrix ADC sends one.
	*/
	Encryptassertion string `json:"encryptassertion,omitempty"`
	/**
	* Name of the SSL certificate of peer/receving party using which Assertion is encrypted.
	*/
	Samlspcertname string `json:"samlspcertname,omitempty"`
	/**
	* Algorithm to be used to encrypt SAML assertion
	*/
	Encryptionalgorithm string `json:"encryptionalgorithm,omitempty"`
	/**
	* This option specifies the number of minutes on either side of current time that the assertion would be valid. For example, if skewTime is 10, then assertion would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.
	*/
	Skewtime *int `json:"skewtime,omitempty"`
	/**
	* Option to sign portions of assertion when Citrix ADC IDP sends one. Based on the user selection, either Assertion or Response or Both or none can be signed
	*/
	Signassertion string `json:"signassertion,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
