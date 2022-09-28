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

package subscriber

/**
* Configuration for RADIUS interface Parameters resource.
*/
type Subscriberradiusinterface struct {
	/**
	* Name of RADIUS LISTENING service that will process RADIUS accounting requests.
	*/
	Listeningservice string `json:"listeningservice,omitempty"`
	/**
	* Treat radius interim message as start radius messages.
	*/
	Radiusinterimasstart string `json:"radiusinterimasstart,omitempty"`

	//------- Read only Parameter ---------;

	Svrstate string `json:"svrstate,omitempty"`

}
