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
* Configuration for dh Parameter resource.
*/
type Ssldhparam struct {
	/**
	* Name of and, optionally, path to the DH key file. /nsconfig/ssl/ is the default path.
	*/
	Dhfile string `json:"dhfile,omitempty"`
	/**
	* Size, in bits, of the DH key being generated.
	*/
	Bits int `json:"bits,omitempty"`
	/**
	* Random number required for generating the DH key. Required as part of the DH key generation algorithm.
	*/
	Gen string `json:"gen,omitempty"`

}
