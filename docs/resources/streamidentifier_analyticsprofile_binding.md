---
subcategory: "Stream"
---

# Resource: streamidentifier_analyticsprofile_binding

Binds an analytics profile to a stream identifier so that traffic tracked by the stream identifier is reported through the specified analytics profile on the Citrix ADC. Create this binding when you want the records collected for a given stream identifier to be exported to a particular analytics profile.

This binding is immutable: it can only be created (bound) or deleted (unbound). Changing either attribute forces Terraform to replace the resource.


## Example usage

```hcl

resource "citrixadc_streamidentifier" "tf_streamidentifier" {
  name     = "tf_streamidentifier"
  selectorname = citrixadc_streamselector.tf_streamselector.name
}

resource "citrixadc_streamselector" "tf_streamselector" {
  name = "tf_streamselector"
  rule = ["CLIENT.IP.SRC"]
}

resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
  name = "tf_analyticsprofile"
  type = "webinsight"
}

resource "citrixadc_streamidentifier_analyticsprofile_binding" "tf_binding" {
  name             = citrixadc_streamidentifier.tf_streamidentifier.name
  analyticsprofile = citrixadc_analyticsprofile.tf_analyticsprofile.name
}

```


## Argument Reference

* `name` - (Required) The name of the stream identifier to which the analytics profile is bound. Maximum length = 127
* `analyticsprofile` - (Required) Name of the analytics profile to bind to the stream identifier. Maximum length = 127


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the streamidentifier_analyticsprofile_binding. It is a system-generated identifier built from the unique attributes as a comma-separated list of `key:value` pairs (the values are URL-encoded), in the form `analyticsprofile:<analyticsprofile>,name:<name>`.


## Import

A streamidentifier_analyticsprofile_binding can be imported using its id, which is the comma-separated list of `key:value` pairs described above, e.g.

```shell
terraform import citrixadc_streamidentifier_analyticsprofile_binding.tf_binding "analyticsprofile:tf_analyticsprofile,name:tf_streamidentifier"
```
