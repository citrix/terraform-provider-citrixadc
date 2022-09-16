---
subcategory: "Stream"
---

# Resource: streamselector

The streamselector resource is used to create streamselector.


## Example usage

```hcl
resource "citrixadc_streamselector" "tf_streamselector" {
  name = "my_streamselector"
  rule = ["HTTP.REQ.URL", "CLIENT.IP.SRC"]
}
```


## Argument Reference

* `name` - (Required) Name for the selector. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. If the name includes one or more spaces, and you are using the Citrix ADC CLI, enclose the name in double or single quotation marks (for example, "my selector" or 'my selector').
* `rule` - (Required) Set of up to five expressions. Maximum length: 7499 characters. Each expression must identify a specific request characteristic, such as the client's IP address (with CLIENT.IP.SRC) or requested server resource (with HTTP.REQ.URL). Note: If two or more selectors contain the same expressions in different order, a separate set of records is created for each selector.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the streamselector. It has the same value as the `name` attribute.


## Import

A streamselector can be imported using its name, e.g.

```shell
terraform import citrixadc_streamselector.tf_streamselector my_streamselector
```
