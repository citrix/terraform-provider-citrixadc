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

package authentication

/**
* Configuration for authentication virtual server resource.
*/
type Authenticationvserver struct {
	/**
	* Name for the new authentication virtual server.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the authentication virtual server is added by using the rename authentication vserver command.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Protocol type of the authentication virtual server. Always SSL.
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* IP address of the authentication virtual server, if a single IP address is assigned to the virtual server.
	*/
	Ipv46 string `json:"ipv46,omitempty"`
	/**
	* If you are creating a series of virtual servers with a range of IP addresses assigned to them, the length of the range.
		The new range of authentication virtual servers will have IP addresses consecutively numbered, starting with the primary address specified with the IP Address parameter.
	*/
	Range int `json:"range,omitempty"`
	/**
	* TCP port on which the virtual server accepts connections.
	*/
	Port int `json:"port,omitempty"`
	/**
	* Initial state of the new virtual server.
	*/
	State string `json:"state,omitempty"`
	/**
	* Require users to be authenticated before sending traffic through this virtual server.
	*/
	Authentication string `json:"authentication,omitempty"`
	/**
	* The domain of the authentication cookie set by Authentication vserver
	*/
	Authenticationdomain string `json:"authenticationdomain,omitempty"`
	/**
	* Any comments associated with this virtual server.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td int `json:"td,omitempty"`
	/**
	* Log AppFlow flow information.
	*/
	Appflowlog string `json:"appflowlog,omitempty"`
	/**
	* Maximum Number of login Attempts
	*/
	Maxloginattempts int `json:"maxloginattempts,omitempty"`
	/**
	* Number of minutes an account will be locked if user exceeds maximum permissible attempts
	*/
	Failedlogintimeout int `json:"failedlogintimeout,omitempty"`
	/**
	* Name of the certificate key that was bound to the corresponding SSL virtual server as the Certificate Authority for the device certificate
	*/
	Certkeynames string `json:"certkeynames,omitempty"`
	/**
	* SameSite attribute value for Cookies generated in AAATM context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite
	*/
	Samesite string `json:"samesite,omitempty"`
	/**
	* New name of the authentication virtual server.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, 'my authentication policy' or "my authentication policy").
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Ip string `json:"ip,omitempty"`
	Value string `json:"value,omitempty"`
	Type string `json:"type,omitempty"`
	Curstate string `json:"curstate,omitempty"`
	Status string `json:"status,omitempty"`
	Cachetype string `json:"cachetype,omitempty"`
	Redirect string `json:"redirect,omitempty"`
	Precedence string `json:"precedence,omitempty"`
	Redirecturl string `json:"redirecturl,omitempty"`
	Curaaausers string `json:"curaaausers,omitempty"`
	Policy string `json:"policy,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Weight string `json:"weight,omitempty"`
	Cachevserver string `json:"cachevserver,omitempty"`
	Backupvserver string `json:"backupvserver,omitempty"`
	Clttimeout string `json:"clttimeout,omitempty"`
	Somethod string `json:"somethod,omitempty"`
	Sothreshold string `json:"sothreshold,omitempty"`
	Sopersistence string `json:"sopersistence,omitempty"`
	Sopersistencetimeout string `json:"sopersistencetimeout,omitempty"`
	Priority string `json:"priority,omitempty"`
	Downstateflush string `json:"downstateflush,omitempty"`
	Bindpoint string `json:"bindpoint,omitempty"`
	Disableprimaryondown string `json:"disableprimaryondown,omitempty"`
	Listenpolicy string `json:"listenpolicy,omitempty"`
	Listenpriority string `json:"listenpriority,omitempty"`
	Tcpprofilename string `json:"tcpprofilename,omitempty"`
	Httpprofilename string `json:"httpprofilename,omitempty"`
	Vstype string `json:"vstype,omitempty"`
	Ngname string `json:"ngname,omitempty"`
	Secondary string `json:"secondary,omitempty"`
	Groupextraction string `json:"groupextraction,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
