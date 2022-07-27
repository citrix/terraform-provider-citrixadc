---
subcategory: "SNMP"
---

# Resource: snmpengineid

The snmpengineid resource is used to create snmpengineid.


## Example usage

```hcl
resource "citrixadc_snmpengineid" "citrixadc_snmpengineid" {
    engineid = "1234567890ABCDEF"
}
```


## Argument Reference

* `engineid` - (Required) A hexadecimal value of at least 10 characters, uniquely identifying the engineid
* `ownernode` - (Optional) ID of the cluster node for which you are setting the engineid


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the snmpengineid. It is a unique string prefixed with "tf-snmpengineid-".


