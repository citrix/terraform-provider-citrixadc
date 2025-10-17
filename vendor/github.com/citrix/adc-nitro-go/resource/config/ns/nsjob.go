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
* Configuration for job resource.
*/
type Nsjob struct {
	/**
	* Running job id
	*/
	Id *int `json:"id,omitempty"`
	/**
	* Running job Name
	*/
	Name string `json:"name,omitempty"`

	//------- Read only Parameter ---------;

	Status string `json:"status,omitempty"`
	Progress string `json:"progress,omitempty"`
	Timeelapsed string `json:"timeelapsed,omitempty"`
	Errorcode string `json:"errorcode,omitempty"`
	Message string `json:"message,omitempty"`
	Response string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
