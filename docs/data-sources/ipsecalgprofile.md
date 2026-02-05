---
subcategory: "Network"
---

# Data Source `ipsecalgprofile`

The ipsecalgprofile data source allows you to retrieve information about IPSec ALG profiles.


## Example usage

```terraform
data "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile" {
  name = "my_ipsecalgprofile"
}

output "ikesessiontimeout" {
  value = data.citrixadc_ipsecalgprofile.tf_ipsecalgprofile.ikesessiontimeout
}

output "espsessiontimeout" {
  value = data.citrixadc_ipsecalgprofile.tf_ipsecalgprofile.espsessiontimeout
}
```


## Argument Reference

* `name` - (Required) The name of the ipsec alg profile

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `connfailover` - Mode in which the connection failover feature must operate for the IPSec Alg. After a failover, established UDP connections and ESP packet flows are kept active and resumed on the secondary appliance. Recomended setting is ENABLED.
* `espgatetimeout` - Timeout ESP in seconds as no ESP packets are seen after IKE negotiation
* `espsessiontimeout` - ESP session timeout in minutes.
* `ikesessiontimeout` - IKE session timeout in minutes
* `id` - The id of the ipsecalgprofile. It has the same value as the `name` attribute.
