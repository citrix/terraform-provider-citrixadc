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
* Configuration for rpc node resource.
*/
type Nsrpcnode struct {
	/**
	* IP address of the node. This has to be in the same subnet as the NSIP address.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Password to be used in authentication with the peer system node.
	*/
	Password string `json:"password,omitempty"`
	/**
	* Source IP address to be used to communicate with the peer system node. The default value is 0, which means that the appliance uses the NSIP address as the source IP address.
	*/
	Srcip string `json:"srcip,omitempty"`
	/**
	* State of the channel when talking to the node.
	*/
	Secure string `json:"secure,omitempty"`
	/**
	* validate the server certificate for secure SSL connections
	*/
	Validatecert string `json:"validatecert,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
