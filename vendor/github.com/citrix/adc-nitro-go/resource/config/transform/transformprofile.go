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
* Configuration for URL Transformation profile resource.
*/
type Transformprofile struct {
	/**
	* Name for the URL transformation profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the URL transformation profile is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, ^A"my transform profile^A" or ^A'my transform profile^A').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Type of transformation. Always URL for URL Transformation profiles.
	*/
	Type string `json:"type,omitempty"`
	/**
	* In the HTTP body, transform only absolute URLs. Relative URLs are ignored.
	*/
	Onlytransformabsurlinbody string `json:"onlytransformabsurlinbody,omitempty"`
	/**
	* Any comments to preserve information about this URL Transformation profile.
	*/
	Comment string `json:"comment,omitempty"`

	//------- Read only Parameter ---------;

	Regexforfindingurlinjavascript string `json:"regexforfindingurlinjavascript,omitempty"`
	Regexforfindingurlincss string `json:"regexforfindingurlincss,omitempty"`
	Regexforfindingurlinxcomponent string `json:"regexforfindingurlinxcomponent,omitempty"`
	Regexforfindingurlinxml string `json:"regexforfindingurlinxml,omitempty"`
	Additionalreqheaderslist string `json:"additionalreqheaderslist,omitempty"`
	Additionalrespheaderslist string `json:"additionalrespheaderslist,omitempty"`

}
