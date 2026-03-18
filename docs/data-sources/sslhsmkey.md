---
subcategory: "SSL"
---

# Data Source: sslhsmkey

The sslhsmkey data source allows you to retrieve information about SSL HSM keys.

## Example usage

```terraform
data "citrixadc_sslhsmkey" "tf_hsmkey1" {
  hsmkeyname = "hsmkey1"
}

output "hsmtype" {
  value = data.citrixadc_sslhsmkey.tf_hsmkey1.hsmtype
}

output "serialnum" {
  value = data.citrixadc_sslhsmkey.tf_hsmkey1.serialnum
}
```

## Argument Reference

* `hsmkeyname` - (Required) Name for the HSM key. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `hsmtype` - Type of HSM.
* `key` - Name of the key. optionally, for Thales, path to the HSM key file; /var/opt/nfast/kmdata/local/ is the default path. Applies when HSMTYPE is THALES or KEYVAULT.
* `keystore` - Name of keystore object representing HSM where key is stored. For example, name of keyvault object or azurekeyvault authentication object. Applies only to KEYVAULT type HSM.
* `password` - Password for a partition. Applies only to SafeNet HSM.
* `serialnum` - Serial number of the partition on which the key is present. Applies only to SafeNet HSM.
* `id` - The id of the sslhsmkey. It has the same value as the `hsmkeyname` attribute.

## Import

A sslhsmkey can be imported using its hsmkeyname, e.g.

```shell
terraform import citrixadc_sslhsmkey.tf_hsmkey1 hsmkey1
```
