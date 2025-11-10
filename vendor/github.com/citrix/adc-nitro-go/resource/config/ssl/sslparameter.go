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
* Configuration for SSL parameter resource.
*/
type Sslparameter struct {
	/**
	* Amount of data to collect before the data is pushed to the crypto hardware for encryption. For large downloads, a larger quantum size better utilizes the crypto resources.
	*/
	Quantumsize string `json:"quantumsize,omitempty"`
	/**
	* Maximum memory size to use for certificate revocation lists (CRLs). This parameter reserves memory for a CRL but sets a limit to the maximum memory that the CRLs loaded on the appliance can consume.
	*/
	Crlmemorysizemb *int `json:"crlmemorysizemb,omitempty"`
	/**
	* Enable strict CA certificate checks on the appliance.
	*/
	Strictcachecks string `json:"strictcachecks,omitempty"`
	/**
	* Time, in milliseconds, after which encryption is triggered for transactions that are not tracked on the Citrix ADC because their length is not known. There can be a delay of up to 10ms from the specified timeout value before the packet is pushed into the queue.
	*/
	Ssltriggertimeout *int `json:"ssltriggertimeout,omitempty"`
	/**
	* Send an SSL Close-Notify message to the client at the end of a transaction.
	*/
	Sendclosenotify string `json:"sendclosenotify,omitempty"`
	/**
	* Maximum number of queued packets after which encryption is triggered. Use this setting for SSL transactions that send small packets from server to Citrix ADC.
	*/
	Encrypttriggerpktcount *int `json:"encrypttriggerpktcount,omitempty"`
	/**
	* Deny renegotiation in specified circumstances. Available settings function as follows:
		* NO - Allow SSL renegotiation.
		* FRONTEND_CLIENT - Deny secure and nonsecure SSL renegotiation initiated by the client.
		* FRONTEND_CLIENTSERVER - Deny secure and nonsecure SSL renegotiation initiated by the client or the Citrix ADC during policy-based client authentication.
		* ALL - Deny all secure and nonsecure SSL renegotiation.
		* NONSECURE - Deny nonsecure SSL renegotiation. Allows only clients that support RFC 5746.
	*/
	Denysslreneg string `json:"denysslreneg,omitempty"`
	/**
	* Encoding method used to insert the subject or issuer's name in HTTP requests to servers.
	*/
	Insertionencoding string `json:"insertionencoding,omitempty"`
	/**
	* Size, per packet engine, in megabytes, of the OCSP cache. A maximum of 10% of the packet engine memory can be assigned. Because the maximum allowed packet engine memory is 4GB, the maximum value that can be assigned to the OCSP cache is approximately 410 MB.
	*/
	Ocspcachesize *int `json:"ocspcachesize,omitempty"`
	/**
	* Insert PUSH flag into decrypted, encrypted, or all records. If the PUSH flag is set to a value other than 0, the buffered records are forwarded on the basis of the value of the PUSH flag. Available settings function as follows:
		0 - Auto (PUSH flag is not set.)
		1 - Insert PUSH flag into every decrypted record.
		2 -Insert PUSH flag into every encrypted record.
		3 - Insert PUSH flag into every decrypted and encrypted record.
	*/
	Pushflag *int `json:"pushflag,omitempty"`
	/**
	* Host header check for SNI enabled sessions. If this check is enabled and the HTTP request does not contain the host header for SNI enabled sessions(i.e vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension), the request is dropped.
	*/
	Dropreqwithnohostheader string `json:"dropreqwithnohostheader,omitempty"`
	/**
	* Controls how the HTTP 'Host' header value is validated. These checks are performed only if the session is SNI enabled (i.e when vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension) and HTTP request contains 'Host' header.
		Available settings function as follows:
		CERT   - Request is forwarded if the 'Host' value is covered
		by the certificate used to establish this SSL session.
		Note: 'CERT' matching mode cannot be applied in
		TLS 1.3 connections established by resuming from a
		previous TLS 1.3 session. On these connections, 'STRICT'
		matching mode will be used instead.
		STRICT - Request is forwarded only if value of 'Host' header
		in HTTP is identical to the 'Server name' value passed
		in 'Client Hello' of the SSL connection.
		NO     - No validation is performed on the HTTP 'Host'
		header value.
	*/
	Snihttphostmatch string `json:"snihttphostmatch,omitempty"`
	/**
	* PUSH encryption trigger timeout value. The timeout value is applied only if you set the Push Encryption Trigger parameter to Timer in the SSL virtual server settings.
	*/
	Pushenctriggertimeout *int `json:"pushenctriggertimeout,omitempty"`
	/**
	* Limit to the number of disabled SSL chips after which the ADC restarts. A value of zero implies that the ADC does not automatically restart.
	*/
	Cryptodevdisablelimit *int `json:"cryptodevdisablelimit,omitempty"`
	/**
	* Name of the undefined built-in control action: CLIENTAUTH, NOCLIENTAUTH, NOOP, RESET, or DROP.
	*/
	Undefactioncontrol string `json:"undefactioncontrol,omitempty"`
	/**
	* Name of the undefined built-in data action: NOOP, RESET or DROP.
	*/
	Undefactiondata string `json:"undefactiondata,omitempty"`
	/**
	* Global parameter used to enable default profile feature.
	*/
	Defaultprofile string `json:"defaultprofile,omitempty"`
	/**
	* Citrix ADC CPU utilization threshold (in percentage) beyond which crypto operations are not done in software.
		A value of zero implies that CPU is not utilized for doing crypto in software.
	*/
	Softwarecryptothreshold *int `json:"softwarecryptothreshold,omitempty"`
	/**
	* When this mode is enabled, system will use additional crypto hardware to accelerate symmetric crypto operations.
	*/
	Hybridfipsmode string `json:"hybridfipsmode,omitempty"`
	/**
	* Signature Digest Algorithms that are supported by appliance. Default value is "ALL" and it will enable the following algorithms depending on the platform.
		On VPX: ECDSA-SHA1 ECDSA-SHA224 ECDSA-SHA256 ECDSA-SHA384 ECDSA-SHA512 RSA-SHA1 RSA-SHA224 RSA-SHA256 RSA-SHA384 RSA-SHA512 DSA-SHA1 DSA-SHA224 DSA-SHA256 DSA-SHA384 DSA-SHA512
		On MPX with Nitrox-III and coleto cards: RSA-SHA1 RSA-SHA224 RSA-SHA256 RSA-SHA384 RSA-SHA512 ECDSA-SHA1 ECDSA-SHA224 ECDSA-SHA256 ECDSA-SHA384 ECDSA-SHA512
		Others: RSA-SHA1 RSA-SHA224 RSA-SHA256 RSA-SHA384 RSA-SHA512.
		Note:ALL doesnot include RSA-MD5 for any platform.
	*/
	Sigdigesttype []string `json:"sigdigesttype,omitempty"`
	/**
	* Enable or disable dynamically learning and caching the learned information to make the subsequent interception or bypass decision. When enabled, NS does the lookup of this cached data to do early bypass.
	*/
	Sslierrorcache string `json:"sslierrorcache,omitempty"`
	/**
	* Specify the maximum memory that can be used for caching the learned data. This memory is used as a LRU cache so that the old entries gets replaced with new entry once the set memory limit is fully utilised. A value of 0 decides the limit automatically.
	*/
	Sslimaxerrorcachemem *int `json:"sslimaxerrorcachemem,omitempty"`
	/**
	* To insert space between lines in the certificate header of request
	*/
	Insertcertspace string `json:"insertcertspace,omitempty"`
	/**
	* Determines whether or not additional checks are carried out during a TLS handshake when validating an X.509 certificate received from the peer.
		Settings apply as follows:
		YES - (1) During certificate verification, ignore the
		Common Name field (inside the subject name) if
		Subject Alternative Name X.509 extension is present
		in the certificate for backend connection.
		(2) Verify the Extended Key Usage X.509 extension
		server/client leaf certificate received over the wire
		is consistent with the peer's role.
		(applicable for frontend and backend connections)
		(3) Verify the Basic Constraint CA field set to TRUE
		for non-leaf certificates. (applicable for frontend,
		backend connections and CAs bound to the Citrix ADC.
		NO  - (1) Verify the Common Name field (inside the subject name)
		irrespective of Subject Alternative Name X.509
		extension.
		(2) Ignore the Extended Key Usage X.509 extension
		for server/client leaf certificate.
		(3) Do not verify the Basic Constraint CA true flag
		for non-leaf certificates.
	*/
	Ndcppcompliancecertcheck string `json:"ndcppcompliancecertcheck,omitempty"`
	/**
	* To support both cavium and coleto based platforms in cluster environment, this mode has to be enabled.
	*/
	Heterogeneoussslhw string `json:"heterogeneoussslhw,omitempty"`
	/**
	* Limit in percentage of capacity of the crypto operations queue beyond which new SSL connections are not accepted until the queue is reduced.
	*/
	Operationqueuelimit *int `json:"operationqueuelimit,omitempty"`

	//------- Read only Parameter ---------;

	Svctls1112disable string `json:"svctls1112disable,omitempty"`
	Montls1112disable string `json:"montls1112disable,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
