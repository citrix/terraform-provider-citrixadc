---
subcategory: "Application Firewall"
---

# Data Source: citrixadc_appfwpolicylabel

The `citrixadc_appfwpolicylabel` data source is used to retrieve information about an existing Application Firewall Policy Label configured on a Citrix ADC appliance.

## Example usage

```hcl
# Retrieve an application firewall policy label by name
data "citrixadc_appfwpolicylabel" "example" {
  labelname = "demo_appfwpolicylabel"
}

# Use the retrieved data in other resources
output "policy_label_type" {
  value = data.citrixadc_appfwpolicylabel.example.policylabeltype
}

```

## Argument Reference

The following arguments are required:

* `labelname` - (Required) Name of the application firewall policy label to retrieve. This is the unique identifier for the policy label.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the application firewall policy label. It has the same value as the `labelname` attribute.
* `policylabeltype` - Type of transformations allowed by the policies bound to the label. Always `http_req` for application firewall policy labels.
* `newname` - The new name of the application firewall policy label (if it has been renamed).
