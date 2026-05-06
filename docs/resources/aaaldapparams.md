---
subcategory: "AAA"
---

# Resource: aaaldapparams

The aaaldapparams resource is used to update aaaldapparams.


## Example usage

### Using ldapbinddnpassword (sensitive attribute - persisted in state)

```hcl
variable "aaaldapparams_ldapbinddnpassword" {
  type      = string
  sensitive = true
}

resource "citrixadc_aaaldapparams" "tf_aaaldapparams" {
  serverip           = "10.222.74.158"
  ldapbinddn         = "cn=admin,dc=example,dc=com"
  ldapbinddnpassword = var.aaaldapparams_ldapbinddnpassword
  authtimeout        = 5
  passwdchange       = "DISABLED"
}
```

### Using ldapbinddnpassword_wo (write-only/ephemeral - NOT persisted in state)

The `ldapbinddnpassword_wo` attribute provides an ephemeral path for providing the LDAP bind password. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `ldapbinddnpassword_wo_version`.

```hcl
variable "aaaldapparams_ldapbinddnpassword" {
  type      = string
  sensitive = true
}

resource "citrixadc_aaaldapparams" "tf_aaaldapparams" {
  serverip                      = "10.222.74.158"
  ldapbinddn                    = "cn=admin,dc=example,dc=com"
  ldapbinddnpassword_wo         = var.aaaldapparams_ldapbinddnpassword
  ldapbinddnpassword_wo_version = 1
  authtimeout                   = 5
  passwdchange                  = "DISABLED"
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_aaaldapparams" "tf_aaaldapparams" {
  serverip                      = "10.222.74.158"
  ldapbinddn                    = "cn=admin,dc=example,dc=com"
  ldapbinddnpassword_wo         = var.aaaldapparams_ldapbinddnpassword
  ldapbinddnpassword_wo_version = 2  # Bumped to trigger update
  authtimeout                   = 5
  passwdchange                  = "DISABLED"
}
```


## Argument Reference

* `authtimeout` - (Optional) Maximum number of seconds that the Citrix ADC waits for a response from the LDAP server. Minimum value = 1
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups. Maximum length = 64
* `groupattrname` - (Optional) Attribute name used for group extraction from the LDAP server.
* `groupnameidentifier` - (Optional) LDAP-group attribute that uniquely identifies the group. No two groups on one LDAP server can have the same group name identifier.
* `groupsearchattribute` - (Optional) LDAP-group attribute that designates the parent group of the specified group. Use this attribute to search for a group's parent group.
* `groupsearchfilter` - (Optional) Search-expression that can be specified for sending group-search requests to the LDAP server.
* `groupsearchsubattribute` - (Optional) LDAP-group subattribute that designates the parent group of the specified group. Use this attribute to search for a group's parent group.
* `ldapbase` - (Optional) Base (the server and location) from which LDAP search commands should start. If the LDAP server is running locally, the default value of base is dc=netscaler, dc=com.
* `ldapbinddn` - (Optional) Complete distinguished name (DN) string used for binding to the LDAP server.
* `ldapbinddnpassword` - (Optional, Sensitive) Password for binding to the LDAP server. The value is persisted in Terraform state (encrypted). See also `ldapbinddnpassword_wo` for an ephemeral alternative. Minimum length = 1
* `ldapbinddnpassword_wo` - (Optional, Sensitive, WriteOnly) Same as `ldapbinddnpassword`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `ldapbinddnpassword_wo_version`. If both `ldapbinddnpassword` and `ldapbinddnpassword_wo` are set, `ldapbinddnpassword_wo` takes precedence.
* `ldapbinddnpassword_wo_version` - (Optional) An integer version tracker for `ldapbinddnpassword_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `ldaploginname` - (Optional) Name attribute that the Citrix ADC uses to query the external LDAP server or an Active Directory.
* `maxnestinglevel` - (Optional) Number of levels up to which the system can query nested LDAP groups. Minimum value = 2
* `nestedgroupextraction` - (Optional) Queries the external LDAP server to determine whether the specified group belongs to another group. Possible values: [ on, off ]
* `passwdchange` - (Optional) Accept password change requests. Possible values: [ ENABLED, DISABLED ]
* `searchfilter` - (Optional) String to be combined with the default LDAP user search string to form the value to use when executing an LDAP search. For example, the following values: vpnallowed=true, ldaploginame=""samaccount"" when combined with the user-supplied username ""bob"", yield the following LDAP search string: ""(&(vpnallowed=true)(samaccount=bob)"". Minimum length = 1
* `sectype` - (Optional) Type of security used for communications between the Citrix ADC and the LDAP server. For the PLAINTEXT setting, no encryption is required. Possible values: [ PLAINTEXT, TLS, SSL ]
* `serverip` - (Optional) IP address of your LDAP server.
* `serverport` - (Optional) Port number on which the LDAP server listens for connections. Minimum value = 1
* `ssonameattribute` - (Optional) Attribute used by the Citrix ADC to query an external LDAP server or Active Directory for an alternative username. This alternative username is then used for single sign-on (SSO).
* `subattributename` - (Optional) Subattribute name used for group extraction from the LDAP server.
* `svrtype` - (Optional) The type of LDAP server. Possible values: [ AD, NDS ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaaldapparams. It is a unique string prefixed with `aaaldapparams-config`.