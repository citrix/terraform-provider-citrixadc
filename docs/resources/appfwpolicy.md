---
subcategory: "Application Firewall"
---

# Resource: appfwpolicy

The `appfwpolicy` resource is used to create Applicatin Firewall Policy resource.

## Example usage

``` hcl
resource citrixadc_appfwpolicy demo_appfwpolicy1 {
    name = "demo_appfwpolicy1"
    profilename = citrixadc_appfwprofile.demo_appfwprofile.name
    rule = "true"
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
```

## Argument Reference

* `name` - Name for the policy. Must begin with a letter, number, or the underscore character \(_\), and must contain only letters, numbers, and the hyphen \(-\), period \(.\) pound \(\#\), space \( \), at (@), equals \(=\), colon \(:\), and underscore characters. Can be changed after the policy is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my policy" or 'my policy'\).
* `rule` - (Optional) Name of the Citrix ADC named rule, or a Citrix ADC expression, that the policy uses to determine whether to filter the connection through the application firewall with the designated profile.
* `profilename` - (Optional) Name of the application firewall profile to use if the policy matches.
* `comment` - (Optional) Any comments to preserve information about the policy for later reference.
* `logaction` - (Optional) Where to log information for connections that match this policy.
* `newname` - (Optional) New name for the policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwpolicy`. It has the same value as the `name` attribute.
