---
subcategory: "System"
---

# Resource: systemkek_change

Rotates the appliance Key Encryption Key (KEK) on a Citrix ADC. The KEK is the
master key used to protect other secrets stored on the appliance. Applying this
resource backs up the existing keys and generates brand-new ones, which is
typically done as part of a security-hardening or key-rotation policy.

~> **WARNING: This operation is IRREVERSIBLE and NON-IDEMPOTENT.** Each `terraform
apply` that creates (or recreates) this resource ROTATES the appliance KEK. The
old keys are backed up and new keys are generated every time. There is no way to
roll a KEK rotation back. Apply this resource only when you intend to rotate the
KEK.

~> **NOTE:** NITRO exposes no GET endpoint for this resource, so the rotation
cannot be read back and Terraform cannot detect drift. The `level` attribute is
marked `RequiresReplace`: any change forces the resource to be destroyed and
recreated, which triggers a fresh KEK rotation. Destroying the resource only
removes it from Terraform state; it does not undo the rotation on the appliance.

~> **NOTE:** Using `level = "extended"` additionally rewrites the configuration
database (`ns.conf`, `nscfg.db`, and all `ns.conf` files for the same release)
across every partition. While the extended rotation runs, the appliance BLOCKS
all configuration changes until the operation completes.


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

* `id` - A synthetic constant identifier for the resource. It is always set to
  `systemkek_change`. Because there is no GET endpoint, this value is not derived
  from the appliance.
