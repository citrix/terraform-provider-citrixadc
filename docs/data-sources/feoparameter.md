---
subcategory: "Front End Optimization"
---

# Data Source: citrixadc_feoparameter

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
