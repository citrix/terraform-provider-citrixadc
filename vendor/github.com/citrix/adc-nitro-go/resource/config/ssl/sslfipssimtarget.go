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

package ssl

/**
* Configuration for FIPS SIM Target resource.
*/
type Sslfipssimtarget struct {
	/**
	* Name of and, optionally, path to the target FIPS appliance's key vector. /nsconfig/ssl/ is the default path.
	*/
	Keyvector string `json:"keyvector,omitempty"`
	/**
	* Name of and, optionally, path to the source FIPS appliance's secret data. /nsconfig/ssl/ is the default path.
	*/
	Sourcesecret string `json:"sourcesecret,omitempty"`
	/**
	* Name of and, optionally, path to the source FIPS appliance's certificate file. /nsconfig/ssl/ is the default path.
	*/
	Certfile string `json:"certfile,omitempty"`
	/**
	* Name for and, optionally, path to the target FIPS appliance's secret data. The default input path for the secret data is /nsconfig/ssl/.
	*/
	Targetsecret string `json:"targetsecret,omitempty"`

}
