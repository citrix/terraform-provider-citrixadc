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

package dns

/**
* Binding class showing the gslbservice that can be bound to dnsview.
*/
type Dnsviewgslbservicebinding struct {
	/**
	* Service name of the service using this view.
	*/
	Gslbservicename string `json:"gslbservicename,omitempty"`
	/**
	* IP of the service corresponding to the given view.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Name of the view to display.
	*/
	Viewname string `json:"viewname,omitempty"`


}