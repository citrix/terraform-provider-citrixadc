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
* Configuration for FIPS key resource.
*/
type Sslfipskey struct {
	/**
	* Name for the FIPS key. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the FIPS key is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my fipskey" or 'my fipskey').
	*/
	Fipskeyname string `json:"fipskeyname,omitempty"`
	/**
	* Only RSA key and ECDSA Key are supported.
	*/
	Keytype string `json:"keytype,omitempty"`
	/**
	* Exponent value for the FIPS key to be created. Available values function as follows:
		3=3 (hexadecimal)
		F4=10001 (hexadecimal)
	*/
	Exponent string `json:"exponent,omitempty"`
	/**
	* Modulus, in multiples of 64, of the FIPS key to be created.
	*/
	Modulus int `json:"modulus,omitempty"`
	/**
	* Only p_256 (prime256v1) and P_384 (secp384r1) are supported.
	*/
	Curve string `json:"curve,omitempty"`
	/**
	* Name of and, optionally, path to the key file to be imported.
		/nsconfig/ssl/ is the default path.
	*/
	Key string `json:"key,omitempty"`
	/**
	* Input format of the key file. Available formats are:
		SIM - Secure Information Management; select when importing a FIPS key. If the external FIPS key is encrypted, first decrypt it, and then import it.
		PEM - Privacy Enhanced Mail; select when importing a non-FIPS key.
	*/
	Inform string `json:"inform,omitempty"`
	/**
	* Name of the wrap key to use for importing the key. Required for importing a non-FIPS key.
	*/
	Wrapkeyname string `json:"wrapkeyname,omitempty"`
	/**
	* Initialization Vector (IV) to use for importing the key. Required for importing a non-FIPS key.
	*/
	Iv string `json:"iv,omitempty"`

	//------- Read only Parameter ---------;

	Size string `json:"size,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
