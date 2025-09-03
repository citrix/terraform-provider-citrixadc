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
* Configuration for Negotiate action resource.
*/
type Authenticationnegotiateaction struct {
	/**
	* Name for the AD KDC server profile (negotiate action).
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after AD KDC server profile is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication action" or 'my authentication action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Domain name of the service principal that represnts Citrix ADC.
	*/
	Domain string `json:"domain,omitempty"`
	/**
	* User name of the account that is mapped with Citrix ADC principal. This can be given along with domain and password when keytab file is not available. If username is given along with keytab file, then that keytab file will be searched for this user's credentials.
	*/
	Domainuser string `json:"domainuser,omitempty"`
	/**
	* Password of the account that is mapped to the Citrix ADC principal.
	*/
	Domainuserpasswd string `json:"domainuserpasswd,omitempty"`
	/**
	* Active Directory organizational units (OU) attribute.
	*/
	Ou string `json:"ou,omitempty"`
	/**
	* This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
	*/
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	/**
	* The path to the keytab file that is used to decrypt kerberos tickets presented to Citrix ADC. If keytab is not available, domain/username/password can be specified in the negotiate action configuration
	*/
	Keytab string `json:"keytab,omitempty"`
	/**
	* The path to the site that is enabled for NTLM authentication, including FQDN of the server. This is used when clients fallback to NTLM.
	*/
	Ntlmpath string `json:"ntlmpath,omitempty"`

	//------- Read only Parameter ---------;

	Kcdspn string `json:"kcdspn,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
