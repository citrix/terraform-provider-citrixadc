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
* Configuration for file resource.
*/
type Systemfile struct {
	/**
	* Name of the file. It should not include filepath.
	*/
	Filename string `json:"filename,omitempty"`
	/**
	* file content in Base64 format.
	*/
	Filecontent string `json:"filecontent,omitempty"`
	/**
	* location of the file on Citrix ADC.
	*/
	Filelocation string `json:"filelocation,omitempty"`
	/**
	* encoding type of the file content.
	*/
	Fileencoding string `json:"fileencoding,omitempty"`

	//------- Read only Parameter ---------;

	Fileaccesstime string `json:"fileaccesstime,omitempty"`
	Filemodifiedtime string `json:"filemodifiedtime,omitempty"`
	Filemode string `json:"filemode,omitempty"`
	Filesize string `json:"filesize,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
