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

package cs

/**
* Binding class showing the domain that can be bound to csvserver.
*/
type Csvserverdomainbinding struct {
	/**
	* Domain name for which to change the time to live (TTL) and/or backup service IP address.
	*/
	Domainname string `json:"domainname,omitempty"`
	Ttl *int `json:"ttl,omitempty"`
	Backupip string `json:"backupip,omitempty"`
	Cookiedomain string `json:"cookiedomain,omitempty"`
	Cookietimeout *int `json:"cookietimeout,omitempty"`
	Sitedomainttl *int `json:"sitedomainttl,omitempty"`
	/**
	* Enable logging appflow flow information
	*/
	Appflowlog string `json:"appflowlog,omitempty"`
	/**
	* Name of the content switching virtual server to which the content switching policy applies.
	*/
	Name string `json:"name,omitempty"`


}