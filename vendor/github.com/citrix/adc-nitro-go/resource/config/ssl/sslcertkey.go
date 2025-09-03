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
* Configuration for certificate key resource.
*/
type Sslcertkey struct {
	/**
	* Name for the certificate and private-key pair. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the certificate-key pair is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cert" or 'my cert').
	*/
	Certkey string `json:"certkey,omitempty"`
	/**
	* Name of and, optionally, path to the X509 certificate file that is used to form the certificate-key pair. The certificate file should be present on the appliance's hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.
	*/
	Cert string `json:"cert,omitempty"`
	/**
	* Name of and, optionally, path to the private-key file that is used to form the certificate-key pair. The certificate file should be present on the appliance's hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.
	*/
	Key string `json:"key,omitempty"`
	/**
	* Passphrase that was used to encrypt the private-key. Use this option to load encrypted private-keys in PEM format.
	*/
	Password bool `json:"password,omitempty"`
	/**
	* Name of the FIPS key that was created inside the Hardware Security Module (HSM) of a FIPS appliance, or a key that was imported into the HSM.
	*/
	Fipskey string `json:"fipskey,omitempty"`
	/**
	* Name of the HSM key that was created in the External Hardware Security Module (HSM) of a FIPS appliance.
	*/
	Hsmkey string `json:"hsmkey,omitempty"`
	/**
	* Input format of the certificate and the private-key files. The three formats supported by the appliance are:
		PEM - Privacy Enhanced Mail
		DER - Distinguished Encoding Rule
		PFX - Personal Information Exchange
	*/
	Inform string `json:"inform,omitempty"`
	/**
	* Pass phrase used to encrypt the private-key. Required when adding an encrypted private-key in PEM format.
	*/
	Passplain string `json:"passplain,omitempty"`
	/**
	* Issue an alert when the certificate is about to expire.
	*/
	Expirymonitor string `json:"expirymonitor,omitempty"`
	/**
	* Time, in number of days, before certificate expiration, at which to generate an alert that the certificate is about to expire.
	*/
	Notificationperiod int `json:"notificationperiod,omitempty"`
	/**
	* Parse the certificate chain as a single file after linking the server certificate to its issuer's certificate within the file.
	*/
	Bundle string `json:"bundle,omitempty"`
	/**
	* This option is used to automatically delete certificate/key files from physical device when the added certkey is removed. When deleteCertKeyFilesOnRemoval option is used at rm certkey command, it overwrites the deleteCertKeyFilesOnRemoval setting used at add/set certkey command
	*/
	Deletecertkeyfilesonremoval string `json:"deletecertkeyfilesonremoval,omitempty"`
	/**
	* Delete cert/key file from file system.
	*/
	Deletefromdevice bool `json:"deletefromdevice,omitempty"`
	/**
	* Name of the Certificate Authority certificate-key pair to which to link a certificate-key pair.
	*/
	Linkcertkeyname string `json:"linkcertkeyname,omitempty"`
	/**
	* Override the check for matching domain names during a certificate update operation.
	*/
	Nodomaincheck bool `json:"nodomaincheck,omitempty"`
	/**
	* Clear cached ocspStapling response in certkey.
	*/
	Ocspstaplingcache bool `json:"ocspstaplingcache,omitempty"`

	//------- Read only Parameter ---------;

	Signaturealg string `json:"signaturealg,omitempty"`
	Certificatetype string `json:"certificatetype,omitempty"`
	Serial string `json:"serial,omitempty"`
	Issuer string `json:"issuer,omitempty"`
	Clientcertnotbefore string `json:"clientcertnotbefore,omitempty"`
	Clientcertnotafter string `json:"clientcertnotafter,omitempty"`
	Daystoexpiration string `json:"daystoexpiration,omitempty"`
	Subject string `json:"subject,omitempty"`
	Publickey string `json:"publickey,omitempty"`
	Publickeysize string `json:"publickeysize,omitempty"`
	Version string `json:"version,omitempty"`
	Priority string `json:"priority,omitempty"`
	Status string `json:"status,omitempty"`
	Passcrypt string `json:"passcrypt,omitempty"`
	Data string `json:"data,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Sandns string `json:"sandns,omitempty"`
	Sanipadd string `json:"sanipadd,omitempty"`
	Ocspresponsestatus string `json:"ocspresponsestatus,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Certkeydigest string `json:"certkeydigest,omitempty"`
	Certificatesource string `json:"certificatesource,omitempty"`
	Certkeystatus string `json:"certkeystatus,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
