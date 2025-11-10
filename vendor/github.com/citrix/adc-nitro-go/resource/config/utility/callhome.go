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

package utility

/**
* Configuration for callhome resource.
*/
type Callhome struct {
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid *int `json:"nodeid,omitempty"`
	/**
	* CallHome mode of operation
	*/
	Mode string `json:"mode,omitempty"`
	/**
	* Email address of the contact administrator.
	*/
	Emailaddress string `json:"emailaddress,omitempty"`
	/**
	* Interval (in days) between CallHome heartbeats
	*/
	Hbcustominterval *int `json:"hbcustominterval,omitempty"`
	/**
	* Enables or disables the proxy mode. The proxy server can be set by either specifying the IP address of the server or the name of the service representing the proxy server.
	*/
	Proxymode string `json:"proxymode,omitempty"`
	/**
	* IP address of the proxy server.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Name of the service that represents the proxy server.
	*/
	Proxyauthservice string `json:"proxyauthservice,omitempty"`
	/**
	* HTTP port on the Proxy server. This is a mandatory parameter for both IP address and service name based configuration.
	*/
	Port *int `json:"port,omitempty"`

	//------- Read only Parameter ---------;

	Sslcardfirstfailure string `json:"sslcardfirstfailure,omitempty"`
	Sslcardlatestfailure string `json:"sslcardlatestfailure,omitempty"`
	Powfirstfail string `json:"powfirstfail,omitempty"`
	Powlatestfailure string `json:"powlatestfailure,omitempty"`
	Hddfirstfail string `json:"hddfirstfail,omitempty"`
	Hddlatestfailure string `json:"hddlatestfailure,omitempty"`
	Flashfirstfail string `json:"flashfirstfail,omitempty"`
	Flashlatestfailure string `json:"flashlatestfailure,omitempty"`
	Rlfirsthighdrop string `json:"rlfirsthighdrop,omitempty"`
	Rllatesthighdrop string `json:"rllatesthighdrop,omitempty"`
	Restartlatestfail string `json:"restartlatestfail,omitempty"`
	Memthrefirstanomaly string `json:"memthrefirstanomaly,omitempty"`
	Memthrelatestanomaly string `json:"memthrelatestanomaly,omitempty"`
	Callhomestatus string `json:"callhomestatus,omitempty"`
	Anomalydetection string `json:"anomalydetection,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
