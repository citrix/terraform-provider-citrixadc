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
* Configuration for system log parameters resource.
*/
type Auditsyslogparams struct {
	/**
	* IP address of the syslog server.
	*/
	Serverip string `json:"serverip,omitempty"`
	/**
	* Port on which the syslog server accepts connections.
	*/
	Serverport int `json:"serverport,omitempty"`
	/**
	* Format of dates in the logs.
		Supported formats are: 
		* MMDDYYYY - U.S. style month/date/year format.
		* DDMMYYYY. European style  -date/month/year format.
		* YYYYMMDD - ISO style year/month/date format.
	*/
	Dateformat string `json:"dateformat,omitempty"`
	/**
	* Types of information to be logged. 
		Available settings function as follows: 
		* ALL - All events.
		* EMERGENCY - Events that indicate an immediate crisis on the server.
		* ALERT - Events that might require action.
		* CRITICAL - Events that indicate an imminent server crisis.
		* ERROR - Events that indicate some type of error.
		* WARNING - Events that require action in the near future.
		* NOTICE - Events that the administrator should know about.
		* INFORMATIONAL - All but low-level events.
		* DEBUG - All events, in extreme detail.
		* NONE - No events.
	*/
	Loglevel []string `json:"loglevel,omitempty"`
	/**
	* Facility value, as defined in RFC 3164, assigned to the log message. 
		Log facility values are numbers 0 to 7 (LOCAL0 through LOCAL7). Each number indicates where a specific message originated from, such as the Citrix ADC itself, the VPN, or external.
	*/
	Logfacility string `json:"logfacility,omitempty"`
	/**
	* Log TCP messages.
	*/
	Tcp string `json:"tcp,omitempty"`
	/**
	* Log access control list (ACL) messages.
	*/
	Acl string `json:"acl,omitempty"`
	/**
	* Time zone used for date and timestamps in the logs. 
		Available settings function as follows: 
		* GMT_TIME - Coordinated Universal Time.
		* LOCAL_TIME  Use the server's timezone setting.
	*/
	Timezone string `json:"timezone,omitempty"`
	/**
	* Log user-configurable log messages to syslog. 
		Setting this parameter to NO causes audit to ignore all user-configured message actions. Setting this parameter to YES causes audit to log user-configured message actions that meet the other logging criteria.
	*/
	Userdefinedauditlog string `json:"userdefinedauditlog,omitempty"`
	/**
	* Export log messages to AppFlow collectors.
		Appflow collectors are entities to which log messages can be sent so that some action can be performed on them.
	*/
	Appflowexport string `json:"appflowexport,omitempty"`
	/**
	* Log the LSN messages
	*/
	Lsn string `json:"lsn,omitempty"`
	/**
	* Log the ALG messages
	*/
	Alg string `json:"alg,omitempty"`
	/**
	* Log subscriber session event information
	*/
	Subscriberlog string `json:"subscriberlog,omitempty"`
	/**
	* Log DNS related syslog messages
	*/
	Dns string `json:"dns,omitempty"`
	/**
	* Log SSL Interceptionn event information
	*/
	Sslinterception string `json:"sslinterception,omitempty"`
	/**
	* Log URL filtering event information
	*/
	Urlfiltering string `json:"urlfiltering,omitempty"`
	/**
	* Log Content Inspection event ifnormation
	*/
	Contentinspectionlog string `json:"contentinspectionlog,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
