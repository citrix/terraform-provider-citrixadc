---
subcategory: "Ipsecalg"
---

# Resource: ipsecalgprofile

The ipsecalgprofile resource is used to create ipsecalgprofile.


## Example usage

```hcl
resource "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile" {
  name              = "my_ipsecalgprofile"
  ikesessiontimeout = 50
  espsessiontimeout = 20
  connfailover      = "DISABLED"
}
```


## Argument Reference

* `name` - (Required) The name of the ipsec alg profile. Minimum length =  1 Maximum length =  32
* `ikesessiontimeout` - (Optional) IKE session timeout in minutes. Minimum value =  1 Maximum value =  1440
* `espsessiontimeout` - (Optional) ESP session timeout in minutes. Minimum value =  1 Maximum value =  1440
* `espgatetimeout` - (Optional) Timeout ESP in seconds as no ESP packets are seen after IKE negotiation. Minimum value =  3 Maximum value =  1200
* `connfailover` - (Optional) Mode in which the connection failover feature must operate for the IPSec Alg. After a failover, established UDP connections and ESP packet flows are kept active and resumed on the secondary appliance. Recomended setting is ENABLED. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ipsecalgprofile. It has the same value as the `name` attribute.


## Import

A ipsecalgprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_ipsecalgprofile.tf_ipsecalgprofile my_ipsecalgprofile
```
