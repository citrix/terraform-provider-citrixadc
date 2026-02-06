---
subcategory: "Appflow"
---

# Resource: appflowaction_analyticsprofile_binding

The appflowaction_analyticsprofile_binding resource is used to create appflowaction_analyticsprofile_binding.


## Example usage

```hcl

resource "citrixadc_appflowaction_analyticsprofile_binding" "tf_appflowaction_analyticsprofile_binding" {
  name      = citrixadc_appflowaction.tf_appflowaction.name
  analyticsprofile = "ns_analytics_global_profile"
}

resource "citrixadc_appflowaction" "tf_appflowaction" {
  name = "test_action"
  collectors = [citrixadc_appflowcollector.tf_appflowcollector.name,
                citrixadc_appflowcollector.tf_appflowcollector2.name,]
  securityinsight = "ENABLED"
  botinsight      = "ENABLED"
  videoanalytics  = "ENABLED"
}
resource "citrixadc_appflowcollector" "tf_appflowcollector" {
  name      = "tf_collector"
  ipaddress = "192.168.2.2"
  port      = 80
}
resource "citrixadc_appflowcollector" "tf_appflowcollector2" {
  name      = "tf2_collector"
  ipaddress = "192.168.2.3"
  port      = 80
}
```


## Argument Reference

* `analyticsprofile` - (Required) Analytics profile to be bound to the appflow action
* `name` - (Required) Name for the action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow action" or 'my appflow action').


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appflowaction_analyticsprofile_binding is the concatenation of the `name` and `analyticsprofile` attributes separated by a comma.


## Import

A appflowaction_analyticsprofile_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_appflowaction_analyticsprofile_binding.tf_appflowaction_analyticsprofile_binding test_action,ns_analytics_global_profile
```