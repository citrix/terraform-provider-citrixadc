# Sample configuration for the citrixadc_opoption resource.
#
# NOTE: citrixadc_opoption is an ALIAS of citrixadc_snmpoption. It manages the
# same NITRO object (snmp/snmpoption), which is an UNNAMED SINGLETON global SNMP
# option object: there is no name/key. Terraform "creates"/"updates" it with an
# HTTP PUT (UpdateUnnamedResource) and "deletes" it by only dropping it from
# Terraform state (there is no NITRO DELETE for the singleton, so the appliance
# configuration is left untouched on destroy).
#
# All attributes are Optional + Computed: any option you omit is populated by the
# appliance (avoiding perpetual diffs), and any you set is applied.

terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}

provider "citrixadc" {
  # Endpoint and credentials are supplied via NS_URL / NS_LOGIN / NS_PASSWORD.
}

resource "citrixadc_opoption" "tf_opoption" {
  # Log SNMP trap events even when no trap listeners are configured.
  snmptraplogging = "ENABLED"

  # Audit log level of SNMP trap logs (valid ADC enum value).
  snmptraplogginglevel = "INFORMATIONAL"
}
