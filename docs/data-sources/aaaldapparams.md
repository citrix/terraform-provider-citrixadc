---
subcategory: "AAA"
---

# Data Source `aaaldapparams`

The aaaldapparams data source allows you to retrieve information about AAA LDAP parameters configuration.


## Example usage

```terraform
data "citrixadc_aaaldapparams" "tf_aaaldapparams" {
}

output "serverip" {
  value = data.citrixadc_aaaldapparams.tf_aaaldapparams.serverip
}

output "ldapbase" {
  value = data.citrixadc_aaaldapparams.tf_aaaldapparams.ldapbase
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `serverip` - IP address of the LDAP server.
* `serverport` - Port number on which the LDAP server listens for connections.
* `authtimeout` - Maximum number of seconds that the Citrix ADC waits for a response from the LDAP server.
* `ldapbase` - Base (the server and directory) from which LDAP search commands should start.
* `ldapbinddn` - Complete distinguished name (DN) string used for binding to the LDAP server.
* `ldaploginname` - Name attribute that the Citrix ADC uses to query the external LDAP server or an Active Directory.
* `searchfilter` - String to be combined with the default LDAP user search string to form the value to use when executing an LDAP search.
* `groupattrname` - Attribute name used for group extraction from the LDAP server.
* `subattributename` - Subattribute name used for group extraction from the LDAP server.
* `sectype` - Type of security used for communications between the Citrix ADC and the LDAP server. Possible values: [ PLAINTEXT, TLS, SSL ]
* `svrtype` - The type of LDAP server. Possible values: [ AD, NDS ]
* `ssonameattribute` - Attribute used by the Citrix ADC to query an external LDAP server or Active Directory for an alternative username.
* `passwdchange` - Enable or disable password change requests. Possible values: [ ENABLED, DISABLED ]
* `nestedgroupextraction` - Enable nested group extraction. Possible values: [ ON, OFF ]
* `maxnestinglevel` - Number of levels up to which the group extraction is performed.
* `groupnameidentifier` - Name attribute that the Citrix ADC uses to query the external LDAP server or an Active Directory.
* `groupsearchattribute` - LDAP group search attribute.
* `groupsearchsubattribute` - LDAP group search subattribute.
* `groupsearchfilter` - String to be combined with the default LDAP group search string to form the value to use when executing an LDAP search.
* `defaultauthenticationgroup` - This is the default group that is chosen when the authentication succeeds in addition to extracted groups.

## Attribute Reference

* `id` - The id of the aaaldapparams. It is a system-generated identifier.
