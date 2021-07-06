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

package cmp

/**
* Configuration for CMP parameter resource.
*/
type Cmpparameter struct {
	/**
	* Specify a compression level. Available settings function as follows:
		* Optimal - Corresponds to a gzip GZIP level of 5-7.
		* Best speed - Corresponds to a gzip level of 1.
		* Best compression - Corresponds to a gzip level of 9.
	*/
	Cmplevel string `json:"cmplevel,omitempty"`
	/**
	* Minimum quantum of data to be filled before compression begins.
	*/
	Quantumsize int `json:"quantumsize,omitempty"`
	/**
	* Allow the server to send compressed data to the Citrix ADC. With the default setting, the Citrix ADC appliance handles all compression.
	*/
	Servercmp string `json:"servercmp,omitempty"`
	/**
	* Heuristic basefile expiry.
	*/
	Heurexpiry string `json:"heurexpiry,omitempty"`
	/**
	* Threshold compression ratio for heuristic basefile expiry, multiplied by 100. For example, to set the threshold ratio to 1.25, specify 125.
	*/
	Heurexpirythres int `json:"heurexpirythres,omitempty"`
	/**
	* For heuristic basefile expiry, weightage to be given to historical delta compression ratio, specified as percentage.  For example, to give 25% weightage to historical ratio (and therefore 75% weightage to the ratio for current delta compression transaction), specify 25.
	*/
	Heurexpiryhistwt int `json:"heurexpiryhistwt,omitempty"`
	/**
	* Smallest response size, in bytes, to be compressed.
	*/
	Minressize int `json:"minressize,omitempty"`
	/**
	* Citrix ADC CPU threshold after which compression is not performed. Range: 0 - 100
	*/
	Cmpbypasspct int `json:"cmpbypasspct,omitempty"`
	/**
	* Citrix ADC does not wait for the quantum to be filled before starting to compress data. Upon receipt of a packet with a PUSH flag, the appliance immediately begins compression of the accumulated packets.
	*/
	Cmponpush string `json:"cmponpush,omitempty"`
	/**
	* Type of policy. Available settings function as follows:
		* Classic -  Classic policies evaluate basic characteristics of traffic and other data. Deprecated.
		* Advanced -  Advanced policies (which have been renamed as default syntax policies) can perform the same type of evaluations as classic policies. They also enable you to analyze more data (for example, the body of an HTTP request) and to configure more operations in the policy rule (for example, transforming data in the body of a request into an HTTP header).
	*/
	Policytype string `json:"policytype,omitempty"`
	/**
	* Control insertion of the Vary header in HTTP responses compressed by Citrix ADC. Intermediate caches store different versions of the response for different values of the headers present in the Vary response header.
	*/
	Addvaryheader string `json:"addvaryheader,omitempty"`
	/**
	* The value of the HTTP Vary header for compressed responses. If this argument is not specified, a default value of "Accept-Encoding" will be used.
	*/
	Varyheadervalue string `json:"varyheadervalue,omitempty"`
	/**
	* Enable insertion of  Cache-Control: private response directive to indicate response message is intended for a single user and must not be cached by a shared or proxy cache.
	*/
	Externalcache string `json:"externalcache,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
