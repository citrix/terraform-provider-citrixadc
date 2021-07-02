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
* Configuration for fips resource.
*/
type Sslfips struct {
	/**
	* FIPS initialization level. The appliance currently supports Level-2 (FIPS 140-2).
	*/
	Inithsm string `json:"inithsm,omitempty"`
	/**
	* Security officer password that will be in effect after you have configured the HSM.
	*/
	Sopassword string `json:"sopassword,omitempty"`
	/**
	* Old password for the security officer.
	*/
	Oldsopassword string `json:"oldsopassword,omitempty"`
	/**
	* The Hardware Security Module's (HSM) User password.
	*/
	Userpassword string `json:"userpassword,omitempty"`
	/**
	* Label to identify the Hardware Security Module (HSM).
	*/
	Hsmlabel string `json:"hsmlabel,omitempty"`
	/**
	* Path to the FIPS firmware file.
	*/
	Fipsfw string `json:"fipsfw,omitempty"`

	//------- Read only Parameter ---------;

	Erasedata string `json:"erasedata,omitempty"`
	Serial string `json:"serial,omitempty"`
	Majorversion string `json:"majorversion,omitempty"`
	Minorversion string `json:"minorversion,omitempty"`
	Fipshwmajorversion string `json:"fipshwmajorversion,omitempty"`
	Fipshwminorversion string `json:"fipshwminorversion,omitempty"`
	Fipshwversionstring string `json:"fipshwversionstring,omitempty"`
	Flashmemorytotal string `json:"flashmemorytotal,omitempty"`
	Flashmemoryfree string `json:"flashmemoryfree,omitempty"`
	Sramtotal string `json:"sramtotal,omitempty"`
	Sramfree string `json:"sramfree,omitempty"`
	Status string `json:"status,omitempty"`
	Flag string `json:"flag,omitempty"`
	Serialno string `json:"serialno,omitempty"`
	Model string `json:"model,omitempty"`
	State string `json:"state,omitempty"`
	Firmwarereleasedate string `json:"firmwarereleasedate,omitempty"`
	Coresmax string `json:"coresmax,omitempty"`
	Coresenabled string `json:"coresenabled,omitempty"`

}
