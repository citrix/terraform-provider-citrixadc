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
* Configuration for SSL action resource.
*/
type Sslaction struct {
	/**
	* Name for the SSL action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the action is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Perform client certificate authentication.
	*/
	Clientauth string `json:"clientauth,omitempty"`
	/**
	* Client certificate verification is mandatory or optional.
	*/
	Clientcertverification string `json:"clientcertverification,omitempty"`
	/**
	* The name of the ssllogprofile.
	*/
	Ssllogprofile string `json:"ssllogprofile,omitempty"`
	/**
	* Insert the entire client certificate into the HTTP header of the request being sent to the web server. The certificate is inserted in ASCII (PEM) format.
	*/
	Clientcert string `json:"clientcert,omitempty"`
	/**
	* Name of the header into which to insert the client certificate.
	*/
	Certheader string `json:"certheader,omitempty"`
	/**
	* Insert the entire client serial number into the HTTP header of the request being sent to the web server.
	*/
	Clientcertserialnumber string `json:"clientcertserialnumber,omitempty"`
	/**
	* Name of the header into which to insert the client serial number.
	*/
	Certserialheader string `json:"certserialheader,omitempty"`
	/**
	* Insert the client certificate subject, also known as the distinguished name (DN), into the HTTP header of the request being sent to the web server.
	*/
	Clientcertsubject string `json:"clientcertsubject,omitempty"`
	/**
	* Name of the header into which to insert the client certificate subject.
	*/
	Certsubjectheader string `json:"certsubjectheader,omitempty"`
	/**
	* Insert the certificate's signature into the HTTP header of the request being sent to the web server. The signature is the value extracted directly from the X.509 certificate signature field. All X.509 certificates contain a signature field.
	*/
	Clientcerthash string `json:"clientcerthash,omitempty"`
	/**
	* Name of the header into which to insert the client certificate signature (hash).
	*/
	Certhashheader string `json:"certhashheader,omitempty"`
	/**
	* Insert the certificate's fingerprint into the HTTP header of the request being sent to the web server. The fingerprint is derived by computing the specified hash value (SHA256, for example) of the DER-encoding of the client certificate.
	*/
	Clientcertfingerprint string `json:"clientcertfingerprint,omitempty"`
	/**
	* Name of the header into which to insert the client certificate fingerprint.
	*/
	Certfingerprintheader string `json:"certfingerprintheader,omitempty"`
	/**
	* Digest algorithm used to compute the fingerprint of the client certificate.
	*/
	Certfingerprintdigest string `json:"certfingerprintdigest,omitempty"`
	/**
	* Insert the certificate issuer details into the HTTP header of the request being sent to the web server.
	*/
	Clientcertissuer string `json:"clientcertissuer,omitempty"`
	/**
	* Name of the header into which to insert the client certificate issuer details.
	*/
	Certissuerheader string `json:"certissuerheader,omitempty"`
	/**
	* Insert the SSL session ID into the HTTP header of the request being sent to the web server. Every SSL connection that the client and the Citrix ADC share has a unique ID that identifies the specific connection.
	*/
	Sessionid string `json:"sessionid,omitempty"`
	/**
	* Name of the header into which to insert the Session ID.
	*/
	Sessionidheader string `json:"sessionidheader,omitempty"`
	/**
	* Insert the cipher suite that the client and the Citrix ADC negotiated for the SSL session into the HTTP header of the request being sent to the web server. The appliance inserts the cipher-suite name, SSL protocol, export or non-export string, and cipher strength bit, depending on the type of browser connecting to the SSL virtual server or service (for example, Cipher-Suite: RC4- MD5 SSLv3 Non-Export 128-bit).
	*/
	Cipher string `json:"cipher,omitempty"`
	/**
	* Name of the header into which to insert the name of the cipher suite.
	*/
	Cipherheader string `json:"cipherheader,omitempty"`
	/**
	* Insert the date from which the certificate is valid into the HTTP header of the request being sent to the web server. Every certificate is configured with the date and time from which it is valid.
	*/
	Clientcertnotbefore string `json:"clientcertnotbefore,omitempty"`
	/**
	* Name of the header into which to insert the date and time from which the certificate is valid.
	*/
	Certnotbeforeheader string `json:"certnotbeforeheader,omitempty"`
	/**
	* Insert the date of expiry of the certificate into the HTTP header of the request being sent to the web server. Every certificate is configured with the date and time at which the certificate expires.
	*/
	Clientcertnotafter string `json:"clientcertnotafter,omitempty"`
	/**
	* Name of the header into which to insert the certificate's expiry date.
	*/
	Certnotafterheader string `json:"certnotafterheader,omitempty"`
	/**
	* If the appliance is in front of an Outlook Web Access (OWA) server, insert a special header field, FRONT-END-HTTPS: ON, into the HTTP requests going to the OWA server. This header communicates to the server that the transaction is HTTPS and not HTTP.
	*/
	Owasupport string `json:"owasupport,omitempty"`
	/**
	* This action takes an argument a vserver name, to this vserver one will be able to forward all the packets.
	*/
	Forward string `json:"forward,omitempty"`
	/**
	* This action will allow to pick CA(s) from the specific CA group, to verify the client certificate.
	*/
	Cacertgrpname string `json:"cacertgrpname,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Referencecount string `json:"referencecount,omitempty"`
	Description string `json:"description,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
