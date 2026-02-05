---
subcategory: "AppQoE"
---

# Data Source `appqoecustomresp`

The appqoecustomresp data source allows you to retrieve information about an existing appqoecustomresp.


## Example usage

```terraform
data "citrixadc_appqoecustomresp" "tf_appqoecustomresp" {
  name = "my_appqoecustomresp"
}

output "name" {
  value = data.citrixadc_appqoecustomresp.tf_appqoecustomresp.name
}

output "src" {
  value = data.citrixadc_appqoecustomresp.tf_appqoecustomresp.src
}
```


## Argument Reference

* `name` - (Required) Indicates name of the custom response HTML page to import/update.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appqoecustomresp. It has the same value as the `name` attribute.
* `src` - Source location of the custom response HTML page file to import. The file can be stored on the local file system or a remote server.
