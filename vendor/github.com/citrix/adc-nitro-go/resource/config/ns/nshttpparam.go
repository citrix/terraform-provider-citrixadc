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
* Configuration for HTTP parameter resource.
*/
type Nshttpparam struct {
	/**
	* Drop invalid HTTP requests or responses.
	*/
	Dropinvalreqs string `json:"dropinvalreqs,omitempty"`
	/**
	* Mark HTTP/0.9 requests as invalid.
	*/
	Markhttp09inval string `json:"markhttp09inval,omitempty"`
	/**
	* Mark CONNECT requests as invalid.
	*/
	Markconnreqinval string `json:"markconnreqinval,omitempty"`
	/**
	* Enable or disable Citrix ADC server header insertion for Citrix ADC generated HTTP responses.
	*/
	Insnssrvrhdr string `json:"insnssrvrhdr,omitempty"`
	/**
	* The server header value to be inserted. If no explicit header is specified then NSBUILD.RELEASE is used as default server header.
	*/
	Nssrvrhdr string `json:"nssrvrhdr,omitempty"`
	/**
	* Server header value to be inserted.
	*/
	Logerrresp string `json:"logerrresp,omitempty"`
	/**
	* Reuse server connections for requests from more than one client connections.
	*/
	Conmultiplex string `json:"conmultiplex,omitempty"`
	/**
	* Maximum limit on the number of connections, from the Citrix ADC to a particular server that are kept in the reuse pool. This setting is helpful for optimal memory utilization and for reducing the idle connections to the server just after the peak time.
	*/
	Maxreusepool int `json:"maxreusepool"` // Zero is a valid value
	/**
	* Enable/Disable HTTP/2 on server side
	*/
	Http2serverside string `json:"http2serverside,omitempty"`
	/**
	* Ignore Coding scheme in CONNECT request.
	*/
	Ignoreconnectcodingscheme string `json:"ignoreconnectcodingscheme,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
