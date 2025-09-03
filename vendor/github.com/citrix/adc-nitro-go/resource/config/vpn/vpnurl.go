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
* Configuration for VPN URL resource.
*/
type Vpnurl struct {
	/**
	* Name of the bookmark link.
	*/
	Urlname string `json:"urlname,omitempty"`
	/**
	* Description of the bookmark link. The description appears in the Access Interface.
	*/
	Linkname string `json:"linkname,omitempty"`
	/**
	* Web address for the bookmark link.
	*/
	Actualurl string `json:"actualurl,omitempty"`
	/**
	* Name of the associated LB/CS vserver
	*/
	Vservername string `json:"vservername,omitempty"`
	/**
	* If clientless access to the resource hosting the link is allowed, also use clientless access for the bookmarked web address in the Secure Client Access based session. Allows single sign-on and other HTTP processing on Citrix Gateway for HTTPS resources.
	*/
	Clientlessaccess string `json:"clientlessaccess,omitempty"`
	/**
	* Any comments associated with the bookmark link.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* URL to fetch icon file for displaying this resource.
	*/
	Iconurl string `json:"iconurl,omitempty"`
	/**
	* Single sign on type for unified gateway
	*/
	Ssotype string `json:"ssotype,omitempty"`
	/**
	* The type of application this VPN URL represents. Possible values are CVPN/SaaS/VPN
	*/
	Applicationtype string `json:"applicationtype,omitempty"`
	/**
	* Profile to be used for doing SAML SSO
	*/
	Samlssoprofile string `json:"samlssoprofile,omitempty"`
	/**
	* To store the template details in the json format.
	*/
	Appjson string `json:"appjson,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
