---
subcategory: "SNMP"
---

# Data Source `snmpview`

The snmpview data source allows you to retrieve information about an SNMPv3 view.


## Example usage

```terraform
data "citrixadc_snmpview" "tf_snmpview" {
  name    = "my_snmpview"
  subtree = "1.2.4.7"
}

output "type" {
  value = data.citrixadc_snmpview.tf_snmpview.type
}
```


## Argument Reference

* `name` - (Required) Name for the SNMPv3 view. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters. You should choose a name that helps identify the SNMPv3 view.
* `subtree` - (Required) A particular branch (subtree) of the MIB tree that you want to associate with this SNMPv3 view. You must specify the subtree as an SNMP OID.

## Attribute Reference

The following attributes are available:

* `name` - Name for the SNMPv3 view.
* `subtree` - A particular branch (subtree) of the MIB tree that you want to associate with this SNMPv3 view.
* `type` - Include or exclude the subtree, specified by the subtree parameter, in or from this view. This setting can be useful when you have included a subtree, such as A, in an SNMPv3 view and you want to exclude a specific subtree of A, such as B, from the SNMPv3 view.
* `id` - The id of the snmpview. It is a system-generated identifier.
