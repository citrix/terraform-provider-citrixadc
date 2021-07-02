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

package ns

/**
* Configuration for Extension resource.
*/
type Nsextension struct {
	/**
	* Local path to and name of, or URL (protocol, host, path, and file name) for, the file in which to store the imported extension.
		NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.
	*/
	Src string `json:"src,omitempty"`
	/**
	* Name to assign to the extension object on the Citrix ADC.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Any comments to preserve information about the extension object.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Overwrites the existing file
	*/
	Overwrite bool `json:"overwrite,omitempty"`
	/**
	* Enables tracing to the NS log file of extension execution:
		off   - turns off tracing (equivalent to unset ns extension <extension-name> -trace)
		calls - traces extension function calls with arguments and function returns with the first return value
		lines - traces the above plus line numbers for executed extension lines
		all   - traces the above plus local variables changed by executed extension lines
		Note that the DEBUG log level must be enabled to see extension tracing.
		This can be done by set audit syslogParams -loglevel ALL or -loglevel DEBUG.
	*/
	Trace string `json:"trace,omitempty"`
	/**
	* Comma-separated list of extension functions to trace. By default, all extension functions are traced.
	*/
	Tracefunctions string `json:"tracefunctions,omitempty"`
	/**
	* Comma-separated list of variables (in traced extension functions) to trace. By default, all variables are traced.
	*/
	Tracevariables string `json:"tracevariables,omitempty"`
	/**
	* Show detail for extension function.
	*/
	Detail string `json:"detail,omitempty"`

	//------- Read only Parameter ---------;

	Type string `json:"type,omitempty"`
	Functionhits string `json:"functionhits,omitempty"`
	Functionundefhits string `json:"functionundefhits,omitempty"`
	Functionhaltcount string `json:"functionhaltcount,omitempty"`

}
