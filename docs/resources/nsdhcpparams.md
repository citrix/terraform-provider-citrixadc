---
subcategory: "NS"
---

# Resource: nsdhcpparams

The nsdhcpparams resource is used to configure DHCP parameters resource.


## Example usage

```hcl
resource "citrixadc_nsdhcpparams" "tf_nsdhcpparams" {
  dhcpclient = "OFF"
  saveroute  = "OFF"
}
```


## Argument Reference

* `dhcpclient` - (Optional) Enables DHCP client to acquire IP address from the DHCP server in the next boot. When set to OFF, disables the DHCP client in the next boot. Possible values: [ on, off ]
* `saveroute` - (Optional) DHCP acquired routes are saved on the Citrix ADC. Possible values: [ on, off ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsdhcpparams. It is a unique string prefixed with "tf-nsdhcpparams-"

