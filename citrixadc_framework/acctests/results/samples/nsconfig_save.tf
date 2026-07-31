# nsconfig_save.tf
#
# What this does:
#   citrixadc_nsconfig_save is an ACTION-ONLY resource. On create it issues a
#   single NITRO call: POST /nitro/v1/config/nsconfig?action=save. This persists
#   the current RUNNING configuration to the STARTUP configuration (the classic
#   "save ns config"). It is SAFE: it only saves; it does NOT clear, reset, or
#   overwrite the running config. There is no GET/update inverse for the save
#   action, so Read/Update are no-ops; `timestamp` is a provider-supplied value
#   used only as the Terraform resource ID. Bump `timestamp` to re-run the save
#   (all attributes are ForceNew / RequiresReplace).
#
#   Delete: by default removing the resource just drops it from state (no ADC
#   call). If save_on_destroy = true, destroy performs one more "save ns config".
#
# Attributes (per resource_nsconfig_save.go / nitro_rest/ns/nsconfig.html):
#   timestamp               (Required, ForceNew) - marker string; becomes the ID.
#   all                     (Optional, Computed, ForceNew, default false) -
#                             saveconfig for all partitions (NITRO `all`). Newer
#                             ADC builds only; omit on 12.0.
#
#   Provider-side-only knobs (no NITRO field; control the concurrent-save retry
#   loop that tolerates errorcode 293 "another save in progress"):
#   concurrent_save_ok      (Optional, Computed, ForceNew, default true) -
#                             tolerate a concurrent save (293) by retrying.
#   concurrent_save_retries (Optional, Computed, ForceNew, default 0) - number of
#                             retry attempts when a concurrent save is in progress.
#                             0 = tolerate 293 without retrying (treated as saved).
#   concurrent_save_timeout (Optional, Computed, ForceNew, default "5m") - total
#                             Go-duration budget for the retry loop.
#   concurrent_save_interval(Optional, Computed, ForceNew, default "10s") - Go
#                             duration between retries.
#   save_on_destroy         (Optional, Computed, ForceNew, default false) - also
#                             save ns config when the resource is destroyed.
#
# How to run:
#   export NS_URL="http://10.101.132.153/"   # any reachable ADC (standalone/HA/cluster)
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

resource "citrixadc_nsconfig_save" "foo" {
  # Bump this marker to force a fresh "save ns config".
  timestamp = "2020-03-24T12:37:06Z"

  # all = true  # only on ADC builds that support saveconfig for all partitions

  # Concurrent-save handling: tolerate a racing save (NITRO errorcode 293) and
  # retry up to 3 times, 10s apart, within a 5m budget.
  concurrent_save_ok       = true
  concurrent_save_retries  = 3
  concurrent_save_interval = "10s"
  concurrent_save_timeout  = "5m"

  # Leave startup config saved on destroy (default just drops from TF state).
  save_on_destroy = false
}
