---
subcategory: "SNMP"
---

# Data Source: citrixadc_snmpuser

The `citrixadc_snmpuser` data source is used to retrieve information about an SNMPv3 user configured on the Citrix ADC.

## Example usage

```hcl
data "citrixadc_snmpuser" "example" {
  name = "snmpuser1"
}
```

## Argument Reference

* `name` - (Required) Name for the SNMPv3 user. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the SNMP user.
* `authpasswd` - Plain-text pass phrase to be used by the authentication algorithm specified by the authType (Authentication Type) parameter.
* `authtype` - Authentication algorithm used by the Citrix ADC and the SNMPv3 user for authenticating the communication between them. You must specify the same authentication algorithm when you configure the SNMPv3 user in the SNMP manager.
* `group` - Name of the configured SNMPv3 group to which to bind this SNMPv3 user. The access rights (bound SNMPv3 views) and security level set for this group are assigned to this user.
* `privpasswd` - Encryption key to be used by the encryption algorithm specified by the privType (Encryption Type) parameter.
* `privtype` - Encryption algorithm used by the Citrix ADC and the SNMPv3 user for encrypting the communication between them. You must specify the same encryption algorithm when you configure the SNMPv3 user in the SNMP manager.
