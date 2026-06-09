---
subcategory: "High Availability"
---

# Resource: hafiles

Synchronizes configuration files (such as SSL certificates, Access Gateway bookmarks, application firewall imports, and license files) from the current node to its peer across a Citrix ADC high-availability (HA) pair. This keeps both HA nodes in sync so the secondary node is ready to take over with the same file-based configuration as the primary.

This is an **action-only** resource: applying it triggers a one-time `sync` action on the ADC. It does **not** manage persistent server-side state. Because NITRO exposes no GET endpoint for this action, the resource performs no read-back (drift cannot be detected), there is no update, and Read/Delete are state-only no-ops. Changing the `mode` argument forces the resource to be re-created, which re-runs the sync.

This resource requires an HA setup and operates on an HA node.


## Example usage

Synchronize all file types:

```hcl
resource "citrixadc_hafiles" "sync_all" {
  mode = ["all"]
}
```

Synchronize only specific file types:

```hcl
resource "citrixadc_hafiles" "sync_selected" {
  mode = ["ssl", "bookmarks"]
}
```

Omit `mode` to synchronize all file types (equivalent to `mode = ["all"]`):

```hcl
resource "citrixadc_hafiles" "sync_default" {}
```


## Argument Reference

* `mode` - (Optional) A list of synchronization modes that select which file types are synchronized across the HA pair. If omitted, all file types are synchronized. Possible values: [ all, bookmarks, ssl, imports, misc, dns, krb, AAA, app_catalog, all_plus_misc, all_minus_misc ]. Selected values:
  * `all` - Synchronize files related to system configuration, Access Gateway bookmarks, SSL certificates, SSL CRL lists, and application firewall XML objects.
  * `bookmarks` - Synchronize all Access Gateway bookmarks.
  * `ssl` - Synchronize all certificates, keys, and CRLs for the SSL feature.
  * `imports` - Synchronize all XML objects (for example, WSDLs, schemas, error pages) configured for the application firewall.
  * `misc` - Synchronize all license files and the `rc.conf` file.
  * `all_plus_misc` - Synchronize files related to system configuration, Access Gateway bookmarks, SSL certificates, SSL CRL lists, application firewall XML objects, licenses, and the `rc.conf` file.

  Changing `mode` forces a new resource to be created (the sync action is re-run).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier with the constant value `"hafiles"`. This action-only resource has no server-assigned identity.
