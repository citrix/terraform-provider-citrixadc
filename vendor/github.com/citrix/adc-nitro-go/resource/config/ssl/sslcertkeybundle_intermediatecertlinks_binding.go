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
* Binding class showing the intermediatecertlinks that can be bound to sslcertkeybundle.
*/
type Sslcertkeybundleintermediatecertlinksbinding struct {
	/**
	* Subject name.
	*/
	Subject string `json:"subject,omitempty"`
	/**
	* Serial number.
	*/
	Serial string `json:"serial,omitempty"`
	/**
	* Issuer name.
	*/
	Issuer string `json:"issuer,omitempty"`
	/**
	* Public key algorithm.
	*/
	Publickey string `json:"publickey,omitempty"`
	/**
	* Size of the public key.
	*/
	Publickeysize int `json:"publickeysize,omitempty"`
	/**
	* Subject Alternative Name (SAN) is an extension to X.509 that allows various values to be associated with a security certificate using a subjectAltName field. These values are called "Subject Alternative Names" (SAN). This field is for DNS names
	*/
	Sandns string `json:"sandns,omitempty"`
	/**
	* Subject Alternative Name (SAN) is an extension to X.509 that allows various values to be associated with a security certificate using a subjectAltName field. These values are called "Subject Alternative Names" (SAN). This field is for IP address
	*/
	Sanipadd string `json:"sanipadd,omitempty"`
	/**
	* Not-Before date.
	*/
	Clientcertnotbefore string `json:"clientcertnotbefore,omitempty"`
	/**
	* Not-After date.
	*/
	Clientcertnotafter string `json:"clientcertnotafter,omitempty"`
	/**
	* Days remaining for the certificate to expire.
	*/
	Daystoexpiration int `json:"daystoexpiration,omitempty"`
	/**
	* Signature algorithm.
	*/
	Signaturealg string `json:"signaturealg,omitempty"`
	/**
	* Status of the certificate.
	*/
	Status string `json:"status,omitempty"`
	/**
	* Name given to the cerKeyBundle. The name will be used to bind/unbind certkey bundle to vip. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my file" or 'my file').
	*/
	Certkeybundlename string `json:"certkeybundlename,omitempty"`


}