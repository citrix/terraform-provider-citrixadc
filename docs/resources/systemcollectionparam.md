---
subcategory: "System"
---

# Resource: systemcollectionparam

The systemcollectionparam resource is used to update systemcollectionparam.


## Example usage

```hcl
resource "citrixadc_systemcollectionparam" "tf_systemcollectionparam" {
  loglevel      = "WARNING"
}
```


## Argument Reference

* `communityname` - (Optional) SNMPv1 community name for authentication.
* `loglevel` - (Optional) specify the log level. Possible values CRITICAL,WARNING,INFO,DEBUG1,DEBUG2. Minimum length =  1
* `datapath` - (Optional) specify the data path to the database. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemcollectionparam. It is a unique string prefixed with  `tf-systemcollectionparam-` attribute.