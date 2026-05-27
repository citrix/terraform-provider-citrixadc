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

package api

/**
* Configuration for API specification resource.
*/
type Apispec struct {
	/**
	* Name for the spec. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the spec is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my spec" or 'my spec').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of and, optionally, path to the api spec file. The spec file should be present on the appliance's hard-disk drive or solid-state drive. Storing a spec file in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/apispec/ is the default path.
	*/
	File string `json:"file,omitempty"`
	/**
	* Input format of the spec file. The three formats supported by the appliance are:
		PROTO 
		OAS/Swagger
		GRAPHQL
	*/
	Type string `json:"type,omitempty"`
	/**
	* Disabling openapi spec validation while adding it
	*/
	Skipvalidation string `json:"skipvalidation,omitempty"`
	/**
	* Specify the encrypted API spec. Must be in NetScaler format
	*/
	Encrypted bool `json:"encrypted,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Ready string `json:"ready,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
