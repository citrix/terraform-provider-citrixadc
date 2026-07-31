# systembackup_create.tf
#
# What this does:
#   citrixadc_systembackup_create is an ACTION-with-read-delete resource. On
#   create it issues POST /nitro/v1/config/systembackup?action=create, which
#   writes a persistent backup archive <filename>.tgz onto the appliance. Read
#   looks the archive up via GET /nitro/v1/config/systembackup/<filename>.tgz
#   (clearing state when the file is gone), and Delete removes that archive
#   (rm systembackup). This mirrors the SDK v2 citrixadc_systembackup_create
#   resource exactly.
#
#   SAFE: it only creates a backup file and later deletes that same file; it does
#   not restore or otherwise mutate the running configuration. Works on a
#   standalone ADC (no cluster/HA gating).
#
# Attributes (per resource_systembackup_create.go / nitro_rest/system/systembackup.html):
#   filename         (Required, ForceNew) - base name of the backup file; the
#                       appliance stores it as <filename>.tgz. Becomes part of
#                       the resource ID.
#   level            (Optional, ForceNew) - amount of data to back up. Values:
#                       basic, full.
#   includekernel    (Optional, ForceNew) - add the kernel to the backup file.
#                       Values: NO, YES.
#   uselocaltimezone (Optional, ForceNew) - stamp the backup with the appliance
#                       local timezone instead of UTC.
#   comment          (Optional, ForceNew) - free-text comment stored with the
#                       backup file.
#
#   Every attribute is ForceNew (RequiresReplace): any change destroys and
#   recreates the backup archive.
#
# How to run:
#   export NS_URL="http://10.101.132.151/"     # any standalone ADC
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

resource "citrixadc_systembackup_create" "tf_systembackup_create" {
  filename         = "my_backup_file"
  level            = "basic"
  uselocaltimezone = "true"
  comment          = "nightly config backup created by terraform"
}
