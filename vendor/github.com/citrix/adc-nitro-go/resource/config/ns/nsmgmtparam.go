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
* Configuration for Management related configuration resource.
*/
type Nsmgmtparam struct {
	/**
	* This allow the configuration of management HTTP port.
	*/
	Mgmthttpport int `json:"mgmthttpport,omitempty"`
	/**
	* This allows the configuration of management HTTPS port.
	*/
	Mgmthttpsport int `json:"mgmthttpsport,omitempty"`
	/**
	* This enables setting the HTTPD Max Clients value in the httpd.conf file. You can configure either Max Clients or Max Request Workers. The allowable range is from a minimum of 1 to a maximum of 255
	*/
	Httpdmaxclients int `json:"httpdmaxclients,omitempty"`
	/**
	* This enables setting the HTTPD Max Request Workers value in the httpd.conf file. You can configure either Max Clients or Max Request Workers. The allowable range is from a minimum of 1 to a maximum of 255
	*/
	Httpdmaxreqworkers int `json:"httpdmaxreqworkers,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
