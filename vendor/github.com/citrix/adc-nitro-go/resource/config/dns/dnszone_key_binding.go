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

package dns

/**
* Binding class showing the key that can be bound to dnszone.
*/
type Dnszonekeybinding struct {
	/**
	* Name of the public/private DNS key pair with which to sign the zone. You can sign a zone with up to four keys.
	*/
	Keyname []string `json:"keyname,omitempty"`
	/**
	* The time when sign was done with this key.
	*/
	Siginceptiontime []uint32 `json:"siginceptiontime,omitempty"`
	/**
	* Integer which denote status of keys.
	*/
	Signed uint32 `json:"signed,omitempty"`
	/**
	* Time period for which to consider the key valid, after the key is used to sign a zone.
	*/
	Expires uint32 `json:"expires,omitempty"`
	/**
	* Name of the zone. Mutually exclusive with the type parameter.
	*/
	Zonename string `json:"zonename,omitempty"`


}