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

package basic

/**
* Configuration for location parameter resource.
*/
type Locationparameter struct {
	/**
	* Context for describing locations. In geographic context, qualifier labels are assigned by default in the following sequence: Continent.Country.Region.City.ISP.Organization. In custom context, the qualifiers labels can have any meaning that you designate.
	*/
	Context string `json:"context,omitempty"`
	/**
	* Label specifying the meaning of the first qualifier. Can be specified for custom context only.
	*/
	Q1label string `json:"q1label,omitempty"`
	/**
	* Label specifying the meaning of the second qualifier. Can be specified for custom context only.
	*/
	Q2label string `json:"q2label,omitempty"`
	/**
	* Label specifying the meaning of the third qualifier. Can be specified for custom context only.
	*/
	Q3label string `json:"q3label,omitempty"`
	/**
	* Label specifying the meaning of the fourth qualifier. Can be specified for custom context only.
	*/
	Q4label string `json:"q4label,omitempty"`
	/**
	* Label specifying the meaning of the fifth qualifier. Can be specified for custom context only.
	*/
	Q5label string `json:"q5label,omitempty"`
	/**
	* Label specifying the meaning of the sixth qualifier. Can be specified for custom context only.
	*/
	Q6label string `json:"q6label,omitempty"`
	/**
	* Indicates whether wildcard qualifiers should match any other
		qualifier including non-wildcard while evaluating
		location based expressions.
		Possible values: Yes, No, Expression.
		Yes - Wildcard qualifiers match any other qualifiers.
		No  - Wildcard qualifiers do not match non-wildcard
		qualifiers, but match other wildcard qualifiers.
		Expression - Wildcard qualifiers in an expression
		match any qualifier in an LDNS location,
		wildcard qualifiers in the LDNS location do not match
		non-wildcard qualifiers in an expression
	*/
	Matchwildcardtoany string `json:"matchwildcardtoany,omitempty"`

	//------- Read only Parameter ---------;

	Locationfile string `json:"Locationfile,omitempty"`
	Format string `json:"format,omitempty"`
	Custom string `json:"custom,omitempty"`
	Static string `json:"Static,omitempty"`
	Lines string `json:"lines,omitempty"`
	Errors string `json:"errors,omitempty"`
	Warnings string `json:"warnings,omitempty"`
	Entries string `json:"entries,omitempty"`
	Locationfile6 string `json:"locationfile6,omitempty"`
	Format6 string `json:"format6,omitempty"`
	Custom6 string `json:"custom6,omitempty"`
	Static6 string `json:"static6,omitempty"`
	Lines6 string `json:"lines6,omitempty"`
	Errors6 string `json:"errors6,omitempty"`
	Warnings6 string `json:"warnings6,omitempty"`
	Entries6 string `json:"entries6,omitempty"`
	Flags string `json:"flags,omitempty"`
	Status string `json:"status,omitempty"`
	Databasemode string `json:"databasemode,omitempty"`
	Flushing string `json:"flushing,omitempty"`
	Loading string `json:"loading,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
