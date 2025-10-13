---
subcategory: "SSL"
---

# Resource: sslhsmkey

The sslhsmkey resource is used to create SSL HSM key.


## Example usage

```hcl
resource "citrixadc_sslhsmkey" "demo_sslhsmkey" {
    hsmkeyname = "hsmk1"
    hsmtype = "SAFENET"
    serialnum = "116877xxxx465464"
    password = "xxxxxxx"
}
```


## Argument Reference

* `hsmkeyname` - (Required) Name for the HSM key. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the HSM key is created. 
* `hsmtype` - (Optional) Type of HSM. Possible values: THALES, SAFENET Default value: THALES
* `key` - (Optional) Name of the key. optionally, for THALES, path to the HSM key file; /var/opt/nfast/kmdata/local/ is the default path. Applies when HSMTYPE is THALES or KEYVAULT.
* `keystore` - (Optional) Name of keystore object representing HSM where key is stored. For example, name of keyvault object or azurekeyvault authentication object. Applies only to KEYVAULT type HSM.
* `password` - (Optional) Password for a partition. Applies only to SAFENET HSM.
* `serialnum` - (Optional) Serial number of the partition on which the key is present. Applies only to SAFENET HSM.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslhsmkey. It has the same value as the `hsmkeyname` attribute.
