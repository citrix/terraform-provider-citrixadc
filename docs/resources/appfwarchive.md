---
subcategory: "Application Firewall"
---

# Resource: appfwarchive

The `appfwarchive` resource is used to import an Application Firewall tar archive onto the Citrix ADC. It models the NITRO `appfwarchive ?action=Import` action.

Note: The NITRO API exposes no update endpoint for `appfwarchive`. Every attribute is therefore `RequiresReplace` — any change to a configured value forces Terraform to destroy and re-create the archive. To export an existing archive to a target path, use the sibling `citrixadc_appfwarchive_export` resource.


## Example usage

```hcl
resource "citrixadc_appfwarchive" "tf_appfwarchive" {
  name    = "tf_appfwarchive"
  src     = "http://archive.example.com/appfw/tf_appfwarchive.tar"
  comment = "Imported application firewall archive for tf demo"
}
```


## Argument Reference

* `name` - (Required) Name of the tar archive. Forces replacement on change.
* `src` - (Required) Indicates the source of the tar archive file as a URL of the form `<protocol>://<host>[:<port>][/<path>]`. `<protocol>` is `http` or `https`. `<host>` is the DNS name or IP address of the http or https server. `<port>` is the port number of the server; if omitted, the default port for http or https is used. `<path>` is the path of the file on the server. Import will fail if an https server requires client certificate authentication. Forces replacement on change.
* `comment` - (Optional) Comments associated with this archive. Forces replacement on change.
* `target` - (Optional) Path to the file to be exported. This attribute belongs to the export action and is ignored by Import; it is preserved here only for backward compatibility. Use the `citrixadc_appfwarchive_export` resource to export an archive. Forces replacement on change.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwarchive`. It has the same value as the `name` attribute.
