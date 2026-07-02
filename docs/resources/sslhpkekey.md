---
subcategory: "SSL"
---

# Resource: sslhpkekey

The sslhpkekey resource installs a Hybrid Public Key Encryption (HPKE) key on the Citrix ADC. This key is used by the Encrypted Client Hello (ECH) feature to decrypt the inner ClientHello, allowing the ADC to terminate TLS connections in which the client has concealed the server name and other handshake details.

All attributes are immutable: changing any of them forces the key to be recreated.


## Example usage

```hcl
resource "citrixadc_sslhpkekey" "tf_hpkekey" {
  hpkekeyname = "ech_hpkekey1"
  file        = "/nsconfig/ssl/ech_hpke.key"
  dhkem       = "X_25519"
}
```


## Argument Reference

* `hpkekeyname` - (Required) The name of the HPKE key configured on the appliance that is used to decrypt ECH. Changing this attribute forces a new resource to be created.
* `file` - (Required) Name of (and, optionally, path to) the HPKE key file present on the appliance filesystem. Changing this attribute forces a new resource to be created.
* `dhkem` - (Required) Type of curve used for HPKE. Possible values: [ X_25519 ]. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslhpkekey. It has the same value as the `hpkekeyname` attribute.


## Import

An sslhpkekey can be imported using its hpkekeyname, e.g.

```shell
terraform import citrixadc_sslhpkekey.tf_hpkekey ech_hpkekey1
```
