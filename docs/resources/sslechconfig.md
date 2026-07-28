---
subcategory: "SSL"
---

# Resource: sslechconfig

This resource is used to manage Encrypted Client Hello (ECH) configurations on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_sslechconfig" "tf_sslechconfig" {
  echconfigname = "echconfig1"
  echcipher     = "AES_128_GCM"
  hpkekeyname   = "hpkekey1"
  echpublicname = "public.example.com"
  echconfigid   = 1
  version       = 65037
}
```

The `hpkekeyname` argument references an existing `citrixadc_sslhpkekey` resource. A more complete example wiring the dependency:

```hcl
resource "citrixadc_sslhpkekey" "tf_hpkekey" {
  name = "hpkekey1"
  # ... HPKE key configuration ...
}

resource "citrixadc_sslechconfig" "tf_sslechconfig" {
  echconfigname = "echconfig1"
  echcipher     = "AES_128_GCM"
  hpkekeyname   = citrixadc_sslhpkekey.tf_hpkekey.name
  echpublicname = "public.example.com"
  echconfigid   = 1
}
```


## Argument Reference

* `echconfigname` - (Required) The name of the ECH configuration. Changing this value forces a new resource to be created.
* `echcipher` - (Required) The supported cipher suite that encrypts the client Hello message. Changing this value forces a new resource to be created.
* `hpkekeyname` - (Required) The name of the configured HPKE key whose public key is used to encrypt the ClientHello. References an `sslhpkekey` resource. Changing this value forces a new resource to be created.
* `echpublicname` - (Required) The public name of the ECH config, expressed as an FQDN or any string. This is the cleartext outer SNI presented to on-path observers. Changing this value forces a new resource to be created.
* `echconfigid` - (Required) The config id of the ECH config. Minimum value = `0`. Maximum value = `255`. Changing this value forces a new resource to be created.
* `version` - (Optional) The version of ECH for which this configuration is used. Defaults to `65037`. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslechconfig. It has the same value as the `echconfigname` attribute.


## Import

An sslechconfig can be imported using its echconfigname, e.g.

```shell
terraform import citrixadc_sslechconfig.tf_sslechconfig echconfig1
```
