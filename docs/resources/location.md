---
subcategory: "Basic"
---

# Resource: location

The location resource is used to create location resource.


## Example usage

```hcl
resource "citrixadc_location" "tf_location" {
  ipfrom            = "2.3.2.4"
  ipto              = "7.6.5.4"
  preferredlocation = "city"
}
```


## Argument Reference

* `ipfrom` - (Required) First IP address in the range, in dotted decimal notation.
* `ipto` - (Required) Last IP address in the range, in dotted decimal notation.
* `preferredlocation` - (Required) String of qualifiers, in dotted notation, describing the geographical location of the IP address range. Each qualifier is more specific than the one that precedes it, as in continent.country.region.city.isp.organization. For example, "NA.US.CA.San Jose.ATT.citrix".  Note: A qualifier that includes a dot (.) or space ( ) must be enclosed in double quotation marks.
* `latitude` - (Optional) Numerical value, in degrees, specifying the latitude of the geographical location of the IP address-range.  Note: Longitude and latitude parameters are used for selecting a service with the static proximity GSLB method. If they are not specified, selection is based on the qualifiers specified for the location.
* `longitude` - (Optional) Numerical value, in degrees, specifying the longitude of the geographical location of the IP address-range.  Note: Longitude and latitude parameters are used for selecting a service with the static proximity GSLB method. If they are not specified, selection is based on the qualifiers specified for the location.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the location. It has the same value as the `ipfrom` attribute.


## Import

A location can be imported using its ipfrom, e.g.

```shell
terraform import citrixadc_location.tf_location 2.3.2.4
```
