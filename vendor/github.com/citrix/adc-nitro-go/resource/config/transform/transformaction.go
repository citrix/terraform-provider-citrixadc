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

package transform

/**
* Configuration for transform action resource.
*/
type Transformaction struct {
	/**
	* Name for the URL transformation action.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the URL Transformation action is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, ^A"my transform action^A" or ^A'my transform action).
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the URL Transformation profile with which to associate this action.
	*/
	Profilename string `json:"profilename,omitempty"`
	/**
	* Positive integer specifying the priority of the action within the profile. A lower number specifies a higher priority. Must be unique within the list of actions bound to the profile. Policies are evaluated in the order of their priority numbers, and the first policy that matches is applied.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Enable or disable this action.
	*/
	State string `json:"state,omitempty"`
	/**
	* PCRE-format regular expression that describes the request URL pattern to be transformed.
	*/
	Requrlfrom string `json:"requrlfrom,omitempty"`
	/**
	* PCRE-format regular expression that describes the transformation to be performed on URLs that match the reqUrlFrom pattern.
	*/
	Requrlinto string `json:"requrlinto,omitempty"`
	/**
	* PCRE-format regular expression that describes the response URL pattern to be transformed.
	*/
	Resurlfrom string `json:"resurlfrom,omitempty"`
	/**
	* PCRE-format regular expression that describes the transformation to be performed on URLs that match the resUrlFrom pattern.
	*/
	Resurlinto string `json:"resurlinto,omitempty"`
	/**
	* Pattern that matches the domain to be transformed in Set-Cookie headers.
	*/
	Cookiedomainfrom string `json:"cookiedomainfrom,omitempty"`
	/**
	* PCRE-format regular expression that describes the transformation to be performed on cookie domains that match the cookieDomainFrom pattern. 
		NOTE: The cookie domain to be transformed is extracted from the request.
	*/
	Cookiedomaininto string `json:"cookiedomaininto,omitempty"`
	/**
	* Any comments to preserve information about this URL Transformation action.
	*/
	Comment string `json:"comment,omitempty"`

	//------- Read only Parameter ---------;

	Continuematching string `json:"continuematching,omitempty"`

}
