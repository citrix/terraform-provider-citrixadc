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
* Configuration for cerificate resource.
*/
type Sslcert struct {
	/**
	* Name for and, optionally, path to the generated certificate file. /nsconfig/ssl/ is the default path.
	*/
	Certfile string `json:"certfile,omitempty"`
	/**
	* Name for and, optionally, path to the certificate-signing request (CSR). /nsconfig/ssl/ is the default path.
	*/
	Reqfile string `json:"reqfile,omitempty"`
	/**
	* Type of certificate to generate. Specify one of the following:
		* ROOT_CERT - Self-signed Root-CA certificate. You must specify the key file name. The generated Root-CA certificate can be used for signing end-user client or server certificates or to create Intermediate-CA certificates.
		* INTM_CERT - Intermediate-CA certificate.
		* CLNT_CERT - End-user client certificate used for client authentication.
		* SRVR_CERT - SSL server certificate used on SSL servers for end-to-end encryption.
	*/
	Certtype string `json:"certtype,omitempty"`
	/**
	* Name for and, optionally, path to the private key. You can either use an existing RSA or DSA key that you own or create a new private key on the Citrix ADC. This file is required only when creating a self-signed Root-CA certificate. The key file is stored in the /nsconfig/ssl directory by default.
		If the input key specified is an encrypted key, you are prompted to enter the PEM pass phrase that was used for encrypting the key.
	*/
	Keyfile string `json:"keyfile,omitempty"`
	/**
	* Format in which the key is stored on the appliance.
	*/
	Keyform string `json:"keyform,omitempty"`
	Pempassphrase string `json:"pempassphrase,omitempty"`
	/**
	* Number of days for which the certificate will be valid, beginning with the time and day (system time) of creation.
	*/
	Days uint32 `json:"days,omitempty"`
	/**
	* Subject Alternative Name (SAN) is an extension to X.509 that allows various values to be associated with a security certificate using a subjectAltName field. These values are called "Subject Alternative Names" (SAN). Names include:
		1. Email addresses
		2. IP addresses
		3. URIs
		4. DNS names (This is usually also provided as the Common Name RDN within the Subject field of the main certificate.)
		5. directory names (alternative Distinguished Names to that given in the Subject)
	*/
	Subjectaltname string `json:"subjectaltname,omitempty"`
	/**
	* Format in which the certificate is stored on the appliance.
	*/
	Certform string `json:"certform,omitempty"`
	/**
	* Name of the CA certificate file that issues and signs the Intermediate-CA certificate or the end-user client and server certificates.
	*/
	Cacert string `json:"cacert,omitempty"`
	/**
	* Format of the CA certificate.
	*/
	Cacertform string `json:"cacertform,omitempty"`
	/**
	* Private key, associated with the CA certificate that is used to sign the Intermediate-CA certificate or the end-user client and server certificate. If the CA key file is password protected, the user is prompted to enter the pass phrase that was used to encrypt the key.
	*/
	Cakey string `json:"cakey,omitempty"`
	/**
	* Format for the CA certificate.
	*/
	Cakeyform string `json:"cakeyform,omitempty"`
	/**
	* Serial number file maintained for the CA certificate. This file contains the serial number of the next certificate to be issued or signed by the CA. If the specified file does not exist, a new file is created, with /nsconfig/ssl/ as the default path. If you do not specify a proper path for the existing serial file, a new serial file is created. This might change the certificate serial numbers assigned by the CA certificate to each of the certificates it signs.
	*/
	Caserial string `json:"caserial,omitempty"`

}
