---
subcategory: "Basic"
---

# Resource: locationfile_import

The locationfile_import resource is used to import locationfile.


## Example usage

```hcl
resource "citrixadc_locationfile_import" "tf_locationfile_import" {
  locationfile = "my_file"
  src          = "local://my_location_file"
}
```


## Argument Reference

* `locationfile` - (Required) Name of the location file, with or without absolute path. If the path is not included, the default path (/var/netscaler/locdb) is assumed. In a high availability setup, the static database must be stored in the same location on both Citrix ADCs.
* `src` - (Required) URL \(protocol, host, path, and file name\) from where the location file will be imported.             NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the locationfile. It is a unique string prefixed with  `tf-locationfile-e` attribute.
