---
subcategory: "NS"
---

# Resource: nsspparams

The nsspparams resource is used to create Surge Protection parameter resource.


## Example usage

```hcl
resource "citrixadc_nsspparams" "tf_nsspparams" {
  basethreshold = 200
  throttle      = "Aggressive"
}
```


## Argument Reference

* `basethreshold` - (Optional) Maximum number of server connections that can be opened before surge protection is activated. Minimum value =  0 Maximum value =  32767
* `throttle` - (Optional) Rate at which the system opens connections to the server. Possible values: [ Aggressive, Normal, Relaxed ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsspparams. It is a unique string prefixed with "tf-nsspparams-"

