---
subcategory: "Cluster"
---

# Resource: clusterfiles

Synchronizes configuration files (SSL certificates, bookmarks, DNS records, Kerberos files, and other on-disk artifacts) across all nodes of a Citrix ADC cluster, copying them from the cluster configuration coordinator (CCO) to the remaining cluster nodes.

This is an imperative action resource, not a persistent configuration object. Applying it triggers a one-time file synchronization (the equivalent of the `sync cluster files` CLI command); it does not manage any remote state. Because NITRO exposes no GET or DELETE endpoint for this action, Read and Delete are no-ops (state-only) and changing the `mode` argument forces the resource to be re-created, which re-runs the sync.

This resource must be applied against a cluster node, specifically the cluster configuration coordinator (CCO).


## Example usage

```hcl
resource "citrixadc_clusterfiles" "files_sync" {
  mode = ["all"]
}
```

Synchronize only a subset of file categories:

```hcl
resource "citrixadc_clusterfiles" "ssl_bookmarks_sync" {
  mode = ["ssl", "bookmarks"]
}
```


## Argument Reference

* `mode` - (Optional) A list of strings specifying which directories and files are synchronized across the cluster nodes. When omitted, the ADC defaults to synchronizing `all` files. Possible values: [ all, bookmarks, ssl, imports, misc, dns, krb, AAA, app_catalog, all_plus_misc, all_minus_misc ]. Changing this value forces the resource to be re-created (re-syncing the selected files).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier with the constant value `clusterfiles`. The sync action has no NITRO identity of its own, so this ID does not reference any server-assigned key.
