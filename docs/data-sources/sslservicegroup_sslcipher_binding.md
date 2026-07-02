---
subcategory: "SSL"
---

# Data Source: sslservicegroup_sslcipher_binding

The sslservicegroup_sslcipher_binding data source allows you to retrieve information about an SSL cipher binding on an SSL service group.

## Example usage

```terraform
data "citrixadc_sslservicegroup_sslcipher_binding" "tf_binding" {
  servicegroupname = "tf_servicegroup"
  ciphername       = "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384"
}

output "cipheraliasname" {
  value = data.citrixadc_sslservicegroup_sslcipher_binding.tf_binding.cipheraliasname
}

output "description" {
  value = data.citrixadc_sslservicegroup_sslcipher_binding.tf_binding.description
}
```

## Argument Reference

* `servicegroupname` - (Required) The name of the SSL service group to which the cipher is bound.
* `ciphername` - (Required) A cipher-suite can consist of an individual cipher name, the system predefined cipher-alias name, or a user defined cipher-group name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservicegroup_sslcipher_binding. It is the concatenation of the `servicegroupname` and `ciphername` attributes separated by a comma.
* `cipheraliasname` - The name of the cipher group/alias/name configured for the SSL service group.
* `description` - The description of the cipher.
