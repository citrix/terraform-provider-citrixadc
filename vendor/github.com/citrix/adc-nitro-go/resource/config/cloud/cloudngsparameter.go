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

package cloud

/**
* Configuration for cloud ngsparameter resource.
*/
type Cloudngsparameter struct {
	/**
	* Enables blocking connections authenticated with a ticket createdby by an entity not whitelisted in allowedngstktprofile
	*/
	Blockonallowedngstktprof string `json:"blockonallowedngstktprof,omitempty"`
	/**
	* Enables the required UDT version to EDT connections in the CGS deployment
	*/
	Allowedudtversion string `json:"allowedudtversion,omitempty"`
	/**
	* Enables Decoupling CSVSERVER state from Ticketing Service state in the CGS deployment
	*/
	Csvserverticketingdecouple string `json:"csvserverticketingdecouple,omitempty"`
	/**
	* Enables DTLS1.2 for client connections on CGS
	*/
	Allowdtls12 string `json:"allowdtls12,omitempty"`
	/**
	* Configure WAF protection for CGS deployment. The available options are:
		* EDGE - Edge APIs are WAF protected.
		* DEFAULT - NetScaler decides which endpoints have WAF protection enabled for CGS deployment.
		* DISABLED - WAF protection is disabled.
	*/
	Wafprotection string `json:"wafprotection,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
