---
subcategory: "NS"
---

# Data Source: nsaigwprofile

The `citrixadc_nsaigwprofile` data source is used to retrieve information about a specific AI GW (AI Gateway) profile configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_nsaigwprofile" "example" {
  name = "my_nsaigwprofile"
}
```

## Argument Reference

* `name` - (Required) Name of the AIGW Profile.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the AI GW profile. It has the same value as the `name` attribute.
* `endpointtype` - The type of AI GW endpoint type.
* `profiletype` - The binding entity for the aigw profile.
* `tokenquota` - Token capacity of the backend server.
* `quotarefreshfrequency` - Quota refresh rate, in minutes.
* `authtoken` - Authentication token/API Key for the AI GW Endpoint (sensitive).
