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

package kafka

/**
* Configuration for Kafka cluster resource.
 */
type Kafkacluster struct {
	/**
	* Name for the Kafka cluster.
	 */
	Name string `json:"name,omitempty"`
	/**
	* Total active services bound to servicegroup.
	 */
	Activesvc float64 `json:"activesvc,omitempty"`
	/**
	* Total services bound to servicegroup.
	 */
	Totalsvc float64 `json:"totalsvc,omitempty"`
	/**
	* Topic of the servicegroup.
	 */
	Topicname string `json:"topicname,omitempty"`
	/**
	* Total number of topic servicegroups bound.
	 */
	Numtopics float64 `json:"numtopics,omitempty"`
}
