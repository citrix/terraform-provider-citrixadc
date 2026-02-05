---
subcategory: "Basic"
---

# Data Source `location`

The location data source allows you to retrieve information about a location configuration on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_location" "tf_location" {
  ipfrom = "8.8.8.8"
}

output "ipfrom" {
  value = data.citrixadc_location.tf_location.ipfrom
}

output "ipto" {
  value = data.citrixadc_location.tf_location.ipto
}

output "preferredlocation" {
  value = data.citrixadc_location.tf_location.preferredlocation
}
```


## Argument Reference

* `ipfrom` - (Required) First IP address in the range, in dotted decimal notation.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the location. It has the same value as the `ipfrom` attribute.
* `ipto` - Last IP address in the range, in dotted decimal notation.
* `latitude` - Numerical value, in degrees, specifying the latitude of the geographical location of the IP address-range. Note: Longitude and latitude parameters are used for selecting a service with the static proximity GSLB method. If they are not specified, selection is based on the qualifiers specified for the location.
* `longitude` - Numerical value, in degrees, specifying the longitude of the geographical location of the IP address-range. Note: Longitude and latitude parameters are used for selecting a service with the static proximity GSLB method. If they are not specified, selection is based on the qualifiers specified for the location.
* `preferredlocation` - String of qualifiers, in dotted notation, describing the geographical location of the IP address range. Each qualifier is more specific than the one that precedes it, as in continent.country.region.city.isp.organization. For example, "NA.US.CA.San Jose.ATT.citrix". Note: A qualifier that includes a dot (.) or space ( ) must be enclosed in double quotation marks.
