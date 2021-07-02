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
* Configuration for pkcs12 resource.
*/
type Sslpkcs12 struct {
	/**
	* Name for and, optionally, path to, the output file that contains the certificate and the private key after converting from PKCS#12 to PEM format. /nsconfig/ssl/ is the default path.
		If importing, the certificate-key pair is stored in PEM format. If exporting, the certificate-key pair is stored in PKCS#12 format.
	*/
	Outfile string `json:"outfile,omitempty"`
	/**
	* Convert the certificate and private-key from PKCS#12 format to PEM format.
	*/
	Import bool `json:"Import,omitempty"`
	/**
	* Name for and, optionally, path to, the PKCS#12 file. If importing, specify the input file name that contains the certificate and the private key in PKCS#12 format. If exporting, specify the output file name that contains the certificate and the private key after converting from PEM to
		PKCS#12 format. /nsconfig/ssl/ is the default path.
		During the import operation, if the key is encrypted, you are prompted to enter the pass phrase used for encrypting the key.
	*/
	Pkcs12file string `json:"pkcs12file,omitempty"`
	/**
	* Encrypt the private key by using the DES algorithm in CBC mode during the import operation. On the command line, you are prompted to enter the pass phrase.
	*/
	Des bool `json:"des,omitempty"`
	/**
	* Encrypt the private key by using the Triple-DES algorithm in EDE CBC mode (168-bit key) during the import operation. On the command line, you are prompted to enter the pass phrase.
	*/
	Des3 bool `json:"des3,omitempty"`
	/**
	* Encrypt the private key by using the AES algorithm (256-bit key) during the import operation. On the command line, you are prompted to enter the pass phrase.
	*/
	Aes256 bool `json:"aes256,omitempty"`
	/**
	* Convert the certificate and private key from PEM format to PKCS#12 format. On the command line, you are prompted to enter the pass phrase.
	*/
	Export bool `json:"export,omitempty"`
	/**
	* Certificate file to be converted from PEM to PKCS#12 format.
	*/
	Certfile string `json:"certfile,omitempty"`
	/**
	* Name of the private key file to be converted from PEM to PKCS#12 format. If the key file is encrypted, you are prompted to enter the pass phrase used for encrypting the key.
	*/
	Keyfile string `json:"keyfile,omitempty"`
	Password string `json:"password,omitempty"`
	Pempassphrase string `json:"pempassphrase,omitempty"`

}
