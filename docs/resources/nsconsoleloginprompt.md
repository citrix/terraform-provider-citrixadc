---
subcategory: "NS"
---

# Resource: nsconsoleloginprompt

The nsconsoleloginprompt resource is used to create console prompt resource.


## Example usage

```hcl
resource "citrixadc_nsconsoleloginprompt" "tf_nsconsoleloginprompt" {
  promptstring = "tf_nsconsoleloginprompt"
}
```


## Argument Reference

* `promptstring` - (Required) Console login prompt string.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsconsoleloginprompt. It is a unique string prefixed with "tf-nsconsoleloginprompt-"

