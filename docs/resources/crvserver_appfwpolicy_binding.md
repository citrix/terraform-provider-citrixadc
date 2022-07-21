---
subcategory: "Cache Redirection"
---

# Resource: crvserver_appfwpolicy_binding

The crvserver_appfwpolicy_binding resource is used to create CRvserver Appfwpolicy binding.


## Example usage

```hcl
resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}
resource "citrixadc_appfwprofile" "demo_appfwprofile" {
    name = "demo_appfwprofile"
    bufferoverflowaction = ["none"]
    contenttypeaction = ["none"]
    cookieconsistencyaction = ["none"]
    creditcard = ["none"]
    creditcardaction = ["none"]
    crosssitescriptingaction = ["none"]
    csrftagaction = ["none"]
    denyurlaction = ["none"]
    dynamiclearning = ["none"]
    fieldconsistencyaction = ["none"]
    fieldformataction = ["none"]
    fileuploadtypesaction = ["none"]
    inspectcontenttypes = ["none"]
    jsondosaction = ["none"]
    jsonsqlinjectionaction = ["none"]
    jsonxssaction = ["none"]
    multipleheaderaction = ["none"]
    sqlinjectionaction = ["none"]
    starturlaction = ["none"]
    type = ["HTML"]
    xmlattachmentaction = ["none"]
    xmldosaction = ["none"]
    xmlformataction = ["none"]
    xmlsoapfaultaction = ["none"]
    xmlsqlinjectionaction = ["none"]
    xmlvalidationaction = ["none"]
    xmlwsiaction = ["none"]
    xmlxssaction = ["none"]
}

resource "citrixadc_appfwpolicy" "demo_appfwpolicy1" {
    name = "demo_appfwpolicy1"
    profilename = citrixadc_appfwprofile.demo_appfwprofile.name
    rule = "true"
}

resource "citrixadc_crvserver_appfwpolicy_binding" "crvserver_appfwpolicy_binding" {
    name = citrixadc_crvserver.crvserver.name
    policyname = citrixadc_appfwpolicy.demo_appfwpolicy1.name
    priority = 20
}
```


## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke flag.
* `labelname` - (Optional) Name of the label invoked.
* `labeltype` - (Optional) The invocation type.
* `policyname` - (Optional) Policies bound to this vserver.
* `priority` - (Optional) The priority for the policy.
* `targetvserver` - (Optional) Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_appfwpolicy_binding. It has the same value as the `name` attribute.


## Import

A crvserver_appfwpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_crvserver_appfwpolicy_binding.crvserver_appfwpolicy_binding my_vserver,demo_appfwpolicy1
```
