terraform {
  # Write-only attributes (filecontent_wo, passplain_wo) require Terraform 1.11+.
  required_version = ">= 1.11.0"

  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
      # Pin to a provider release that includes citrixadc_systemfile.filecontent_wo.
      # version = ">= x.y.z"
    }

    # OPTIONAL: uncomment to source the certificate material from HashiCorp Vault.
    # See README.md and the commented data source in resources.tf.
    # vault = {
    #   source = "hashicorp/vault"
    # }
  }
}

provider "citrixadc" {
  endpoint = var.adc_endpoint
  username = var.adc_username
  password = var.adc_password
}
