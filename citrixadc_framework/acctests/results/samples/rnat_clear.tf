# rnat_clear.tf
#
# What this does:
#   citrixadc_rnat_clear manages a *set* of RNAT (Reverse NAT) route rules under a
#   single synthetic handle (rnatsname). Its scope is RNAT configuration ONLY - it
#   does not touch anything else on the appliance.
#     - Create applies every rule in the `rnat` set with a NITRO PUT
#       /nitro/v1/config/rnat (UpdateUnnamedResource), then stores `rnatsname`
#       (or a generated tf-rnat-* value) as the Terraform resource ID.
#     - Update diffs the old/new sets: rules that were removed are cleared with
#       POST /nitro/v1/config/rnat?action=clear, added rules are (re)applied.
#     - Delete clears every rule in state with POST /nitro/v1/config/rnat?action=clear.
#   Read is a no-op (state-preserving), mirroring the legacy SDK v2 behavior.
#
#   The `action=clear` call scope is limited to the RNAT rule described by its
#   payload keys (network/netmask/aclname/td/ownergroup) - it clears the matching
#   RNAT configuration, not the whole box.
#
# Attributes (per resource_rnat_clear.go / nitro_rest/network/rnat.html):
#   rnatsname (Optional, Computed) - name handle for this set of RNAT rules;
#                                    if omitted a unique tf-rnat-* value is generated.
#                                    Becomes the resource ID.
#   rnat      (Optional, Computed) - set of RNAT rules to apply. Each element:
#       network      (Optional) - network address for the RNAT entry (e.g. 192.168.96.0)
#       netmask      (Optional) - subnet mask for the network address (e.g. 255.255.240.0)
#       natip        (Optional) - a NetScaler-owned IPv4 address (except the NSIP)
#                                 used to replace the source IP of server packets.
#       aclname      (Optional) - an extended ACL defined for the RNAT entry.
#       td           (Optional) - traffic-domain id (0-4094) the entry belongs to.
#       natip2       (Optional) - provider-side only; NOT part of the NITRO rnat object.
#       redirectport (Optional) - kept for backward compatibility; a bool in this
#                                 legacy schema and NOT sent to NITRO.
#
# How to run:
#   export NS_URL="http://10.101.132.154/"     # a standalone test ADC
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

resource "citrixadc_rnat_clear" "rnat_clear_example" {
  rnatsname = "tf_rnat_clear"

  # A route-based RNAT rule (network + netmask). Add natip to have the ADC
  # replace server-generated source IPs with a NetScaler-owned address.
  rnat = [
    {
      network = "192.168.96.0"
      netmask = "255.255.240.0"
      # natip = "192.0.2.10"   # uncomment with a real NetScaler-owned IP
      # aclname = "my_rnat_acl"
      # td      = 0
    },
  ]
}
