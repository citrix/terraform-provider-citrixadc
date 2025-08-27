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
* Configuration for LB group resource.
*/
type Lbgroup struct {
	/**
	* Name of the load balancing virtual server group.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Type of persistence for the group. Available settings function as follows:
		* SOURCEIP - Create persistence sessions based on the client IP.
		* COOKIEINSERT - Create persistence sessions based on a cookie in client requests. The cookie is inserted by a Set-Cookie directive from the server, in its first response to a client.
		* RULE - Create persistence sessions based on a user defined rule.
		* NONE - Disable persistence for the group.
	*/
	Persistencetype string `json:"persistencetype,omitempty"`
	/**
	* Type of backup persistence for the group.
	*/
	Persistencebackup string `json:"persistencebackup,omitempty"`
	/**
	* Time period, in minutes, for which backup persistence is in effect.
	*/
	Backuppersistencetimeout int `json:"backuppersistencetimeout,omitempty"`
	/**
	* Persistence mask to apply to source IPv4 addresses when creating source IP based persistence sessions.
	*/
	Persistmask string `json:"persistmask,omitempty"`
	/**
	* Use this parameter to specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.
	*/
	Cookiename string `json:"cookiename,omitempty"`
	/**
	* Persistence mask to apply to source IPv6 addresses when creating source IP based persistence sessions.
	*/
	V6persistmasklen int `json:"v6persistmasklen,omitempty"`
	/**
	* Domain attribute for the HTTP cookie.
	*/
	Cookiedomain string `json:"cookiedomain,omitempty"`
	/**
	* Time period for which a persistence session is in effect.
	*/
	Timeout int `json:"timeout,omitempty"`
	/**
	* Expression, or name of a named expression, against which traffic is evaluated.
		The following requirements apply only to the Citrix ADC CLI:
		* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.
		* If the expression itself includes double quotation marks, escape the quotations by using the \ character.
		* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* When USE_VSERVER_PERSISTENCE is enabled, one can use this setting to designate a member vserver as master which is responsible to create the persistence sessions
	*/
	Mastervserver string `json:"mastervserver,omitempty"`
	/**
	* Use this parameter to enable vserver level persistence on group members. This allows member vservers to have their own persistence, but need to be compatible with other members persistence rules. When this setting is enabled persistence sessions created by any of the members can be shared by other member vservers.
	*/
	Usevserverpersistency string `json:"usevserverpersistency,omitempty"`
	/**
	* New name for the load balancing virtual server group.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Td string `json:"td,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
