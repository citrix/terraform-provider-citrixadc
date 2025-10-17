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
* Configuration for OAuth Identity Provider (IdP) profile resource.
*/
type Authenticationoauthidpprofile struct {
	/**
	* Name for the new OAuth Identity Provider (IdP) single sign-on profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Unique identity of the relying party requesting for authentication.
	*/
	Clientid string `json:"clientid,omitempty"`
	/**
	* Unique secret string to authorize relying party at authorization server.
	*/
	Clientsecret string `json:"clientsecret,omitempty"`
	/**
	* URL endpoint on relying party to which the OAuth token is to be sent.
	*/
	Redirecturl string `json:"redirecturl,omitempty"`
	/**
	* The name to be used in requests sent from	Citrix ADC to IdP to uniquely identify Citrix ADC.
	*/
	Issuer string `json:"issuer,omitempty"`
	/**
	* Name of the entity that is used to obtain configuration for the current authentication request. It is used only in Citrix Cloud.
	*/
	Configservice string `json:"configservice,omitempty"`
	/**
	* Audience for which token is being sent by Citrix ADC IdP. This is typically entity name or url that represents the recipient
	*/
	Audience string `json:"audience,omitempty"`
	/**
	* This option specifies the duration for which the token sent by Citrix ADC IdP is valid. For example, if skewTime is 10, then token would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.
	*/
	Skewtime *int `json:"skewtime,omitempty"`
	/**
	* This group will be part of AAA session's internal group list. This will be helpful to admin in Nfactor flow to decide right AAA configuration for Relaying Party. In authentication policy AAA.USER.IS_MEMBER_OF("<default_auth_group>")  is way to use this feature.
	*/
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	/**
	* This is the endpoint at which Citrix ADC IdP can get details about Relying Party (RP) being configured. Metadata response should include endpoints for jwks_uri for RP public key(s).
	*/
	Relyingpartymetadataurl string `json:"relyingpartymetadataurl,omitempty"`
	/**
	* Interval at which Relying Party metadata is refreshed.
	*/
	Refreshinterval *int `json:"refreshinterval,omitempty"`
	/**
	* Option to encrypt token when Citrix ADC IDP sends one.
	*/
	Encrypttoken string `json:"encrypttoken,omitempty"`
	/**
	* Name of the service in cloud used to sign the data. This is applicable only if signature if offloaded to cloud.
	*/
	Signatureservice string `json:"signatureservice,omitempty"`
	/**
	* Algorithm to be used to sign OpenID tokens.
	*/
	Signaturealg string `json:"signaturealg,omitempty"`
	/**
	* Name-Value pairs of attributes to be inserted in idtoken. Configuration format is name=value_expr@@@name2=value2_expr@@@.
		'@@@' is used as delimiter between Name-Value pairs. name is a literal string whose value is 127 characters and does not contain '=' character.
		Value is advanced policy expression terminated by @@@ delimiter. Last value need not contain the delimiter.
	*/
	Attributes string `json:"attributes,omitempty"`
	/**
	* Option to send encrypted password in idtoken.
	*/
	Sendpassword string `json:"sendpassword,omitempty"`

	//------- Read only Parameter ---------;

	Oauthstatus string `json:"oauthstatus,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
