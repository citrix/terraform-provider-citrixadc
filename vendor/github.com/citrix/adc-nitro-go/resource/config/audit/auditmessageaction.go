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

package audit

/**
* Configuration for message action resource.
*/
type Auditmessageaction struct {
	/**
	* Name of the audit message action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the message action is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my message action" or 'my message action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Audit log level, which specifies the severity level of the log message being generated.. 
		The following loglevels are valid: 
		* EMERGENCY - Events that indicate an immediate crisis on the server.
		* ALERT - Events that might require action.
		* CRITICAL - Events that indicate an imminent server crisis.
		* ERROR - Events that indicate some type of error.
		* WARNING - Events that require action in the near future.
		* NOTICE - Events that the administrator should know about.
		* INFORMATIONAL - All but low-level events.
		* DEBUG - All events, in extreme detail.
	*/
	Loglevel string `json:"loglevel,omitempty"`
	/**
	* Default-syntax expression that defines the format and content of the log message.
	*/
	Stringbuilderexpr string `json:"stringbuilderexpr,omitempty"`
	/**
	* Send the message to the new nslog.
	*/
	Logtonewnslog string `json:"logtonewnslog,omitempty"`
	/**
	* Bypass the safety check and allow unsafe expressions.
	*/
	Bypasssafetycheck string `json:"bypasssafetycheck,omitempty"`

	//------- Read only Parameter ---------;

	Loglevel1 string `json:"loglevel1,omitempty"`
	Hits string `json:"hits,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Referencecount string `json:"referencecount,omitempty"`

}
