---
subcategory: "SNMP"
---

# citrixadc_snmpengineid (Data Source)

Data source for querying Citrix ADC SNMP engine ID. This data source retrieves information about the SNMP engine ID configured on the ADC appliance, which is a unique identifier used in SNMP v3 communications.

## Example Usage

```hcl
data "citrixadc_snmpengineid" "example" {
  ownernode = -1
}

# Output engine ID
output "snmp_engineid" {
  value = data.citrixadc_snmpengineid.example.engineid
}
```

## Argument Reference

The following arguments are supported:

* `ownernode` - (Required) ID of the cluster node for which you are querying the engineid. Use -1 for standalone or primary node.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the snmpengineid datasource.
* `engineid` - A hexadecimal value of at least 10 characters, uniquely identifying the engineid.

### Read-only snmpengineid metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_snmpengineid` resource). They are Computed / GET-only and are `null` when the appliance does not return them.

* `defaultengineid` - Unique identifier to assign to the SNMPv3 engine. Should be a hexadecimal value with a minimum length of 10 hex characters.

