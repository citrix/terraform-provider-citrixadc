---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_restvalidation_binding

The appfwprofile_restvalidation_binding data source allows you to retrieve information about a REST/API schema validation relaxation rule bound to a Citrix ADC application firewall profile.

## Example usage

```terraform
data "citrixadc_appfwprofile_restvalidation_binding" "tf_binding" {
  name                   = "tf_appfwprofile"
  rest_validation_action = "log"
  restvalidation         = "GET:/v1/bookstore/viewbooks"
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_restvalidation_binding.tf_binding.alertonly
}

output "isautodeployed" {
  value = data.citrixadc_appfwprofile_restvalidation_binding.tf_binding.isautodeployed
}

output "resourceid" {
  value = data.citrixadc_appfwprofile_restvalidation_binding.tf_binding.resourceid
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `rest_validation_action` - (Required) Action to be taken for traffic matching the configured relaxation rule. Possible values: [ log, none ]
* `restvalidation` - (Required) Exempt REST endpoints or any other URLs matching the given pattern from the API schema validation check. Example: `GET:/v1/bookstore/viewbooks`.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_restvalidation_binding. It is the concatenation of the `name`, `rest_validation_action` and `restvalidation` attributes separated by a comma.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Indicates whether the rule was auto deployed by a dynamic profile. Possible values: [ AUTODEPLOYED, NOTAUTODEPLOYED ]
* `resourceid` - A server-assigned identifier that identifies the rule.
* `state` - Enabled. Possible values: [ ENABLED, DISABLED ]
