---
subcategory: "NTP"
---

# Resource: ntpparam

The ntpparam resource is used to update ntpparam.


## Example usage

```hcl
resource "citrixadc_ntpparam" "tf_ntpparam" {
  authentication = "YES"
  trustedkey     = [123, 456]
  autokeylogsec  = 15
  revokelogsec   = 20
}
```


## Argument Reference

* `authentication` - (Optional) Apply NTP authentication, which enables the NTP client (Citrix ADC) to verify that the server is in fact known and trusted. Possible values: [ YES, NO ]
* `trustedkey` - (Optional) Key identifiers that are trusted for server authentication with symmetric key cryptography in the keys file. Minimum value =  1 Maximum value =  65534
* `autokeylogsec` - (Optional) Autokey protocol requires the keys to be refreshed periodically. This parameter specifies the interval between regenerations of new session keys. In seconds, expressed as a power of 2. Minimum value =  0 Maximum value =  32
* `revokelogsec` - (Optional) Interval between re-randomizations of the autokey seeds to prevent brute-force attacks on the autokey algorithms. Minimum value =  0 Maximum value =  32


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ntpparam. It is a unique string prefixed with  `tf-ntpparam-`.