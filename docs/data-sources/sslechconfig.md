---
subcategory: "SSL"
---

# Data Source: sslechconfig

The sslechconfig data source allows you to retrieve information about an Encrypted Client Hello (ECH) configuration defined on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_sslechconfig" "example" {
  echconfigname = "echconfig1"
}

output "ech_publicname" {
  value = data.citrixadc_sslechconfig.example.echpublicname
}
```


## Argument Reference

* `echconfigname` - (Required) The name of the ECH configuration to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslechconfig. It has the same value as the `echconfigname` attribute.
* `echcipher` - The supported cipher suite that encrypts the client Hello message.
* `echconfigid` - The config id of the ECH config.
* `echpublicname` - The public name of the ECH config (FQDN or any string).
* `hpkekeyname` - The name of the configured HPKE key.
* `version` - The version of ECH for which this configuration is used.
