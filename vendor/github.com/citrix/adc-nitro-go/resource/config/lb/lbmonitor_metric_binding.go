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

package lb

/**
* Binding class showing the metric that can be bound to lbmonitor.
*/
type Lbmonitormetricbinding struct {
	/**
	* Metric name in the metric table, whose setting is changed. A value zero disables the metric and it will not be used for load calculation
	*/
	Metric string `json:"metric,omitempty"`
	/**
	* Metric table to which to bind metrics.
	*/
	Metrictable string `json:"metrictable,omitempty"`
	/**
	* Giving the unit of the metric
	*/
	Metricunit string `json:"metric_unit,omitempty"`
	/**
	* The weight for the specified service metric with respect to others.
	*/
	Metricweight *int `json:"metricweight,omitempty"`
	/**
	* Threshold to be used for that metric.
	*/
	Metricthreshold *int `json:"metricthreshold,omitempty"`
	/**
	* Name of the monitor.
	*/
	Monitorname string `json:"monitorname,omitempty"`


}