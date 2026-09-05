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
* Configuration for Zerotouch params resource.
*/
type Sslzerotouchparam struct {
	/**
	* Timeout(in minutes) for caching the OCSP response.
	*/
	Ocspcachetimeout *int `json:"ocspcachetimeout,omitempty"`
	/**
	* Number of certificates to batch together into one OCSP request. Batching avoids overloading the OCSP responder. A value of 1 signifies that each request is queried independently. For a value greater than 1, specify a timeout (batching delay) to avoid inordinately delaying the processing of a single certificate.
	*/
	Ocspbatchingdepth *int `json:"ocspbatchingdepth,omitempty"`
	/**
	* Maximum time, in milliseconds, to wait to accumulate OCSP requests to batch. Does not apply if the Batching Depth is 1.
	*/
	Ocspbatchingdelay *int `json:"ocspbatchingdelay,omitempty"`
	/**
	* Time, in milliseconds, to wait for an OCSP response. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server. Includes Batching Delay time.
	*/
	Ocspresptimeout *int `json:"ocspresptimeout,omitempty"`
	/**
	* Time, in milliseconds, to wait for an OCSP URL Resolution. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server.
	*/
	Ocspurlresolvetimeout *int `json:"ocspurlresolvetimeout,omitempty"`
	/**
	* If trustResponder is set to YES, signature verification will be skipped for the OCSP response
	*/
	Ocsptrustresponder string `json:"ocsptrustresponder,omitempty"`
	/**
	* Time, in seconds, for which the Citrix ADC waits before considering the response as invalid. The response is considered invalid if the Produced At time stamp in the OCSP response exceeds or precedes the current Citrix ADC clock time by the amount of time specified.
	*/
	Ocspproducedattimeskew *int `json:"ocspproducedattimeskew,omitempty"`
	/**
	* Enable the OCSP nonce extension, which is designed to prevent replay attacks.
	*/
	Ocspusenonce string `json:"ocspusenonce,omitempty"`
	/**
	* HTTP method used to send ocsp request. POST is the default httpmethod. If request length is > 255, POST wil be used even if GET is set as httpMethod
	*/
	Ocsphttpmethod string `json:"ocsphttpmethod,omitempty"`

	//------- Read only Parameter ---------;

	Zerotouch string `json:"zerotouch,omitempty"`
	Remoteserverip string `json:"remoteserverip,omitempty"`
	Keyfilename string `json:"keyfilename,omitempty"`
	Passphrase string `json:"passphrase,omitempty"`
	Admconnectivitystatus string `json:"admconnectivitystatus,omitempty"`
	Httpstatuscode string `json:"httpstatuscode,omitempty"`
	Requesttype string `json:"requesttype,omitempty"`
	Requesttimestamp string `json:"requesttimestamp,omitempty"`
	Nextrequesttime string `json:"nextrequesttime,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
