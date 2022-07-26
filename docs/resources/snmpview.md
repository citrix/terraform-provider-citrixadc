---
subcategory: "SNMP"
---

# Resource: snmpview

The snmpview resource is used to create snmpview.


## Example usage

```hcl
resource "citrixadc_snmpview" "tf_snmpview" {
  name    = "test_name"
  subtree = "1.2.4.7"
  type    = "excluded"
}
```


## Argument Reference

* `name` - (Required) Name for the SNMPv3 view. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters. You should choose a name that helps identify the SNMPv3 view.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose it in double or single quotation marks (for example, "my view" or 'my view').
* `subtree` - (Required) A particular branch (subtree) of the MIB tree that you want to associate with this SNMPv3 view. You must specify the subtree as an SNMP OID.
* `type` - (Optional) Include or exclude the subtree, specified by the subtree parameter, in or from this view. This setting can be useful when you have included a subtree, such as A, in an SNMPv3 view and you want to exclude a specific subtree of A, such as B, from the SNMPv3 view.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the snmpview. It has the same value as the `name` attribute.


## Import

A snmpview can be imported using its name, e.g.

```shell
terraform import citrixadc_snmpview.tf_snmpview test_name
```