---
subcategory: "System"
---

# Resource: systemkek_change

This resource is used to rotate the appliance Key Encryption Key (KEK) on the Citrix ADC.

!> **WARNING:** This operation is IRREVERSIBLE and NON-IDEMPOTENT — each apply backs up the old keys and generates new ones. A KEK rotation cannot be rolled back.


## Example usage

```hcl
resource "citrixadc_systemkek_change" "tf_systemkek_change" {
  level = "basic"
}
```

To perform an extended rotation that also rewrites the configuration database
across all partitions:

```hcl
resource "citrixadc_systemkek_change" "tf_systemkek_change" {
  level = "extended"
}
```


## Argument Reference

* `level` - (Required) Type of update KEK to be performed. Changing this value
  forces a new resource to be created, which rotates the KEK again.
  * `basic` - Backs up the old keys, creates new keys, and responds back.
  * `extended` - Backs up the old keys, creates new keys, and updates `ns.conf`,
    `nscfg.db`, and all `ns.conf` files for the same release in all partitions.
    All configuration changes are blocked while this runs.

  Possible values: [ basic, extended ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemkek_change resource. It is set to `systemkek_change`.
