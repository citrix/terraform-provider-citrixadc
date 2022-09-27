---
subcategory: "Autoscale"
---

# Resource: autoscaleprofile

The autoscaleprofile resource is used to create autoscaleprofile.


## Example usage

```hcl
resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
  name         = "my_autoscaleprofile"
  type         = "CLOUDSTACK"
  apikey       = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
  url          = "www.service.example.com"
  sharedsecret = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
}
```


## Argument Reference

* `apikey` - (Required) api key for authentication with service
* `name` - (Required) AutoScale profile name.
* `sharedsecret` - (Required) shared secret for authentication with service
* `type` - (Required) The type of profile.
* `url` - (Required) URL providing the service


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the autoscaleprofile. It has the same value as the `name` attribute.


## Import

A autoscaleprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_autoscaleprofile.tf_autoscaleprofile my_autoscaleprofile
```
