---
subcategory: "AAA"
---

# Resource: aaaldapparams

The aaaldapparams resource is used to update aaaldapparams.


## Example usage

```hcl
resource "citrixadc_aaaldapparams" "tf_aaaldapparams" {
  authtimeout  = 5
  serverip     = "10.222.74.158"
  passwdchange = "DISABLED"
}
```


## Argument Reference

* `serverip` - (Optional) IP address of your LDAP server.
* `serverport` - (Optional) Port number on which the LDAP server listens for connections. Minimum value =  1
* `authtimeout` - (Optional) Maximum number of seconds that the Citrix ADC waits for a response from the LDAP server. Minimum value =  1
* `ldapbase` - (Optional) Base (the server and location) from which LDAP search commands should start. If the LDAP server is running locally, the default value of base is dc=netscaler, dc=com.
* `ldapbinddn` - (Optional) Complete distinguished name (DN) string used for binding to the LDAP server.
* `ldapbinddnpassword` - (Optional) Password for binding to the LDAP server. Minimum length =  1
* `ldaploginname` - (Optional) Name attribute that the Citrix ADC uses to query the external LDAP server or an Active Directory.
* `searchfilter` - (Optional) String to be combined with the default LDAP user search string to form the value to use when executing an LDAP search. For example, the following values: vpnallowed=true, ldaploginame=""samaccount"" when combined with the user-supplied username ""bob"", yield the following LDAP search string: ""(&(vpnallowed=true)(samaccount=bob)"". Minimum length =  1
* `groupattrname` - (Optional) Attribute name used for group extraction from the LDAP server.
* `subattributename` - (Optional) Subattribute name used for group extraction from the LDAP server.
* `sectype` - (Optional) Type of security used for communications between the Citrix ADC and the LDAP server. For the PLAINTEXT setting, no encryption is required. Possible values: [ PLAINTEXT, TLS, SSL ]
* `svrtype` - (Optional) The type of LDAP server. Possible values: [ AD, NDS ]
* `ssonameattribute` - (Optional) Attribute used by the Citrix ADC to query an external LDAP server or Active Directory for an alternative username. This alternative username is then used for single sign-on (SSO).
* `passwdchange` - (Optional) Accept password change requests. Possible values: [ ENABLED, DISABLED ]
* `nestedgroupextraction` - (Optional) Queries the external LDAP server to determine whether the specified group belongs to another group. Possible values: [ on, off ]
* `maxnestinglevel` - (Optional) Number of levels up to which the system can query nested LDAP groups. Minimum value =  2
* `groupnameidentifier` - (Optional) LDAP-group attribute that uniquely identifies the group. No two groups on one LDAP server can have the same group name identifier.
* `groupsearchattribute` - (Optional) LDAP-group attribute that designates the parent group of the specified group. Use this attribute to search for a group's parent group.
* `groupsearchsubattribute` - (Optional) LDAP-group subattribute that designates the parent group of the specified group. Use this attribute to search for a group's parent group.
* `groupsearchfilter` - (Optional) Search-expression that can be specified for sending group-search requests to the LDAP server.
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups. Maximum length =  64


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaaldapparams. It is a unique string prefixed with `tf-aaaldapparams-`.