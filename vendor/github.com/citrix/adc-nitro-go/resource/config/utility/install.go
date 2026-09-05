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

package utility


type Install struct {
	/**
	* Url for the build file. Must be in the following formats:
		http://[user]:[password]@host/path/to/file
		https://[user]:[password]@host/path/to/file
		sftp://[user]:[password]@host/path/to/file
		scp://[user]:[password]@host/path/to/file
		ftp://[user]:[password]@host/path/to/file
		file://path/to/file
	*/
	Url string `json:"url,omitempty"`
	/**
	* Use this flag to return the install id when the nitro api request is sent.
		The id can be used later to track the installation progress via show ns job <id> command.
		For the cli request of install the flag is by default set as false as the installation progress details can be tracked via cli
	*/
	Async bool `json:"Async,omitempty"`
	/**
	* Do not prompt for yes/no before rebooting.
	*/
	Y bool `json:"y,omitempty"`
	/**
	* Use this flag to enable callhome.
	*/
	L bool `json:"l,omitempty"`
	/**
	* Use this flag to enable Citrix ADM Service Connect. This feature helps you discover your Citrix ADC instances effortlessly on Citrix ADM service and get insights and curated machine learning based recommendations for applications and Citrix ADC infrastructure. This feature lets the Citrix ADC instance automatically send system, usage and telemetry data to Citrix ADM service. View here [https://docs.citrix.com/en-us/citrix-adc/13/data-governance.html] to learn more about this feature. Use of this feature is subject to the Citrix End User ServiceAgreement. View here [https://www.citrix.com/buy/licensing/agreements.html].
	*/
	A bool `json:"a,omitempty"`
	/**
	* Use this flag for upgrading from/to enhancement mode.
	*/
	Enhancedupgrade bool `json:"enhancedupgrade,omitempty"`
	/**
	* Use this flag to perform FIPS installation.
	*/
	Fipsinstall bool `json:"fipsinstall,omitempty"`
	/**
	* Use this flag to answer yes to all prompts.
	*/
	Answeryestoall bool `json:"answeryestoall,omitempty"`
	/**
	* Use this flag to prevent reboot after installation when answerYesToAll is true.
	*/
	Dontreboot bool `json:"dontreboot,omitempty"`
	/**
	* Use this flag to skip ns.conf version equivalence check during downgrade.
	*/
	Dontchecknsconf bool `json:"dontchecknsconf,omitempty"`
	/**
	* Use this flag to delete all signature files and associated kernel images during installation.
	*/
	Deletesigfiles bool `json:"deletesigfiles,omitempty"`
	/**
	* Use this flag to ignore nsapimgr symbols not found error(s) during installation.
	*/
	Ignorensapimgrerrors bool `json:"ignorensapimgrerrors,omitempty"`
	/**
	* Use this flag to exit on license server connectivity errors.
	*/
	Exitonlicserverconnerror bool `json:"exitonlicserverconnerror,omitempty"`
	/**
	* Use this flag to run all installation pre-checks in a single step
	*/
	Precheck bool `json:"precheck,omitempty"`
	/**
	* Use this string to pass extra flags which are not yet supported.
		Example: -flag1 -flag2 -flag3
	*/
	Advancedoptions string `json:"advancedoptions,omitempty"`
	/**
	* Use this flag to ignore unsaved config check during build update.
	*/
	Ignoreunsavedconfig bool `json:"ignoreunsavedconfig,omitempty"`
	/**
	* Use this flag to ignore unsynced HA config check during build update.
	*/
	Ignoreunsyncedconfig bool `json:"ignoreunsyncedconfig,omitempty"`
	/**
	* Use this flag to ignore certificate digest verification errors during build update.
	*/
	Ignorecertcheckerrors bool `json:"ignorecertcheckerrors,omitempty"`
	/**
	* Use this flag to change swap size on ONLY 64bit nCore/MCNS/VMPE systems NON-VPX systems.
	*/
	Resizeswapvar bool `json:"resizeswapvar,omitempty"`

	//------- Read only Parameter ---------;

	Id string `json:"id,omitempty"`

}
