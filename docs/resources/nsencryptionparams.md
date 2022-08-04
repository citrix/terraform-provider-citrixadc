---
subcategory: "NS"
---

# Resource: nsencryptionparams

The nsencryptionparams resource is used to create nsencryptionparams.


## Example usage

```hcl
resource "citrixadc_nsencryptionparams" "tf_nsencryptionparams" {
   method = "DES-CFB"
   keyvalue = ""
}
```


## Argument Reference

* `method` - (Required) Cipher method (and key length) to be used to encrypt and decrypt content. The default value is AES256. Possible values: [ NONE, RC4, DES3, AES128, AES192, AES256, DES, DES-CBC, DES-CFB, DES-OFB, DES-ECB, DES3-CBC, DES3-CFB, DES3-OFB, DES3-ECB, AES128-CBC, AES128-CFB, AES128-OFB, AES128-ECB, AES192-CBC, AES192-CFB, AES192-OFB, AES192-ECB, AES256-CBC, AES256-CFB, AES256-OFB, AES256-ECB ]
* `keyvalue` - (Required) The base64-encoded key generation number, method, and key value. Note: * Do not include this argument if you are changing the encryption method. * To generate a new key value for the current encryption method, specify an empty string \(""\) as the value of this parameter. The parameter is passed implicitly, with its automatically generated value, to the Citrix ADC packet engines even when it is not included in the command. Passing the parameter to the packet engines enables the appliance to save the key value to the configuration file and to propagate the key value to the secondary appliance in a high availability setup.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsencryptionparams. It is a unique string prefixed with `tf-nsencryptionparams-`.


