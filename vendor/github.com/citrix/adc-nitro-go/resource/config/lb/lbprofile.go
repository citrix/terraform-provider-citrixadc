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

package lb

/**
* Configuration for LB profile resource.
*/
type Lbprofile struct {
	/**
	* Name of the LB profile.
	*/
	Lbprofilename string `json:"lbprofilename,omitempty"`
	/**
	* Enable database specific load balancing for MySQL and MSSQL service types.
	*/
	Dbslb string `json:"dbslb,omitempty"`
	/**
	* By turning on this option packets destined to a vserver in a cluster will not under go any steering. Turn this option for single pa
		cket request response mode or when the upstream device is performing a proper RSS for connection based distribution.
	*/
	Processlocal string `json:"processlocal,omitempty"`
	/**
	* Include the HttpOnly attribute in persistence cookies. The HttpOnly attribute limits the scope of a cookie to HTTP requests and helps mitigate the risk of cross-site scripting attacks.
	*/
	Httponlycookieflag string `json:"httponlycookieflag,omitempty"`
	/**
	* Use this parameter to specify the passphrase used to generate secured persistence cookie value. It specifies the passphrase with a maximum of 31 characters.
	*/
	Cookiepassphrase string `json:"cookiepassphrase,omitempty"`
	/**
	* Encode persistence cookie values using SHA2 hash.
	*/
	Usesecuredpersistencecookie string `json:"usesecuredpersistencecookie,omitempty"`
	/**
	* Encode persistence cookie values using SHA2 hash.
	*/
	Useencryptedpersistencecookie string `json:"useencryptedpersistencecookie,omitempty"`
	/**
	* String configured as LiteralADCCookieAttribute will be appended as attribute for Citrix ADC cookie (for example: LB cookie persistence , GSLB site persistence, CS cookie persistence, LB group cookie persistence).
		Sample usage -
		add lb profile lbprof -LiteralADCCookieAttribute ";SameSite=None"
	*/
	Literaladccookieattribute string `json:"literaladccookieattribute,omitempty"`
	/**
	* ComputedADCCookieAttribute accepts ns variable as input in form of string starting with $ (to understand how to configure ns variable, please check man add ns variable). policies can be configured to modify this variable for every transaction and the final value of the variable after policy evaluation will be appended as attribute to Citrix ADC cookie (for example: LB cookie persistence , GSLB sitepersistence, CS cookie persistence, LB group cookie persistence). Only one of ComputedADCCookieAttribute, LiteralADCCookieAttribute can be set.
		Sample usage -
		add ns variable lbvar -type TEXT(100) -scope Transaction
		add ns assignment lbassign -variable $lbvar -set "\\";SameSite=Strict\\""
		add rewrite policy lbpol <valid policy expression> lbassign
		bind rewrite global lbpol 100 next -type RES_OVERRIDE
		add lb profile lbprof -ComputedADCCookieAttribute "$lbvar"
		For incoming client request, if above policy evaluates TRUE, then SameSite=Strict will be appended to ADC generated cookie
	*/
	Computedadccookieattribute string `json:"computedadccookieattribute,omitempty"`
	/**
	* This option allows to store the MQTT clientid and username in transactional logs
	*/
	Storemqttclientidandusername string `json:"storemqttclientidandusername,omitempty"`
	/**
	* This option dictates the hashing algorithm used for hash based LB methods (URLHASH, DOMAINHASH, SOURCEIPHASH, DESTINATIONIPHASH, SRCIPDESTIPHASH, SRCIPSRCPORTHASH, TOKEN, USER_TOKEN, CALLIDHASH).
	*/
	Lbhashalgorithm string `json:"lbhashalgorithm,omitempty"`
	/**
	* This option is used to specify the number of fingers to be used in PRAC and JARH algorithms for hash based LB methods. Increasing the number of fingers might give better distribution of traffic at the expense of additional memory.
	*/
	Lbhashfingers int `json:"lbhashfingers,omitempty"`

	//------- Read only Parameter ---------;

	Vsvrcount string `json:"vsvrcount,omitempty"`
	Adccookieattributewarningmsg string `json:"adccookieattributewarningmsg,omitempty"`
	Lbhashalgowinsize string `json:"lbhashalgowinsize,omitempty"`

}
