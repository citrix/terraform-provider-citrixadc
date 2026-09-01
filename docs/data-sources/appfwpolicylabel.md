---
subcategory: "Application Firewall"
---

# Data Source: appfwpolicylabel

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

### Read-only appfwpolicylabel metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_appfwpolicylabel` resource). They are GET-only/Computed and are `null` when the appliance does not return them.

* `numpol` - Number of polices bound to label.
* `hits` - Number of times policy label was invoked.
* `priority` - Positive integer specifying the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `labeltype` - Type of policy label to invoke if the current policy evaluates to TRUE and the invoke parameter is set.
* `invoke_labelname` - Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.
* `description` - Description of the policylabel.
* `policytype` - Policy type (for example `Classic Policy`, `Advanced Policy`).
