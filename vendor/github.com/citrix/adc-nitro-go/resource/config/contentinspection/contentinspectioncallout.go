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

package contentinspection

/**
* Configuration for Content Inspection callout resource.
*/
type Contentinspectioncallout struct {
	/**
	* Name for the Content Inspection callout. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, stringmap, or callout.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Type of the Content Inspection callout. It must be one of the following:
		* ICAP - Sends ICAP request to the configured ICAP server.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Name of the Content Inspection profile. The type of the configured profile must match the type specified using -type argument.
	*/
	Profilename string `json:"profilename,omitempty"`
	/**
	* Name of the load balancing or content switching virtual server or service to which the Content Inspection request is issued. Mutually exclusive with server IP address and port parameters. The service type must be TCP or SSL_TCP. If there are vservers and services with the same name, then vserver is selected.
	*/
	Servername string `json:"servername,omitempty"`
	/**
	* IP address of Content Inspection server. Mutually exclusive with the server name parameter.
	*/
	Serverip string `json:"serverip,omitempty"`
	/**
	* Port of the Content Inspection server.
	*/
	Serverport *int `json:"serverport,omitempty"`
	/**
	* Type of data that the target callout agent returns in response to the callout.
		Available settings function as follows:
		* TEXT - Treat the returned value as a text string.
		* NUM - Treat the returned value as a number.
		* BOOL - Treat the returned value as a Boolean value.
		Note: You cannot change the return type after it is set.
	*/
	Returntype string `json:"returntype,omitempty"`
	/**
	* Expression that extracts the callout results from the response sent by the CI callout agent. Must be a response based expression, that is, it must begin with ICAP.RES. The operations in this expression must match the return type. For example, if you configure a return type of TEXT, the result expression must be a text based expression, as in the following example: icap.res.header("ISTag")
	*/
	Resultexpr string `json:"resultexpr,omitempty"`
	/**
	* Any comments to preserve information about this Content Inspection callout.
	*/
	Comment string `json:"comment,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Undefreason string `json:"undefreason,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
