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

package responder

/**
* Configuration for Responder HTML page resource.
*/
type Responderhtmlpage struct {
	/**
	* Local path to and name of, or URL \(protocol, host, path, and file name\) for, the file in which to store the imported HTML page.
		NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.
	*/
	Src string `json:"src,omitempty"`
	/**
	* Name to assign to the HTML page object on the Citrix ADC.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Any comments to preserve information about the HTML page object.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Overwrites the existing file
	*/
	Overwrite bool `json:"overwrite,omitempty"`
	/**
	* CA certificate file name which will be used to verify the peer's certificate. The certificate should be imported using "import ssl certfile" CLI command or equivalent in API or GUI. If certificate name is not configured, then default root CA certificates are used for peer's certificate verification.
	*/
	Cacertfile string `json:"cacertfile,omitempty"`

	//------- Read only Parameter ---------;

	Response string `json:"response,omitempty"`

}
