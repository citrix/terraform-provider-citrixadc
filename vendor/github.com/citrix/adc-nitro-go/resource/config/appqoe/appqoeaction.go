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

package appqoe

/**
* Configuration for AppQoS action resource.
*/
type Appqoeaction struct {
	/**
	* Name for the AppQoE action. Must begin with a letter, number, or the underscore symbol (_). Other characters allowed, after the first character, are the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), and colon (:) characters. This is a mandatory argument
	*/
	Name string `json:"name,omitempty"`
	/**
	* Priority for queuing the request. If server resources are not available for a request that matches the configured rule, this option specifies a priority for queuing the request until the server resources are available again. If priority is not configured then Lowest priority will be used to queue the request.
	*/
	Priority string `json:"priority,omitempty"`
	/**
	* Responder action to be taken when the threshold is reached. Available settings function as follows:
		ACS - Serve content from an alternative content service
		Threshold : maxConn or delay
		NS - Serve from the Citrix ADC (built-in response)
		Threshold : maxConn or delay
	*/
	Respondwith string `json:"respondwith,omitempty"`
	/**
	* name of the HTML page object to use as the response
	*/
	Customfile string `json:"customfile,omitempty"`
	/**
	* Name of the alternative content service to be used in the ACS
	*/
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	/**
	* Path to the alternative content service to be used in the ACS
	*/
	Altcontentpath string `json:"altcontentpath,omitempty"`
	/**
	* Policy queue depth threshold value. When the policy queue size (number of requests queued for the policy binding this action is attached to) increases to the specified polqDepth value, subsequent requests are dropped to the lowest priority level.
	*/
	Polqdepth int `json:"polqdepth"` // Zero is a valid value
	/**
	* Queue depth threshold value per priorirty level. If the queue size (number of requests in the queue of that particular priorirty) on the virtual server to which this policy is bound, increases to the specified qDepth value, subsequent requests are dropped to the lowest priority level.
	*/
	Priqdepth int `json:"priqdepth"` // Zero is a valid value
	/**
	* Maximum number of concurrent connections that can be open for requests that matches with rule.
	*/
	Maxconn int `json:"maxconn,omitempty"`
	/**
	* Delay threshold, in microseconds, for requests that match the policy's rule. If the delay statistics gathered for the matching request exceed the specified delay, configured action triggered for that request, if there is no action then requests are dropped to the lowest priority level
	*/
	Delay int `json:"delay,omitempty"`
	/**
	* Optional expression to add second level check to trigger DoS actions. Specifically used for Analytics based DoS response generation
	*/
	Dostrigexpression string `json:"dostrigexpression,omitempty"`
	/**
	* DoS Action to take when vserver will be considered under DoS attack and corresponding rule matches. Mandatory if AppQoE actions are to be used for DoS attack prevention.
	*/
	Dosaction string `json:"dosaction,omitempty"`
	/**
	* Bind TCP Profile based on L2/L3/L7 parameters.
	*/
	Tcpprofile string `json:"tcpprofile,omitempty"`
	/**
	* Retry on TCP Reset
	*/
	Retryonreset string `json:"retryonreset,omitempty"`
	/**
	* Retry on request Timeout(in millisec) upon sending request to backend servers
	*/
	Retryontimeout int `json:"retryontimeout,omitempty"`
	/**
	* Retry count
	*/
	Numretries int `json:"numretries"` // Zero is a valid value

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
