# systemcollectionparam.tf
#
# citrixadc_systemcollectionparam configures the ADC "collection parameter"
# object (the data-collection / SNMP polling params used by the perf-collector).
#
# It is an UNNAMED SINGLETON config object: there is exactly one
# systemcollectionparam per appliance and it has NO key. Consequently the
# resource has no ADD or DELETE NITRO endpoint, and the CRUD maps to:
#
#   Create / Update -> UpdateUnnamedResource (`set systemcollectionparam`,
#                      NITRO PUT to /nitro/v1/config/systemcollectionparam).
#   Read            -> FindResource (GET systemcollectionparam), reading the
#                      datapath/loglevel params back (communityname is NOT read
#                      back by design - the SNMPv1 community name is not
#                      reliably returned, so its configured value is preserved
#                      to avoid a perpetual diff).
#   Delete          -> NO-OP. systemcollectionparam has no delete API, so
#                      Delete only drops the resource from Terraform state; the
#                      ADC object (and whatever value it currently holds) is
#                      left untouched on the box.
#
# Because Delete is a no-op, `terraform destroy` will NOT revert the loglevel
# on the appliance - it simply forgets the resource. This is intentional and
# mirrors the SDK v2 resource it replaces.
#
# The ID is a synthetic constant ("systemcollectionparam-config") since the
# object is keyless; Read/Delete never depend on its value.
#
# Attributes (per ../../custom_resources/systemcollectionparam/resource_systemcollectionparam.go
#             and the NITRO systemcollectionparam spec):
#   communityname (Optional, Computed) - SNMPv1 community name for
#                 authentication. OMITTED here (write-only in practice; not read
#                 back). Safe to set, but left out to keep the sample minimal.
#   datapath      (Optional, Computed) - path to the collection database.
#                 OMITTED - changing the data path is environment-specific and
#                 not a "safe" knob to flip on a shared lab box.
#   loglevel      (Optional, Computed) - collector log verbosity. Allowed
#                 values: CRITICAL, WARNING, INFO, DEBUG1, DEBUG2. This is the
#                 safe, self-contained knob exercised below.
#
# How to run (dev_overrides; NO `terraform init`):
#   export TF_CLI_CONFIG_FILE=/tmp/shard/dev.tfrc
#   export NS_URL="http://10.101.132.152/"
#   export NS_LOGIN="nsroot"
#   export NS_PASSWORD='CADS123$%^'
#   terraform validate
#   terraform apply -auto-approve
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

resource "citrixadc_systemcollectionparam" "tf_systemcollectionparam" {
  # Safe, self-contained param: set the collector log level to WARNING.
  # (Create/Update issue `set systemcollectionparam`; Delete is a no-op so
  # destroy will not revert this value on the appliance.)
  loglevel = "WARNING"
}
