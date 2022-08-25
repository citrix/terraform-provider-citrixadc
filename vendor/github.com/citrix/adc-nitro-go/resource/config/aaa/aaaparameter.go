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

package aaa

/**
* Configuration for AAA parameter resource.
*/
type Aaaparameter struct {
	/**
	* The default state of VPN Static Page caching. If nothing is specified, the default value is set to YES.
	*/
	Enablestaticpagecaching string `json:"enablestaticpagecaching,omitempty"`
	/**
	* Enhanced auth feedback provides more information to the end user about the reason for an authentication failure.  The default value is set to NO.
	*/
	Enableenhancedauthfeedback string `json:"enableenhancedauthfeedback,omitempty"`
	/**
	* The default authentication server type.
	*/
	Defaultauthtype string `json:"defaultauthtype,omitempty"`
	/**
	* Maximum number of concurrent users allowed to log on to VPN simultaneously.
	*/
	Maxaaausers int `json:"maxaaausers,omitempty"`
	/**
	* Maximum Number of login Attempts
	*/
	Maxloginattempts int `json:"maxloginattempts,omitempty"`
	/**
	* Number of minutes an account will be locked if user exceeds maximum permissible attempts
	*/
	Failedlogintimeout int `json:"failedlogintimeout,omitempty"`
	/**
	* Source IP address to use for traffic that is sent to the authentication server.
	*/
	Aaadnatip string `json:"aaadnatip,omitempty"`
	/**
	* Enables/Disables stickiness to authentication servers
	*/
	Enablesessionstickiness string `json:"enablesessionstickiness,omitempty"`
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
	Aaasessionloglevel string `json:"aaasessionloglevel,omitempty"`
	/**
	* AAAD log level, which specifies the types of AAAD events to log in nsvpn.log. 
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
	Aaadloglevel string `json:"aaadloglevel,omitempty"`
	/**
	* Set by the DHCP client when the IP address was fetched dynamically.
	*/
	Dynaddr string `json:"dynaddr,omitempty"`
	/**
	* First time user mode determines which configuration options are shown by default when logging in to the GUI. This setting is controlled by the GUI.
	*/
	Ftmode string `json:"ftmode,omitempty"`
	/**
	* This will set the maximum deflate size in case of SAML Redirect binding.
	*/
	Maxsamldeflatesize int `json:"maxsamldeflatesize,omitempty"`
	/**
	* Persistent storage of unsuccessful user login attempts
	*/
	Persistentloginattempts string `json:"persistentloginattempts,omitempty"`
	/**
	* This will set the threshold time in days for password expiry notification. Default value is 0, which means no notification is sent
	*/
	Pwdexpirynotificationdays int `json:"pwdexpirynotificationdays,omitempty"`
	/**
	* This will set maximum number of Questions to be asked for KB Validation. Default value is 2, Max Value is 6
	*/
	Maxkbquestions int `json:"maxkbquestions,omitempty"`
	/**
	* Parameter to encrypt login information for nFactor flow
	*/
	Loginencryption string `json:"loginencryption,omitempty"`
	/**
	* SameSite attribute value for Cookies generated in AAATM context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite
	*/
	Samesite string `json:"samesite,omitempty"`
	/**
	* Option to enable/disable API cache feature.
	*/
	Apitokencache string `json:"apitokencache,omitempty"`
	/**
	* Frequency at which a token must be verified at the Authorization Server (AS) despite being found in cache.
	*/
	Tokenintrospectioninterval int `json:"tokenintrospectioninterval,omitempty"`
	/**
	* Parameter to enable/disable default CSP header
	*/
	Defaultcspheader string `json:"defaultcspheader,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
