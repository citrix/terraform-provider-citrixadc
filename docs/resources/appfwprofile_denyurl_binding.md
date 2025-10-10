---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_denyurl_binding

The `appfwprofile_denyurl_binding` resource is used to add bindings between Application Firewall Profile and DenyURL relaxation rule.

## Example usage

``` hcl
resource citrixadc_appfwprofile_denyurl_binding appfwprofile_denyurl1 {
    name = citrixadc_appfwprofile.demo_appfw.name
    denyurl = "debug[.][^/?]*(|[?].*)$"

    # Below attributes are to be provided to maintain idempotency
    alertonly      = "OFF"
    isautodeployed = "NOTAUTODEPLOYED"
    state          = "ENABLED"
}

resource citrixadc_appfwprofile_denyurl_binding appfwprofile_denyurl2 {
    name = citrixadc_appfwprofile.demo_appfw.name
    denyurl = "^[^?]*/default[.]ida[?]N+"

    # Below attributes are to be provided to maintain idempotency
    alertonly      = "OFF"
    isautodeployed = "NOTAUTODEPLOYED"
    state          = "ENABLED"
}
resource citrixadc_appfwprofile demo_appfw {
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

```

## Argument Reference

* `name` - Name of the profile to which to bind an exemption or rule.
* `denyurl` - A regular expression that designates a URL on the Deny URL list.
* `state` - (Optional) Enabled. Possible values: [ ENABLED, DISABLED ]
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?. Possible values: [ AUTODEPLOYED, NOTAUTODEPLOYED ]
* `alertonly` - (Optional) Send SNMP alert?. Possible values: [ on, off ]
* `resourceid` - (Optional) A unique id that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwprofile_denyurl_binding`. It has the value of the string `name,denyurl`.
