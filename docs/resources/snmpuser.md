---
subcategory: "SNMP"
---

# Resource: snmpuser

The snmpuser resource is used to create snmpuser.


## Example usage

```hcl
resource "citrixadc_snmpuser" "tf_snmpuser" {
  name       = "test_user"
  group      = "test_group"
  authtype   = "MD5"
  authpasswd = "this_is_my_password"
  privtype   = "DES"
  privpasswd = "this_is_my_password2"
}


resource "citrixadc_snmpgroup" "tf_snmpgroup" {
  name    = "test_group"
  securitylevel = "authNoPriv"
  readviewname = "test_readviewname"
}
```


## Argument Reference

* `name` - (Required) Name for the SNMPv3 user. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose it in double or single quotation marks (for example, "my user" or 'my user').
* `group` - (Required) Name of the configured SNMPv3 group to which to bind this SNMPv3 user. The access rights (bound SNMPv3 views) and security level set for this group are assigned to this user.
* `authpasswd` - (Optional) Plain-text pass phrase to be used by the authentication algorithm specified by the authType (Authentication Type) parameter. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  The following requirement applies only to the Citrix ADC CLI: If the pass phrase includes one or more spaces, enclose it in double or single quotation marks (for example, "my phrase" or 'my phrase').
* `authtype` - (Optional) Authentication algorithm used by the Citrix ADC and the SNMPv3 user for authenticating the communication between them. You must specify the same authentication algorithm when you configure the SNMPv3 user in the SNMP manager.
* `privpasswd` - (Optional) Encryption key to be used by the encryption algorithm specified by the privType (Encryption Type) parameter. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  The following requirement applies only to the Citrix ADC CLI: If the key includes one or more spaces, enclose it in double or single quotation marks (for example, "my key" or 'my key').
* `privtype` - (Optional) Encryption algorithm used by the Citrix ADC and the SNMPv3 user for encrypting the communication between them. You must specify the same encryption algorithm when you configure the SNMPv3 user in the SNMP manager.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the snmpuser. It has the same value as the `name` attribute.


## Import

A snmpuser can be imported using its name, e.g.

```shell
terraform import citrixadc_snmpuser.tf_snmpuser test_user
```
