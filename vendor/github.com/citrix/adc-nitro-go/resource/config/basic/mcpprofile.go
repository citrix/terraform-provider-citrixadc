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
* Configuration for mcpProfile resource.
*/
type Mcpprofile struct {
	/**
	* Name for the mcp profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my mcp profile").
	*/
	Name string `json:"name,omitempty"`
	/**
	* Proxy mode for the MCP profile. FORWARD mode replaces Host and URL in backend requests. REVERSE mode passes requests as-is.
	*/
	Proxymode string `json:"proxymode,omitempty"`
	/**
	* Type of MCP profile. Frontend profiles apply to the entity that receives requests from a client. Backend profiles apply to the entity that sends client requests to a server.
	*/
	Profiletype string `json:"profiletype,omitempty"`
	/**
	* Whether the Host header should be replaced with the backend MCP server FQDN in FORWARD proxy mode. If mcpProxyMode is FORWARD, this parameter is ENABLED bydefault. If mcpProxyMode is REVERSE, this parameter is DISABLED and cannot be ENABLED.
	*/
	Hostreplacement string `json:"hostreplacement,omitempty"`
	/**
	* Whether the URL should be replaced with the backend MCP server URL in FORWARD proxy mode. If mcpProxyMode is FORWARD, this parameter is ENABLED bydefault. If mcpProxyMode is REVERSE, this parameter is DISABLED and cannot be ENABLED.
	*/
	Urlreplacement string `json:"urlreplacement,omitempty"`
	/**
	* MCP protocol version to advertise during monitoring of a mcp server.
	*/
	Protocolversion string `json:"protocolversion,omitempty"`
	/**
	* If you like to insert Bearer or API token, configure this parameter with full header.
	*/
	Tokenorapi string `json:"tokenorapi,omitempty"`
	/**
	* Any information about the MCP profile.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Whether mcp_token_or_api configuration will be used for MCP requests coming from client.
	*/
	Insertheaderinclientrequest string `json:"insertheaderinclientrequest,omitempty"`
	/**
	* New name for the mcpProfile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
