---
subcategory: "NS"
---

# Data Source `nsencryptionkey`

The nsencryptionkey data source allows you to retrieve information about encryption keys.


## Example usage

```terraform
data "citrixadc_nsencryptionkey" "my_encryptionkey" {
  name = "my_encryptionkey"
}

output "method" {
  value = data.citrixadc_nsencryptionkey.my_encryptionkey.method
}

output "comment" {
  value = data.citrixadc_nsencryptionkey.my_encryptionkey.comment
}
```


## Argument Reference

* `name` - (Required) Key name. This follows the same syntax rules as other expression entity names: It must begin with an alpha character (A-Z or a-z) or an underscore (_). The rest of the characters must be alpha, numeric (0-9) or underscores. It cannot be re or xp (reserved for regular and XPath expressions). It cannot be an expression reserved word (e.g. SYS or HTTP). It cannot be used for an existing expression object (HTTP callout, patset, dataset, stringmap, or named expression).

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Comments associated with this encryption key.
* `iv` - The initialization vector (IV) for a block cipher, one block of data used to initialize the encryption. The best practice is to not specify an IV, in which case a new random IV will be generated for each encryption. The format must be iv_data or keyid_iv_data to include the generated IV in the encrypted data. The IV should only be specified if it cannot be included in the encrypted data. The IV length is the cipher block size.
* `keyvalue` - The hex-encoded key value. The length is determined by the cipher method. Note that the keyValue will be encrypted when it is saved. There is a special key value AUTO which generates a new random key for the specified method. This kind of key is intended for use cases where the NetScaler both encrypts and decrypts the same data, such as an HTTP header.
* `method` - Cipher method to be used to encrypt and decrypt content. Possible values include NONE, RC4, DES, DES3, AES128, AES192, AES256 (with various modes like CBC, CFB, OFB, ECB).
* `padding` - Enables or disables the padding of plaintext to meet the block size requirements of block ciphers. Valid values are ON and OFF. Default is DEFAULT.
* `id` - The id of the nsencryptionkey. It has the same value as the `name` attribute.
