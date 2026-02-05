---
subcategory: "NS"
---

# Data Source `nsencryptionparams`

The nsencryptionparams data source allows you to retrieve information about NetScaler encryption parameters configuration.


## Example usage

```terraform
data "citrixadc_nsencryptionparams" "tf_nsencryptionparams" {
}

output "method" {
  value = data.citrixadc_nsencryptionparams.tf_nsencryptionparams.method
}

output "keyvalue" {
  value = data.citrixadc_nsencryptionparams.tf_nsencryptionparams.keyvalue
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `method` - Cipher method (and key length) to be used to encrypt and decrypt content. The default value is AES256.
* `keyvalue` - The base64-encoded key generation number, method, and key value. Note: Do not include this argument if you are changing the encryption method. To generate a new key value for the current encryption method, specify an empty string ("") as the value of this parameter. The parameter is passed implicitly, with its automatically generated value, to the Citrix ADC packet engines even when it is not included in the command. Passing the parameter to the packet engines enables the appliance to save the key value to the configuration file and to propagate the key value to the secondary appliance in a high availability setup.

## Attribute Reference

* `id` - The id of the nsencryptionparams. It is a system-generated identifier.
