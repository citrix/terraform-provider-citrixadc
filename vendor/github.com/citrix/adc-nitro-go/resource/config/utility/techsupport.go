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
* Configuration for tech support resource.
*/
type Techsupport struct {
	/**
	* Use this option to gather data on the present node, all cluster nodes, or for the specified partitions. The CLUSTER scope generates smaller abbreviated archives for all nodes. The PARTITION scope collects the admin partition in addition to those specified. The partitionName option is only required for the PARTITION scope.
	*/
	Scope string `json:"scope,omitempty"`
	/**
	* Name of the partition
	*/
	Partitionname string `json:"partitionname,omitempty"`
	/**
	* Securely upload the collector archive to Citrix Technical Support using SSL. MyCitrix credentials will be required. If used with the -file option, no new collector archive is generated. Instead, the specified archive is uploaded. Note that the upload operation time depends on the size of the archive file, and the connection bandwidth.
	*/
	Upload bool `json:"upload,omitempty"`
	/**
	* Specifies the proxy server to be used when uploading a collector archive. Use this parameter if the Citrix ADC does not have direct internet connectivity. The basic format of the proxy string is: "proxy_IP:<proxy_port>" (without quotes). If the proxy requires authentication the format is: "username:password@proxy_IP:<proxy_port>"
	*/
	Proxy string `json:"proxy,omitempty"`
	/**
	* Specifies the associated case or service request number if it has already been opened with Citrix Technical Support.
	*/
	Casenumber string `json:"casenumber,omitempty"`
	/**
	* Specifies the name (with full path) of the collector archive file to be uploaded. If this is specified, no new collector archive is generated.
	*/
	File string `json:"file,omitempty"`
	/**
	* Provides a text description for the the upload, and can be used for logging purposes.
	*/
	Description string `json:"description,omitempty"`
	/**
	* Specifies My Citrix user name, which is used to login to Citrix upload server
	*/
	Username string `json:"username,omitempty"`
	/**
	* Specifies My Citrix password, which is used to login to Citrix upload server
	*/
	Password string `json:"password,omitempty"`

	//------- Read only Parameter ---------;

	Response string `json:"response,omitempty"`
	Servername string `json:"servername,omitempty"`

}
