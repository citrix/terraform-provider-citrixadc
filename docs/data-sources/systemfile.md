---
subcategory: "System"
---

# Data Source: systemfile

The systemfile data source allows you to retrieve information about files on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_systemfile" "example" {
  filelocation = "/var/tmp"
  filename     = "hello.txt"
}

output "fileencoding" {
  value = data.citrixadc_systemfile.example.fileencoding
}

output "filecontent" {
  value = data.citrixadc_systemfile.example.filecontent
}
```

## Argument Reference

* `filelocation` - (Required) Location of the file on Citrix ADC.
* `filename` - (Required) Name of the file. It should not include filepath.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `filecontent` - File content in Base64 format.
* `fileencoding` - Encoding type of the file content.
* `id` - The id of the systemfile. It has a composite value of `<filelocation>,<filename>`.
