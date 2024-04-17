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
* Configuration for application firewall signatures XML configuration resource.
*/
type Appfwsignatures struct {
	/**
	* Name of the signature object.
	*/
	Name string `json:"name,omitempty"`
	/**
	* URL (protocol, host, path, and file name) for the location at which to store the imported signatures object.
		NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.
	*/
	Src string `json:"src,omitempty"`
	/**
	* XSLT file source.
	*/
	Xslt string `json:"xslt,omitempty"`
	/**
	* Any comments to preserve information about the signatures object.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Overwrite any existing signatures object of the same name.
	*/
	Overwrite bool `json:"overwrite,omitempty"`
	/**
	* Merges the existing Signature with new signature rules
	*/
	Merge bool `json:"merge,omitempty"`
	/**
	* preserves def actions of signature rules
	*/
	Preservedefactions bool `json:"preservedefactions,omitempty"`
	/**
	* File path for sha1 file to validate signature file
	*/
	Sha1 string `json:"sha1,omitempty"`
	/**
	* Third party vendor type for which WAF signatures has to be generated.
	*/
	Vendortype string `json:"vendortype,omitempty"`
	/**
	* Merges signature file with default signature file.
	*/
	Mergedefault bool `json:"mergedefault,omitempty"`
	/**
	* Flag used to enable/disable auto enable new signatures.
	*/
	Autoenablenewsignatures string `json:"autoenablenewsignatures,omitempty"`
	/**
	* Signature rule IDs to be Enabled/Disabled.
	*/
	Ruleid []int `json:"ruleid,omitempty"`
	/**
	* Signature category to be Enabled/Disabled.
	*/
	Category string `json:"category,omitempty"`
	/**
	* Flag used to enable/disable enable signature rule IDs/Signature Category.
	*/
	Enabled string `json:"enabled,omitempty"`
	/**
	* Signature action.
	*/
	Action []string `json:"action,omitempty"`

	//------- Read only Parameter ---------;

	Response string `json:"response,omitempty"`

}
