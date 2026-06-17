---
subcategory: "NS"
---

# Data Source: nslimitsessions

The nslimitsessions data source allows you to retrieve the active rate-limit session statistics tracked on the Citrix ADC for a given rate-limit identifier, such as the accumulated hit and drop counters and the matched selectlet values.


## Example usage

```terraform
data "citrixadc_nslimitsessions" "example" {
  limitidentifier = "myratelimit"
}

output "nslimitsessions_hits" {
  value = data.citrixadc_nslimitsessions.example.hits
}
```


## Argument Reference

* `limitidentifier` - (Required) Name of the rate limit identifier for which to display the sessions.
* `detail` - (Optional) Show the individual hash values.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the data source. It has the same value as the `limitidentifier` attribute.
* `timeout` - The time remaining on the session before a flush can be attempted.
* `hits` - The number of times this entry was hit.
* `drop` - The number of times action was taken.
* `name` - The string formed by gathering selectlet values.
* `unit` - Total computed hash of the matched selectlets.
