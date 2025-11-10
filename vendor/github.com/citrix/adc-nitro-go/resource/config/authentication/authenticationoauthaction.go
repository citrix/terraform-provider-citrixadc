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
* Configuration for OAuth authentication action resource.
*/
type Authenticationoauthaction struct {
	/**
	* Name for the OAuth Authentication action.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication action" or 'my authentication action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Type of the OAuth implementation. Default value is generic implementation that is applicable for most deployments.
	*/
	Oauthtype string `json:"oauthtype,omitempty"`
	/**
	* Authorization endpoint/url to which unauthenticated user will be redirected. Citrix ADC redirects user to this endpoint by adding query parameters including clientid. If this parameter not specified then as default value we take Token Endpoint/URL value. Please note that Authorization Endpoint or Token Endpoint is mandatory for oauthAction
	*/
	Authorizationendpoint string `json:"authorizationendpoint,omitempty"`
	/**
	* URL to which OAuth token will be posted to verify its authenticity. User obtains this token from Authorization server upon successful authentication. Citrix ADC will validate presented token by posting it to the URL configured
	*/
	Tokenendpoint string `json:"tokenendpoint,omitempty"`
	/**
	* URL to which obtained idtoken will be posted to get a decrypted user identity. Encrypted idtoken will be obtained by posting OAuth token to token endpoint. In order to decrypt idtoken, Citrix ADC posts request to the URL configured
	*/
	Idtokendecryptendpoint string `json:"idtokendecryptendpoint,omitempty"`
	/**
	* Unique identity of the client/user who is getting authenticated. Authorization server infers client configuration using this ID
	*/
	Clientid string `json:"clientid,omitempty"`
	/**
	* Secret string established by user and authorization server
	*/
	Clientsecret string `json:"clientsecret,omitempty"`
	/**
	* This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
	*/
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	/**
	* Option to set/unset miscellaneous feature flags.
		Available values function as follows:
		* Base64Encode_Authorization_With_Padding - On setting this value, for endpoints (token and introspect), basic authorization header will be base64 encoded with padding.
		* EnableJWTRequest - By enabling this field, Authorisation request to IDP will have jwt signed 'request' parameter
	*/
	Oauthmiscflags []string `json:"oauthmiscflags,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute1
	*/
	Attribute1 string `json:"attribute1,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute2
	*/
	Attribute2 string `json:"attribute2,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute3
	*/
	Attribute3 string `json:"attribute3,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute4
	*/
	Attribute4 string `json:"attribute4,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute5
	*/
	Attribute5 string `json:"attribute5,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute6
	*/
	Attribute6 string `json:"attribute6,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute7
	*/
	Attribute7 string `json:"attribute7,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute8
	*/
	Attribute8 string `json:"attribute8,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute9
	*/
	Attribute9 string `json:"attribute9,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute10
	*/
	Attribute10 string `json:"attribute10,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute11
	*/
	Attribute11 string `json:"attribute11,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute12
	*/
	Attribute12 string `json:"attribute12,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute13
	*/
	Attribute13 string `json:"attribute13,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute14
	*/
	Attribute14 string `json:"attribute14,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute15
	*/
	Attribute15 string `json:"attribute15,omitempty"`
	/**
	* Name of the attribute to be extracted from OAuth Token and to be stored in the attribute16
	*/
	Attribute16 string `json:"attribute16,omitempty"`
	/**
	* List of attribute names separated by ',' which needs to be extracted.
		Note that preceding and trailing spaces will be removed.
		Attribute name can be 127 bytes and total length of this string should not cross 1023 bytes.
		These attributes have multi-value support separated by ',' and stored as key-value pair in AAA session
	*/
	Attributes string `json:"attributes,omitempty"`
	/**
	* TenantID of the application. This is usually specific to providers such as Microsoft and usually refers to the deployment identifier.
	*/
	Tenantid string `json:"tenantid,omitempty"`
	/**
	* URL of the Graph API service to learn Enterprise Mobility Services (EMS) endpoints.
	*/
	Graphendpoint string `json:"graphendpoint,omitempty"`
	/**
	* Interval at which services are monitored for necessary configuration.
	*/
	Refreshinterval *int `json:"refreshinterval,omitempty"`
	/**
	* URL of the endpoint that contains JWKs (Json Web Key) for JWT (Json Web Token) verification.
	*/
	Certendpoint string `json:"certendpoint,omitempty"`
	/**
	* Audience for which token sent by Authorization server is applicable. This is typically entity name or url that represents the recipient
	*/
	Audience string `json:"audience,omitempty"`
	/**
	* Attribute in the token from which username should be extracted.
	*/
	Usernamefield string `json:"usernamefield,omitempty"`
	/**
	* This option specifies the allowed clock skew in number of minutes that Citrix ADC allows on an incoming token. For example, if skewTime is 10, then token would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.
	*/
	Skewtime *int `json:"skewtime,omitempty"`
	/**
	* Identity of the server whose tokens are to be accepted.
	*/
	Issuer string `json:"issuer,omitempty"`
	/**
	* URL to which OAuth access token will be posted to obtain user information.
	*/
	Userinfourl string `json:"userinfourl,omitempty"`
	/**
	* Path to the file that contains JWKs (Json Web Key) for JWT (Json Web Token) verification.
	*/
	Certfilepath string `json:"certfilepath,omitempty"`
	/**
	* Grant type support. value can be code or password
	*/
	Granttype string `json:"granttype,omitempty"`
	/**
	* If authentication is disabled, password is not sent in the request. 
	*/
	Authentication string `json:"authentication,omitempty"`
	/**
	* URL to which access token would be posted for validation
	*/
	Introspecturl string `json:"introspecturl,omitempty"`
	/**
	* Multivalued option to specify allowed token verification algorithms. 
	*/
	Allowedalgorithms []string `json:"allowedalgorithms,omitempty"`
	/**
	* Option to enable/disable PKCE flow during authentication. 
	*/
	Pkce string `json:"pkce,omitempty"`
	/**
	* Option to select the variant of token authentication method. This method is used while exchanging code with IdP. 
	*/
	Tokenendpointauthmethod string `json:"tokenendpointauthmethod,omitempty"`
	/**
	* Well-known configuration endpoint of the Authorization Server. Citrix ADC fetches server details from this endpoint. 
	*/
	Metadataurl string `json:"metadataurl,omitempty"`
	/**
	* Resource URL for Oauth configuration.
	*/
	Resourceuri string `json:"resourceuri,omitempty"`
	/**
	* Name-Value pairs of attributes to be inserted in request parameter. Configuration format is name=value_expr@@@name2=value2_expr@@@.
		'@@@' is used as delimiter between Name-Value pairs. name is a literal string whose value is 127 characters and does not contain '=' character.
		Value is advanced policy expression terminated by @@@ delimiter. Last value need not contain the delimiter.
	*/
	Requestattribute string `json:"requestattribute,omitempty"`
	/**
	* The expression that will be evaluated to obtain IntuneDeviceId for compliance check against IntuneNAC device compliance endpoint. The expression is applicable when the OAuthType is INTUNE. The maximum length allowed to be used as IntuneDeviceId for the device compliance check from the computed response after the expression evaluation is 41.
		Examples:
		add authentication oauthAction <actionName> -intuneDeviceIdExpression 'AAA.LOGIN.INTUNEURI.AFTER_STR("IntuneDeviceId://")'
	*/
	Intunedeviceidexpression string `json:"intunedeviceidexpression,omitempty"`

	//------- Read only Parameter ---------;

	Oauthstatus string `json:"oauthstatus,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
