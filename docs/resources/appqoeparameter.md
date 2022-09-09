---
subcategory: "Appqoe"
---

# Resource: appqoeparameter

The appqoeparameter resource is used to update appqoeparameter.


## Example usage

```hcl
resource "citrixadc_appqoeparameter" "tf_appqoeparameter" {
  sessionlife         = 300
  avgwaitingclient    = 400
  maxaltrespbandwidth = 50
  dosattackthresh     = 100
}
```


## Argument Reference

* `avgwaitingclient` - (Optional) average number of client connections, that can sit in service waiting queue
* `dosattackthresh` - (Optional) average number of client connection that can queue up on vserver level without triggering DoS mitigation module
* `maxaltrespbandwidth` - (Optional) maximum bandwidth which will determine whether to send alternate content response
* `sessionlife` - (Optional) Time, in seconds, between the first time and the next time the AppQoE alternative content window is displayed. The alternative content window is displayed only once during a session for the same browser accessing a configured URL, so this parameter determines the length of a session.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appqoeparameter. It is a unique string prefixed with `tf-appqoeparameter-`.