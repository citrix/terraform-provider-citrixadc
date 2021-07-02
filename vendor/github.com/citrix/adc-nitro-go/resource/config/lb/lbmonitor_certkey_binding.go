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

package lb

/**
* Binding class showing the certkey that can be bound to lbmonitor.
*/
type Lbmonitorcertkeybinding struct {
	/**
	* The name of the certificate bound to the monitor.
	*/
	Certkeyname string `json:"certkeyname,omitempty"`
	/**
	* The rule for use of CRL corresponding to this CA certificate during client authentication. If crlCheck is set to Mandatory, the system will deny all SSL clients if the CRL is missing, expired - NextUpdate date is in the past, or is incomplete with remote CRL refresh enabled. If crlCheck is set to optional, the system will allow SSL clients in the above error cases.However, in any case if the client certificate is revoked in the CRL, the SSL client will be denied access.
	*/
	Ca bool `json:"ca,omitempty"`
	/**
	* The state of the CRL check parameter. (Mandatory/Optional)
	*/
	Crlcheck string `json:"crlcheck,omitempty"`
	/**
	* The state of the OCSP check parameter. (Mandatory/Optional)
	*/
	Ocspcheck string `json:"ocspcheck,omitempty"`
	/**
	* Name of the monitor.
	*/
	Monitorname string `json:"monitorname,omitempty"`


}