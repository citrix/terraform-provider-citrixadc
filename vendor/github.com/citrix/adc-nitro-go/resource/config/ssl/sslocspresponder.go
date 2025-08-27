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
* Configuration for OCSP responser resource.
*/
type Sslocspresponder struct {
	/**
	* Name for the OCSP responder. Cannot begin with a hash (#) or space character and must contain only ASCII alphanumeric, underscore (_), hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the responder is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my responder" or 'my responder').
	*/
	Name string `json:"name,omitempty"`
	/**
	* URL of the OCSP responder.
	*/
	Url string `json:"url,omitempty"`
	/**
	* Enable caching of responses. Caching of responses received from the OCSP responder enables faster responses to the clients and reduces the load on the OCSP responder.
	*/
	Cache string `json:"cache,omitempty"`
	/**
	* Timeout for caching the OCSP response. After the timeout, the Citrix ADC sends a fresh request to the OCSP responder for the certificate status. If a timeout is not specified, the timeout provided in the OCSP response applies.
	*/
	Cachetimeout int `json:"cachetimeout,omitempty"`
	/**
	* Number of client certificates to batch together into one OCSP request. Batching avoids overloading the OCSP responder. A value of 1 signifies that each request is queried independently. For a value greater than 1, specify a timeout (batching delay) to avoid inordinately delaying the processing of a single certificate.
	*/
	Batchingdepth int `json:"batchingdepth,omitempty"`
	/**
	* Maximum time, in milliseconds, to wait to accumulate OCSP requests to batch.  Does not apply if the Batching Depth is 1.
	*/
	Batchingdelay int `json:"batchingdelay,omitempty"`
	/**
	* Time, in milliseconds, to wait for an OCSP response. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server. Includes Batching Delay time.
	*/
	Resptimeout int `json:"resptimeout,omitempty"`
	/**
	* Time, in milliseconds, to wait for an OCSP URL Resolution. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server.
	*/
	Ocspurlresolvetimeout int `json:"ocspurlresolvetimeout,omitempty"`
	Respondercert string `json:"respondercert,omitempty"`
	/**
	* A certificate to use to validate OCSP responses.  Alternatively, if -trustResponder is specified, no verification will be done on the reponse.  If both are omitted, only the response times (producedAt, lastUpdate, nextUpdate) will be verified.
	*/
	Trustresponder bool `json:"trustresponder,omitempty"`
	/**
	* Time, in seconds, for which the Citrix ADC waits before considering the response as invalid. The response is considered invalid if the Produced At time stamp in the OCSP response exceeds or precedes the current Citrix ADC clock time by the amount of time specified.
	*/
	Producedattimeskew int `json:"producedattimeskew,omitempty"`
	/**
	* Certificate-key pair that is used to sign OCSP requests. If this parameter is not set, the requests are not signed.
	*/
	Signingcert string `json:"signingcert,omitempty"`
	/**
	* Enable the OCSP nonce extension, which is designed to prevent replay attacks.
	*/
	Usenonce string `json:"usenonce,omitempty"`
	/**
	* Include the complete client certificate in the OCSP request.
	*/
	Insertclientcert string `json:"insertclientcert,omitempty"`
	/**
	* HTTP method used to send ocsp request. POST is the default httpmethod. If request length is > 255, POST wil be used even if GET is set as httpMethod
	*/
	Httpmethod string `json:"httpmethod,omitempty"`

	//------- Read only Parameter ---------;

	Ocspaiarefcount string `json:"ocspaiarefcount,omitempty"`
	Ocspipaddrstr string `json:"ocspipaddrstr,omitempty"`
	Port string `json:"port,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
