---
subcategory: "SNMP"
---

# Resource: snmpcommunity

The snmpcommunity resource is used to create snmpcommunity.


## Example usage

```hcl
resource "citrixadc_snmpcommunity" "tf_snmpcommunity" {
  communityname = "test_community"
  permissions   = "GET_BULK"
}
```


## Argument Reference

* `communityname` - (Required) The SNMP community string. Can consist of 1 to 31 characters that include uppercase and lowercase letters,numbers and special characters.  The following requirement applies only to the Citrix ADC CLI: If the string includes one or more spaces, enclose the name in double or single quotation marks (for example, "my string" or 'my string').
* `permissions` - (Required) The SNMP V1 or V2 query-type privilege that you want to associate with this SNMP community. Possible values : GET, GET_NEXT, GET_BULK, SET, ALL


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the snmpcommunity. It has the same value as the `communityname` attribute.


## Import

A snmpcommunity can be imported using its name, e.g.

```shell
terraform import citrixadc_snmpcommunity.tf_snmpcommunity test_community
```