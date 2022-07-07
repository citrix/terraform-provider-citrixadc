data "citrixadc_nitro_info" "sslcipher_info" {
    workflow = yamldecode(file("../workflows/sslcipher.yaml"))
    primary_id = "tf_sslcipher"
}

output "nitro_object" {
  value = data.citrixadc_nitro_info.sslcipher_info.nitro_object
}

output "nitro_object_length" {
  value = length(data.citrixadc_nitro_info.sslcipher_info.nitro_object)
}

resource "citrixadc_sslcipher" "tfsslcipher" {
  ciphergroupname = "tf_sslcipher"

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
