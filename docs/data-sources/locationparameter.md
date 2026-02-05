---
subcategory: "Basic"
---

# Data Source: citrixadc_locationparameter

The `citrixadc_locationparameter` data source allows you to retrieve information about the global location parameters configuration. These parameters control the behavior of location-based policy evaluation on the Citrix ADC.

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
