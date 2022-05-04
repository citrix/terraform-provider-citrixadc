---
subcategory: "Application Firewall"
---

# Resource: appfwlearningsettings

The appfwlearningsettings resource is used to update appfw learning settings.


## Example usage

```hcl
resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name                     = "tf_appfwprofile"
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
resource "citrixadc_appfwlearningsettings" "tf_learningsetting" {
  profilename                        = citrixadc_appfwprofile.tf_appfwprofile.name
  starturlminthreshold               = 9
  starturlpercentthreshold           = 10
  cookieconsistencyminthreshold      = 2
  cookieconsistencypercentthreshold  = 1
  csrftagminthreshold                = 2
  csrftagpercentthreshold            = 10
  fieldconsistencyminthreshold       = 20
  fieldconsistencypercentthreshold   = 8
  crosssitescriptingminthreshold     = 10
  crosssitescriptingpercentthreshold = 1
  sqlinjectionminthreshold           = 10
  sqlinjectionpercentthreshold       = 1
  fieldformatminthreshold            = 10
  fieldformatpercentthreshold        = 1
  creditcardnumberminthreshold       = 1
  creditcardnumberpercentthreshold   = 0
  contenttypeminthreshold            = 1
  contenttypepercentthreshold        = 0
}
```


## Argument Reference

* `profilename` - (Required) Name of the profile.
* `contenttypeautodeploygraceperiod` - (Optional) The number of minutes after the threshold hit alert the learned rule will be deployed
* `contenttypeminthreshold` - (Optional) Minimum threshold to learn Content Type information.
* `contenttypepercentthreshold` - (Optional) Minimum threshold in percent to learn Content Type information.
* `cookieconsistencyautodeploygraceperiod` - (Optional) The number of minutes after the threshold hit alert the learned rule will be deployed
* `cookieconsistencyminthreshold` - (Optional) Minimum number of application firewall sessions that the learning engine must observe to learn cookies.
* `cookieconsistencypercentthreshold` - (Optional) Minimum percentage of application firewall sessions that must contain a particular cookie pattern for the learning engine to learn that cookie.
* `creditcardnumberminthreshold` - (Optional) Minimum threshold to learn Credit Card information.
* `creditcardnumberpercentthreshold` - (Optional) Minimum threshold in percent to learn Credit Card information.
* `crosssitescriptingautodeploygraceperiod` - (Optional) The number of minutes after the threshold hit alert the learned rule will be deployed
* `crosssitescriptingminthreshold` - (Optional) Minimum number of application firewall sessions that the learning engine must observe to learn HTML cross-site scripting patterns.
* `crosssitescriptingpercentthreshold` - (Optional) Minimum percentage of application firewall sessions that must contain a particular cross-site scripting pattern for the learning engine to learn that cross-site scripting pattern.
* `csrftagautodeploygraceperiod` - (Optional) The number of minutes after the threshold hit alert the learned rule will be deployed
* `csrftagminthreshold` - (Optional) Minimum number of application firewall sessions that the learning engine must observe to learn cross-site request forgery (CSRF) tags.
* `csrftagpercentthreshold` - (Optional) Minimum percentage of application firewall sessions that must contain a particular CSRF tag for the learning engine to learn that CSRF tag.
* `fieldconsistencyautodeploygraceperiod` - (Optional) The number of minutes after the threshold hit alert the learned rule will be deployed
* `fieldconsistencyminthreshold` - (Optional) Minimum number of application firewall sessions that the learning engine must observe to learn field consistency information.
* `fieldconsistencypercentthreshold` - (Optional) Minimum percentage of application firewall sessions that must contain a particular field consistency pattern for the learning engine to learn that field consistency pattern.
* `fieldformatautodeploygraceperiod` - (Optional) The number of minutes after the threshold hit alert the learned rule will be deployed
* `fieldformatminthreshold` - (Optional) Minimum number of application firewall sessions that the learning engine must observe to learn field formats.
* `fieldformatpercentthreshold` - (Optional) Minimum percentage of application firewall sessions that must contain a particular web form field pattern for the learning engine to recommend a field format for that form field.
* `sqlinjectionautodeploygraceperiod` - (Optional) The number of minutes after the threshold hit alert the learned rule will be deployed
* `sqlinjectionminthreshold` - (Optional) Minimum number of application firewall sessions that the learning engine must observe to learn HTML SQL injection patterns.
* `sqlinjectionpercentthreshold` - (Optional) Minimum percentage of application firewall sessions that must contain a particular HTML SQL injection pattern for the learning engine to learn that HTML SQL injection pattern.
* `starturlautodeploygraceperiod` - (Optional) The number of minutes after the threshold hit alert the learned rule will be deployed
* `starturlminthreshold` - (Optional) Minimum number of application firewall sessions that the learning engine must observe to learn start URLs.
* `starturlpercentthreshold` - (Optional) Minimum percentage of application firewall sessions that must contain a particular start URL pattern for the learning engine to learn that start URL.
* `xmlattachmentminthreshold` - (Optional) Minimum number of application firewall sessions that the learning engine must observe to learn XML attachment patterns.
* `xmlattachmentpercentthreshold` - (Optional) Minimum percentage of application firewall sessions that must contain a particular XML attachment pattern for the learning engine to learn that XML attachment pattern.
* `xmlwsiminthreshold` - (Optional) Minimum number of application firewall sessions that the learning engine must observe to learn web services interoperability (WSI) information.
* `xmlwsipercentthreshold` - (Optional) Minimum percentage of application firewall sessions that must contain a particular pattern for the learning engine to learn a web services interoperability (WSI) pattern.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwlearningsettings. It has the same value as the `profilename` attribute.


## Import

A appfwlearningsettings can be imported using its profilename, e.g.

```shell
terraform import citrixadc_appfwlearningsettings.tf_learningsetting tf_appfwprofile
```
