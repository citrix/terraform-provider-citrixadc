resource "citrixadc_sslcipher" "tfsslcipher" {
  ciphergroupname = "tfsslcipher"

  # ciphersuitebinding is MANDATORY attribute
  # Any change in the ciphersuitebinding will result in re-creation of the whole sslcipher resource.
  ciphersuitebinding {
    ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
    cipherpriority = 1
  }
  ciphersuitebinding {
    ciphername     = "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384"
    cipherpriority = 2
  }
  ciphersuitebinding {
    ciphername     = "TLS1.2-ECDHE-RSA-AES-128-SHA256"
    cipherpriority = 3
  }
}
