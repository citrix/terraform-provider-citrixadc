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

package cloud

/**
* Configuration for cloud parameter resource.
*/
type Cloudparameter struct {
	/**
	* FQDN of the controller to which the Citrix ADC SDProxy Connects
	*/
	Controllerfqdn string `json:"controllerfqdn,omitempty"`
	/**
	* Port number of the controller to which the Citrix ADC SDProxy connects
	*/
	Controllerport int `json:"controllerport,omitempty"`
	/**
	* Instance ID of the customer provided by Trust
	*/
	Instanceid string `json:"instanceid,omitempty"`
	/**
	* Customer ID of the citrix cloud customer
	*/
	Customerid string `json:"customerid,omitempty"`
	/**
	* Resource Location of the customer provided by Trust
	*/
	Resourcelocation string `json:"resourcelocation,omitempty"`
	/**
	* Activation code for the NGS Connector instance
	*/
	Activationcode string `json:"activationcode,omitempty"`
	/**
	* Describes if the customer is a Staging/Production or Dev Citrix Cloud customer
	*/
	Deployment string `json:"deployment,omitempty"`
	/**
	* Identifies whether the connector is located Onprem, Aws or Azure
	*/
	Connectorresidence string `json:"connectorresidence,omitempty"`

	//------- Read only Parameter ---------;

	Controlconnectionstatus string `json:"controlconnectionstatus,omitempty"`

}
