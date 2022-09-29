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
    ciphergroupname = "External"

    ciphersuitebinding {
        ciphername     = "TLS1.2-AES256-GCM-SHA384"
        cipherpriority = 1
    }
    ciphersuitebinding {
        ciphername     = "TLS1.2-AES128-GCM-SHA256"
        cipherpriority = 2
    }
    ciphersuitebinding {
        ciphername     = "TLS1.2-AES-256-SHA256"
        cipherpriority = 3
    }
    ciphersuitebinding {
        ciphername     = "TLS1.2-AES-128-SHA256"
        cipherpriority = 4
    }
     ciphersuitebinding {
        ciphername     = "TLS1-AES-128-CBC-SHA"
        cipherpriority = 5
    }
    ciphersuitebinding {
        ciphername     = "TLS1-AES-256-CBC-SHA"
        cipherpriority = 6
    }
}

