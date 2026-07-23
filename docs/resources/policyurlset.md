---
subcategory: "Policy"
---

# Resource: policyurlset

This resource is used to manage policy URL sets.


## Example usage

### Using url (sensitive attribute - persisted in state)

```hcl
variable "policyurlset_url" {
  type      = string
  sensitive = true
}

resource "citrixadc_policyurlset" "tf_policyurlset" {
  name      = "blocklist_urlset"
  url       = var.policyurlset_url
  comment   = "Corporate URL block list, refreshed hourly"
  interval  = 3600
  overwrite = true
}
```

with a `terraform.tfvars` such as:

```hcl
policyurlset_url = "https://files.example.com/urlsets/blocklist.csv"
```

### Using url_wo (write-only/ephemeral - NOT persisted in state)

The `url_wo` attribute provides an ephemeral path for the source URL. The value is sent to the Citrix ADC but is **not stored in Terraform state**, reducing the risk of exposing the (potentially credential-bearing) source location. To trigger a re-import when the value changes, increment `url_wo_version`.

```hcl
variable "policyurlset_url" {
  type      = string
  sensitive = true
}

resource "citrixadc_policyurlset" "tf_policyurlset" {
  name           = "blocklist_urlset"
  url_wo         = var.policyurlset_url
  url_wo_version = 1
  comment        = "Corporate URL block list, refreshed hourly"
  interval       = 3600
  overwrite      = true
}
```

To re-import from a new source, update the variable value and bump the version:

```hcl
resource "citrixadc_policyurlset" "tf_policyurlset" {
  name           = "blocklist_urlset"
  url_wo         = var.policyurlset_url
  url_wo_version = 2  # Bumped to trigger a re-import
  comment        = "Corporate URL block list, refreshed hourly"
  interval       = 3600
  overwrite      = true
}
```


## Argument Reference

* `name` - (Required) Unique name of the url set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters. Must not be the name of an existing named expression, pattern set, dataset, string map, or HTTP callout. Changing this attribute forces a new resource to be created.
* `url` - (Optional, Sensitive) URL (protocol, host, path and file name) from where the CSV (comma separated) file will be imported. Each record/line becomes one entry within the urlset; the first field contains the URL pattern and subsequent fields contain the metadata, if available. HTTP, HTTPS and FTP protocols are supported. Note: the operation fails if the destination HTTPS server requires client certificate authentication for access. The value is persisted in Terraform state (encrypted). See also `url_wo` for an ephemeral alternative. Either `url` or `url_wo` must be specified. Changing this attribute forces a new resource to be created.
* `url_wo` - (Optional, Sensitive, WriteOnly) Same as `url`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `url_wo_version`. If both `url` and `url_wo` are set, `url_wo` takes precedence. Either `url` or `url_wo` must be specified. Changing this attribute forces a new resource to be created.
* `url_wo_version` - (Optional) An integer version tracker for `url_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger a re-import. Defaults to `1`. Changing this attribute forces a new resource to be created.
* `overwrite` - (Optional) Overwrites the existing file. Changing this attribute forces a new resource to be created.
* `delimiter` - (Optional) CSV file record delimiter. Defaults to `44` (the ASCII code for a comma). Changing this attribute forces a new resource to be created.
* `rowseparator` - (Optional) CSV file row separator. Defaults to `10` (the ASCII code for a newline). Changing this attribute forces a new resource to be created.
* `interval` - (Optional) The interval, in seconds, rounded down to the nearest 15 minutes, at which the update of the urlset occurs. Defaults to `0` (no periodic update). Changing this attribute forces a new resource to be created.
* `privateset` - (Optional) Prevent this urlset from being exported. Changing this attribute forces a new resource to be created.
* `subdomainexactmatch` - (Optional) Force exact subdomain matching. For example, given an entry `google.com` in the urlset, a request to `news.google.com` will not match when this is set. Changing this attribute forces a new resource to be created.
* `matchedid` - (Optional) An ID that is sent to AppFlow to indicate which URLSet was the last one that matched the requested URL. Defaults to `1`. Changing this attribute forces a new resource to be created.
* `canaryurl` - (Optional) Add this URL to the urlset. Used for testing when the contents of the urlset are kept confidential. Changing this attribute forces a new resource to be created.
* `comment` - (Optional) Any comments to preserve information about this url set. Changing this attribute forces a new resource to be created.
* `imported` - (Optional) When set, the display shows all imported urlsets. This is a query-only filter argument; it is not part of the import payload.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policyurlset. It has the same value as the `name` attribute.


## Import

A policyurlset can be imported using its name, e.g.

```shell
terraform import citrixadc_policyurlset.tf_policyurlset blocklist_urlset
```

Note: the source `url` is a write-only secret that the NITRO API does not return, so it is **not recoverable on import**. After importing, supply the original `url` (or `url_wo` and `url_wo_version`) in your configuration to keep Terraform from re-importing the set on the next apply.
