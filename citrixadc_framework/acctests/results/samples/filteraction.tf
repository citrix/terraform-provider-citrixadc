# Sample config for the citrixadc_filteraction resource.
#
# What it does:
#   Creates a Filter Action named "tf_filteraction" with the qualifier "drop".
#   The DROP qualifier silently deletes matching HTTP requests without sending
#   any response to the client. `drop` (and `reset`) are SYSTEM built-in
#   qualifiers, so they work even when the (deprecated) `filter` feature is not
#   licensed/enabled.
#
# `qual` (the action, Required) is one of the NITRO filteraction qualifiers:
#   ADD       - adds an HTTP header            (value = "<header_name>:<header_value>")
#   RESET     - terminates the connection                                (SYSTEM built-in)
#   FORWARD   - redirects the request to a service   (needs `servicename` OR `page`)
#   DROP      - silently drops the request                               (SYSTEM built-in) <-- used here
#   CORRUPT   - corrupts a designated HTTP header     (value = "<header_name>")
#   ERRORCODE - returns an HTTP error code          (needs `respcode`, optional `page`)
#
# IMPORTANT - ADC build requirements:
#   * `filteraction` belongs to the DEPRECATED `filter` feature and only exists on
#     older ADC builds (e.g. NS13.1). It does NOT exist on the 14.1 pool.
#   * On the NS13.1 build this sample was validated against (build 61.26,
#     10.101.132.122) the filter feature is only PARTIALLY present:
#       - Only the SYSTEM qualifiers `drop`/`reset` are accepted; the
#         header/redirect/errorcode qualifiers (add/corrupt/forward/errorcode)
#         are rejected with NITRO errorcode 1097 ("Invalid argument value")
#         because the filter feature cannot be enabled on this build.
#       - The NITRO `set`/`rm` (update/delete) commands for filteraction have
#         been removed (errorcode 1088, "No such command [filteraction]"), so an
#         in-place update or `terraform destroy` will not succeed on this build.
#     To exercise the full CRUD lifecycle (e.g. qual = "corrupt") use an older
#     ADC build where the `filter` feature is fully present and enabled.

terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}

# Credentials/endpoint are taken from the environment:
#   NS_URL, NS_LOGIN, NS_PASSWORD
provider "citrixadc" {
  insecure_skip_verify = true
}

resource "citrixadc_filteraction" "tf_filteraction" {
  name = "tf_filteraction"
  qual = "drop"
}
