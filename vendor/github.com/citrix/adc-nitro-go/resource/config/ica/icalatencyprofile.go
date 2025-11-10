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

package ica

/**
* Configuration for Profile for Latency monitoring resource.
*/
type Icalatencyprofile struct {
	/**
	* Name for the ICA latencyprofile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and
		the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA latency profile is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my ica l7latencyprofile" or 'my ica l7latencyprofile').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Enable/Disable L7 Latency monitoring for L7 latency notifications
	*/
	L7latencymonitoring string `json:"l7latencymonitoring,omitempty"`
	/**
	* L7 Latency threshold factor. This is the factor by which the active latency should be greater than the minimum observed value to determine that the latency is high and may need to be reported
	*/
	L7latencythresholdfactor *int `json:"l7latencythresholdfactor,omitempty"`
	/**
	* L7 Latency Wait time. This is the time for which the Citrix ADC waits after the threshold is exceeded before it sends out a Notification to the Insight Center.
	*/
	L7latencywaittime *int `json:"l7latencywaittime,omitempty"`
	/**
	* L7 Latency Notify Interval. This is the interval at which the Citrix ADC sends out notifications to the Insight Center after the wait time has passed.
	*/
	L7latencynotifyinterval *int `json:"l7latencynotifyinterval,omitempty"`
	/**
	* L7 Latency Max notify Count. This is the upper limit on the number of notifications sent to the Insight Center within an interval where the Latency is above the threshold.
	*/
	L7latencymaxnotifycount *int `json:"l7latencymaxnotifycount,omitempty"`

	//------- Read only Parameter ---------;

	Refcnt string `json:"refcnt,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Isdefault string `json:"isdefault,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
