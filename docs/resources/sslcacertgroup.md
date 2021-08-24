---
subcategory: "SSL"
---

# Resource: sslcacertgroup

The sslcacertgroup resource is used to configure a Group of CA certificate-key pairs resource.


## Example usage

```hcl
resource "citrixadc_sslcacertgroup" "ns_callout_certs1" {
    cacertgroupname = "ns_callout_certs1"
}
```


## Argument Reference

* `cacertgroupname` - (Required) Name given to the CA certificate group. The name will be used to add the CA certificates to the group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my file" or 'my file').


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcacertgroup. It has the same value as the `cacertgroupname` attribute.

## Import

A sslaction can be imported using its name, e.g.

```shell
terraform import citrixadc_sslcacertgroup.tf_sslcacertgroup tf_sslcacertgroup
```
