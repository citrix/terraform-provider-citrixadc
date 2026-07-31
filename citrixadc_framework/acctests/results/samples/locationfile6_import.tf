# locationfile6_import.tf
#
# What this does:
#   citrixadc_locationfile6_import is an ACTION-ONLY resource. On create it issues
#   a single NITRO call: POST /nitro/v1/config/locationfile6?action=import with a
#   payload of {"locationfile6":{"Locationfile":<name>,"src":<url>}}. This imports
#   a static IPv6 GeoIP location database from `src` into the ADC's location DB
#   directory (/var/netscaler/locdb) under the name `locationfile`. There is no
#   GET/update/delete inverse for the import action, so Read/Update/Delete are
#   no-ops and every attribute is ForceNew; the resource ID is a provider-supplied
#   unique handle ("tf-locationfile6-<nanos>").
#
#   `src` protocols: the import action accepts http, https, ftp, and local://
#   sources (file:// is rejected with NITRO errorcode 3234). The local:// scheme
#   resolves to /var/tmp on the appliance, so you can import a file that is already
#   on the box without any external download by first placing it in /var/tmp.
#   The two-resource pattern below does exactly that: it uploads a small netscaler6
#   IPv6 location DB to /var/tmp via the citrixadc_systemfile resource, then imports
#   it via src = "local://<filename>". depends_on guarantees ordering.
#
#   NOTE (rerun/cleanup): the import has no inverse. `terraform destroy` removes the
#   /var/tmp source (the systemfile), but the imported copy in /var/netscaler/locdb
#   and its entry in the import index (/var/netscaler/locdb/mapping-spdbfile) remain.
#   Re-importing the same `locationfile` name while that index entry exists fails
#   with errorcode 3198 ("Object already exists"), so remove the locdb copy and the
#   mapping-spdbfile index (via NITRO) between runs if you reuse the same name.
#
# Attributes (per resource_locationfile6_import.go / nitro_rest/basic/locationfile6.html):
#   locationfile (Required, ForceNew) - destination name of the IPv6 location file
#                (with or without absolute path; default path /var/netscaler/locdb).
#   src          (Required, ForceNew) - http/https/ftp/local:// URL to import from.
#   format       (Optional, ForceNew) - file format; possible values:
#                netscaler6, geoip-country6 (appliance default: netscaler6).
#                NOTE: format is a schema attribute but is NOT sent in the import
#                payload (mirrors the SDK v2 resource, which only sent
#                Locationfile + src); the import action rejects a `format` argument
#                with errorcode 278.
#
# How to run:
#   export NS_URL="http://10.101.132.152/"     # any standalone ADC
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

# Stage a small valid netscaler6-format IPv6 location DB in /var/tmp so the import
# has a real, on-box source (no external download). filecontent is raw text; the
# systemfile resource base64-encodes it (fileencoding defaults to BASE64).
resource "citrixadc_systemfile" "tf_locationfile6_src" {
  filename     = "tf_locationfile6_src"
  filelocation = "/var/tmp"
  filecontent  = "\"2001:db8::-2001:db8::ffff\",\"North America.United States.California.San Jose\"\n\"2001:db9::-2001:db9::ffff\",\"Europe.United Kingdom.England.London\"\n"
}

resource "citrixadc_locationfile6_import" "tf_locationfile6_import" {
  locationfile = "tf_locationfile6_import"
  src          = "local://tf_locationfile6_src" # local:// resolves to /var/tmp
  format       = "netscaler6"
  depends_on   = [citrixadc_systemfile.tf_locationfile6_src]
}
