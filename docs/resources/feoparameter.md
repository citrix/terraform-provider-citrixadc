---
subcategory: "Front-end-optimization"
---

# Resource: feoparameter

The feoparameter resource is used to update feoparameter.


## Example usage

```hcl
resource "citrixadc_feoparameter" "tf_feoparameter" {
  jpegqualitypercent = 10
  cssinlinethressize = 100
  jsinlinethressize  = 50
  imginlinethressize = 1
}

```


## Argument Reference

* `jpegqualitypercent` - (Optional) The percentage value of a JPEG image quality to be reduced. Range: 0 - 100. Minimum value =  0 Maximum value =  100 Default value = 75
* `cssinlinethressize` - (Optional) Threshold value of the file size (in bytes) for converting external CSS files to inline CSS files. Minimum value =  1 Maximum value =  2048
* `jsinlinethressize` - (Optional) Threshold value of the file size (in bytes), for converting external JavaScript files to inline JavaScript files. Minimum value =  1 Maximum value =  2048
* `imginlinethressize` - (Optional) Maximum file size of an image (in bytes), for coverting linked images to inline images. Minimum value =  1 Maximum value =  2048


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the feoparameter. It is a unique string prefixed with  `tf-feoparameter-` attribute.
