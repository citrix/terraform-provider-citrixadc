---
subcategory: "Reputation"
---

# Resource: reputationsettings

The reputationsettings resource is used to create reputationsettings.


## Example usage

```hcl
resource "citrixadc_reputationsettings" "tf_reputationsettings" {
  proxyserver = "my_proxyserver"
  proxyport   = 3500
}
```


## Argument Reference

* `proxyport` - (Optional) Proxy server port.
* `proxyserver` - (Optional) Proxy server IP to get Reputation data.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the reputationsettings. It is a unique string prefixed with `tf-reputationsettings-` attribute.