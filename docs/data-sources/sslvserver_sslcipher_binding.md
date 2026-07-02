---
subcategory: "SSL"
---

# Data Source: sslvserver_sslcipher_binding

The sslvserver_sslcipher_binding data source allows you to retrieve information about a cipher binding to an SSL virtual server.

## Example usage

```terraform
data "citrixadc_sslvserver_sslcipher_binding" "tf_binding" {
  vservername = "tf_lbvserver"
  ciphername  = "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384"
}

output "vservername" {
  value = data.citrixadc_sslvserver_sslcipher_binding.tf_binding.vservername
}

output "description" {
  value = data.citrixadc_sslvserver_sslcipher_binding.tf_binding.description
}
```

## Argument Reference

The following arguments are required:

* `vservername` - (Required) Name of the SSL virtual server.
* `ciphername` - (Required) Name of the individual cipher, user-defined cipher group, or predefined (built-in) cipher alias.

The following arguments are optional:

* `cipheraliasname` - (Optional) The name of the cipher group/alias/individual cipher bindings.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslvserver_sslcipher_binding. It is the concatenation of the `vservername` and `ciphername` attributes separated by a comma.
* `description` - The cipher suite description. This is a read-only value returned by the Citrix ADC.
