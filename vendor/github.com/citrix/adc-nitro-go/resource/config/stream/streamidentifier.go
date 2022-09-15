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

package stream

/**
* Configuration for identifier resource.
*/
type Streamidentifier struct {
	/**
	* The name of stream identifier.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the selector to use with the stream identifier.
	*/
	Selectorname string `json:"selectorname,omitempty"`
	/**
	* Number of minutes of data to use when calculating session statistics (number of requests, bandwidth, and response times). The interval is a moving window that keeps the most recently collected data. Older data is discarded at regular intervals.
	*/
	Interval int `json:"interval,omitempty"`
	/**
	* Size of the sample from which to select a request for evaluation. The smaller the sample count, the more accurate is the statistical data. To evaluate all requests, set the sample count to 1. However, such a low setting can result in excessive consumption of memory and processing resources.
	*/
	Samplecount int `json:"samplecount,omitempty"`
	/**
	* Sort stored records by the specified statistics column, in descending order. Performed during data collection, the sorting enables real-time data evaluation through Citrix ADC policies (for example, compression and caching policies) that use functions such as IS_TOP(n).
	*/
	Sort string `json:"sort,omitempty"`
	/**
	* Enable/disable SNMP trap for stream identifier
	*/
	Snmptrap string `json:"snmptrap,omitempty"`
	/**
	* Enable/disable Appflow logging for stream identifier
	*/
	Appflowlog string `json:"appflowlog,omitempty"`
	/**
	* Track ack only packets as well. This setting is applicable only when packet rate limiting is being used.
	*/
	Trackackonlypackets string `json:"trackackonlypackets,omitempty"`
	/**
	* Track transactions exceeding configured threshold. Transaction tracking can be enabled for following metric: ResponseTime.
		By default transaction tracking is disabled
	*/
	Tracktransactions string `json:"tracktransactions,omitempty"`
	/**
	* Maximum per transcation value of metric. Metric to be tracked is specified by tracktransactions attribute.
	*/
	Maxtransactionthreshold int `json:"maxtransactionthreshold,omitempty"`
	/**
	* Minimum per transcation value of metric. Metric to be tracked is specified by tracktransactions attribute.
	*/
	Mintransactionthreshold int `json:"mintransactionthreshold,omitempty"`
	/**
	* Non-Breaching transactions to Total transactions threshold expressed in percent.
		Maximum of 6 decimal places is supported.
	*/
	Acceptancethreshold string `json:"acceptancethreshold,omitempty"`
	/**
	* Breaching transactions threshold calculated over interval.
	*/
	Breachthreshold int `json:"breachthreshold,omitempty"`

	//------- Read only Parameter ---------;

	Rule string `json:"rule,omitempty"`

}
