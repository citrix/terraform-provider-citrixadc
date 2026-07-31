# pinger.tf
#
# What this does:
#   citrixadc_pinger is a ONE-SHOT ACTION resource: it makes the ADC run the NITRO
#   `ping` utility action against a target address and records a synthetic handle as
#   the Terraform resource ID. Its scope is limited to firing a single ping from the
#   appliance - it creates no persistent configuration object on the ADC.
#     - Create fires the ping via NITRO POST /nitro/v1/config/ping?action=ping
#       (ActOnResource("ping", ...)), then stores a generated tf-pinger-* value as
#       the resource ID.
#     - Read is a no-op (state-preserving): NITRO has no GET endpoint for a ping.
#     - Update is a no-op: every input attribute is ForceNew (RequiresReplace), so a
#       changed value re-runs the ping through resource REPLACEMENT, never Update.
#     - Delete is a no-op: there is no inverse of a ping; it only drops the resource
#       from Terraform state (no ADC-side change).
#
#   Pinging is harmless - it sends ICMP echo requests only. This sample pings the
#   ADC's own loopback (127.0.0.1), which is always reachable, so apply succeeds.
#
# Attributes (per resource_pinger.go / citrixadc/resource_citrixadc_pinger.go):
#   hostname (Optional, Computed, ForceNew) - address of the host to ping.
#   c        (Optional, Computed, ForceNew) - number of packets to send.
#   i        (Optional, Computed, ForceNew) - waiting time, in seconds (default 1).
#   n        (Optional, Computed, ForceNew) - numeric output only; no name resolution.
#   p        (Optional, Computed, ForceNew) - pattern to fill packets (up to 16 bytes).
#   q        (Optional, Computed, ForceNew) - quiet output; only the summary is printed.
#   s        (Optional, Computed, ForceNew) - data size, in bytes (default 56).
#   t        (Optional, Computed, ForceNew) - traffic domain id.
#   forcenew_id_set (Optional, Computed, ForceNew) - helper set; changing it forces a
#                                                    new ping. NOT sent to NITRO.
#
# How to run:
#   export NS_URL="http://10.101.132.154/"     # a standalone test ADC
#   export NS_LOGIN="nsroot"
#   export NS_PASSWORD='CADS123$%^'
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

resource "citrixadc_pinger" "pinger_example" {
  # Ping the ADC's own loopback - a safe, always-reachable target.
  hostname = "127.0.0.1"

  # Send a bounded number of packets so the action finishes quickly.
  c = 2

  # Numeric output only (no DNS lookup) and quiet summary output.
  n = true
  q = true
}
