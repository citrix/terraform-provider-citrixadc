---
subcategory: "Application Firewall"
---

# Data Source: appfwprotofile

The `citrixadc_appfwprotofile` data source is used to retrieve information about an existing Application Firewall gRPC schema (proto) file imported on the Citrix ADC.

The underlying NITRO `get` endpoint only echoes back the `name` and `src` fields; the original `comment` and `overwrite` inputs are not returned by the ADC and therefore are not available through this data source.

The configured `name` must refer to a gRPC schema object that already exists on the ADC. Reference the corresponding `citrixadc_appfwprotofile` resource attribute (as shown below) so Terraform creates the dependency before the data source is read.


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
