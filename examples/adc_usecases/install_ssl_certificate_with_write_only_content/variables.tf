variable "adc_endpoint" {
  description = "NetScaler NITRO endpoint, e.g. http://10.0.0.1 (maps to NS_URL)."
  type        = string
}

variable "adc_username" {
  description = "NetScaler user (maps to NS_LOGIN)."
  type        = string
  default     = "nsroot"
}

variable "adc_password" {
  description = "NetScaler password (maps to NS_PASSWORD). Prefer TF_VAR_adc_password."
  type        = string
  sensitive   = true
}

# --- Certificate material (Vault-style) -------------------------------------
# In production these come from a secret store (see README.md). They are marked
# sensitive and are delivered to the appliance through WRITE-ONLY attributes, so
# their values never enter Terraform state, plan output, or logs.

variable "cert_pem" {
  description = "X.509 certificate in PEM format."
  type        = string
  sensitive   = true
}

variable "key_pem" {
  description = "Private key in PEM format (the secret)."
  type        = string
  sensitive   = true
}

variable "key_passphrase" {
  description = "Passphrase for an encrypted private key. Leave empty for an unencrypted key."
  type        = string
  sensitive   = true
  default     = ""
}

# --- Rotation control -------------------------------------------------------
# Increment this (in the same apply that supplies the new cert_pem/key_pem
# values) to re-upload the files and reload the certificate on the appliance in
# place. See README.md -> "Rotating the certificate".
variable "cert_material_version" {
  description = "Rotation counter for the certificate material. Bump on every rotation."
  type        = number
  default     = 1
}
