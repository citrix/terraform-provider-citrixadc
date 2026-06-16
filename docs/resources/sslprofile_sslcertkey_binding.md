---
subcategory: "SSL"
---

# Resource: sslprofile\_sslcertkey\_binding

The sslprofile\_sslcertkey\_binding resource is used to bind an SSL certificate key to an SSL profile.


## Example usage

```hcl
resource "citrixadc_sslprofile" "tf_sslprofile" {
  name = "tfUnit_sslprofile-hello"

  // `ecccurvebindings` is a REQUIRED attribute.
  // The default ecccurvebindings will be DELETED and only the explicitly given ecccurvebindings will be retained.
  // To unbind all the ecccurvebindings, an empty list `[]` is to be assigned to the `ecccurvebindings` attribute.
  ecccurvebindings = ["P_256"]
  sslinterception  = "ENABLED"
}

resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert    = "/nsconfig/ssl/ns-root.cert"
  key     = "/nsconfig/ssl/ns-root.key"
}

resource "citrixadc_sslprofile_sslcertkey_binding" "tf_binding" {
  name          = citrixadc_sslprofile.tf_sslprofile.name
  sslicacertkey = citrixadc_sslcertkey.tf_sslcertkey.certkey
}
```


## Argument Reference

* `name` - (Required) Name of the SSL profile.
* `sslicacertkey` - (Required) The certkey (CA certificate + private key) to be used for SSL interception.
* `cipherpriority` - (Optional) Priority of the cipher binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslprofile\_sslcertkey\_binding. It is the concatenation of the `name` and `sslicacertkey` attributes separated by a comma.


## Import

A sslprofile\_sslcertkey\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslprofile_sslcertkey_binding.tf_binding tfUnit_sslprofile-hello,tf_sslcertkey
```
