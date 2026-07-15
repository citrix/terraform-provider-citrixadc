---
subcategory: "System"
---

# Resource: systemcpuparam

Configures the global packet-engine (PE) CPU parameters on the Citrix ADC. Use this resource to control how the appliance distributes processing weight across packet engines by selecting the PE mode (DEFAULT or CPUBOUND). Setting `CPUBOUND` distributes the PE weights equally, which can help on CPU-bound workloads.

This is a global singleton: a single configuration object always exists on the appliance. Creating the resource applies your settings, and destroying it only removes the object from Terraform state (the underlying configuration on the ADC is left in place).


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

* `id` - The ID of the systemcpuparam resource. Because this is a global singleton, the ID is a synthetic constant string `systemcpuparam-config`.


## Import

A systemcpuparam can be imported using its id (a synthetic constant, because this is a global singleton), e.g.

```shell
terraform import citrixadc_systemcpuparam.tf_systemcpuparam systemcpuparam-config
```
