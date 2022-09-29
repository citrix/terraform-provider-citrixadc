terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}
terraform {
  required_version = ">= 0.12"
}
resource "citrixadc_sslcipher" "tf_sslcipher" {
    ciphergroupname = "Internal"

    ciphersuitebinding {
        ciphername     = "SSL3-RC4-SHA"
        cipherpriority = 1
    }
    ciphersuitebinding {
        ciphername     = "SSL3-EXP-RC4-MD5"
        cipherpriority = 2
    }
    ciphersuitebinding {
        ciphername     = "TLS1-AES-128-CBC-SHA"
        cipherpriority = 3
    }
    ciphersuitebinding {
        ciphername     = "TLS1-AES-256-CBC-SHA"
        cipherpriority = 4
    }
     ciphersuitebinding {
        ciphername     = "SSL3-DES-CBC3-SHA"
        cipherpriority = 5
    }
    ciphersuitebinding {
        ciphername     = "TLS1.2-AES-256-SHA256"
        cipherpriority = 6
    }
    ciphersuitebinding {
        ciphername     = "TLS1.2-AES-128-SHA256"
        cipherpriority = 7
    }
}

