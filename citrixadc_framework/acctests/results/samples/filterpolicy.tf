# filterpolicy.tf
#
# What this does:
#   citrixadc_filterpolicy is a real config-CRUD resource that maps 1:1 to the
#   NITRO `filterpolicy` object (part of the classic Content Filtering / "filter"
#   feature). Create = add filterpolicy, Read = get by name, Update =
#   update filterpolicy, Delete = rm filterpolicy. The resource id == the policy
#   name. This is a backward-compatible migration of the legacy SDK v2
#   citrixadc_filterpolicy resource (same type name, schema, and id scheme).
#
# Attributes (per resource_filterpolicy.go / NITRO filterpolicy spec):
#   name      (Required)          - Name of the filter policy. Becomes the id.
#   rule      (Optional+Computed) - Citrix ADC *classic* expression selecting the
#                                     connections that match this policy. A rule
#                                     that evaluates REQ.* pairs with reqaction;
#                                     a rule that evaluates RES.* pairs with
#                                     resaction (the two are mutually exclusive on
#                                     a single policy).
#   reqaction (Optional+Computed) - Action for matching *requests*. Per the NITRO
#                                     spec this accepts EITHER the name of a
#                                     citrixadc_filteraction you created OR a
#                                     built-in action: "RESET", "DROP", "NOOP",
#                                     "FORWARD:<vserver>", "ADD:<hdr>",
#                                     "CORRUPT:<hdr>", "ERRORCODE:<code>".
#   resaction (Optional+Computed) - Action for matching *responses* (built-in or
#                                     a named filteraction). Not set here.
#
#   Because a built-in request action ("DROP") is a valid reqaction, this policy
#   does NOT require a separate citrixadc_filteraction, so the primary config
#   below is fully self-contained. If you instead want the policy to reference a
#   CUSTOM filter action, use the commented variant at the bottom of this file
#   (distinct names tf_sample_fp_action / tf_sample_filterpolicy + a depends_on).
#
# 13.1 REQUIREMENT (important):
#   The `filterpolicy` object only exists when the classic Content Filtering
#   feature ("CF", CLI: `enable ns feature CF`) is supported AND enabled on the
#   appliance. Content Filtering was removed from the modern firmware line: it is
#   absent on the 14.1 pool, and even on NS13.1 build 61.26 the feature flag is
#   deprecated (enabling it returns NITRO warning 558 "Content Filtering is no
#   longer supported"). On an appliance where CF cannot be enabled, `add
#   filterpolicy` fails with NITRO errorcode 1232 "Invalid object name
#   [filterpolicy]". Run this sample only against an ADC where `enable ns feature
#   CF` actually turns the feature on and `GET /nitro/v1/config/filterpolicy`
#   succeeds.
#
# How to run:
#   export NS_URL="http://<adc-with-CF-enabled>/"
#   export NS_LOGIN="nsroot"
#   export NS_PASSWORD='CADS123$%^'
#   terraform validate
#   terraform plan
#   terraform apply
#   (endpoint/username/password are read from the environment by the provider.)

terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}

provider "citrixadc" {
  # endpoint (NS_URL), username (NS_LOGIN) and password (NS_PASSWORD) come from
  # the environment. insecure_skip_verify allows self-signed ADC certs.
  insecure_skip_verify = true
}

# Self-contained filter policy using a BUILT-IN request action ("DROP"): any
# request whose URL contains the blocked host is silently dropped. No separate
# filter action object is needed because "DROP" is a built-in reqaction.
resource "citrixadc_filterpolicy" "tf_sample_filterpolicy" {
  name      = "tf_sample_filterpolicy"
  rule      = "REQ.HTTP.URL CONTAINS http://blocked.example.com"
  reqaction = "DROP"
}

# ----------------------------------------------------------------------------
# VARIANT (custom filter action): uncomment to have the policy reference a
# CUSTOM citrixadc_filteraction instead of a built-in. Distinct names avoid
# colliding with the standalone filteraction sample; depends_on guarantees the
# action exists before the policy that references it is created. Requires the
# same CF feature as above.
#
# resource "citrixadc_filteraction" "tf_sample_fp_action" {
#   name = "tf_sample_fp_action"
#   qual = "RESET" # RESET terminates matching request connections
# }
#
# resource "citrixadc_filterpolicy" "tf_sample_filterpolicy_custom" {
#   name       = "tf_sample_filterpolicy_custom"
#   rule       = "REQ.HTTP.URL CONTAINS http://blocked.example.com"
#   reqaction  = citrixadc_filteraction.tf_sample_fp_action.name
#   depends_on = [citrixadc_filteraction.tf_sample_fp_action]
# }
# ----------------------------------------------------------------------------
