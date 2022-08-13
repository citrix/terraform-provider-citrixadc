---
subcategory: "NS"
---

# Resource: nslicenseproxyserver

The nslicenseproxyserver resource is used to create nslicenseproxyserver.


## Example usage

```hcl
resource "citrixadc_nslicenseproxyserver" "tf_nslicenseproxyserver" {
  servername = "www.example.com"
  port       = 80
}

```


## Argument Reference

* `serverip` - (Optional) IP address of the License proxy server. Minimum length =  1
* `servername` - (Optional) Fully qualified domain name of the License proxy server.
* `port` - (Required) License proxy server port.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nslicenseproxyserver. It has the same value as the `servername` or `serverip` attribute.


## Import

A nslicenseproxyserver can be imported using its name, e.g.

```shell
terraform import citrixadc_nslicenseproxyserver.tf_nslicenseproxyserver www.example.com
```
