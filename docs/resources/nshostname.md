---
subcategory: "NS"
---

# Resource: nshostname

The nshostname resource is used to create nshostname.


## Example usage

```hcl
resource "citrixadc_nshostname" "tf_nshostname" {
   hostname = "mycitrix_adc"
}
```


## Argument Reference

* `hostname` - (Required) Host name for the Citrix ADC. Minimum length =  1 Maximum length =  255
* `ownernode` - (Optional) ID of the cluster node for which you are setting the hostname. Can be configured only through the cluster IP address. Minimum value =  0 Maximum value =  31


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nshostname. It has the same value as the `name` attribute.
