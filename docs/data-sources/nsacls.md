---
subcategory: "NS"
---

# Data Source `nsacls`

The nsacls data source allows you to retrieve information about the ACL configuration.


## Example usage

```terraform
resource "citrixadc_nsacls" "tf_nsacls" {
  type = "CLASSIC"
}

data "citrixadc_nsacls" "tf_nsacls_data" {
  depends_on = [citrixadc_nsacls.tf_nsacls]
}

output "nsacls_type" {
  value = data.citrixadc_nsacls.tf_nsacls_data.type
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `type` - Type of the ACL, default will be CLASSIC. Available options as follows:
  * CLASSIC - specifies the regular extended acls.
  * DFD - cluster specific acls, specifies hashmethod for steering of the packet in cluster.
* `id` - The id of the nsacls configuration. It has a static value of `nsacls-config`.
