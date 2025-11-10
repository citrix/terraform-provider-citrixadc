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

package lsn

/**
* Configuration for LSN RTSPALG Profile resource.
*/
type Lsnrtspalgprofile struct {
	/**
	* The name of the RTSPALG Profile.
	*/
	Rtspalgprofilename string `json:"rtspalgprofilename,omitempty"`
	/**
	* Idle timeout for the rtsp sessions in seconds.
	*/
	Rtspidletimeout *int `json:"rtspidletimeout,omitempty"`
	/**
	* port for the RTSP
	*/
	Rtspportrange string `json:"rtspportrange,omitempty"`
	/**
	* RTSP ALG Profile transport protocol type.
	*/
	Rtsptransportprotocol string `json:"rtsptransportprotocol,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
