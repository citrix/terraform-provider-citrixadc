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
* Configuration for LDAP parameter resource.
*/
type Aaaldapparams struct {
	/**
	* IP address of your LDAP server.
	*/
	Serverip string `json:"serverip,omitempty"`
	/**
	* Port number on which the LDAP server listens for connections.
	*/
	Serverport int `json:"serverport,omitempty"`
	/**
	* Maximum number of seconds that the Citrix ADC waits for a response from the LDAP server.
	*/
	Authtimeout int `json:"authtimeout,omitempty"`
	/**
	* Base (the server and location) from which LDAP search commands should start.
		If the LDAP server is running locally, the default value of base is dc=netscaler, dc=com.
	*/
	Ldapbase string `json:"ldapbase,omitempty"`
	/**
	* Complete distinguished name (DN) string used for binding to the LDAP server.
	*/
	Ldapbinddn string `json:"ldapbinddn,omitempty"`
	/**
	* Password for binding to the LDAP server.
	*/
	Ldapbinddnpassword string `json:"ldapbinddnpassword,omitempty"`
	/**
	* Name attribute that the Citrix ADC uses to query the external LDAP server or an Active Directory.
	*/
	Ldaploginname string `json:"ldaploginname,omitempty"`
	/**
	* String to be combined with the default LDAP user search string to form the value to use when executing an LDAP search.
		For example, the following values:
		vpnallowed=true,
		ldaploginame=""samaccount""
		when combined with the user-supplied username ""bob"", yield the following LDAP search string:
		""(&(vpnallowed=true)(samaccount=bob)""
	*/
	Searchfilter string `json:"searchfilter,omitempty"`
	/**
	* Attribute name used for group extraction from the LDAP server.
	*/
	Groupattrname string `json:"groupattrname,omitempty"`
	/**
	* Subattribute name used for group extraction from the LDAP server.
	*/
	Subattributename string `json:"subattributename,omitempty"`
	/**
	* Type of security used for communications between the Citrix ADC and the LDAP server. For the PLAINTEXT setting, no encryption is required.
	*/
	Sectype string `json:"sectype,omitempty"`
	/**
	* The type of LDAP server.
	*/
	Svrtype string `json:"svrtype,omitempty"`
	/**
	* Attribute used by the Citrix ADC to query an external LDAP server or Active Directory for an alternative username.
		This alternative username is then used for single sign-on (SSO).
	*/
	Ssonameattribute string `json:"ssonameattribute,omitempty"`
	/**
	* Accept password change requests.
	*/
	Passwdchange string `json:"passwdchange,omitempty"`
	/**
	* Queries the external LDAP server to determine whether the specified group belongs to another group.
	*/
	Nestedgroupextraction string `json:"nestedgroupextraction,omitempty"`
	/**
	* Number of levels up to which the system can query nested LDAP groups.
	*/
	Maxnestinglevel int `json:"maxnestinglevel,omitempty"`
	/**
	* LDAP-group attribute that uniquely identifies the group. No two groups on one LDAP server can have the same group name identifier.
	*/
	Groupnameidentifier string `json:"groupnameidentifier,omitempty"`
	/**
	* LDAP-group attribute that designates the parent group of the specified group. Use this attribute to search for a group's parent group.
	*/
	Groupsearchattribute string `json:"groupsearchattribute,omitempty"`
	/**
	* LDAP-group subattribute that designates the parent group of the specified group. Use this attribute to search for a group's parent group.
	*/
	Groupsearchsubattribute string `json:"groupsearchsubattribute,omitempty"`
	/**
	* Search-expression that can be specified for sending group-search requests to the LDAP server.
	*/
	Groupsearchfilter string `json:"groupsearchfilter,omitempty"`
	/**
	* This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
	*/
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`

	//------- Read only Parameter ---------;

	Groupauthname string `json:"groupauthname,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
