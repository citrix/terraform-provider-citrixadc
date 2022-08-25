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
* Configuration for Kerberos constrained delegation account resource.
*/
type Aaakcdaccount struct {
	/**
	* The name of the KCD account.
	*/
	Kcdaccount string `json:"kcdaccount,omitempty"`
	/**
	* The path to the keytab file. If specified other parameters in this command need not be given
	*/
	Keytab string `json:"keytab,omitempty"`
	/**
	* Kerberos Realm.
	*/
	Realmstr string `json:"realmstr,omitempty"`
	/**
	* Username that can perform kerberos constrained delegation.
	*/
	Delegateduser string `json:"delegateduser,omitempty"`
	/**
	* Password for Delegated User.
	*/
	Kcdpassword string `json:"kcdpassword,omitempty"`
	/**
	* SSL Cert (including private key) for Delegated User.
	*/
	Usercert string `json:"usercert,omitempty"`
	/**
	* CA Cert for UserCert or when doing PKINIT backchannel.
	*/
	Cacert string `json:"cacert,omitempty"`
	/**
	* Realm of the user
	*/
	Userrealm string `json:"userrealm,omitempty"`
	/**
	* Enterprise Realm of the user. This should be given only in certain KDC deployments where KDC expects Enterprise username instead of Principal Name
	*/
	Enterpriserealm string `json:"enterpriserealm,omitempty"`
	/**
	* Service SPN. When specified, this will be used to fetch kerberos tickets. If not specified, Citrix ADC will construct SPN using service fqdn
	*/
	Servicespn string `json:"servicespn,omitempty"`

	//------- Read only Parameter ---------;

	Principle string `json:"principle,omitempty"`
	Kcdspn string `json:"kcdspn,omitempty"`

}
