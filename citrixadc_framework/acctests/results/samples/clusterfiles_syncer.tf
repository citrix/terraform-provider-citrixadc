# clusterfiles_syncer.tf
#
# What this does:
#   citrixadc_clusterfiles_syncer is an ACTION-ONLY resource. On create it issues
#   a single NITRO call: POST /nitro/v1/config/clusterfiles?action=sync with a
#   payload of {"clusterfiles":{"mode":[...]}}. This triggers the Citrix ADC
#   cluster to synchronize the listed directories/files from the configuration
#   coordinator to the other cluster nodes. There is no GET/update/delete inverse
#   for the sync action, so Read/Update/Delete are no-ops; `timestamp` is a
#   provider-supplied value used only as the Terraform resource ID.
#
#   NOTE: This resource is CLUSTER-only. It must be applied against a cluster CLIP
#   (e.g. 10.101.132.133) with active nodes; it is not meaningful on a standalone
#   or HA-pair ADC.
#
# Attributes (per resource_clusterfiles_syncer.go / nitro_rest/cluster/clusterfiles.html):
#   timestamp (Required, ForceNew) - arbitrary marker string; becomes the resource ID.
#   mode      (Required, ForceNew) - set of directories/files to sync.
#                Possible values = all, bookmarks, ssl, imports, misc, dns, krb,
#                AAA, app_catalog, all_plus_misc, all_minus_misc
#
# How to run:
#   export NS_URL="http://10.101.132.133/"     # cluster CLIP
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

resource "citrixadc_clusterfiles_syncer" "syncer" {
  timestamp = "2020-03-24T12:37:06Z"
  mode      = ["all", "misc"]
}
