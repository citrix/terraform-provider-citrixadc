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

package cmp

/**
* Configuration for compression action resource.
*/
type Cmpaction struct {
	/**
	* Name of the compression action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the action is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cmp action" or 'my cmp action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Type of compression performed by this action.
		Available settings function as follows:
		* COMPRESS - Apply GZIP or DEFLATE compression to the response, depending on the request header. Prefer GZIP.
		* GZIP - Apply GZIP compression.
		* DEFLATE - Apply DEFLATE compression.
		* NOCOMPRESS - Do not compress the response if the request matches a policy that uses this action.
	*/
	Cmptype string `json:"cmptype,omitempty"`
	/**
	* Control insertion of the Vary header in HTTP responses compressed by Citrix ADC. Intermediate caches store different versions of the response for different values of the headers present in the Vary response header.
	*/
	Addvaryheader string `json:"addvaryheader,omitempty"`
	/**
	* The value of the HTTP Vary header for compressed responses.
	*/
	Varyheadervalue string `json:"varyheadervalue,omitempty"`
	/**
	* The type of delta action (if delta type compression action is defined).
	*/
	Deltatype string `json:"deltatype,omitempty"`
	/**
	* New name for the compression action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at
		(@), equals (=), and hyphen (-) characters.
		Choose a name that can be correlated with the function that the action performs.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cmp action" or 'my cmp action').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Isdefault string `json:"isdefault,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
