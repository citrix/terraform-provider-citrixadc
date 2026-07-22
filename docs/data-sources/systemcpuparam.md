---
subcategory: "System"
---

# Data Source: systemcpuparam

The `citrixadc_systemcpuparam` data source is used to retrieve the global packet-engine (PE) CPU parameter configuration from the Citrix ADC.


## Example usage

```hcl
data "citrixadc_systemcpuparam" "example" {}

output "systemcpuparam_pemode" {
  value = data.citrixadc_systemcpuparam.example.pemode
}
```


## Argument Reference

This data source is a singleton and does not require any lookup arguments. It retrieves the current packet-engine CPU parameter configuration from the Citrix ADC.

* `pemode` - (Optional) Set PEmode to DEFAULT/CPUBOUND. Distribute the PE weights equally if PEmode is set to CPUBOUND. Possible values: [ DEFAULT, CPUBOUND ].


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the systemcpuparam data source. It is set to `systemcpuparam-config`.
* `pemode` - The currently configured PE mode on the appliance.
