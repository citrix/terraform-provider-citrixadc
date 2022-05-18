---
subcategory: "Basic"
---

# Resource: locationparameter

The locationparameter resource is used to create location parameter resource.


## Example usage

```hcl
resource "citrixadc_locationparameter" "tf_locationpara" {
  context            = "geographic"
  q1label            = "asia"
  matchwildcardtoany = "YES"
}
```


## Argument Reference

* `context` - (Optional) Context for describing locations. In geographic context, qualifier labels are assigned by default in the following sequence: Continent.Country.Region.City.ISP.Organization. In custom context, the qualifiers labels can have any meaning that you designate. Possible values: [ geographic, custom ]
* `q1label` - (Optional) Label specifying the meaning of the first qualifier. Can be specified for custom context only. Minimum length =  1
* `q2label` - (Optional) Label specifying the meaning of the second qualifier. Can be specified for custom context only. Minimum length =  1
* `q3label` - (Optional) Label specifying the meaning of the third qualifier. Can be specified for custom context only. Minimum length =  1
* `q4label` - (Optional) Label specifying the meaning of the fourth qualifier. Can be specified for custom context only. Minimum length =  1
* `q5label` - (Optional) Label specifying the meaning of the fifth qualifier. Can be specified for custom context only. Minimum length =  1
* `q6label` - (Optional) Label specifying the meaning of the sixth qualifier. Can be specified for custom context only. Minimum length =  1
* `matchwildcardtoany` - (Optional) Indicates whether wildcard qualifiers should match any other qualifier including non-wildcard while evaluating location based expressions. Yes - Wildcard qualifiers match any other qualifiers. No  - Wildcard qualifiers do not match non-wildcard qualifiers, but match other wildcard qualifiers. Expression - Wildcard qualifiers in an expression match any qualifier in an LDNS location, wildcard qualifiers in the LDNS location do not match non-wildcard qualifiers in an expression. Possible values: [ YES, NO, Expression ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the locationparameter. It is a unique string prefixed with "tf-locationparameter-"
