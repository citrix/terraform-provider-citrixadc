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

package network

/**
* Binding class showing the interface that can be bound to channel.
*/
type Channelinterfacebinding struct {
	/**
	* Interfaces to be bound to the LA channel of a Citrix ADC or to the LA channel of a cluster configuration.
		For an LA channel of a Citrix ADC, specify an interface in C/U notation (for example, 1/3).
		For an LA channel of a cluster configuration, specify an interface in N/C/U notation (for example, 2/1/3).
		where C can take one of the following values:
		* 0 - Indicates a management interface.
		* 1 - Indicates a 1 Gbps port.
		* 10 - Indicates a 10 Gbps port.
		U is a unique integer for representing an interface in a particular port group.
		N is the ID of the node to which an interface belongs in a cluster configuration.
		Use spaces to separate multiple entries.
	*/
	Ifnum []string `json:"ifnum,omitempty"`
	/**
	* The  mode(AUTO/MANNUAL) for the LA channel.
	*/
	Lamode string `json:"lamode,omitempty"`
	/**
	* State of the member interfaces.
	*/
	Slavestate *int `json:"slavestate,omitempty"`
	/**
	* Media type of the member interfaces.
	*/
	Slavemedia *int `json:"slavemedia,omitempty"`
	/**
	* Speed of the member interfaces.
	*/
	Slavespeed *int `json:"slavespeed,omitempty"`
	/**
	* Duplex of the member interfaces.
	*/
	Slaveduplex *int `json:"slaveduplex,omitempty"`
	/**
	* Flowcontrol of the member interfaces.
	*/
	Slaveflowctl *int `json:"slaveflowctl,omitempty"`
	/**
	* UP time of the member interfaces.
	*/
	Slavetime *int `json:"slavetime,omitempty"`
	/**
	* LR set member interface state(active/inactive).
	*/
	Lractiveintf *int `json:"lractiveintf,omitempty"`
	/**
	* New attribute added to identify the source of cmd, when SVM fires the nitro cmd, it will set the value of SVMCMD to be 1. 
	*/
	Svmcmd *int `json:"svmcmd,omitempty"`
	/**
	* ID of the LA channel or the cluster LA channel to which you want to bind interfaces. Specify an LA channel in LA/x notation, where x can range from 1 to 8 or a cluster LA channel in CLA/x notation or  Link redundant channel in LR/x notation , where x can range from 1 to 4.
	*/
	Id string `json:"id,omitempty"`


}