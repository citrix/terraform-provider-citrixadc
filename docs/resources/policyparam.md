---
subcategory: "Policy"
---

# Resource: policyparam

The policyparam resource is used to configure the given policy parameter.


## Example usage

```hcl
resource "citrixadc_policyparam" "tf_policyparam" {
	timeout = 5
}
```


## Argument Reference

* `timeout` - (Optional) 
* `maxeventsize` - (Optional) Maximum event size in kilobytes that the policy engine will process. When event data exceeds this limit, the action specified by maxEventSizeExceedAction is taken. This parameter helps prevent resource exhaustion from processing extremely large events.
* `maxeventsizeexceedaction` - (Optional) Action to take when event data exceeds maxEventSize:


## Import

A policyparam can be imported using its id, e.g.

```shell
terraform import citrixadc_policyparam.tf_policyparam policyparam-config
```
