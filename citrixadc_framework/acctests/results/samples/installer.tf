# installer.tf
#
# ############################################################################
# ##                          !!!  DANGER  !!!                              ##
# ##                                                                        ##
# ##  APPLYING THIS CONFIG PERFORMS A REAL FIRMWARE INSTALL / UPGRADE OF    ##
# ##  THE TARGET NetScaler (ADC) APPLIANCE.                                 ##
# ##                                                                        ##
# ##  `terraform apply` on citrixadc_installer will:                        ##
# ##    - download the build image from `url`,                              ##
# ##    - run the NITRO `install` action (equivalent to `install ns`),      ##
# ##    - REBOOT the appliance, and                                         ##
# ##    - upgrade/downgrade the running firmware to that build.             ##
# ##                                                                        ##
# ##  A wrong/incompatible `url`, or interrupting the install, can BRICK    ##
# ##  the box. Only ever `apply` this INTENTIONALLY, against an appliance   ##
# ##  you positively intend to upgrade, with a real, verified build .tgz    ##
# ##  URL that the appliance can reach.                                     ##
# ##                                                                        ##
# ##  This sample exists for `terraform validate` + `terraform plan` ONLY.  ##
# ##  The `url` below is a NON-FUNCTIONAL PLACEHOLDER - do NOT apply it.    ##
# ############################################################################
#
# What this resource does:
#   citrixadc_installer is a one-shot, side-effect ACTION resource. On create it
#   POSTs the NITRO `install` action to upgrade the appliance to the build image
#   at `url`. NITRO has no GET endpoint reporting install state and there is no
#   inverse API, so Read/Update/Delete are no-ops and EVERY attribute is
#   ForceNew (RequiresReplace). The install typically reboots the box; the
#   provider tolerates the resulting TCP reset / EOF from the NITRO call.
#
#   When wait_until_reachable = true, create blocks after issuing the install and
#   polls the appliance's nslicense endpoint until it responds again (i.e. the
#   box has rebooted and come back), bounded by reachable_timeout.
#
# Attributes (per resource_installer.go / resource_citrixadc_installer.go):
#   url                     (Optional, Computed, ForceNew) - URL of the build .tgz
#                             image to install. Must be reachable from the ADC.
#   enhancedupgrade         (Optional, Computed, ForceNew) - upgrade from/to
#                             enhancement mode.
#   l                       (Optional, Computed, ForceNew) - enable callhome.
#   resizeswapvar           (Optional, Computed, ForceNew) - change swap size on
#                             64-bit nCore/MCNS/VMPE (NON-VPX) systems only.
#   y                       (Optional, Computed, ForceNew) - do not prompt for
#                             yes/no before rebooting.
#   wait_until_reachable    (REQUIRED, ForceNew) - block until the appliance is
#                             reachable again after the (rebooting) install.
#   reachable_timeout       (Optional, Computed, ForceNew) - overall timeout for
#                             the wait_until_reachable poll loop. Default "10m".
#   reachable_poll_delay    (Optional, Computed, ForceNew) - initial delay before
#                             the first reachability poll. Default "60s".
#   reachable_poll_interval (Optional, Computed, ForceNew) - interval between
#                             reachability polls. Default "60s".
#   reachable_poll_timeout  (Optional, Computed, ForceNew) - per-poll HTTP timeout
#                             for a single reachability check. Default "20s".
#
# How to validate/plan (SAFE - makes no CRUD call against the ADC):
#   export NS_URL="http://10.101.132.155/"
#   export NS_LOGIN="nsroot"
#   export NS_PASSWORD='CADS123$%^'
#   terraform validate
#   terraform plan
#
#   Do NOT run `terraform apply` unless you truly intend to upgrade the box.

terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}

# endpoint (NS_URL), username (NS_LOGIN) and password (NS_PASSWORD) are taken
# from the environment. insecure_skip_verify allows self-signed ADC certs.
provider "citrixadc" {
  insecure_skip_verify = true
}

resource "citrixadc_installer" "tf_installer" {
  # PLACEHOLDER build image URL - replace with a real, ADC-reachable build .tgz
  # before ever applying. Applying with this value would fail to download / would
  # NOT perform a valid upgrade; it exists only so validate/plan have a value.
  url = "http://build-server.example.com/builds/build-14.1-XX.YY_nc_64.tgz"

  # Build-install flags (all default to false).
  enhancedupgrade = false
  l               = false
  resizeswapvar   = false
  y               = true # do not prompt for yes/no before rebooting

  # Block until the appliance reboots and is reachable again (required).
  wait_until_reachable = true

  # Reachability poll-loop tuning (defaults shown).
  reachable_timeout       = "10m"
  reachable_poll_delay    = "60s"
  reachable_poll_interval = "60s"
  reachable_poll_timeout  = "20s"
}
