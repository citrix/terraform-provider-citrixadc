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
* Configuration for HTML injection parameter resource.
*/
type Filterhtmlinjectionparameter struct {
	/**
	* For a rate of x, HTML injection is done for 1 out of x policy matches.
	*/
	Rate int `json:"rate,omitempty"`
	/**
	* For a frequency of x, HTML injection is done at least once per x milliseconds.
	*/
	Frequency int `json:"frequency,omitempty"`
	/**
	* Searching for <html> tag. If this parameter is enabled, HTML injection does not insert the prebody or postbody content unless the <html> tag is found.
	*/
	Strict string `json:"strict,omitempty"`
	/**
	* Number of characters, in the HTTP body, in which to search for the <html> tag if strict mode is set.
	*/
	Htmlsearchlen int `json:"htmlsearchlen,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
