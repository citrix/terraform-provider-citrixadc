---
subcategory: "Content Switching"
---

# Resource: csparameter

The csparameter resource is used to modify the status of the state update parameter.


## Example usage

```hcl
resource "citrixadc_csparameter" "tf_csparameter" {
	stateupdate = "ENABLED"
}
```


## Argument Reference

* `stateupdate` - (Optional) Specifies whether the virtual server checks the attached load balancing server for state information. Possible values: [ ENABLED, DISABLED ]
