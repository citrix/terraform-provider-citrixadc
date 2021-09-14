---
subcategory: "SSL"
---

# Resource: sslprofile_sslcertkey_binding

The sslprofile_sslcertkey_binding resource is used to create bindings between sslprofile and sslcertkey.


## Example usage

```hcl
resource "citrixadc_sslprofile" "tf_sslprofile" {
  name = "tfUnit_sslprofile-hello"

  // `ecccurvebindings` is REQUIRED attribute.
  // The default ecccurvebindings will be DELETED and only the explicitly given ecccurvebindings will be retained
  // To unbind all the ecccurvebindings, an empty list `[]` is to be assinged to `ecccurvebindings` attribute
  ecccurvebindings = ["P_256"]
  sslinterception = "ENABLED"

}

resource "citrixadc_sslcertkey" "tf_sslcertkey" {
	certkey = "tf_sslcertkey"
	cert = "/nsconfig/ssl/ns-root.cert"
	key = "/nsconfig/ssl/ns-root.key"
}
  
resource "citrixadc_sslprofile_sslcertkey_binding" "demo_sslprofile_sslcertkey_binding" {
	name = citrixadc_sslprofile.tf_sslprofile.name
	sslicacertkey = citrixadc_sslcertkey.tf_sslcertkey.certkey 
}
```


## Argument Reference

* `sslicacertkey` - (Required) The certkey (CA certificate + private key) to be used for SSL interception.
* `name` - (Required) Name of the SSL profile.
* `cipherpriority` - (Optional) Priority of the cipher binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslprofile_sslcertkey_binding. IIt has is the conatenation of the `name` and `sslicacertkey` attributes.


## Import

A sslprofile_sslcertkey_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslprofile_sslcertkey_binding.tf_sslprofile_sslcertkey_binding tf_sslprofile_sslcertkey_binding
```
