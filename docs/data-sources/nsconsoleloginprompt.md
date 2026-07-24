---
subcategory: "NS"
---

# Data Source: nsconsoleloginprompt

The nsconsoleloginprompt data source allows you to retrieve information about the console login prompt configuration.


## Example usage

```terraform
data "citrixadc_nsconsoleloginprompt" "tf_nsconsoleloginprompt" {
}

output "promptstring" {
  value = data.citrixadc_nsconsoleloginprompt.tf_nsconsoleloginprompt.promptstring
}
```

## Argument Reference

This datasource does not require any arguments. It retrieves the global console login prompt configuration.

## Attribute Reference

The following attributes are exported:

* `id` - The id of the nsconsoleloginprompt. It is a system-generated identifier.

* `promptstring` - Console login prompt string. This string is displayed on the console login screen.
