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

package policy

/**
* Configuration for map policy resource.
*/
type Policymap struct {
	/**
	* Name for the map policy. Must begin with a letter, number, or the underscore (_) character and must consist only of letters, numbers, and the hash (#), period (.), colon (:), space ( ), at (@), equals (=), hyphen (-), and underscore (_) characters.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my map" or 'my map').
	*/
	Mappolicyname string `json:"mappolicyname,omitempty"`
	/**
	* Publicly known source domain name. This is the domain name with which a client request arrives at a reverse proxy virtual server for cache redirection. If you specify a source domain, you must specify a target domain.
	*/
	Sd string `json:"sd,omitempty"`
	/**
	* Source URL. Specify all or part of the source URL, in the following format: /[[prefix] [*]] [.suffix].
	*/
	Su string `json:"su,omitempty"`
	/**
	* Target domain name sent to the server. The source domain name is replaced with this domain name.
	*/
	Td string `json:"td,omitempty"`
	/**
	* Target URL. Specify the target URL in the following format: /[[prefix] [*]][.suffix].
	*/
	Tu string `json:"tu,omitempty"`

	//------- Read only Parameter ---------;

	Targetname string `json:"targetname,omitempty"`

}
