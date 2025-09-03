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

package policy

/**
* Configuration for PAT set resource.
*/
type Policypatset struct {
	/**
	* Unique name of the pattern set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters. Must not be the name of an existing named expression, pattern set, dataset, string map, or HTTP callout.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Any comments to preserve information about this patset or a pattern bound to this patset.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* File which contains list of patterns that needs to be bound to the patset. A patsetfile cannot be associated with multiple patsets.
	*/
	Patsetfile string `json:"patsetfile,omitempty"`
	/**
	* This is used to populate internal patset information so that the patset can also be used dynamically in an expression. Here dynamically means the patset name can also be derived using an expression. For example for a given patset name "allow_test" it can be used dynamically as http.req.url.contains_any("allow_" + http.req.url.path.get(1)). This cannot be used with default patsets.
	*/
	Dynamic string `json:"dynamic,omitempty"`
	/**
	* Shows only dynamic patsets when set true.
	*/
	Dynamiconly bool `json:"dynamiconly,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
