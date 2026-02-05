---
subcategory: "Content Inspection"
---

# Data Source: citrixadc_contentinspectionparameter

The `citrixadc_contentinspectionparameter` data source allows you to retrieve information about the global Content Inspection parameters configuration. These parameters control the behavior of content inspection on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_contentinspectionparameter" "tf_contentinspectionparameter" {
}

output "undefaction" {
  value = data.citrixadc_contentinspectionparameter.tf_contentinspectionparameter.undefaction
}
```

## Argument Reference

This datasource does not require any arguments. It retrieves the global content inspection parameters configuration.

## Attribute Reference

The following attributes are exported:

* `id` - The id of the contentinspectionparameter. It is a system-generated identifier.

* `undefaction` - Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an error condition in evaluating the expression. Available settings function as follows:
  * `NOINSPECTION` - Do not Inspect the traffic.
  * `RESET` - Reset the connection and notify the user's browser, so that the user can resend the request.
  * `DROP` - Drop the message without sending a response to the user.
