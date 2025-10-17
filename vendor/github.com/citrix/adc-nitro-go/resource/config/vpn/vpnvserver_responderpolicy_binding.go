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

package vpn

/**
* Binding class showing the responderpolicy that can be bound to vpnvserver.
*/
type Vpnvserverresponderpolicybinding struct {
	/**
	* The name of the policy, if any, bound to the VPN virtual server.
	*/
	Policy string `json:"policy,omitempty"`
	/**
	* Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.
	*/
	Priority *int `json:"priority,omitempty"`
	/**
	* Next priority expression.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Bindpoint to which the policy is bound.
	*/
	Bindpoint string `json:"bindpoint,omitempty"`
	/**
	* Name of the virtual server.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.
	*/
	Secondary bool `json:"secondary,omitempty"`
	/**
	* Binds the authentication policy to a tertiary chain which will be used only for group extraction.  The user will not authenticate against this server, and this will only be called if primary and/or secondary authentication has succeeded.
	*/
	Groupextraction bool `json:"groupextraction,omitempty"`


}