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
* Configuration for location resource.
*/
type Location struct {
	/**
	* First IP address in the range, in dotted decimal notation.
	*/
	Ipfrom string `json:"ipfrom,omitempty"`
	/**
	* Last IP address in the range, in dotted decimal notation.
	*/
	Ipto string `json:"ipto,omitempty"`
	/**
	* String of qualifiers, in dotted notation, describing the geographical location of the IP address range. Each qualifier is more specific than the one that precedes it, as in continent.country.region.city.isp.organization. For example, "NA.US.CA.San Jose.ATT.citrix".
		Note: A qualifier that includes a dot (.) or space ( ) must be enclosed in double quotation marks.
	*/
	Preferredlocation string `json:"preferredlocation,omitempty"`
	/**
	* Numerical value, in degrees, specifying the longitude of the geographical location of the IP address-range.
		Note: Longitude and latitude parameters are used for selecting a service with the static proximity GSLB method. If they are not specified, selection is based on the qualifiers specified for the location.
	*/
	Longitude *int `json:"longitude,omitempty"`
	/**
	* Numerical value, in degrees, specifying the latitude of the geographical location of the IP address-range.
		Note: Longitude and latitude parameters are used for selecting a service with the static proximity GSLB method. If they are not specified, selection is based on the qualifiers specified for the location.
	*/
	Latitude *int `json:"latitude,omitempty"`

	//------- Read only Parameter ---------;

	Q1label string `json:"q1label,omitempty"`
	Q2label string `json:"q2label,omitempty"`
	Q3label string `json:"q3label,omitempty"`
	Q4label string `json:"q4label,omitempty"`
	Q5label string `json:"q5label,omitempty"`
	Q6label string `json:"q6label,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
