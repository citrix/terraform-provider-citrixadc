---
subcategory: "Content Switching"
---

# Resource: csvserver_appfwpolicy_binding

The csvserver_appfwpolicy_binding resource is used to add csvserver_appfwpolicy_binding.

## Example usage

``` hcl
resource citrixadc_csvserver_appfwpolicy_binding demo_csvserver_appfwpolicy_binding {
  name                   = citrixadc_csvserver.demo_cs.name
  priority               = 100
  policyname             = citrixadc_appfwpolicy.demo_appfwpolicy1.name
  gotopriorityexpression = "END"
}
resource "citrixadc_csvserver" "demo_cs" {
  ipv46       = "10.10.10.33"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource citrixadc_appfwprofile demo_appfwprofile {
  name                     = "demo_appfwprofile"
  bufferoverflowaction     = ["none"]
  contenttypeaction        = ["none"]
  cookieconsistencyaction  = ["none"]
  creditcard               = ["none"]
  creditcardaction         = ["none"]
  crosssitescriptingaction = ["none"]
  csrftagaction            = ["none"]
  denyurlaction            = ["none"]
  dynamiclearning          = ["none"]
  fieldconsistencyaction   = ["none"]
  fieldformataction        = ["none"]
  fileuploadtypesaction    = ["none"]
  inspectcontenttypes      = ["none"]
  jsondosaction            = ["none"]
  jsonsqlinjectionaction   = ["none"]
  jsonxssaction            = ["none"]
  multipleheaderaction     = ["none"]
  sqlinjectionaction       = ["none"]
  starturlaction           = ["none"]
  type                     = ["HTML"]
  xmlattachmentaction      = ["none"]
  xmldosaction             = ["none"]
  xmlformataction          = ["none"]
  xmlsoapfaultaction       = ["none"]
  xmlsqlinjectionaction    = ["none"]
  xmlvalidationaction      = ["none"]
  xmlwsiaction             = ["none"]
  xmlxssaction             = ["none"]
}

resource citrixadc_appfwpolicy demo_appfwpolicy1 {
  name        = "demo_appfwpolicy1"
  profilename = citrixadc_appfwprofile.demo_appfwprofile.name
  rule        = "true"
}
```

## Argument Reference

* `name` - Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - Policies bound to this vserver.
* `priority` - (Optional) Priority for the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]
* `invoke` - (Optional) Invoke flag.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `targetlbvserver` - (Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_appfwpolicy_binding. It has the same value as the `name` attribute.
