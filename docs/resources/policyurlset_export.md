---
subcategory: "Policy"
---

# Resource: policyurlset_export

The policyurlset_export resource exports the entries of an existing URL set on the Citrix ADC to an external CSV file. Use it when you want to back up or share the contents of a named `policyurlset` object by writing them to a remote location over HTTP, HTTPS or FTP.

~> **One-shot action.** This resource maps to the NITRO `export` action; it does not manage a persistent object, so each `terraform apply` that creates or replaces this resource performs the export once, and changing `name` or `url` forces a new export (replacement).


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
