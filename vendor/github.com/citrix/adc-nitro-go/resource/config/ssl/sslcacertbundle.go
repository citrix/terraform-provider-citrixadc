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
* Configuration for CA certbundle resource.
*/
type Sslcacertbundle struct {
	/**
	* Name given to the CA certbundle. The name will be used for bind/unbind/update operations. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my file" or 'my file').
	*/
	Cacertbundlename string `json:"cacertbundlename,omitempty"`
	/**
	* Name of and, optionally, path to the X509 CA certificate bundle file that is used to form cacertbundle entity. The CA certificate bundle file should be present on the appliance's hard-disk drive or solid-state drive. /nsconfig/ssl/ is the default path. The CA certificate bundle file consists of list of certificates.
	*/
	Bundlefile string `json:"bundlefile,omitempty"`

	//------- Read only Parameter ---------;

	Servername string `json:"servername,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
