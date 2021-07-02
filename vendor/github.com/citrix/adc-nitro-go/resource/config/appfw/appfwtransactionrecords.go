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

package appfw

/**
* Configuration for Application firewall transaction record resource.
*/
type Appfwtransactionrecords struct {
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid uint32 `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Httptransactionid string `json:"httptransactionid,omitempty"`
	Packetengineid string `json:"packetengineid,omitempty"`
	Appfwsessionid string `json:"appfwsessionid,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Url string `json:"url,omitempty"`
	Clientip string `json:"clientip,omitempty"`
	Destip string `json:"destip,omitempty"`
	Starttime string `json:"starttime,omitempty"`
	Endtime string `json:"endtime,omitempty"`
	Requestcontentlength string `json:"requestcontentlength,omitempty"`
	Requestyields string `json:"requestyields,omitempty"`
	Requestmaxprocessingtime string `json:"requestmaxprocessingtime,omitempty"`
	Responsecontentlength string `json:"responsecontentlength,omitempty"`
	Responseyields string `json:"responseyields,omitempty"`
	Responsemaxprocessingtime string `json:"responsemaxprocessingtime,omitempty"`

}
