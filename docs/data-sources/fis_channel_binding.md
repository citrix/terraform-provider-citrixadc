---
subcategory: "Network"
---

# Data Source: fis_channel_binding

The fis_channel_binding data source allows you to retrieve information about a channel bound to a Failover Interface Set (FIS) on the Citrix ADC. Look up a binding by the FIS name and the channel interface number.


## Example usage

```terraform
data "citrixadc_fis_channel_binding" "example" {
  name  = "fis1"
  ifnum = "LA/1"
}

output "fis_channel_binding_id" {
  value = data.citrixadc_fis_channel_binding.example.id
}
```


## Argument Reference

* `name` - (Required) The name of the FIS to which the channel is bound.
* `ifnum` - (Required) Channel bound to the FIS, specified in slot/port notation (for example, `LA/1`).
* `ownernode` - (Optional) ID of the cluster node for which the FIS was created. Can be configured only through the cluster IP address.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the fis_channel_binding. It is a composite key of the form `name:<name>,ifnum:<ifnum>`, where each value is URL-encoded (for example, `name:fis1,ifnum:LA%2F1`).
* `name` - The name of the FIS to which the channel is bound.
* `ifnum` - The channel bound to the FIS.
* `ownernode` - ID of the cluster node for which the FIS was created.
