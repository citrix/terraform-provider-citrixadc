---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_cookieconsistency_binding

The `appfwprofile_cookieconsistency_binding` resource is used to add bindings between appfwprofile and cookieconsistency relaxation rules.

## Example usage

``` hcl
resource citrixadc_appfwprofile_cookieconsistency_binding demo_binding {
    name              = citrixadc_appfwprofile.demo_appfw.name
    cookieconsistency = "^logon_[0-9A-Za-z]{2,15}$"
}

resource citrixadc_appfwprofile demo_appfw {
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
```

## Argument Reference

* `name` - Name of the profile to which to bind an exemption or rule.
* `cookieconsistency` - The name of the cookie to be checked.
* `isregex` - (Optional) Is the cookie name a regular expression?. Possible values: [ REGEX, NOTREGEX ]
* `state` - (Optional) Enabled. Possible values: [ ENABLED, DISABLED ]
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?. Possible values: [ AUTODEPLOYED, NOTAUTODEPLOYED ]
* `alertonly` - (Optional) Send SNMP alert?. Possible values: [ on, off ]

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwprofile_cookieconsistency_binding`. It has the same value as the `name` attribute.
