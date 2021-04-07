---
subcategory: "Load Balancing"
---

# Resource: lbvserver_appfwpolicy_binding

The lbvserver_appfwpolicy_binding resource is used to add AppFw policies to lbvserver.

## Example usage

``` hcl
resource citrixadc_lbvserver_appfwpolicy_binding demo_binding {
    name = citrixadc_lbvserver.demo_lb.name
    priority = 100
    bindpoint = "REQUEST"
    policyname  = citrixadc_appfwpolicy.demo_appfwpolicy.name
    labelname = citrixadc_lbvserver.demo_lb.name
    gotopriorityexpression = "END"
    invoke = true
    labeltype = "reqvserver"
}

resource citrixadc_lbvserver demo_lb {
name        = "demo_lb"
ipv46       = "1.1.1.1"
port        = "80"
servicetype = "HTTP"
}

resource citrixadc_appfwprofile demo_appfwprofile {
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

resource citrixadc_appfwpolicy demo_appfwpolicy {
    name = "demo_appfwpolicy"
    profilename = citrixadc_appfwprofile.demo_appfwprofile.name
    rule = "true"
}
```

## Argument Reference

* `name` - Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). .
* `policyname` - Name of the policy bound to the LB vserver.
* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE ]
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_appfwpolicy_binding. It has the same value as the `name` attribute.
