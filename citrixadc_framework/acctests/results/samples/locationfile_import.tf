# locationfile_import.tf
#
# What this does:
#   citrixadc_locationfile_import is an ACTION-ONLY resource. On create it issues
#   POST /nitro/v1/config/locationfile?action=import, which copies a static GeoIP
#   location database file from a source (`src`) into the appliance's location DB
#   directory (/var/netscaler/locdb) under the chosen `locationfile` name.
#
#   It mirrors the SDK v2 citrixadc_locationfile_import resource exactly:
#     - Create  -> POST locationfile?action=import  (payload: Locationfile + src)
#     - Read    -> no-op  (the import action has no GET-backed object keyed by ID)
#     - Update  -> no-op  (every attribute is ForceNew/RequiresReplace)
#     - Delete  -> no-op  (NITRO has no inverse of the import action)
#   The Terraform ID is a synthetic per-apply handle ("tf-locationfile-<n>"),
#   identical to the SDK v2 scheme; there is no queryable managed object.
#
#   NOTE: `format` is present in the schema for backward compatibility but, like
#   SDK v2, is intentionally NOT sent in the import payload (the NITRO import
#   action only consumes Locationfile + src). Loading/parsing of the file's
#   format happens at "add locationfile" time, not at import time.
#
# Attributes (per ../locationfile/resource_locationfile_import.go and
#             ../../nitro_rest/basic/locationfile.html):
#   locationfile (Required, ForceNew) - name of the location file (with or without
#                   an absolute path). If no path is given, /var/netscaler/locdb
#                   is assumed. This is the destination name on the appliance.
#   src          (Required, ForceNew) - URL (protocol, host, path, file name) the
#                   file is imported from. Supported schemes include
#                   http://, https://, ftp:// and local:// . The local:// scheme
#                   refers to a file already present in /var/tmp on the appliance
#                   (e.g. one uploaded via citrixadc_systemfile or the GUI).
#   format       (Optional, ForceNew) - format used to read the file. Values:
#                   netscaler (default), ip-country, ip-country-isp,
#                   ip-country-region-city, ip-country-region-city-isp,
#                   geoip-country, geoip-region, geoip-city, geoip-country-org,
#                   geoip-country-isp, geoip-city-isp-org.
#
#   Every attribute is ForceNew (RequiresReplace): any change re-imports the file.
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

# Companion upload: place a small netscaler-format location DB file in /var/tmp so
# the import below has a genuinely-available local:// source (no external download
# required). In production your `src` would typically be an http(s)/ftp URL to a
# real GeoIP database, or a file you have already uploaded to the appliance.
resource "citrixadc_systemfile" "tf_locationfile_src" {
  filename     = "tf_locationfile_src"
  filelocation = "/var/tmp"

  # netscaler-format location DB: "fromIP-toIP","geographical.qualifiers"
  filecontent = <<-EOT
    "1.0.0.0-1.0.0.255","North America.United States.California.San Jose"
    "2.0.0.0-2.0.0.255","Europe.United Kingdom.England.London"
  EOT
}

resource "citrixadc_locationfile_import" "tf_locationfile_import" {
  locationfile = "tf_locationfile_import"
  src          = "local://tf_locationfile_src"
  format       = "netscaler"

  # import must run after the source file is present in /var/tmp.
  depends_on = [citrixadc_systemfile.tf_locationfile_src]
}
