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

package system

/**
* Configuration for system parameter resource.
*/
type Systemparameter struct {
	/**
	* Enable or disable Role-Based Authentication (RBA) on responses.
	*/
	Rbaonresponse string `json:"rbaonresponse,omitempty"`
	/**
	* String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (_), and the following variables:
		* %u - Will be replaced by the user name.
		* %h - Will be replaced by the hostname of the Citrix ADC.
		* %t - Will be replaced by the current time in 12-hour format.
		* %T - Will be replaced by the current time in 24-hour format.
		* %d - Will be replaced by the current date.
		* %s - Will be replaced by the state of the Citrix ADC.
		Note: The 63-character limit for the length of the string does not apply to the characters that replace the variables.
	*/
	Promptstring string `json:"promptstring,omitempty"`
	/**
	* Flush the system if the number of Network Address Translation Protocol Control Blocks (NATPCBs) exceeds this value.
	*/
	Natpcbforceflushlimit int `json:"natpcbforceflushlimit,omitempty"`
	/**
	* Send a reset signal to client and server connections when their NATPCBs time out. Avoids the buildup of idle TCP connections on both the sides.
	*/
	Natpcbrstontimeout string `json:"natpcbrstontimeout,omitempty"`
	/**
	* CLI session inactivity timeout, in seconds. If Restrictedtimeout argument is enabled, Timeout can have values in the range [300-86400] seconds.
		If Restrictedtimeout argument is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.
	*/
	Timeout int `json:"timeout,omitempty"`
	/**
	* When enabled, local users can access Citrix ADC even when external authentication is configured. When disabled, local users are not allowed to access the Citrix ADC, Local users can access the Citrix ADC only when the configured external authentication servers are unavailable. This parameter is not applicable to SSH Key-based authentication
	*/
	Localauth string `json:"localauth,omitempty"`
	/**
	* Minimum length of system user password. When strong password is enabled default minimum length is 8. User entered value can be greater than or equal to 8. Default mininum value is 1 when strong password is disabled. Maximum value is 127 in both cases.
	*/
	Minpasswordlen int `json:"minpasswordlen,omitempty"`
	/**
	* After enabling strong password (enableall / enablelocal - not included in exclude list), all the passwords / sensitive information must have - Atleast 1 Lower case character, Atleast 1 Upper case character, Atleast 1 numeric character, Atleast 1 special character ( ~, `, !, @, #, $, %, ^, &, *, -, _, =, +, {, }, [, ], |, \, :, <, >, /, ., ,, " "). Exclude list in case of enablelocal is - NS_FIPS, NS_CRL, NS_RSAKEY, NS_PKCS12, NS_PKCS8, NS_LDAP, NS_TACACS, NS_TACACSACTION, NS_RADIUS, NS_RADIUSACTION, NS_ENCRYPTION_PARAMS. So no Strong Password checks will be performed on these ObjectType commands for enablelocal case.
	*/
	Strongpassword string `json:"strongpassword,omitempty"`
	/**
	* Enable/Disable the restricted timeout behaviour. When enabled, timeout cannot be configured beyond admin configured timeout  and also it will have the [minimum - maximum] range check. When disabled, timeout will have the old behaviour. By default the value is disabled
	*/
	Restrictedtimeout string `json:"restrictedtimeout,omitempty"`
	/**
	* Use this option to set the FIPS mode for key user-land processes. When enabled, these user-land processes will operate in FIPS mode. In this mode, these processes will use FIPS 140-2 certified crypto algorithms.
		With a FIPS license, it is enabled by default and cannot be disabled.
		Without a FIPS license, it is disabled by default, wherein these user-land processes will not operate in FIPS mode.
	*/
	Fipsusermode string `json:"fipsusermode,omitempty"`
	/**
	* Enable or disable Doppler
	*/
	Doppler string `json:"doppler,omitempty"`
	/**
	* Enable or disable Google analytics
	*/
	Googleanalytics string `json:"googleanalytics,omitempty"`
	/**
	* Total time a request can take for authentication/authorization
	*/
	Totalauthtimeout int `json:"totalauthtimeout,omitempty"`
	/**
	* Audit log level, which specifies the types of events to log for cli executed commands.
		Available values function as follows:
		* EMERGENCY - Events that indicate an immediate crisis on the server.
		* ALERT - Events that might require action.
		* CRITICAL - Events that indicate an imminent server crisis.
		* ERROR - Events that indicate some type of error.
		* WARNING - Events that require action in the near future.
		* NOTICE - Events that the administrator should know about.
		* INFORMATIONAL - All but low-level events.
		* DEBUG - All events, in extreme detail.
	*/
	Cliloglevel string `json:"cliloglevel,omitempty"`
	/**
	* Enable or disable force password change for nsroot user
	*/
	Forcepasswordchange string `json:"forcepasswordchange,omitempty"`
	/**
	* Enable or disable basic authentication for Nitro API.
	*/
	Basicauth string `json:"basicauth,omitempty"`
	/**
	* Enable or disable External user reauthentication when authentication parameter changes
	*/
	Reauthonauthparamchange string `json:"reauthonauthparamchange,omitempty"`
	/**
	* Use this option to remove the sensitive files from the system like authorise keys, public keys etc. The commands which will remove sensitive files when this system paramter is enabled are rm cluster instance, rm cluster node, rm ha node, clear config full, join cluster and add cluster instance.
	*/
	Removesensitivefiles string `json:"removesensitivefiles,omitempty"`
	/**
	* Maximum number of client connection allowed per user.The maxsessionperuser value ranges from 1 to 40
	*/
	Maxsessionperuser int `json:"maxsessionperuser,omitempty"`
	/**
	* Configure WAF protection for endpoints used by NetScaler management interfaces. The available options are:
		* DEFAULT - NetScaler decides which endpoints have WAF protection enabled.
		* GUI - Endpoints used by the Management GUI Interface are WAF protected.
		* DISABLED - WAF protection is disabled.
	*/
	Wafprotection []string `json:"wafprotection,omitempty"`
	/**
	* Enables or disable password expiry feature for system users.
		If the feature is ENABLED, by default the last 6 passwords of users will be maintained and will not be allowed to reuse same.
		When the feature is enabled the daystoexpire, warnpriorndays and pwdhistoryCount will be set with default values. The values can only be set in system
		for system parameter. It cannot be unset. It is possible to set and unset the values for daytoexpire and warnpriorndays in system groups.
		Default values if feature is ENABLED:
		daystoexpire: 30
		warnpriorndays: 5
		pwdhistoryCount: 6
		If the feature is DISABLED the values cannot be set or unset in system parameter and system groups
	*/
	Passwordhistorycontrol string `json:"passwordhistorycontrol,omitempty"`
	/**
	* Password expiry days for all the system users. The daystoexpire value ranges from 30 to 255.
	*/
	Daystoexpire int `json:"daystoexpire,omitempty"`
	/**
	* Number of days before which password expiration warning would be thrown with respect to daystoexpire. The warnpriorndays value ranges from 5 to 40.
	*/
	Warnpriorndays int `json:"warnpriorndays,omitempty"`
	/**
	* Number of passwords to be maintained as history for system users. The pwdhistorycount value ranges from 1 to 10.
	*/
	Pwdhistorycount int `json:"pwdhistorycount,omitempty"`

	//------- Read only Parameter ---------;

	Maxclient string `json:"maxclient,omitempty"`
	Allowdefaultpartition string `json:"allowdefaultpartition,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
