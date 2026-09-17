locals {
  ssl_dir       = "/nsconfig/ssl"
  cert_filename = "tf_ephemeral.crt"
  key_filename  = "tf_ephemeral.key"
}

# ---------------------------------------------------------------------------
# OPTIONAL: source the material from HashiCorp Vault instead of variables.
# Uncomment the vault provider in provider.tf, then feed the data source outputs
# into the resources below (in place of var.cert_pem / var.key_pem), and use the
# secret's version as cert_material_version so rotation is fully automatic:
#
#   data "vault_kv_secret_v2" "tls" {
#     mount = "secret"
#     name  = "netscaler/tls"
#   }
#
#   # then, below:
#   #   filecontent_wo         = data.vault_kv_secret_v2.tls.data["certificate"]
#   #   filecontent_wo_version = data.vault_kv_secret_v2.tls.metadata.version
# ---------------------------------------------------------------------------

# 1) Upload the certificate file. filecontent_wo keeps the content OUT of state.
resource "citrixadc_systemfile" "cert" {
  filename               = local.cert_filename
  filelocation           = local.ssl_dir
  filecontent_wo         = var.cert_pem
  filecontent_wo_version = var.cert_material_version
}

# 2) Upload the private key file (the secret) via the write-only attribute.
resource "citrixadc_systemfile" "key" {
  filename               = local.key_filename
  filelocation           = local.ssl_dir
  filecontent_wo         = var.key_pem
  filecontent_wo_version = var.cert_material_version
}

# 3) Install the cert/key pair and reload it gracefully on rotation.
#
#    cert_hash / key_hash are change-detection tokens (NOT the secret), so they
#    can live in state. We source them from the READABLE filecontent_wo_version
#    (the write-only filecontent_wo itself cannot be referenced in expressions),
#    so a single cert_material_version bump both re-uploads the files AND makes
#    sslcertkey reload the certificate in place. The references also guarantee
#    the files are re-uploaded before the reload.
resource "citrixadc_sslcertkey" "kp" {
  certkey = "tf_ephemeral_kp"
  cert    = format("%s/%s", citrixadc_systemfile.cert.filelocation, citrixadc_systemfile.cert.filename)
  key     = format("%s/%s", citrixadc_systemfile.key.filelocation, citrixadc_systemfile.key.filename)
  inform  = "PEM"

  # Passphrase for an encrypted key, also delivered write-only (never in state).
  # Both are null (unset) when key_passphrase is empty, i.e. an unencrypted key.
  password             = var.key_passphrase != "" ? true : null
  passplain_wo         = var.key_passphrase != "" ? var.key_passphrase : null
  passplain_wo_version = var.cert_material_version

  # Graceful in-place reload trigger (see comment above).
  cert_hash = tostring(citrixadc_systemfile.cert.filecontent_wo_version)
  key_hash  = tostring(citrixadc_systemfile.key.filecontent_wo_version)
}
