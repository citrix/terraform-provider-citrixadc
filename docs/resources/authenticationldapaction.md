---
subcategory: "Authentication"
---

# Resource: authenticationldapaction

The authenticationldapaction resource is used to create LDAP action resource.


## Example usage

```hcl
resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
  name   = "ldapaction"
  serverip = "1.2.3.4"
	serverport = 8080
}
```


## Argument Reference

* `name` - (Required) Name for the new LDAP action.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the LDAP action is added.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication action" or 'my authentication action').
* `alternateemailattr` - (Optional) The NetScaler appliance uses the alternateive email attribute to query the Active Directory for the alternative email id of a user
* `attribute1` - (Optional) Expression that would be evaluated to extract attribute1 from the ldap response
* `attribute10` - (Optional) Expression that would be evaluated to extract attribute10 from the ldap response
* `attribute11` - (Optional) Expression that would be evaluated to extract attribute11 from the ldap response
* `attribute12` - (Optional) Expression that would be evaluated to extract attribute12 from the ldap response
* `attribute13` - (Optional) Expression that would be evaluated to extract attribute13 from the ldap response
* `attribute14` - (Optional) Expression that would be evaluated to extract attribute14 from the ldap response
* `attribute15` - (Optional) Expression that would be evaluated to extract attribute15 from the ldap response
* `attribute16` - (Optional) Expression that would be evaluated to extract attribute16 from the ldap response
* `attribute2` - (Optional) Expression that would be evaluated to extract attribute2 from the ldap response
* `attribute3` - (Optional) Expression that would be evaluated to extract attribute3 from the ldap response
* `attribute4` - (Optional) Expression that would be evaluated to extract attribute4 from the ldap response
* `attribute5` - (Optional) Expression that would be evaluated to extract attribute5 from the ldap response
* `attribute6` - (Optional) Expression that would be evaluated to extract attribute6 from the ldap response
* `attribute7` - (Optional) Expression that would be evaluated to extract attribute7 from the ldap response
* `attribute8` - (Optional) Expression that would be evaluated to extract attribute8 from the ldap response
* `attribute9` - (Optional) Expression that would be evaluated to extract attribute9 from the ldap response
* `attributes` - (Optional) List of attribute names separated by ',' which needs to be fetched from ldap server.  Note that preceeding and trailing spaces will be removed.  Attribute name can be 127 bytes and total length of this string should not cross 2047 bytes. These attributes have multi-value support separated by ',' and stored as key-value pair in AAA session
* `authentication` - (Optional) Perform LDAP authentication. If authentication is disabled, any LDAP authentication attempt returns authentication success if the user is found.  CAUTION! Authentication should be disabled only for authorization group extraction or where other (non-LDAP) authentication methods are in use and either bound to a primary list or flagged as secondary.
* `authtimeout` - (Optional) Number of seconds the Citrix ADC waits for a response from the RADIUS server.
* `cloudattributes` - (Optional) The Citrix ADC uses the cloud attributes to extract additional attributes from LDAP servers required for Citrix Cloud operations
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `email` - (Optional) The Citrix ADC uses the email attribute to query the Active Directory for the email id of a user
* `followreferrals` - (Optional) Setting this option to ON enables following LDAP referrals received from the LDAP server.
* `groupattrname` - (Optional) LDAP group attribute name. Used for group extraction on the LDAP server.
* `groupnameidentifier` - (Optional) Name that uniquely identifies a group in LDAP or Active Directory.
* `groupsearchattribute` - (Optional) LDAP group search attribute.  Used to determine to which groups a group belongs.
* `groupsearchfilter` - (Optional) String to be combined with the default LDAP group search string to form the search value.  For example, the group search filter ""vpnallowed=true"" when combined with the group identifier ""samaccount"" and the group name ""g1"" yields the LDAP search string ""(&(vpnallowed=true)(samaccount=g1)"". If nestedGroupExtraction is ENABLED, the filter is applied on the first level group search as well, otherwise first level groups (of which user is a direct member of) will be fetched without applying this filter. (Be sure to enclose the search string in two sets of double quotation marks; both sets are needed.)
* `groupsearchsubattribute` - (Optional) LDAP group search subattribute.  Used to determine to which groups a group belongs.
* `kbattribute` - (Optional) KnowledgeBasedAuthentication(KBA) attribute on AD. This attribute is used to store and retrieve preconfigured Question and Answer knowledge base used for KBA authentication.
* `ldapbase` - (Optional) Base (node) from which to start LDAP searches.  If the LDAP server is running locally, the default value of base is dc=netscaler, dc=com.
* `ldapbinddn` - (Optional) Full distinguished name (DN) that is used to bind to the LDAP server.  Default: cn=Manager,dc=netscaler,dc=com
* `ldapbinddnpassword` - (Optional) Password used to bind to the LDAP server.
* `ldaphostname` - (Optional) Hostname for the LDAP server.  If -validateServerCert is ON then this must be the host name on the certificate from the LDAP server. A hostname mismatch will cause a connection failure.
* `ldaploginname` - (Optional) LDAP login name attribute.  The Citrix ADC uses the LDAP login name to query external LDAP servers or Active Directories.
* `maxldapreferrals` - (Optional) Specifies the maximum number of nested referrals to follow.
* `maxnestinglevel` - (Optional) If nested group extraction is ON, specifies the number of levels up to which group extraction is performed.
* `mssrvrecordlocation` - (Optional) MSSRV Specific parameter. Used to locate the DNS node to which the SRV record pertains in the domainname. The domainname is appended to it to form the srv record. Example : For "dc._msdcs", the srv record formed is _ldap._tcp.dc._msdcs.<domainname>.
* `nestedgroupextraction` - (Optional) Allow nested group extraction, in which the Citrix ADC queries external LDAP servers to determine whether a group is part of another group.
* `otpsecret` - (Optional) OneTimePassword(OTP) Secret key attribute on AD. This attribute is used to store and retrieve secret key used for OTP check
* `passwdchange` - (Optional) Allow password change requests.
* `pushservice` - (Optional) Name of the service used to send push notifications
* `referraldnslookup` - (Optional) Specifies the DNS Record lookup Type for the referrals
* `requireuser` - (Optional) Require a successful user search for authentication. CAUTION!  This field should be set to NO only if usersearch not required [Both username validation as well as password validation skipped] and (non-LDAP) authentication methods are in use and either bound to a primary list or flagged as secondary.
* `searchfilter` - (Optional) String to be combined with the default LDAP user search string to form the search value. For example, if the search filter "vpnallowed=true" is combined with the LDAP login name "samaccount" and the user-supplied username is "bob", the result is the LDAP search string ""(&(vpnallowed=true)(samaccount=bob)"" (Be sure to enclose the search string in two sets of double quotation marks; both sets are needed.).
* `sectype` - (Optional) Type of security used for communications between the Citrix ADC and the LDAP server. For the PLAINTEXT setting, no encryption is required.
* `serverip` - (Optional) IP address assigned to the LDAP server.
* `servername` - (Optional) LDAP server name as a FQDN.  Mutually exclusive with LDAP IP address.
* `serverport` - (Optional) Port on which the LDAP server accepts connections.
* `sshpublickey` - (Optional) SSH PublicKey is attribute on AD. This attribute is used to retrieve ssh PublicKey for RBA authentication
* `ssonameattribute` - (Optional) LDAP single signon (SSO) attribute.  The Citrix ADC uses the SSO name attribute to query external LDAP servers or Active Directories for an alternate username.
* `subattributename` - (Optional) LDAP group sub-attribute name.  Used for group extraction from the LDAP server.
* `svrtype` - (Optional) The type of LDAP server.
* `validateservercert` - (Optional) When to validate LDAP server certs


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationldapaction. It has the same value as the `name` attribute.


## Import

A authenticationldapaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationldapaction.tf_authenticationldapaction ldapaction
```
