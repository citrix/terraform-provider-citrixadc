---
subcategory: "SSL"
---

# Data Source: sslprofile_sslcertkey_binding

The sslprofile_sslcertkey_binding data source allows you to retrieve information about the binding between an SSL profile and an SSL certificate key.

## Example Usage

```terraform
data "citrixadc_sslprofile_sslcertkey_binding" "demo_sslprofile_sslcertkey_binding" {
  name          = "tfUnit_sslprofile-hello"
  sslicacertkey = "tf_sslcertkey"
}

output "cipherpriority" {
  value = data.citrixadc_sslprofile_sslcertkey_binding.demo_sslprofile_sslcertkey_binding.cipherpriority
}
```

## Argument Reference

* `name` - (Required) Name of the SSL profile.
* `sslicacertkey` - (Required) The certkey (CA certificate + private key) to be used for SSL interception.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `cipherpriority` - Priority of the cipher binding.
* `id` - The id of the sslprofile_sslcertkey_binding. It is a system-generated identifier.
