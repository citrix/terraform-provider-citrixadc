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

package ha

/**
* Configuration for files resource.
*/
type Hafiles struct {
	/**
	* Specify one of the following modes of synchronization.
		* all - Synchronize files related to system configuration, Access Gateway bookmarks, SSL certificates, SSL CRL lists,  and Application Firewall XML objects.
		* bookmarks - Synchronize all Access Gateway bookmarks.
		* ssl - Synchronize all certificates, keys, and CRLs for the SSL feature.
		* imports. Synchronize all XML objects (for example, WSDLs, schemas, error pages) configured for the application firewall.
		* misc - Synchronize all license files and the rc.conf file.
		* all_plus_misc - Synchronize files related to system configuration, Access Gateway bookmarks, SSL certificates, SSL CRL lists, application firewall XML objects, licenses, and the rc.conf file.
	*/
	Mode []string `json:"mode,omitempty"`

}
