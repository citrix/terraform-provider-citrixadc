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

package lsn

/**
* Configuration for LSN SIPALG Profile resource.
*/
type Lsnsipalgprofile struct {
	/**
	* The name of the SIPALG Profile.
	*/
	Sipalgprofilename string `json:"sipalgprofilename,omitempty"`
	/**
	* Idle timeout for the data channel sessions in seconds.
	*/
	Datasessionidletimeout *int `json:"datasessionidletimeout,omitempty"`
	/**
	* SIP control channel session timeout in seconds.
	*/
	Sipsessiontimeout *int `json:"sipsessiontimeout,omitempty"`
	/**
	* SIP registration timeout in seconds.
	*/
	Registrationtimeout *int `json:"registrationtimeout,omitempty"`
	/**
	* Source port range for SIP_UDP and SIP_TCP.
	*/
	Sipsrcportrange string `json:"sipsrcportrange,omitempty"`
	/**
	* Destination port range for SIP_UDP and SIP_TCP.
	*/
	Sipdstportrange string `json:"sipdstportrange,omitempty"`
	/**
	* ENABLE/DISABLE RegisterPinhole creation.
	*/
	Openregisterpinhole string `json:"openregisterpinhole,omitempty"`
	/**
	* ENABLE/DISABLE ContactPinhole creation.
	*/
	Opencontactpinhole string `json:"opencontactpinhole,omitempty"`
	/**
	* ENABLE/DISABLE ViaPinhole creation.
	*/
	Openviapinhole string `json:"openviapinhole,omitempty"`
	/**
	* ENABLE/DISABLE RecordRoutePinhole creation.
	*/
	Openrecordroutepinhole string `json:"openrecordroutepinhole,omitempty"`
	/**
	* SIP ALG Profile transport protocol type.
	*/
	Siptransportprotocol string `json:"siptransportprotocol,omitempty"`
	/**
	* ENABLE/DISABLE RoutePinhole creation.
	*/
	Openroutepinhole string `json:"openroutepinhole,omitempty"`
	/**
	* ENABLE/DISABLE rport.
	*/
	Rport string `json:"rport,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
