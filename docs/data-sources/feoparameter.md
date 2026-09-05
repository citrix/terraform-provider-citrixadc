---
subcategory: "Front-end-optimization"
---

# Data Source: feoparameter

This data source is used to retrieve information about the Front End Optimization (FEO) parameters configuration.

## Example Usage

```hcl
data "citrixadc_feoparameter" "example" {
}
```

## Argument Reference

This data source does not require any arguments.

## Attribute Reference

The following attributes are exported:

* `id` - The ID of the FEO parameter.
* `cssinlinethressize` - Threshold value of the file size (in bytes) for converting external CSS files to inline CSS files.
* `imginlinethressize` - Maximum file size of an image (in bytes), for coverting linked images to inline images.
* `jpegqualitypercent` - The percentage value of a JPEG image quality to be reduced. Range: 0 - 100
* `jsinlinethressize` - Threshold value of the file size (in bytes), for converting external JavaScript files to inline JavaScript files.

### Read-only feoparameter metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_feoparameter` resource). They are GET-only/Computed and are `null` when the appliance does not return them.

* `builtin` - Indicates that a variable is a built-in (SYSTEM INTERNAL) type. A list of strings.
* `feature` - The feature to be checked while applying this config.
