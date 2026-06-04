---
subcategory: "Application Firewall"
---

# Resource: appfwcustomsettings

Exports the Application Firewall custom settings configuration from the Citrix ADC to a target location. Applying this resource executes the NITRO `export` action, which writes the named custom-settings object to the specified target file so it can be archived, version-controlled, or transferred to another appliance. This is an action-only resource: it triggers a one-shot export side effect rather than creating a persistent configuration object on the appliance.

## Example usage

```hcl
resource "citrixadc_appfwcustomsettings" "tf_appfwcustomsettings" {
  name   = "custom_settings1"
  target = "/var/tmp/custom_settings1.export"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the Application Firewall custom settings object to export. Changing this value forces the export action to run again (the resource is replaced).
* `target` - (Required) Path of the file on the Citrix ADC to which the custom settings are exported. Changing this value forces the export action to run again (the resource is replaced).

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwcustomsettings resource. It has the same value as the `name` attribute.

## Behavior Notes

This is an **action-only export resource**. The NITRO API exposes only the `?action=export` operation for `appfwcustomsettings`; there is no GET, update, or inverse (delete) endpoint. As a result the resource lifecycle behaves as follows:

* **Create (apply)** - Executes the `export` action against the Citrix ADC, writing the named custom settings to `target`. This is the only operation that contacts the appliance.
* **Read** - No-op. NITRO provides no query endpoint for the export, so the provider simply preserves the existing Terraform state and never performs drift detection.
* **Update** - No-op. There is no update endpoint, and both `name` and `target` are marked `RequiresReplace`. Changing either attribute destroys and re-creates the resource, which re-runs the export action.
* **Destroy** - State-only removal. The export is a one-shot side effect with no inverse NITRO API, so destroying the resource only removes it from Terraform state; the previously exported file on the appliance is left untouched.

Because there is no GET endpoint, **import is not supported** for this resource.
