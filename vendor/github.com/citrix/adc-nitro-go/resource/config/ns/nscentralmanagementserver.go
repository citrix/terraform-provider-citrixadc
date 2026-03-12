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
* Configuration for centralmanagementserver resource.
*/
type Nscentralmanagementserver struct {
	/**
	* Type of the central management server. Must be either CLOUD or ONPREM depending on whether the server is on the cloud or on premise.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Username for access to central management server. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or
		single quotation marks (for example, "my ns centralmgmtserver" or "my ns centralmgmtserver").
	*/
	Username string `json:"username,omitempty"`
	/**
	* Password for access to central management server. Required for any user account.
	*/
	Password string `json:"password,omitempty"`
	/**
	* Activation code is used to register to ADM service
	*/
	Activationcode string `json:"activationcode,omitempty"`
	/**
	* Ip Address of central management server.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Fully qualified domain name of the central management server or service-url to locate ADM service.
	*/
	Servername string `json:"servername,omitempty"`
	/**
	* validate the server certificate for secure SSL connections.
	*/
	Validatecert string `json:"validatecert,omitempty"`
	/**
	* Device profile is created on ADM and contains the user name and password of the instance(s).
	*/
	Deviceprofilename string `json:"deviceprofilename,omitempty"`
	/**
	* ADC username used to create device profile on ADM
	*/
	Adcusername string `json:"adcusername,omitempty"`
	/**
	* ADC password used to create device profile on ADM
	*/
	Adcpassword string `json:"adcpassword,omitempty"`

	//------- Read only Parameter ---------;

	Instanceid string `json:"instanceid,omitempty"`
	Customerid string `json:"customerid,omitempty"`
	Admserviceenvironment string `json:"admserviceenvironment,omitempty"`
	Admserviceconnectionstatus string `json:"admserviceconnectionstatus,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
