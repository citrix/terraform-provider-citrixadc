# systembackup_restore.tf
#
# ############################################################################
# # ⚠  DANGER — DESTRUCTIVE OPERATION  ⚠
# #
# # citrixadc_systembackup_restore OVERWRITES the appliance's running
# # configuration with the contents of a previously-created backup archive
# # (<filename>.tgz). The restore action can roll the box back to an older
# # config, drop entities you created after the backup, and typically triggers
# # a service reload. THERE IS NO INVERSE / UNDO of the restore action.
# #
# # ONLY apply this against a DISPOSABLE / throw-away ADC whose current config
# # you are willing to lose (e.g. the .155 test box). NEVER point this at a
# # production or shared appliance.
# ############################################################################
#
# What this does:
#   citrixadc_systembackup_restore is an ACTION-ONLY resource. On create it
#   issues a single NITRO call:
#       POST /nitro/v1/config/systembackup?action=restore
#       payload: {"systembackup":{"filename":"<file>.tgz","skipbackup":<bool>}}
#   which restores the appliance from the named backup archive. NITRO exposes no
#   restore-state GET endpoint and no inverse API, so Read/Update/Delete are
#   no-ops. Both arguments are ForceNew (RequiresReplace): any change re-fires
#   the restore action. `id` is a provider-generated marker used only as the
#   Terraform resource ID. This mirrors the SDK v2 citrixadc_systembackup_restore
#   resource exactly.
#
#   This sample first CREATES a backup (citrixadc_systembackup_create), then
#   RESTORES from it. The restore references the create's filename with a ".tgz"
#   suffix because the appliance stores the archive as <filename>.tgz. The
#   depends_on makes the ordering explicit: the archive must exist before restore.
#
# Attributes (per resource_systembackup_restore.go / nitro_rest/system/systembackup.html):
#   filename   (Required, ForceNew) - name of the backup file (*.tgz) to restore.
#   skipbackup (Optional, ForceNew) - skip taking a safety backup during the
#                 restore operation (true/false).
#
# How to run:
#   ############################################################################
#   # ⚠  Use ONLY a disposable ADC. This WILL overwrite its running config.  ⚠
#   ############################################################################
#   export NS_URL="http://10.101.132.155/"     # DISPOSABLE standalone ADC ONLY
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

# Step 1: produce a backup archive named my_backup_file.tgz on the appliance.
resource "citrixadc_systembackup_create" "tf_systembackup_create" {
  filename         = "my_backup_file"
  level            = "basic"
  uselocaltimezone = "true"
}

# Step 2: DESTRUCTIVE — restore the appliance from that archive.
#         This overwrites the running configuration. Disposable box only.
resource "citrixadc_systembackup_restore" "tf_systembackup_restore" {
  filename   = "${citrixadc_systembackup_create.tf_systembackup_create.filename}.tgz"
  skipbackup = "false"

  depends_on = [citrixadc_systembackup_create.tf_systembackup_create]
}
