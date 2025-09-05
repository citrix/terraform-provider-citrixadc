---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_starturl_binding

The `appfwprofile_starturl_binding` resource is used to add bindings between Application Firewall Profile and StartURL relaxation rule.


## Example usage

```hcl

resource citrixadc_appfwprofile_starturl_binding appfwprofile_starturl1 {
    name = citrixadc_appfwprofile.demo_appfw.name
    starturl = "^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pif|pdf|css|csv)$"

    # Below attributes are to be provided to maintain idempotency
    alertonly      = "OFF"
    isautodeployed = "NOTAUTODEPLOYED"
    state          = "ENABLED"
}

resource citrixadc_appfwprofile_starturl_binding appfwprofile_starturl2 {
    name = citrixadc_appfwprofile.demo_appfw.name
    starturl = "^[^?]+[.](cgi|aspx?|jsp|php|pl)([?].*)?$"

    # Below attributes are to be provided to maintain idempotency
    alertonly      = "OFF"
    isautodeployed = "NOTAUTODEPLOYED"
    state          = "ENABLED"
}
resource citrixadc_appfwprofile demo_appfw {
    name = "demo_appfwprofile1"
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
* `starturl` - A regular expression that designates a URL on the Start URL list.
* `state` - (Optional) Enabled. Possible values: [ ENABLED, DISABLED ]
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?. Possible values: [ AUTODEPLOYED, NOTAUTODEPLOYED ]
* `alertonly` - (Optional) Send SNMP alert?. Possible values: [ on, off ]
* `ruletype` - (Optional) Specifies rule type of binding.
* `resourceid` - (Optional) A "id" that identifies the rule.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwprofile_starturl_binding`. It has the value of the string `name,starturl`.
