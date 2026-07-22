---
subcategory: "Stream"
---

# Resource: streamsession_clear

Clears the accumulated stream-session records collected for a stream identifier on the Citrix ADC. Stream analytics builds up per-request session data against a named stream identifier; applying this resource performs an imperative `clear` action that discards that accumulated data, letting administrators reset counters and start a fresh collection window without reconfiguring the stream identifier itself.

This is an action resource: applying it performs the clear; it does not manage a persistent object, so re-applying re-runs the action. The `name` argument is immutable, so changing it re-runs the clear action against the new stream identifier.


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
