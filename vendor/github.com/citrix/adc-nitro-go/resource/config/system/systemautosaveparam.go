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

package system

/**
* Configuration for System Autosaveparam resource.
*/
type Systemautosaveparam struct {
	/**
	* Configure autosave feature. Aavilable options are:
		* DEFAULT - NetScaler decides default option for autosave feature.
		* DISABLED - Autosave feature is disabled.
		* ENABLED - Autosave feature is enabled.
	*/
	Status string `json:"status,omitempty"`
	/**
	* Enable or disable periodic save of autosave configuration. If enabled, saveconfig will be done periodically for all partitions including default
	*/
	Periodicsave string `json:"periodicsave,omitempty"`
	/**
	* Frequency in multiple of 60 minutes for periodic save of autosave configuration. Default value is 720 minutes.
	*/
	Periodicsavefrequency *int `json:"periodicsavefrequency,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
