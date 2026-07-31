# nsconfig_update.tf
#
# ##########################################################################
# ##  DANGER - THIS RESOURCE CAN MAKE THE ADC PERMANENTLY UNREACHABLE.   ##
# ##########################################################################
#
#   citrixadc_nsconfig_update issues `set ns config` (NITRO PUT to
#   /nitro/v1/config/nsconfig) on create/update. Setting `ipaddress` /
#   `netmask` CHANGES THE MANAGEMENT NSIP of the appliance. If you point it
#   at any address other than the box's OWN current NSIP, the management
#   session is severed and the box becomes UNREACHABLE (and Terraform will
#   hang/fail because the endpoint it is talking to just moved).
#   Likewise `nsvlan`/`ifnum`/`tagged` bind an NSVLAN, which can disrupt the
#   management data path.
#
#   RUN THIS ONLY:
#     * on a DISPOSABLE / lab appliance you are willing to lose, and
#     * with `ipaddress` set to that SAME box's CURRENT NSIP and
#       `netmask` set to that box's CURRENT netmask.
#   Under those conditions the `set ns config` is effectively a no-op and
#   CANNOT disconnect the box.
#
#   The values below are the CURRENT NSIP/netmask of the designated
#   disposable box .155 (read live via
#   `GET /nitro/v1/config/nsip?filter=type:NSIP`). Do NOT copy them to any
#   other appliance. If you point this at a different box, first replace them
#   with THAT box's own current NSIP/netmask.
#
# What this does:
#   Create/Update -> UpdateUnnamedResource (`set ns config`, PUT).
#   Read          -> FindResource (GET nsconfig), reads settable params back.
#   Delete        -> no-op (nsconfig has no delete API; state is just dropped;
#                    mirrors the SDK v2 schema.Noop delete).
#
# Attributes (per ../../nsconfig/resource_nsconfig_update.go and
#             ../../../nitro_rest/ns/nsconfig.html):
#   ipaddress (Optional, Computed) - the NSIP of the appliance. SEE DANGER ABOVE.
#   netmask   (Optional, Computed) - netmask for the NSIP.
#   nsvlan    (Optional, Computed) - VLAN id for the NSIP subnet. OMITTED here.
#   ifnum     (Optional, Set of string) - interfaces bound to the NSVLAN. OMITTED.
#   tagged    (Optional, Computed) - add interfaces as 802.1q tagged. OMITTED.
#
# How to run (ONLY on the disposable box .155):
#   export NS_URL="http://10.101.132.155/"
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

resource "citrixadc_nsconfig_update" "foo" {
  # DANGER: these MUST be the running box's OWN current NSIP + netmask so the
  # `set ns config` is a no-op that cannot disconnect the appliance.
  # Values below are the live current NSIP/netmask of disposable box .155.
  ipaddress = "10.101.132.155"
  netmask   = "255.255.255.0"

  # nsvlan / ifnum / tagged intentionally OMITTED - binding an NSVLAN could
  # disrupt the management path. Do not set them on a box you need to keep.
}
