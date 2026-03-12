---
subcategory: "SNMP"
---

# citrixadc_snmpcommunity (Data Source)

Data source for querying Citrix ADC SNMP communities. This data source retrieves information about SNMP community strings configured on the ADC appliance, which control access permissions for SNMP V1 and V2 queries.

## Example Usage

```hcl
data "citrixadc_snmpcommunity" "example" {
  communityname = "public"
}

# Output community attributes
output "community_permissions" {
  value = data.citrixadc_snmpcommunity.example.permissions
}
```

## Argument Reference

The following arguments are supported:

* `communityname` - (Required) The SNMP community string. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers and special characters. If the string includes one or more spaces, enclose the name in double or single quotation marks.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the snmpcommunity datasource.
* `permissions` - The SNMP V1 or V2 query-type privilege that is associated with this SNMP community. Possible values include GET, GET_BULK, GET_NEXT, and ALL.

## Notes

SNMP communities are used for authentication in SNMP V1 and V2c protocols. The community string acts as a password that grants different levels of access to the SNMP agent on the ADC appliance. Different permissions can be assigned to different community strings to control what operations can be performed.
