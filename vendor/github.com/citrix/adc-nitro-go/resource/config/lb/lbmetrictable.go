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

package lb

/**
* Configuration for metric table resource.
*/
type Lbmetrictable struct {
	/**
	* Name for the metric table. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my metrictable" or 'my metrictable').
	*/
	Metrictable string `json:"metrictable,omitempty"`
	/**
	* Name of the metric for which to change the SNMP OID.
	*/
	Metric string `json:"metric,omitempty"`
	/**
	* New SNMP OID of the metric.
	*/
	Snmpoid string `json:"Snmpoid,omitempty"`

	//------- Read only Parameter ---------;

	Metrictype string `json:"metrictype,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
