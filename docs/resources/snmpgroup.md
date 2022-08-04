---
subcategory: "SNMP"
---

# Resource: snmpgroup

The snmpgroup resource is used to create snmpgroup.


## Example usage

```hcl
resource "citrixadc_snmpgroup" "tf_snmpgroup" {
  name    = "test_group"
  securitylevel = "authNoPriv"
  readviewname = "test_name"
}
```


## Argument Reference

* `name` - (Required) Name for the SNMPv3 group. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  You should choose a name that helps identify the SNMPv3 group.               The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose it in double or single quotation marks (for example, "my name" or 'my name').
* `securitylevel` - (Required) Security level required for communication between the Citrix ADC and the SNMPv3 users who belong to the group. Specify one of the following options: noAuthNoPriv. Require neither authentication nor encryption. authNoPriv. Require authentication but no encryption. authPriv. Require authentication and encryption. Note: If you specify authentication, you must specify an encryption algorithm when you assign an SNMPv3 user to the group. If you also specify encryption, you must assign both an authentication and an encryption algorithm for each group member.
* `readviewname` - (Required) Name of the configured SNMPv3 view that you want to bind to this SNMPv3 group. An SNMPv3 user bound to this group can access the subtrees that are bound to this SNMPv3 view as type INCLUDED, but cannot access the ones that are type EXCLUDED. If the Citrix ADC has multiple SNMPv3 view entries with the same name, all such entries are associated with the SNMPv3 group.



## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the snmpgroup. It has the same value as the `name` attribute.


## Import

A snmpgroup can be imported using its name, e.g.

```shell
terraform import citrixadc_snmpgroup.tf_snmpgroup test_group
```
