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

### Read-only systemfile metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_systemfile` resource) and are `null` when the appliance omits them.

* `fileaccesstime` - Last access time of the file.
* `filemodifiedtime` - Last modified time of the file.
* `filemode` - File mode. A list of strings (for example `DIRECTORY`).
* `filesize` - Size of the file in BYTES.
