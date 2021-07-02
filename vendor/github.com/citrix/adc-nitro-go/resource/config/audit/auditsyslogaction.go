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
* Configuration for system log action resource.
*/
type Auditsyslogaction struct {
	/**
	* Name of the syslog action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the syslog action is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my syslog action" or 'my syslog action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* IP address of the syslog server.
	*/
	Serverip string `json:"serverip,omitempty"`
	/**
	* SYSLOG server name as a FQDN.  Mutually exclusive with serverIP/lbVserverName
	*/
	Serverdomainname string `json:"serverdomainname,omitempty"`
	/**
	* Time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the syslog server if the last query failed.
	*/
	Domainresolveretry int32 `json:"domainresolveretry,omitempty"`
	/**
	* Name of the LB vserver. Mutually exclusive with syslog serverIP/serverName
	*/
	Lbvservername string `json:"lbvservername,omitempty"`
	/**
	* Port on which the syslog server accepts connections.
	*/
	Serverport int32 `json:"serverport,omitempty"`
	/**
	* Audit log level, which specifies the types of events to log. 
		Available values function as follows: 
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
		* MMDDYYYY. -U.S. style month/date/year format.
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
		Supported settings are: 
		* GMT_TIME. Coordinated Universal time.
		* LOCAL_TIME. Use the server's timezone setting.
	*/
	Timezone string `json:"timezone,omitempty"`
	/**
	* Log user-configurable log messages to syslog. 
		Setting this parameter to NO causes auditing to ignore all user-configured message actions. Setting this parameter to YES causes auditing to log user-configured message actions that meet the other logging criteria.
	*/
	Userdefinedauditlog string `json:"userdefinedauditlog,omitempty"`
	/**
	* Export log messages to AppFlow collectors.
		Appflow collectors are entities to which log messages can be sent so that some action can be performed on them.
	*/
	Appflowexport string `json:"appflowexport,omitempty"`
	/**
	* Log lsn info
	*/
	Lsn string `json:"lsn,omitempty"`
	/**
	* Log alg info
	*/
	Alg string `json:"alg,omitempty"`
	/**
	* Log subscriber session event information
	*/
	Subscriberlog string `json:"subscriberlog,omitempty"`
	/**
	* Transport type used to send auditlogs to syslog server. Default type is UDP.
	*/
	Transport string `json:"transport,omitempty"`
	/**
	* Name of the TCP profile whose settings are to be applied to the audit server info to tune the TCP connection parameters.
	*/
	Tcpprofilename string `json:"tcpprofilename,omitempty"`
	/**
	* Max size of log data that can be held in NSB chain of server info.
	*/
	Maxlogdatasizetohold uint32 `json:"maxlogdatasizetohold,omitempty"`
	/**
	* Log DNS related syslog messages
	*/
	Dns string `json:"dns,omitempty"`
	/**
	* Log Content Inspection event information
	*/
	Contentinspectionlog string `json:"contentinspectionlog,omitempty"`
	/**
	* Name of the network profile.
		The SNIP configured in the network profile will be used as source IP while sending log messages.
	*/
	Netprofile string `json:"netprofile,omitempty"`
	/**
	* Log SSL Interception event information
	*/
	Sslinterception string `json:"sslinterception,omitempty"`
	/**
	* Log URL filtering event information
	*/
	Urlfiltering string `json:"urlfiltering,omitempty"`
	/**
	* Immediately send a DNS query to resolve the server's domain name.
	*/
	Domainresolvenow bool `json:"domainresolvenow,omitempty"`

	//------- Read only Parameter ---------;

	Ip string `json:"ip,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
