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

package appfw

/**
* Configuration for archive resource.
*/
type Appfwarchive struct {
	/**
	* Name of tar archive
	*/
	Name string `json:"name,omitempty"`
	/**
	* Path to the file to be exported
	*/
	Target string `json:"target,omitempty"`
	/**
	* Indicates the source of the tar archive file as a URL
		of the form
		<protocol>://<host>[:<port>][/<path>]
		<protocol> is http or https.
		<host> is the DNS name or IP address of the http or https server.
		<port> is the port number of the server. If omitted, the
		default port for http or https will be used.
		<path> is the path of the file on the server.
		Import will fail if an https server requires client
		certificate authentication.
		
	*/
	Src string `json:"src,omitempty"`
	/**
	* Comments associated with this archive.
	*/
	Comment string `json:"comment,omitempty"`

	//------- Read only Parameter ---------;

	Response string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
