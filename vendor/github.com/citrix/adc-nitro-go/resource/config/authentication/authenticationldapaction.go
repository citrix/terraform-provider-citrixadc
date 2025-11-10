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
* Configuration for LDAP action resource.
*/
type Authenticationldapaction struct {
	/**
	* Name for the new LDAP action.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the LDAP action is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication action" or 'my authentication action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* IP address assigned to the LDAP server.
	*/
	Serverip string `json:"serverip,omitempty"`
	/**
	* LDAP server name as a FQDN.  Mutually exclusive with LDAP IP address.
	*/
	Servername string `json:"servername,omitempty"`
	/**
	* Port on which the LDAP server accepts connections.
	*/
	Serverport *int `json:"serverport,omitempty"`
	/**
	* Number of seconds the Citrix ADC waits for a response from the RADIUS server.
	*/
	Authtimeout *int `json:"authtimeout,omitempty"`
	/**
	* Base (node) from which to start LDAP searches.
		If the LDAP server is running locally, the default value of base is dc=netscaler, dc=com.
	*/
	Ldapbase string `json:"ldapbase,omitempty"`
	/**
	* Full distinguished name (DN) that is used to bind to the LDAP server.
		Default: cn=Manager,dc=netscaler,dc=com
	*/
	Ldapbinddn string `json:"ldapbinddn,omitempty"`
	/**
	* Password used to bind to the LDAP server.
	*/
	Ldapbinddnpassword string `json:"ldapbinddnpassword,omitempty"`
	/**
	* LDAP login name attribute.
		The Citrix ADC uses the LDAP login name to query external LDAP servers or Active Directories.
	*/
	Ldaploginname string `json:"ldaploginname,omitempty"`
	/**
	* String to be combined with the default LDAP user search string to form the search value. For example, if the search filter "vpnallowed=true" is combined with the LDAP login name "samaccount" and the user-supplied username is "bob", the result is the LDAP search string ""&(vpnallowed=true)(samaccount=bob)"" (Be sure to enclose the search string in two sets of double quotation marks; both sets are needed.).
	*/
	Searchfilter string `json:"searchfilter,omitempty"`
	/**
	* LDAP group attribute name.
		Used for group extraction on the LDAP server.
	*/
	Groupattrname string `json:"groupattrname,omitempty"`
	/**
	* LDAP group sub-attribute name.
		Used for group extraction from the LDAP server.
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
	* LDAP single signon (SSO) attribute.
		The Citrix ADC uses the SSO name attribute to query external LDAP servers or Active Directories for an alternate username.
	*/
	Ssonameattribute string `json:"ssonameattribute,omitempty"`
	/**
	* Perform LDAP authentication.
		If authentication is disabled, any LDAP authentication attempt returns authentication success if the user is found.
		CAUTION! Authentication should be disabled only for authorization group extraction or where other (non-LDAP) authentication methods are in use and either bound to a primary list or flagged as secondary.
	*/
	Authentication string `json:"authentication,omitempty"`
	/**
	* Require a successful user search for authentication.
		CAUTION!  This field should be set to NO only if usersearch not required [Both username validation as well as password validation skipped] and (non-LDAP) authentication methods are in use and either bound to a primary list or flagged as secondary.
	*/
	Requireuser string `json:"requireuser,omitempty"`
	/**
	* Allow password change requests.
	*/
	Passwdchange string `json:"passwdchange,omitempty"`
	/**
	* Allow nested group extraction, in which the Citrix ADC queries external LDAP servers to determine whether a group is part of another group.
	*/
	Nestedgroupextraction string `json:"nestedgroupextraction,omitempty"`
	/**
	* If nested group extraction is ON, specifies the number of levels up to which group extraction is performed.
	*/
	Maxnestinglevel *int `json:"maxnestinglevel,omitempty"`
	/**
	* Setting this option to ON enables following LDAP referrals received from the LDAP server.
	*/
	Followreferrals string `json:"followreferrals,omitempty"`
	/**
	* Specifies the maximum number of nested referrals to follow.
	*/
	Maxldapreferrals *int `json:"maxldapreferrals,omitempty"`
	/**
	* Specifies the DNS Record lookup Type for the referrals
	*/
	Referraldnslookup string `json:"referraldnslookup,omitempty"`
	/**
	* MSSRV Specific parameter. Used to locate the DNS node to which the SRV record pertains in the domainname. The domainname is appended to it to form the srv record.
		Example : For "dc._msdcs", the srv record formed is _ldap._tcp.dc._msdcs.<domainname>.
	*/
	Mssrvrecordlocation string `json:"mssrvrecordlocation,omitempty"`
	/**
	* When to validate LDAP server certs
	*/
	Validateservercert string `json:"validateservercert,omitempty"`
	/**
	* Hostname for the LDAP server.  If -validateServerCert is ON then this must be the host name on the certificate from the LDAP server.
		A hostname mismatch will cause a connection failure.
	*/
	Ldaphostname string `json:"ldaphostname,omitempty"`
	/**
	* Name that uniquely identifies a group in LDAP or Active Directory.
	*/
	Groupnameidentifier string `json:"groupnameidentifier,omitempty"`
	/**
	* LDAP group search attribute.
		Used to determine to which groups a group belongs.
	*/
	Groupsearchattribute string `json:"groupsearchattribute,omitempty"`
	/**
	* LDAP group search subattribute.
		Used to determine to which groups a group belongs.
	*/
	Groupsearchsubattribute string `json:"groupsearchsubattribute,omitempty"`
	/**
	* String to be combined with the default LDAP group search string to form the search value.  For example, the group search filter ""vpnallowed=true"" when combined with the group identifier ""samaccount"" and the group name ""g1"" yields the LDAP search string ""(&(vpnallowed=true)(samaccount=g1)"". If nestedGroupExtraction is ENABLED, the filter is applied on the first level group search as well, otherwise first level groups (of which user is a direct member of) will be fetched without applying this filter. (Be sure to enclose the search string in two sets of double quotation marks; both sets are needed.)
	*/
	Groupsearchfilter string `json:"groupsearchfilter,omitempty"`
	/**
	* This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
	*/
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute1 from the ldap response
	*/
	Attribute1 string `json:"attribute1,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute2 from the ldap response
	*/
	Attribute2 string `json:"attribute2,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute3 from the ldap response
	*/
	Attribute3 string `json:"attribute3,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute4 from the ldap response
	*/
	Attribute4 string `json:"attribute4,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute5 from the ldap response
	*/
	Attribute5 string `json:"attribute5,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute6 from the ldap response
	*/
	Attribute6 string `json:"attribute6,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute7 from the ldap response
	*/
	Attribute7 string `json:"attribute7,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute8 from the ldap response
	*/
	Attribute8 string `json:"attribute8,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute9 from the ldap response
	*/
	Attribute9 string `json:"attribute9,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute10 from the ldap response
	*/
	Attribute10 string `json:"attribute10,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute11 from the ldap response
	*/
	Attribute11 string `json:"attribute11,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute12 from the ldap response
	*/
	Attribute12 string `json:"attribute12,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute13 from the ldap response
	*/
	Attribute13 string `json:"attribute13,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute14 from the ldap response
	*/
	Attribute14 string `json:"attribute14,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute15 from the ldap response
	*/
	Attribute15 string `json:"attribute15,omitempty"`
	/**
	* Expression that would be evaluated to extract attribute16 from the ldap response
	*/
	Attribute16 string `json:"attribute16,omitempty"`
	/**
	* List of attribute names separated by ',' which needs to be fetched from ldap server.
		Note that preceeding and trailing spaces will be removed.
		Attribute name can be 127 bytes and total length of this string should not cross 2047 bytes.
		These attributes have multi-value support separated by ',' and stored as key-value pair in AAA session
	*/
	Attributes string `json:"attributes,omitempty"`
	/**
	* SSH PublicKey is attribute on AD. This attribute is used to retrieve ssh PublicKey for RBA authentication
	*/
	Sshpublickey string `json:"sshpublickey,omitempty"`
	/**
	* Name of the service used to send push notifications
	*/
	Pushservice string `json:"pushservice,omitempty"`
	/**
	* OneTimePassword(OTP) Secret key attribute on AD. This attribute is used to store and retrieve secret key used for OTP check
	*/
	Otpsecret string `json:"otpsecret,omitempty"`
	/**
	* The Citrix ADC uses the email attribute to query the Active Directory for the email id of a user
	*/
	Email string `json:"email,omitempty"`
	/**
	* KnowledgeBasedAuthentication(KBA) attribute on AD. This attribute is used to store and retrieve preconfigured Question and Answer knowledge base used for KBA authentication.
	*/
	Kbattribute string `json:"kbattribute,omitempty"`
	/**
	* The NetScaler appliance uses the alternateive email attribute to query the Active Directory for the alternative email id of a user
	*/
	Alternateemailattr string `json:"alternateemailattr,omitempty"`
	/**
	* The Citrix ADC uses the cloud attributes to extract additional attributes from LDAP servers required for Citrix Cloud operations
	*/
	Cloudattributes string `json:"cloudattributes,omitempty"`

	//------- Read only Parameter ---------;

	Success string `json:"success,omitempty"`
	Failure string `json:"failure,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
