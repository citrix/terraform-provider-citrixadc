---
subcategory: "Application Firewall"
---

# Data Source `citrixadc_appfwlearningsettings`

The appfwlearningsettings data source allows you to retrieve information about Application Firewall learning settings for a specific profile.

## Example usage

```terraform
data "citrixadc_appfwlearningsettings" "tf_learningsetting" {
  profilename = "tf_appfwprofile"
}

output "starturlminthreshold" {
  value = data.citrixadc_appfwlearningsettings.tf_learningsetting.starturlminthreshold
}

output "cookieconsistencyminthreshold" {
  value = data.citrixadc_appfwlearningsettings.tf_learningsetting.cookieconsistencyminthreshold
}

output "csrftagminthreshold" {
  value = data.citrixadc_appfwlearningsettings.tf_learningsetting.csrftagminthreshold
}
```

## Argument Reference

* `profilename` - (Required) Name of the profile.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `contenttypeautodeploygraceperiod` - The number of minutes after the threshold hit alert the learned rule will be deployed
* `contenttypeminthreshold` - Minimum threshold to learn Content Type information.
* `contenttypepercentthreshold` - Minimum threshold in percent to learn Content Type information.
* `cookieconsistencyautodeploygraceperiod` - The number of minutes after the threshold hit alert the learned rule will be deployed
* `cookieconsistencyminthreshold` - Minimum number of application firewall sessions that the learning engine must observe to learn cookies.
* `cookieconsistencypercentthreshold` - Minimum percentage of application firewall sessions that must contain a particular cookie pattern for the learning engine to learn that cookie.
* `creditcardnumberminthreshold` - Minimum threshold to learn Credit Card information.
* `creditcardnumberpercentthreshold` - Minimum threshold in percent to learn Credit Card information.
* `crosssitescriptingautodeploygraceperiod` - The number of minutes after the threshold hit alert the learned rule will be deployed
* `crosssitescriptingminthreshold` - Minimum number of application firewall sessions that the learning engine must observe to learn HTML cross-site scripting patterns.
* `crosssitescriptingpercentthreshold` - Minimum percentage of application firewall sessions that must contain a particular cross-site scripting pattern for the learning engine to learn that cross-site scripting pattern.
* `csrftagautodeploygraceperiod` - The number of minutes after the threshold hit alert the learned rule will be deployed
* `csrftagminthreshold` - Minimum number of application firewall sessions that the learning engine must observe to learn cross-site request forgery (CSRF) tags.
* `csrftagpercentthreshold` - Minimum percentage of application firewall sessions that must contain a particular CSRF tag for the learning engine to learn that CSRF tag.
* `fieldconsistencyautodeploygraceperiod` - The number of minutes after the threshold hit alert the learned rule will be deployed
* `fieldconsistencyminthreshold` - Minimum number of application firewall sessions that the learning engine must observe to learn field consistency information.
* `fieldconsistencypercentthreshold` - Minimum percentage of application firewall sessions that must contain a particular field consistency pattern for the learning engine to learn that field consistency pattern.
* `fieldformatautodeploygraceperiod` - The number of minutes after the threshold hit alert the learned rule will be deployed
* `fieldformatminthreshold` - Minimum number of application firewall sessions that the learning engine must observe to learn field formats.
* `fieldformatpercentthreshold` - Minimum percentage of application firewall sessions that must contain a particular web form field pattern for the learning engine to recommend a field format for that form field.
* `id` - The id of the appfwlearningsettings. It has the same value as the `profilename` attribute.
* `sqlinjectionautodeploygraceperiod` - The number of minutes after the threshold hit alert the learned rule will be deployed
* `sqlinjectionminthreshold` - Minimum number of application firewall sessions that the learning engine must observe to learn HTML SQL injection patterns.
* `sqlinjectionpercentthreshold` - Minimum percentage of application firewall sessions that must contain a particular HTML SQL injection pattern for the learning engine to learn that HTML SQL injection pattern.
* `starturlautodeploygraceperiod` - The number of minutes after the threshold hit alert the learned rule will be deployed
* `starturlminthreshold` - Minimum number of application firewall sessions that the learning engine must observe to learn start URLs.
* `starturlpercentthreshold` - Minimum percentage of application firewall sessions that must contain a particular start URL pattern for the learning engine to learn that start URL.
* `xmlattachmentminthreshold` - Minimum number of application firewall sessions that the learning engine must observe to learn XML attachment patterns.
* `xmlattachmentpercentthreshold` - Minimum percentage of application firewall sessions that must contain a particular XML attachment pattern for the learning engine to learn that XML attachment pattern.
* `xmlwsiminthreshold` - Minimum number of application firewall sessions that the learning engine must observe to learn web services interoperability (WSI) information.
* `xmlwsipercentthreshold` - Minimum percentage of application firewall sessions that must contain a particular pattern for the learning engine to learn a web services interoperability (WSI) pattern.

## Import

A appfwlearningsettings can be imported using its profilename, e.g.

```shell
terraform import citrixadc_appfwlearningsettings.tf_learningsetting tf_appfwprofile
```
