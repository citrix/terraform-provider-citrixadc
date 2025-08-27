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

package gslb

/**
* Configuration for gslb config resource.
*/
type Gslbconfig struct {
	/**
	* Do not synchronize the GSLB sites, but display the commands that would be applied on the slave node upon synchronization. Mutually exclusive with the Save Configuration option.
	*/
	Preview bool `json:"preview,omitempty"`
	/**
	* Generate verbose output when synchronizing the GSLB sites. The Debug option generates more verbose output than the sync gslb config command in which the option is not used, and is useful for analyzing synchronization issues.
	*/
	Debug bool `json:"debug,omitempty"`
	/**
	* Force synchronization of the specified site even if a dependent configuration on the remote site is preventing synchronization or if one or more GSLB entities on the remote site have the same name but are of a different type. You can specify either the name of the remote site that you want to synchronize with the local site, or you can specify All Sites in the configuration utility (the string all-sites in the CLI). If you specify All Sites, all the sites in the GSLB setup are synchronized with the site on the master node.
		Note: If you select the Force Sync option, the synchronization starts without displaying the commands that are going to be executed.
	*/
	Forcesync string `json:"forcesync,omitempty"`
	/**
	* Suppress the warning and the confirmation prompt that are displayed before site synchronization begins. This option can be used in automation scripts that must not be interrupted by a prompt.
	*/
	Nowarn bool `json:"nowarn,omitempty"`
	/**
	* Save the configuration on all the nodes participating in the synchronization process, automatically. The master saves its configuration immediately before synchronization begins. Slave nodes save their configurations after the process of synchronization is complete. A slave node saves its configuration only if the configuration difference was successfully applied to it. Mutually exclusive with the Preview option.
	*/
	Saveconfig bool `json:"saveconfig,omitempty"`
	/**
	* Run the specified command on the master node and then on all the slave nodes. You cannot use this option with the force sync and preview options.
	*/
	Command string `json:"command,omitempty"`

}
