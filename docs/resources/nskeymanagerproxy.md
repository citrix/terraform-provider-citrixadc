---
subcategory: "NS"
---

# Resource: nskeymanagerproxy

This resource is used to manage the Key Manager proxy server configuration.


## Example usage

```hcl
resource "citrixadc_nskeymanagerproxy" "tf_keymanagerproxy" {
  serverip = "192.168.20.30"
  port     = 443
}
```

Using a fully qualified domain name instead of an IP address:

```hcl
resource "citrixadc_nskeymanagerproxy" "tf_keymanagerproxy" {
  servername = "keymanager.example.com"
  port       = 443
}
```


## Argument Reference

* `port` - (Required) Key Manager proxy server port. Changing this attribute forces a new resource to be created.
* `serverip` - (Optional) IP address of the Key Manager proxy server. Exactly one of `serverip` or `servername` must be set, and they are mutually exclusive (setting both is rejected). Changing this attribute forces a new resource to be created.
* `servername` - (Optional) Fully qualified domain name of the Key Manager proxy server. Exactly one of `serverip` or `servername` must be set, and they are mutually exclusive (setting both is rejected). Changing this attribute forces a new resource to be created.
* `nodeid` - (Optional) Unique number that identifies the cluster node. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nskeymanagerproxy. It is the `serverip` value when set; otherwise it falls back to the `servername` value. This is also the key used for the delete operation, where `servername` (if set) is passed as an additional argument.


## Import

A nskeymanagerproxy can be imported using its serverip, e.g.

```shell
terraform import citrixadc_nskeymanagerproxy.tf_keymanagerproxy 192.168.20.30
```
