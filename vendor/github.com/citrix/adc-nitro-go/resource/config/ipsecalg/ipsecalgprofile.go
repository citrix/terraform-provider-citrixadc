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

package ipsecalg

/**
* Configuration for IPSEC ALG profile resource.
*/
type Ipsecalgprofile struct {
	/**
	* The name of the ipsec alg profile
	*/
	Name string `json:"name,omitempty"`
	/**
	* IKE session timeout in minutes
	*/
	Ikesessiontimeout int `json:"ikesessiontimeout,omitempty"`
	/**
	* ESP session timeout in minutes.
	*/
	Espsessiontimeout int `json:"espsessiontimeout,omitempty"`
	/**
	* Timeout ESP in seconds as no ESP packets are seen after IKE negotiation
	*/
	Espgatetimeout int `json:"espgatetimeout,omitempty"`
	/**
	* Mode in which the connection failover feature must operate for the IPSec Alg. After a failover, established UDP connections and ESP packet flows are kept active and resumed on the secondary appliance. Recomended setting is ENABLED.
	*/
	Connfailover string `json:"connfailover,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
