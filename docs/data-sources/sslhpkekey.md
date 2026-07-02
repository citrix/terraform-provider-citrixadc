---
subcategory: "SSL"
---

# Data Source: sslhpkekey

The sslhpkekey data source allows you to retrieve information about an HPKE key configured on the Citrix ADC. The HPKE key is used by the Encrypted Client Hello (ECH) feature to decrypt the inner ClientHello.


## Example usage

```terraform
data "citrixadc_sslhpkekey" "tf_hpkekey" {
  hpkekeyname = "ech_hpkekey1"
}

output "hpkekey_dhkem" {
  value = data.citrixadc_sslhpkekey.tf_hpkekey.dhkem
}
```


## Argument Reference

* `hpkekeyname` - (Required) The name of the HPKE key configured on the appliance that is used to decrypt ECH.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslhpkekey. It has the same value as the `hpkekeyname` attribute.
* `dhkem` - Type of curve used for HPKE.
* `file` - Name of the HPKE key file.
