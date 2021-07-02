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


type Reboot struct {
	/**
	* Restarts the Citrix ADC software without rebooting the underlying operating system. The session terminates and you must log on to the appliance after it has restarted.
		Note: This argument is required only for nCore appliances. Classic appliances ignore this argument.
	*/
	Warm bool `json:"warm,omitempty"`

}
