---
subcategory: "System"
---

# Resource: systemcpuparam

This resource is used to manage global packet-engine CPU parameters on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_systemcpuparam" "tf_systemcpuparam" {
  pemode = "CPUBOUND"
}
```


## Argument Reference

* `pemode` - (Optional) Set PEmode to DEFAULT/CPUBOUND. Distribute the PE weights equally if PEmode is set to CPUBOUND. Possible values: [ DEFAULT, CPUBOUND ]. If not specified, the value applied by the ADC is read back into state.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemcpuparam. It is set to `systemcpuparam-config`.


## Import

A systemcpuparam can be imported using its id (a synthetic constant, because this is a global singleton), e.g.

```shell
terraform import citrixadc_systemcpuparam.tf_systemcpuparam systemcpuparam-config
```
