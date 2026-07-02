---
subcategory: "Stream"
---

# Resource: streamsession

Clears the accumulated stream-session records collected for a stream identifier on the Citrix ADC. Stream analytics builds up per-request session data against a named stream identifier; applying this resource performs an imperative `clear` action that discards that accumulated data, letting administrators reset counters and start a fresh collection window without reconfiguring the stream identifier itself.

This is an action-only resource. It does not manage a persistent server-side object:

* **Apply runs the clear action.** Creating the resource sends a `POST ?action=clear` request to the ADC for the named stream identifier. The accumulated stream-session records are discarded immediately. This is a one-shot side-effect.
* **There is no read-back.** Read is a no-op. Cleared session data is not a queryable managed object (NITRO exposes no get-by-name key for `streamsession`), so the provider cannot re-resolve it or detect drift. State holds only the `name` you supplied plus a synthetic ID.
* **There is no in-place update.** The `name` argument forces resource replacement. Changing it and re-applying re-runs the clear action against the new stream identifier.
* **Delete is state-only.** The `clear` action has no inverse endpoint, so destroying the resource simply removes it from Terraform state; no session data is restored.

Because re-applying re-runs the clear, use this resource for deliberate, one-shot reset workflows rather than for declarative, drift-corrected configuration.


## Example usage

Clear the accumulated stream-session records for a named stream identifier:

```hcl
resource "citrixadc_streamsession" "clear_sessions" {
  name = "streamident1"
}
```


## Argument Reference

* `name` - (Required) Name of the stream identifier whose accumulated stream-session records should be cleared. Maximum length = 127. Changing this value forces the resource to be recreated, which re-runs the clear action against the new stream identifier.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier with the constant value `"streamsession-config"`. The ADC does not assign an ID to the clear action; this value is purely a Terraform state handle and is not a NITRO lookup key.

This resource is action-only and represents no importable server-side object, so importing it is not meaningful and there is no Import section.
