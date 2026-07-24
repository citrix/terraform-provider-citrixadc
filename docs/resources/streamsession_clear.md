---
subcategory: "Stream"
---

# Resource: streamsession_clear

This resource is used to clear accumulated stream-session records for a stream identifier.


## Example usage

Clear the accumulated stream-session records for a named stream identifier:

```hcl
resource "citrixadc_streamsession_clear" "clear_sessions" {
  name = "streamident1"
}
```


## Argument Reference

* `name` - (Required) Name of the stream identifier whose accumulated stream-session records should be cleared. Maximum length = 127. Changing this value forces the resource to be recreated, which re-runs the clear action against the new stream identifier.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the streamsession_clear resource. It is set to `streamsession_clear`.

This resource is action-only and represents no importable server-side object, so importing it is not meaningful and there is no Import section.
