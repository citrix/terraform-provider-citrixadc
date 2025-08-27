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

package ns

/**
* Configuration for HMAC key resource.
*/
type Nshmackey struct {
	/**
	* Key name.  This follows the same syntax rules as other expression entity names:
		It must begin with an alpha character (A-Z or a-z) or an underscore (_).
		The rest of the characters must be alpha, numeric (0-9) or underscores.
		It cannot be re or xp (reserved for regular and XPath expressions).
		It cannot be an expression reserved word (e.g. SYS or HTTP).
		It cannot be used for an existing expression object (HTTP callout, patset, dataset, stringmap, or named expression).
	*/
	Name string `json:"name,omitempty"`
	/**
	* Digest (hash) function to be used in the HMAC computation.
	*/
	Digest string `json:"digest,omitempty"`
	/**
	* The hex-encoded key to be used in the HMAC computation. The key can be any length (up to a Citrix ADC-imposed maximum of 255 bytes). If the length is less than the digest block size, it will be zero padded up to the block size. If it is greater than the block size, it will be hashed using the digest function to the block size. The block size for each digest is:
		MD2    - 16 bytes
		MD4    - 16 bytes
		MD5    - 16 bytes
		SHA1   - 20 bytes
		SHA224 - 28 bytes
		SHA256 - 32 bytes
		SHA384 - 48 bytes
		SHA512 - 64 bytes
		Note that the key will be encrypted when it it is saved
		There is a special key value AUTO which generates a new random key for the specified digest. This kind of key is
		intended for use cases where the NetScaler both generates and verifies an HMAC on  the same data.
	*/
	Keyvalue string `json:"keyvalue,omitempty"`
	/**
	* Comments associated with this encryption key.
	*/
	Comment string `json:"comment,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
