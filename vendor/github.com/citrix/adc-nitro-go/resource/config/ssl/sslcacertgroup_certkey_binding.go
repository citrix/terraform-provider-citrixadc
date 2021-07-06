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
* Binding class showing the certkey that can be bound to sslcacertgroup.
*/
type Sslcacertgroupcertkeybinding struct {
	/**
	* Name for the certkey added to the Citrix ADC. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the certificate-key pair is created.The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cert" or 'my cert').
	*/
	Certkeyname string `json:"certkeyname,omitempty"`
	/**
	* The state of the CRL check parameter. (Mandatory/Optional)
	*/
	Crlcheck string `json:"crlcheck,omitempty"`
	/**
	* The state of the OCSP check parameter. (Mandatory/Optional)
	*/
	Ocspcheck string `json:"ocspcheck,omitempty"`
	/**
	* Name given to the CA certificate group. The name will be used to add the CA certificates to the group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my file" or 'my file').
	*/
	Cacertgroupname string `json:"cacertgroupname,omitempty"`


}