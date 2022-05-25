---
subcategory: "Network"
---

# Resource: mapbmr_bmrv4network_binding

The mapbmr_bmrv4network_binding resource is used to bind bmrv4network that can be bound to mapbmr.


## Example usage

```hcl
resource "citrixadc_mapbmr" "tf_mapbmr" {
  name           = "tf_mapbmr"
  ruleipv6prefix = "2001:db8:abcd:12::/64"
  psidoffset     = 6
  eabitlength    = 16
  psidlength     = 8
}
resource "citrixadc_mapbmr_bmrv4network_binding" "tf_binding" {
  name    = citrixadc_mapbmr.tf_mapbmr.name
  network = "1.2.3.0"
  netmask = "255.255.255.0"
}
```


## Argument Reference

* `name` - (Required) Name for the Basic Mapping Rule. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the  MAP Basic Mapping Rule is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "add network MapBmr bmr1 -natprefix 2005::/64 -EAbitLength 16 -psidoffset 6 -portsharingratio 8" ). 			The Basic Mapping Rule information allows a MAP BR to determine source IPv4 address from the IPv6 packet sent from MAP CE device. 			Also it allows to determine destination IPv6 address of MAP CE before sending packets to MAP CE. Minimum length =  1 Maximum length =  127
* `network` - (Required) IPv4 NAT address range of Customer Edge (CE). Minimum length =  1
* `netmask` - (Optional) Subnet mask for the IPv4 address specified in the Network parameter.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the mapbmr_bmrv4network_binding. It is the concatenation of `name` and `network` attributes separated by comma.


## Import

A mapbmr_bmrv4network_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_mapbmr_bmrv4network_binding.tf_binding tf_mapbmr,1.2.3.0
```
