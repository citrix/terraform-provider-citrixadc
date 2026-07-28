---
subcategory: "Policy"
---

# Data Source: policyurlset

The policyurlset data source allows you to retrieve information about a URL set.


## Example usage

```terraform
data "citrixadc_policyurlset" "example" {
  name = "blocklist_urlset"
}

output "policyurlset_interval" {
  value = data.citrixadc_policyurlset.example.interval
}
```


## Argument Reference

* `name` - (Required) Unique name of the url set to look up.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policyurlset. It has the same value as the `name` attribute.
* `canaryurl` - The URL added to this urlset for testing when its contents are kept confidential.
* `comment` - Any comments preserved about this url set.
* `delimiter` - CSV file record delimiter.
* `rowseparator` - CSV file row separator.
* `interval` - The interval, in seconds, at which the update of the urlset occurs.
* `matchedid` - An ID that is sent to AppFlow to indicate which URLSet was the last one that matched the requested URL.
* `overwrite` - Whether the import overwrites the existing file.
* `privateset` - Whether this urlset is prevented from being exported.
* `subdomainexactmatch` - Whether exact subdomain matching is enforced.
* `imported` - When set, indicates the urlset has been imported.
