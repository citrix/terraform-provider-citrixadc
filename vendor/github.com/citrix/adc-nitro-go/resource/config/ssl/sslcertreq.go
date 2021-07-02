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

package ssl

/**
* Configuration for certificate request resource.
*/
type Sslcertreq struct {
	/**
	* Name for and, optionally, path to the certificate signing request (CSR). /nsconfig/ssl/ is the default path.
	*/
	Reqfile string `json:"reqfile,omitempty"`
	/**
	* Name of and, optionally, path to the private key used to create the certificate signing request, which then becomes part of the certificate-key pair. The private key can be either an RSA or a DSA key. The key must be present in the appliance's local storage. /nsconfig/ssl is the default path.
	*/
	Keyfile string `json:"keyfile,omitempty"`
	/**
	* Subject Alternative Name (SAN) is an extension to X.509 that allows various values to be associated with a security certificate using a subjectAltName field. These values are called "Subject Alternative Names" (SAN). Names include:
		1. Email addresses
		2. IP addresses
		3. URIs
		4. DNS names (this is usually also provided as the Common Name RDN within the Subject field of the main certificate.)
		5. Directory names (alternative Distinguished Names to that given in the Subject)
	*/
	Subjectaltname string `json:"subjectaltname,omitempty"`
	/**
	* Name of the FIPS key used to create the certificate signing request. FIPS keys are created inside the Hardware Security Module of the FIPS card.
	*/
	Fipskeyname string `json:"fipskeyname,omitempty"`
	/**
	* Format in which the key is stored on the appliance.
	*/
	Keyform string `json:"keyform,omitempty"`
	Pempassphrase string `json:"pempassphrase,omitempty"`
	/**
	* Two letter ISO code for your country. For example, US for United States.
	*/
	Countryname string `json:"countryname,omitempty"`
	/**
	* Full name of the state or province where your organization is located.
		Do not abbreviate.
	*/
	Statename string `json:"statename,omitempty"`
	/**
	* Name of the organization that will use this certificate. The organization name (corporation, limited partnership, university, or government agency) must be registered with some authority at the national, state, or city level. Use the legal name under which the organization is registered.
		Do not abbreviate the organization name and do not use the following characters in the name:
		Angle brackets (< >) tilde (~), exclamation mark, at (@), pound (#), zero (0), caret (^), asterisk (*), forward slash (/), square brackets ([ ]), question mark (?).
	*/
	Organizationname string `json:"organizationname,omitempty"`
	/**
	* Name of the division or section in the organization that will use the certificate.
	*/
	Organizationunitname string `json:"organizationunitname,omitempty"`
	/**
	* Name of the city or town in which your organization's head office is located.
	*/
	Localityname string `json:"localityname,omitempty"`
	/**
	* Fully qualified domain name for the company or web site. The common name must match the name used by DNS servers to do a DNS lookup of your server. Most browsers use this information for authenticating the server's certificate during the SSL handshake. If the server name in the URL does not match the common name as given in the server certificate, the browser terminates the SSL handshake or prompts the user with a warning message.
		Do not use wildcard characters, such as asterisk (*) or question mark (?), and do not use an IP address as the common name. The common name must not contain the protocol specifier <http://> or <https://>.
	*/
	Commonname string `json:"commonname,omitempty"`
	/**
	* Contact person's e-mail address. This address is publically displayed as part of the certificate. Provide an e-mail address that is monitored by an administrator who can be contacted about the certificate.
	*/
	Emailaddress string `json:"emailaddress,omitempty"`
	/**
	* Pass phrase, embedded in the certificate signing request that is shared only between the client or server requesting the certificate and the SSL certificate issuer (typically the certificate authority). This pass phrase can be used to authenticate a client or server that is requesting a certificate from the certificate authority.
	*/
	Challengepassword string `json:"challengepassword,omitempty"`
	/**
	* Additional name for the company or web site.
	*/
	Companyname string `json:"companyname,omitempty"`
	/**
	* Digest algorithm used in creating CSR
	*/
	Digestmethod string `json:"digestmethod,omitempty"`

}
