---
subcategory: "Policy"
---

# Resource: policyurlset_export

This resource is used to export the entries of a URL set to an external CSV file.


## Example usage

```hcl
resource "citrixadc_policyurlset_export" "tf_policyurlset_export" {
  name = "top_malware_urls"
  url  = "https://backup.example.com/urlsets/top_malware_urls.csv"
}
```


## Argument Reference

* `name` - (Required) Unique name of the url set to export. Maximum length: 127. Changing this value forces the resource to be recreated (re-running the export action against the new url set).
* `url` - (Required) URL (protocol, host, path and file name) to which the CSV file will be exported. HTTP, HTTPS and FTP protocols are supported. Maximum length: 2047. Changing this value forces the resource to be recreated (re-running the export action against the new destination).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the policyurlset_export resource. It has the format `policyurlset_export-<name>` (for example, `policyurlset_export-top_malware_urls`).
