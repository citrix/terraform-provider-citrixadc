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

package ssl

/**
* Configuration for Certificate Revocation List resource.
*/
type Sslcrl struct {
	/**
	* Name for the Certificate Revocation List (CRL). Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the CRL is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my crl" or 'my crl').
	*/
	Crlname string `json:"crlname,omitempty"`
	/**
	* Path to the CRL file. /var/netscaler/ssl/ is the default path.
	*/
	Crlpath string `json:"crlpath,omitempty"`
	/**
	* Input format of the CRL file. The two formats supported on the appliance are:
		PEM - Privacy Enhanced Mail.
		DER - Distinguished Encoding Rule.
	*/
	Inform string `json:"inform,omitempty"`
	/**
	* Set CRL auto refresh.
	*/
	Refresh string `json:"refresh,omitempty"`
	/**
	* CA certificate that has issued the CRL. Required if CRL Auto Refresh is selected. Install the CA certificate on the appliance before adding the CRL.
	*/
	Cacert string `json:"cacert,omitempty"`
	/**
	* Method for CRL refresh. If LDAP is selected, specify the method, CA certificate, base DN, port, and LDAP server name. If HTTP is selected, specify the CA certificate, method, URL, and port. Cannot be changed after a CRL is added.
	*/
	Method string `json:"method,omitempty"`
	/**
	* IP address of the LDAP server from which to fetch the CRLs.
	*/
	Server string `json:"server,omitempty"`
	/**
	* URL of the CRL distribution point.
	*/
	Url string `json:"url,omitempty"`
	/**
	* Port for the LDAP server.
	*/
	Port int `json:"port,omitempty"`
	/**
	* Base distinguished name (DN), which is used in an LDAP search to search for a CRL. Citrix recommends searching for the Base DN instead of the Issuer Name from the CA certificate, because the Issuer Name field might not exactly match the LDAP directory structure's DN.
	*/
	Basedn string `json:"basedn,omitempty"`
	/**
	* Extent of the search operation on the LDAP server. Available settings function as follows:
		One - One level below Base DN.
		Base - Exactly the same level as Base DN.
	*/
	Scope string `json:"scope,omitempty"`
	/**
	* CRL refresh interval. Use the NONE setting to unset this parameter.
	*/
	Interval string `json:"interval,omitempty"`
	/**
	* Day on which to refresh the CRL, or, if the Interval parameter is not set, the number of days after which to refresh the CRL. If Interval is set to MONTHLY, specify the date. If Interval is set to WEEKLY, specify the day of the week (for example, Sun=0 and Sat=6). This parameter is not applicable if the Interval is set to DAILY.
	*/
	Day int `json:"day,omitempty"`
	/**
	* Time, in hours (1-24) and minutes (1-60), at which to refresh the CRL.
	*/
	Time string `json:"time,omitempty"`
	/**
	* Bind distinguished name (DN) to be used to access the CRL object in the LDAP repository if access to the LDAP repository is restricted or anonymous access is not allowed.
	*/
	Binddn string `json:"binddn,omitempty"`
	/**
	* Password to access the CRL in the LDAP repository if access to the LDAP repository is restricted or anonymous access is not allowed.
	*/
	Password string `json:"password,omitempty"`
	/**
	* Set the LDAP-based CRL retrieval mode to binary.
	*/
	Binary string `json:"binary,omitempty"`
	/**
	* Name of and, optionally, path to the CA certificate file.
		/nsconfig/ssl/ is the default path.
	*/
	Cacertfile string `json:"cacertfile,omitempty"`
	/**
	* Name of and, optionally, path to the CA key file. /nsconfig/ssl/ is the default path
	*/
	Cakeyfile string `json:"cakeyfile,omitempty"`
	/**
	* Name of and, optionally, path to the file containing the serial numbers of all the certificates that are revoked. Revoked certificates are appended to the file. /nsconfig/ssl/ is the default path
	*/
	Indexfile string `json:"indexfile,omitempty"`
	/**
	* Name of and, optionally, path to the certificate to be revoked. /nsconfig/ssl/ is the default path.
	*/
	Revoke string `json:"revoke,omitempty"`
	/**
	* Name of and, optionally, path to the CRL file to be generated. The list of certificates that have been revoked is obtained from the index file. /nsconfig/ssl/ is the default path.
	*/
	Gencrl string `json:"gencrl,omitempty"`

	//------- Read only Parameter ---------;

	Flags string `json:"flags,omitempty"`
	Lastupdatetime string `json:"lastupdatetime,omitempty"`
	Version string `json:"version,omitempty"`
	Signaturealgo string `json:"signaturealgo,omitempty"`
	Issuer string `json:"issuer,omitempty"`
	Lastupdate string `json:"lastupdate,omitempty"`
	Nextupdate string `json:"nextupdate,omitempty"`
	Daystoexpiration string `json:"daystoexpiration,omitempty"`

}
