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

## Notes

The SNMP engine ID is a unique identifier used by SNMP v3 for authentication and privacy. Each SNMP agent must have a unique engine ID. In a cluster environment, each node can have its own engine ID, identified by the ownernode parameter.
