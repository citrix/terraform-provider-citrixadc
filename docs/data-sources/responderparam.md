---
subcategory: "Responder"
---

# Data Source: citrixadc_responderparam

The `citrixadc_responderparam` data source is used to retrieve information about the global responder parameters configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_responderparam" "example" {
}
```

## Argument Reference

No arguments are required for this data source as it retrieves global responder parameters.

## Attribute Reference

The following attributes are exported:

* `id` - The ID of the responder parameters.
* `timeout` - Maximum time in milliseconds to allow for processing all the policies and their selected actions without interruption. Default value: `3900`.
* `undefaction` - Action to perform when policy evaluation creates an UNDEF condition. Possible values:
  * `NOOP` - Send the request to the protected server.
  * `RESET` - Reset the request and notify the user's browser.
  * `DROP` - Drop the request without sending a response to the user.
  Default value: `NOOP`.
