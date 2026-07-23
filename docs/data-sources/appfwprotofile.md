---
subcategory: "Application Firewall"
---

# Data Source: appfwprotofile

The appfwprotofile data source allows you to retrieve information about an Application Firewall gRPC schema (proto) file.

~> **Note:** The NITRO GET endpoint returns only `name` and `src`; the `comment` and `overwrite` inputs are not read back from the ADC.


## Example usage

```hcl
# Look up an appfwprotofile that is managed by Terraform. Referencing the
# resource's name attribute enforces the precondition that the protofile
# exists before this data source attempts to read it.
data "citrixadc_appfwprotofile" "tf_appfwprotofile" {
  name = citrixadc_appfwprotofile.tf_appfwprotofile.name
}

output "appfwprotofile_id" {
  value = data.citrixadc_appfwprotofile.tf_appfwprotofile.id
}

output "appfwprotofile_src" {
  value = data.citrixadc_appfwprotofile.tf_appfwprotofile.src
}
```


## Argument Reference

* `name` - (Required) Name of the gRPC schema object to retrieve.


## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The id of the appfwprotofile. It has the same value as the `name` attribute.
* `src` - Source path of the gRPC schema file that was imported.
* `comment` - Comments associated with this gRPC schema file. Not returned by the NITRO API; present in the schema for symmetry with the resource but typically empty when read.
* `overwrite` - Whether to overwrite any existing gRPC schema object of the same name. Not returned by the NITRO API; present in the schema for symmetry with the resource but typically empty when read.
