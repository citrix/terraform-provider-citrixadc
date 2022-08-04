---
subcategory: "AppFlow"
---

# Resource: appflowpolicy

The appflowpolicy resource is used to create appflowpolicy.


## Example usage

```hcl
resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
  name      = "test_policy"
  action    = "test_action"
  rule      = "client.TCP.DSTPORT.EQ(22)"
}

# -------------------- ADC CLI ----------------------------
#add appflow collector tf_collector -IPAddress 192.168.2.2
#add appflowaction test_action -collectors tf_collector

# ---------------- NOT YET IMPLEMENTED -------------------
# resource "citrixadc_appflowaction" "tf_appflowaction" {
#   name = "test_action"
#   collectors     = [citrixadc_appflowcollector.tf_appflowcollector.name,
#                    citrixadc_appflowcollector.tf_appflowcollector2.name, ]
#   securityinsight = "ENABLED"
#   botinsight      = "ENABLED"
#   videoanalytics  = "ENABLED"
# }
# resource "citrixadc_appflowcollector" "tf_appflowcollector" {
#   name      = "tf_collector"
#   ipaddress = "192.168.2.2"
#   port      = 80
# }
# resource "citrixadc_appflowcollector" "tf_appflowcollector2" {
#   name      = "tf2_collector"
#   ipaddress = "192.168.2.3"
#   port      = 80
# }
```


## Argument Reference

* `name` - (Required) Name for the policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.   The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow policy" or 'my appflow policy').
* `action` - (Optional) Name of the action to be associated with this policy.
* `comment` - (Optional) Any comments about this policy.
* `rule` - (Optional) Expression or other value against which the traffic is evaluated. Must be a Boolean expression.  The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character.  * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `undefaction` - (Optional) Name of the appflow action to be associated with this policy when an undef event occurs.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appflowpolicy. It has the same value as the `name` attribute.


## Import

A appflowpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_appflowpolicy.tf_appflowpolicy test_policy
```