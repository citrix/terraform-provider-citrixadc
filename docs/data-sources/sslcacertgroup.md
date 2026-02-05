---
subcategory: "SSL"
---

# Data Source `sslcacertgroup`

The sslcacertgroup data source allows you to retrieve information about an SSL CA certificate group.


## Example usage

```terraform
data "citrixadc_sslcacertgroup" "tf_sslcacertgroup" {
  cacertgroupname = "my_cacertgroup"
}

output "cacertgroupname" {
  value = data.citrixadc_sslcacertgroup.tf_sslcacertgroup.cacertgroupname
}
```


## Argument Reference

* `cacertgroupname` - (Required) Name given to the CA certificate group. The name will be used to add the CA certificates to the group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

The following attributes are available:

* `cacertgroupname` - Name given to the CA certificate group.
* `id` - The id of the sslcacertgroup. It is a system-generated identifier.
