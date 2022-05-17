---
subcategory: "NS"
---

# Resource: nsappflowcollector

The nsappflowcollector resource is used to create appflowCollector resource.


## Example usage

```hcl
resource "citrixadc_nsappflowcollector" "tf_appflowcollector" {
  name      = "tf_appflowcollector"
  ipaddress = "1.2.4.1"
  port      = 30
}
```


## Argument Reference

* `name` - (Required) Name of the AppFlow collector. Minimum length =  1 Maximum length =  127
* `ipaddress` - (Required) The IPv4 address of the AppFlow collector.
* `port` - (Optional) The UDP port on which the AppFlow collector is listening.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsappflowcollector. It has the same value as the `name` attribute.


## Import

A nsappflowcollector can be imported using its name, e.g.

```shell
terraform import citrixadc_nsappflowcollector.tf_appflowcollector tf_appflowcollector
```
