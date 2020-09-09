---
subcategory: "System"
---

# Resource: systemfile

The systemfile resource is used to upload files to the target ADC.


## Example usage

```hcl
resource "citrixadc_systemfile" "tf_file" {
    filename = "hello.txt"
    filelocation = "/var/tmp"
    filecontent = "hello"
}
```


## Argument Reference

* `filename` - (Optional) Name of the file. It should not include filepath.
* `filecontent` - (Optional) File contents.
* `filelocation` - (Optional) Location of the file on Citrix ADC.
* `fileencoding` - (Optional) Encoding type of the file content. Defaults to `BASE64`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemfile. It is the fullpath of the system file.


## Import

A systemfile can be imported using its full path, e.g.

```shell
terraform import citrixadc_systemfile.tf_file /var/tmp/hello.txt
```
