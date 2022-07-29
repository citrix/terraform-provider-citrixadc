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

package appflow

/**
* Configuration for AppFlow action resource.
*/
type Appflowaction struct {
	/**
	* Name for the action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow action" or 'my appflow action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name(s) of collector(s) to be associated with the AppFlow action.
	*/
	Collectors []string `json:"collectors,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will collect the time required to load and render the mainpage on the client.
	*/
	Clientsidemeasurements string `json:"clientsidemeasurements,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will start tracking the page for waterfall chart by inserting a NS_ESNS cookie in the response.
	*/
	Pagetracking string `json:"pagetracking,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will send the webinsight records to the configured collectors.
	*/
	Webinsight string `json:"webinsight,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will send the security insight records to the configured collectors.
	*/
	Securityinsight string `json:"securityinsight,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will send the bot insight records to the configured collectors.
	*/
	Botinsight string `json:"botinsight,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will send the ContentInspection Insight records to the configured collectors.
	*/
	Ciinsight string `json:"ciinsight,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will send the videoinsight records to the configured collectors.
	*/
	Videoanalytics string `json:"videoanalytics,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will distribute records among the collectors. Else, all records will be sent to all the collectors.
	*/
	Distributionalgorithm string `json:"distributionalgorithm,omitempty"`
	/**
	* If only the stats records are to be exported, turn on this option.
	*/
	Metricslog bool `json:"metricslog,omitempty"`
	/**
	* Log ANOMALOUS or ALL transactions
	*/
	Transactionlog string `json:"transactionlog,omitempty"`
	/**
	* Any comments about this action.  In the CLI, if including spaces between words, enclose the comment in quotation marks. (The quotation marks are not required in the configuration utility.)
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* New name for the AppFlow action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at
		(@), equals (=), and hyphen (-) characters. 
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow action" or 'my appflow action').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Referencecount string `json:"referencecount,omitempty"`
	Description string `json:"description,omitempty"`

}
