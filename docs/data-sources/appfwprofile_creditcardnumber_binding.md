---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_creditcardnumber_binding

The appfwprofile_creditcardnumber_binding data source allows you to retrieve information about creditcard number bindings to an appfw profile.

## Example Usage

```terraform
data "citrixadc_appfwprofile_creditcardnumber_binding" "tf_binding" {
  name                = "tf_appfwprofile"
  creditcardnumber    = "123456789"
  creditcardnumberurl = "www.example.com"
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_creditcardnumber_binding.tf_binding.alertonly
}

output "state" {
  value = data.citrixadc_appfwprofile_creditcardnumber_binding.tf_binding.state
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `creditcardnumber` - (Required) The object expression that is to be excluded from safe commerce check.
* `creditcardnumberurl` - (Required) The url for which the list of credit card numbers are needed to be bypassed from inspection.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_creditcardnumber_binding. It is a system-generated identifier.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
