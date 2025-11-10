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

package filter

/**
* Configuration for filter action resource.
*/
type Filteraction struct {
	/**
	* Name for the filtering action. Must begin with a letter, number, or the underscore character (_). Other characters allowed, after the first character, are the hyphen (-), period (.) hash (#), space ( ), at sign (@), equals (=), and colon (:) characters. Choose a name that helps identify the type of action. The name of a filter action cannot be changed after it is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Qualifier, which is the action to be performed. The qualifier cannot be changed after it is set. The available options function as follows:
		ADD - Adds the specified HTTP header.
		RESET - Terminates the connection, sending the appropriate termination notice to the user's browser.
		FORWARD - Redirects the request to the designated service. You must specify either a service name or a page, but not both.
		DROP - Silently deletes the request, without sending a response to the user's browser. 
		CORRUPT - Modifies the designated HTTP header to prevent it from performing the function it was intended to perform, then sends the request/response to the server/browser.
		ERRORCODE. Returns the designated HTTP error code to the user's browser (for example, 404, the standard HTTP code for a non-existent Web page).
	*/
	Qual string `json:"qual,omitempty"`
	/**
	* Service to which to forward HTTP requests. Required if the qualifier is FORWARD.
	*/
	Servicename string `json:"servicename,omitempty"`
	/**
	* String containing the header_name and header_value. If the qualifier is ADD, specify <header_name>:<header_value>. If the qualifier is CORRUPT, specify only the header_name
	*/
	Value string `json:"value,omitempty"`
	/**
	* Response code to be returned for HTTP requests (for use with the ERRORCODE qualifier).
	*/
	Respcode *int `json:"respcode,omitempty"`
	/**
	* HTML page to return for HTTP requests (For use with the ERRORCODE qualifier).
	*/
	Page string `json:"page,omitempty"`

	//------- Read only Parameter ---------;

	Isdefault string `json:"isdefault,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
