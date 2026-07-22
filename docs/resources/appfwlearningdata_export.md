---
subcategory: "Application Firewall"
---

# Resource: appfwlearningdata_export

The appfwlearningdata_export resource exports the Citrix ADC Application-Firewall learned-data for a specific profile and security check to a file on the appliance. It is an action-only resource: applying it invokes the NITRO `export` action on `appfwlearningdata`. Use it to capture the rules that the App-Firewall learning engine has accumulated (for example, to review, archive, or transfer them) rather than deploying them directly.

This resource is immutable, so changing any input re-runs the export action.


## Example usage

```hcl
resource "citrixadc_appfwlearningdata_export" "export_starturl" {
  profilename   = "my_appfwprofile"
  securitycheck = "startURL"
  target        = "starturl_learneddata.csv"
}
```


## Argument Reference

* `profilename` - (Required) Name of the App-Firewall profile whose learned data is exported. Changing this attribute re-triggers the export.
* `securitycheck` - (Required) Name of the security check whose learned data is exported. Changing this attribute re-triggers the export. Possible values: `startURL`, `cookieConsistency`, `fieldConsistency`, `crossSiteScripting`, `SQLInjection`, `fieldFormat`, `CSRFtag`, `XMLDoSCheck`, `XMLWSICheck`, `XMLAttachmentCheck`, `TotalXMLRequests`, `creditCardNumber`, `ContentType`.
* `target` - (Optional) Target filename for the data to be exported. Changing this attribute re-triggers the export.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwlearningdata_export resource. It is set to `appfwlearningdata_export`.
