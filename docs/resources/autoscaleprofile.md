---
subcategory: "autoscale"
---

# Resource: autoscaleprofile

The `autoscaleprofile` resource is used to create autoscaleprofile.


## Example usage

```hcl
resource "citrixadc_autoscaleprofile" "profile1" {
    name = "profile1"
    type = "CLOUDSTACK"
    apikey = "abc123"
    url = "https://1.1.1.1"
    sharedsecret = "abc123"
}
```


## Argument Reference

* `apikey` - (Optional) api key for authentication with service
* `name` - (Optional) AutoScale profile name.
* `sharedsecret` - (Optional) shared secret for authentication with service
* `type` - (Optional) The type of profile.
* `url` - (Optional) URL providing the service


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `autoscaleprofile`. It has the same value as the `name` attribute.


## Import

A `autoscaleprofile` can be imported using its name, e.g.

```shell
terraform import citrixadc_csaction.tf_csaction tf_csaction
```
