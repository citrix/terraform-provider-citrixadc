---
subcategory: "Responder"
---

# Data Source: responderhtmlpage

The `citrixadc_responderhtmlpage` data source is used to retrieve information about an existing responder HTML page object configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_responderhtmlpage" "example" {
  name = "my_responderhtmlpage"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the responder HTML page object to retrieve.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the responder HTML page (same as name).
* `cacertfile` - CA certificate file name which will be used to verify the peer's certificate.
* `comment` - Any comments to preserve information about the HTML page object.
* `overwrite` - Indicates whether the existing file was overwritten.
* `src` - Local path or URL for the file from which to retrieve the imported HTML page.

### Read-only responderhtmlpage metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_responderhtmlpage` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `response` - The imported HTML page response content returned by the appliance.
