---
subcategory: "NS"
---

# Resource: nsservicefunction

The nsservicefunction resource is used to create service Function resource.


## Example usage

```hcl
resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 20
  aliasname = "Management VLAN"
}
resource "citrixadc_nsservicefunction" "tf_servicefunc" {
  servicefunctionname = "tf_servicefunc"
  ingressvlan         = citrixadc_vlan.tf_vlan.vlanid
}
```


## Argument Reference

* `servicefunctionname` - (Required) Name of the service function to be created. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ). Minimum length =  1
* `ingressvlan` - (Required) VLAN ID on which the traffic from service function reaches Citrix ADC. Minimum value =  1 Maximum value =  4094


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsservicefunction. It has the same value as the `servicefunctionname` attribute.


## Import

A nsservicefunction can be imported using its servicefunctionname, e.g.

```shell
terraform import citrixadc_nsservicefunction.tf_servicefunc tf_servicefunc
```
