---
subcategory: "VPN"
---

# Data Source: vpnsecureprivateaccessprofile

The `citrixadc_vpnsecureprivateaccessprofile` data source is used to retrieve information about a specific Secure Private Access profile configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_vpnsecureprivateaccessprofile" "example" {
  name = "my_spaprofile"
}
```

## Argument Reference

* `name` - (Required) The name of the Secure Private Access profile.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the Secure Private Access profile.
* `url` - Public URL for your Secure Private Access site or load balancing server.
* `customerid` - Customer ID of the citrix cloud customer.
* `chromeenterprisepremiummode` - Secure Private Access Chrome Enterprise Premium mode of operation.
* `googlecustomerid` - Your organization's unique ID on Google's Admin console Profile settings.
* `googlesecuritygatewayid` - The ID of the Google Secure Gateway.
* `forceclienttype` - Automatically configures the session for Citrix Secure Access client connectivity.
