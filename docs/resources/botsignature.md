---
subcategory: "Bot"
---

# Resource: botsignature

The botsignature resource is used Configuration for bot signatures resource.


## Example usage

```hcl
resource "citrixadc_systemfile" "tf_signature" {
  filename     = "bot_signature.json"
  filelocation = "/var/tmp"
  filecontent  = file("${path.module}/bot_signatures.json")
}
resource "citrixadc_botsignature" "tf_botsignature" {
  name       = "tf_botsignature"
  src        = "local://bot_signature.json"
  depends_on = [citrixadc_systemfile.tf_signature]
  comment    = "TestingExample"
}
```


## Argument Reference

* `name` - (Required) Name to assign to the bot signature file object on the Citrix ADC.
* `comment` - (Optional) Any comments to preserve information about the signature file object.
* `overwrite` - (Optional) Overwrites the existing file
* `src` - (Optional) Local path to and name of, or URL (protocol, host, path, and file name) for, the file in which to store the imported signature file. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botsignature. It has the same value as the `name` attribute.


## Import

A botsignature can be imported using its name, e.g.

```shell
terraform import citrixadc_botsignature.tf_botsignature tf_botsignature
```
