# rebooter.tf
#
# ############################################################################
# # !!! DANGER: APPLYING THIS CONFIG REBOOTS THE CITRIX ADC APPLIANCE !!!    #
# #                                                                          #
# # citrixadc_rebooter is a one-shot side-effect ACTION resource. On         #
# # `terraform apply` its Create issues the NITRO `reboot` action, which     #
# # RESTARTS the appliance (warm = ADC software restart; cold = full OS      #
# # reboot). This WILL take the box down / interrupt traffic.                #
# #                                                                          #
# # This file is intended ONLY for `terraform validate` and `terraform plan`.#
# # DO NOT run `terraform apply` against a box you care about.               #
# ############################################################################
#
# What this does (per resource_rebooter.go / resource_citrixadc_rebooter.go):
#   Create issues NITRO `reboot` (warm or cold). When wait_until_reachable is
#   true it polls the appliance's nslicense endpoint until the ADC comes back
#   up. There is no GET endpoint and no inverse API, so Read/Update/Delete are
#   no-ops and every input attribute is ForceNew (RequiresReplace). A synthetic
#   ID (tf-rebooter-<nanos>) keeps the action-only resource addressable.
#
# Attributes:
#   warm                    (Required, bool,   ForceNew) - true  => restart the
#                             ADC software only, without rebooting the underlying
#                             OS. false => full cold reboot of the appliance.
#   timestamp               (Required, string, ForceNew) - trigger value. Change
#                             it to force a new reboot on a subsequent apply.
#   wait_until_reachable    (Required, bool,   ForceNew) - when true, block after
#                             the reboot until the ADC is reachable again.
#   reachable_timeout       (Optional, string, ForceNew, default "10m") - max
#                             duration to wait for the ADC to become reachable.
#   reachable_poll_delay    (Optional, string, ForceNew, default "60s") - delay
#                             before the first reachability poll.
#   reachable_poll_interval (Optional, string, ForceNew, default "60s") - interval
#                             between reachability polls.
#   reachable_poll_timeout  (Optional, string, ForceNew, default "20s") - per-poll
#                             HTTP timeout when checking reachability.
#
# How to VALIDATE / PLAN (safe) - do NOT apply:
#   export NS_URL="http://10.101.132.155/"
#   export NS_LOGIN="nsroot"
#   export NS_PASSWORD='CADS123$%^'
#   terraform validate
#   terraform plan          # shows the resource WOULD be created
#   # terraform apply       # <-- INTENTIONALLY OMITTED: this reboots the ADC

terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}

provider "citrixadc" {
  # endpoint (NS_URL), username (NS_LOGIN) and password (NS_PASSWORD) are taken
  # from the environment. insecure_skip_verify allows self-signed ADC certs.
  insecure_skip_verify = true
}

resource "citrixadc_rebooter" "tf_rebooter" {
  # WARNING: applying this reboots the appliance. Validate/plan only.
  warm                 = true
  timestamp            = "2026-07-31T00:00:00Z"
  wait_until_reachable = true

  # Optional reachability-polling knobs (defaults shown for clarity):
  reachable_timeout       = "10m"
  reachable_poll_delay    = "60s"
  reachable_poll_interval = "60s"
  reachable_poll_timeout  = "20s"
}
