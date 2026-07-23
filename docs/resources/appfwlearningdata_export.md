---
subcategory: "Application Firewall"
---

# Resource: appfwlearningdata_export

This resource is used to export Application Firewall learned data for a profile and security check.


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
