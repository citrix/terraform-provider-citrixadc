# nsconfig_clear.tf
#
# ############################################################################
# ##  !!! DANGER — THIS RESOURCE WIPES THE NETSCALER CONFIGURATION !!!      ##
# ##                                                                        ##
# ##  citrixadc_nsconfig_clear issues `clear ns config` on the target ADC   ##
# ##  (NITRO: POST /nitro/v1/config/nsconfig?action=clear). With            ##
# ##  level = "full" it clears ALL configuration except the NSIP, the       ##
# ##  default route, and interface settings. This is DESTRUCTIVE and        ##
# ##  IRREVERSIBLE. There is NO inverse/undo API.                           ##
# ##                                                                        ##
# ##  ONLY point NS_URL at a DISPOSABLE / throwaway appliance. Never run     ##
# ##  this against a production, shared, or otherwise valuable ADC.         ##
# ############################################################################
#
# What this does:
#   citrixadc_nsconfig_clear is an ACTION-ONLY resource. On create it issues a
#   single NITRO call: POST /nitro/v1/config/nsconfig?action=clear with the
#   configured force/level/rbaconfig attributes. There is no GET/update inverse
#   for the clear action, so Read/Update are no-ops; `timestamp` is a
#   provider-supplied marker used only as the Terraform resource ID. Bump
#   `timestamp` to re-run "clear ns config" (all attributes are ForceNew /
#   RequiresReplace). Delete just drops the resource from state (no ADC call).
#
# Attributes (per resource_nsconfig_clear.go / nitro_rest/ns/nsconfig.html):
#   level     (Required, ForceNew) - types of configuration to clear. One of:
#               "basic"    - clears all config except NSIP, routes, MIP/SNIP,
#                            network settings, cluster/HA, feature/mode, nsroot pw
#               "extended" - as "basic" plus feature and mode settings
#               "full"     - clears all config except NSIP, default route, and
#                            interface settings
#   force     (Optional, ForceNew) - clear without prompting for confirmation.
#   rbaconfig (Optional, ForceNew) - "YES"/"NO". If "NO", RBA configs and TACACS
#               policies bound to system global are preserved. Applies only to
#               the "basic" level. Default on the ADC is "YES".
#   timestamp (Required, ForceNew) - marker string; becomes the Terraform ID.
#
# How to run (ONLY against a disposable box):
#   export NS_URL="http://10.101.132.155/"   # DISPOSABLE appliance ONLY
#   export NS_LOGIN="nsroot"
#   export NS_PASSWORD='CADS123$%^'
#   terraform init
#   terraform apply
#
#   (The provider reads NS_URL / NS_LOGIN / NS_PASSWORD from the environment.)

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

resource "citrixadc_nsconfig_clear" "foo" {
  # !!! DESTRUCTIVE: clears (most of) the running config on the target ADC. !!!
  # Only run this against a disposable appliance.
  force = false
  level = "full"

  # rbaconfig only applies to level = "basic"; shown here for completeness.
  # rbaconfig = "YES"

  # Bump this marker to force a fresh "clear ns config".
  timestamp = "2024-06-01T12:00:00"
}
