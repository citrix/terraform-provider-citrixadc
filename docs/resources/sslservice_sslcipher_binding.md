---
subcategory: "SSL"
---

# Resource: sslservice_sslcipher_binding

Binds an individual cipher, a user-defined cipher group, or a predefined (built-in) cipher alias directly to an SSL service, controlling which cipher suites the service negotiates during the TLS handshake. Use this to tailor the cipher set of a single SSL service independently of the cipher list configured on its SSL profile.

~> Direct cipher binding on an SSL service takes effect only when the default SSL profile feature is disabled. When the default SSL profile is enabled, ciphers are managed through the SSL profile (`sslprofile_sslcipher_binding`) rather than bound directly to the service.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "10.10.10.44"
  port        = 443
  servicetype = "SSL"
}

resource "citrixadc_service" "tf_service" {
  name        = "tf_service"
  servicetype = "SSL"
  port        = 443
  ip          = "10.77.33.22"
  lbvserver   = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_sslservice_sslcipher_binding" "tf_binding" {
  servicename = citrixadc_service.tf_service.name
  ciphername  = "HIGH"
}
```


## Argument Reference

* `servicename` - (Required) Name of the SSL service for which to set advanced configuration. Changing this forces a new resource to be created.
* `ciphername` - (Required) Name of the individual cipher, user-defined cipher group, or predefined (built-in) cipher alias. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservice_sslcipher_binding. It is the concatenation of the `ciphername` and `servicename` unique attributes, formatted as comma-separated `key:value` pairs (for example, `ciphername:HIGH,servicename:tf_service`).
* `cipheraliasname` - (Read-only) The cipher group/alias/individual cipher configuration. This value is returned by the Citrix ADC and cannot be set; the NITRO server rejects it in the bind payload.
* `cipherdefaulton` - (Read-only) Flag indicating whether the bound cipher was the DEFAULT cipher, bound at boot time, or any other cipher from the CLI. This value is returned by the Citrix ADC and cannot be set.
* `description` - (Read-only) The cipher suite description. This value is returned by the Citrix ADC and cannot be set.


## Import

A sslservice_sslcipher_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslservice_sslcipher_binding.tf_binding ciphername:HIGH,servicename:tf_service
```
