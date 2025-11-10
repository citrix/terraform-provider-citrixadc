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
* Configuration for ns log action resource.
*/
type Auditnslogaction struct {
	/**
	* Name of the nslog action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the nslog action is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my nslog action" or 'my nslog action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* IP address of the nslog server.
	*/
	Serverip string `json:"serverip,omitempty"`
	/**
	* Auditserver name as a FQDN. Mutually exclusive with serverIP
	*/
	Serverdomainname string `json:"serverdomainname,omitempty"`
	/**
	* Time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the audit server if the last query failed.
	*/
	Domainresolveretry *int `json:"domainresolveretry,omitempty"`
	/**
	* Port on which the nslog server accepts connections.
	*/
	Serverport *int `json:"serverport,omitempty"`
	/**
	* Audit log level, which specifies the types of events to log.
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
	* Format of dates in the logs.
		Supported formats are:
		* MMDDYYYY - U.S. style month/date/year format.
		* DDMMYYYY - European style date/month/year format.
		* YYYYMMDD - ISO style year/month/date format.
	*/
	Dateformat string `json:"dateformat,omitempty"`
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
		* GMT_TIME. Coordinated Universal Time.
		* LOCAL_TIME. The server's timezone setting.
	*/
	Timezone string `json:"timezone,omitempty"`
	/**
	* Log user-configurable log messages to nslog.
		Setting this parameter to NO causes auditing to ignore all user-configured message actions. Setting this parameter to YES causes auditing to log user-configured message actions that meet the other logging criteria.
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
	* Log SSL Interception event information
	*/
	Sslinterception string `json:"sslinterception,omitempty"`
	/**
	* Log URL filtering event information
	*/
	Urlfiltering string `json:"urlfiltering,omitempty"`
	/**
	* Log Content Inspection event information
	*/
	Contentinspectionlog string `json:"contentinspectionlog,omitempty"`
	/**
	* Log protocol violations
	*/
	Protocolviolations string `json:"protocolviolations,omitempty"`
	/**
	* Immediately send a DNS query to resolve the server's domain name.
	*/
	Domainresolvenow bool `json:"domainresolvenow,omitempty"`

	//------- Read only Parameter ---------;

	Ip string `json:"ip,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
