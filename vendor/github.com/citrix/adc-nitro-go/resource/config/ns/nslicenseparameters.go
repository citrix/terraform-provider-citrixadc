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

package ns

/**
* Configuration for licenseparameters resource.
*/
type Nslicenseparameters struct {
	/**
	* If ADC remains in grace for the configured hours then first grace alert will be raised
	*/
	Alert1gracetimeout int `json:"alert1gracetimeout"` // Zero is a valid value
	/**
	* If ADC remains in grace for the configured hours then major grace alert will be raised
	*/
	Alert2gracetimeout int `json:"alert2gracetimeout,omitempty"`
	/**
	* If ADC license contract expiry date is nearer then GUI/SNMP license expiry alert will be raised
	*/
	Licenseexpiryalerttime int `json:"licenseexpiryalerttime,omitempty"`
	/**
	* Heartbeat between ADC and Licenseserver is configurable and applicable in case of pooled licensing
	*/
	Heartbeatinterval int `json:"heartbeatinterval,omitempty"`
	/**
	* Inventory refresh interval between ADC and Licenseserver is configurable and applicable in case of pooled licensing
	*/
	Inventoryrefreshinterval int `json:"inventoryrefreshinterval,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
