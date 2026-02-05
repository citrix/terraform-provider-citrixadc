---
subcategory: "SNMP"
---

# citrixadc_snmpgroup (Data Source)

Data source for querying Citrix ADC SNMP groups. This data source retrieves information about SNMPv3 groups configured on the ADC appliance, which define security levels and access permissions for SNMPv3 users.

## Example Usage

```hcl
data "citrixadc_snmpgroup" "example" {
  name          = "admin_group"
  securitylevel = "noAuthNoPriv"
}

# Output group attributes
output "group_readview" {
  value = data.citrixadc_snmpgroup.example.readviewname
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the SNMPv3 group. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and special characters like hyphen (-), period (.), pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_).
* `securitylevel` - (Required) Security level required for communication between the Citrix ADC and the SNMPv3 users who belong to the group. Possible values:
  * `noAuthNoPriv` - Require neither authentication nor encryption
  * `authNoPriv` - Require authentication but no encryption
  * `authPriv` - Require authentication and encryption

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the snmpgroup datasource.
* `readviewname` - Name of the configured SNMPv3 view that is bound to this SNMPv3 group. An SNMPv3 user bound to this group can access the subtrees that are bound to this SNMPv3 view as type INCLUDED, but cannot access the ones that are type EXCLUDED.

## Notes

SNMPv3 groups provide a way to organize users and define security policies. Each group has a security level that determines the authentication and encryption requirements for all users in that group. Groups are also associated with SNMP views that control which MIB objects the users can access.
