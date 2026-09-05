---
subcategory: "Basic"
---

# Data Source: locationparameter

The locationparameter data source allows you to retrieve information about the global location parameters configuration.


## Example usage

```terraform

data "citrixadc_locationparameter" "tf_locationpara" {
  depends_on = [citrixadc_locationparameter.tf_locationpara]
}

output "context" {
  value = data.citrixadc_locationparameter.tf_locationpara.context
}

output "q1label" {
  value = data.citrixadc_locationparameter.tf_locationpara.q1label
}

output "matchwildcardtoany" {
  value = data.citrixadc_locationparameter.tf_locationpara.matchwildcardtoany
}
```

## Argument Reference

This datasource does not require any arguments. It retrieves the global location parameters configuration.

## Attribute Reference

The following attributes are available:

* `id` - The id of the locationparameter. It is a system-generated identifier.

* `context` - Context for describing locations. In geographic context, qualifier labels are assigned by default in the following sequence: Continent.Country.Region.City.ISP.Organization. In custom context, the qualifiers labels can have any meaning that you designate. Possible values: `geographic`, `custom`.

* `q1label` - Label specifying the meaning of the first qualifier. Can be specified for custom context only.

* `q2label` - Label specifying the meaning of the second qualifier. Can be specified for custom context only.

* `q3label` - Label specifying the meaning of the third qualifier. Can be specified for custom context only.

* `q4label` - Label specifying the meaning of the fourth qualifier. Can be specified for custom context only.

* `q5label` - Label specifying the meaning of the fifth qualifier. Can be specified for custom context only.

* `q6label` - Label specifying the meaning of the sixth qualifier. Can be specified for custom context only.

* `matchwildcardtoany` - Indicates whether wildcard qualifiers should match any other qualifier including non-wildcard while evaluating location based expressions. Possible values:
  * `YES` - Wildcard qualifiers match any other qualifiers.
  * `NO` - Wildcard qualifiers do not match non-wildcard qualifiers, but match other wildcard qualifiers.
  * `Expression` - Wildcard qualifiers in an expression match any qualifier in an LDNS location, wildcard qualifiers in the LDNS location do not match non-wildcard qualifiers in an expression.

### Read-only locationparameter metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_locationparameter` resource) and are `Computed`/GET-only. Any attribute the appliance does not return is `null`.

* `locationfile` - Currently loaded location database file.
* `format` - Location file format.
* `custom` - Number of configured custom locations.
* `static` - Number of configured locations in the database file (static locations).
* `lines` - Number of lines in the database files.
* `errors` - Number of errors encountered while reading the database file.
* `warnings` - Number of warnings encountered while reading the database file.
* `entries` - Number of successfully added entries.
* `locationfile6` - Currently loaded location database file (IPv6).
* `format6` - Location file format (IPv6).
* `custom6` - Number of configured custom locations (IPv6).
* `static6` - Number of configured locations in the database file (static locations, IPv6).
* `lines6` - Number of lines in the database files (IPv6).
* `errors6` - Number of errors encountered while reading the database file (IPv6).
* `warnings6` - Number of warnings encountered while reading the database file (IPv6).
* `entries6` - Number of successfully added entries (IPv6).
* `flags` - Information needed for display. This argument passes information from the kernel to the user space.
* `status` - Status (success or failure) of database loading.
* `databasemode` - Database mode. Possible values: `File`, `Internal`, `Not applicable`.
* `flushing` - State of flushing. Possible values: `In progress`, `Idle`.
* `loading` - State of loading. Possible values: `In progress`, `Idle`.
* `builtin` - Flags indicating the built-in nature of the configuration. A list of strings. Possible values: `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`.
* `feature` - The feature to be checked while applying this configuration.
