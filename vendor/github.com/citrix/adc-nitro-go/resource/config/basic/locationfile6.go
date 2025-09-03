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

package basic

/**
* Configuration for location file6 resource.
*/
type Locationfile6 struct {
	/**
	* Name of the IPv6 location file, with or without absolute path. If the path is not included, the default path (/var/netscaler/locdb) is assumed. In a high availability setup, the static database must be stored in the same location on both NetScalers.
	*/
	Locationfile string `json:"Locationfile,omitempty"`
	/**
	* Format of the IPv6 location file. Required for the NetScaler to identify how to read the location file.
	*/
	Format string `json:"format,omitempty"`
	/**
	* URL \(protocol, host, path, and file name\) from where the location file will be imported.
		NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.
	*/
	Src string `json:"src,omitempty"`

	//------- Read only Parameter ---------;

	Curlocfilestatus string `json:"curlocfilestatus,omitempty"`
	Prevlocationfile string `json:"prevlocationfile,omitempty"`
	Prevlocfileformat string `json:"prevlocfileformat,omitempty"`
	Prevlocfilestatus string `json:"prevlocfilestatus,omitempty"`
	Locfilestatusstr string `json:"locfilestatusstr,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
