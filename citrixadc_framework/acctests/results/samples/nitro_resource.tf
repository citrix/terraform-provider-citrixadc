# nitro_resource.tf
#
# What this does:
#   citrixadc_nitro_resource is the GENERIC NITRO PASSTHROUGH resource. Unlike
#   every other citrixadc_* resource, it has NO fixed schema for the managed
#   object. Instead its behaviour is driven entirely by:
#     - workflows_file : path to a YAML file describing available NITRO workflows
#     - workflow       : the key (inside that YAML) selecting one workflow
#     - attributes / non_updateable_attributes : string->string maps carrying the
#       actual NITRO object fields.
#   The selected workflow's `lifecycle` picks the CRUD flavour:
#     - object / non_updateable_object -> named-object CRUD
#       (AddResource / FindResourceArrayWithParams / UpdateResource / DeleteResource)
#     - binding                        -> binding CRUD
#       (UpdateResource add / Find read / DeleteResourceWithArgs remove)
#     - object_by_args                 -> unnamed/args-keyed object CRUD
#       (AddResource / Find-by-args / UpdateUnnamedResource / DeleteResourceWithArgsMap)
#
#   THIS SAMPLE manages a BENIGN object: an "nsacl" (Extended ACL) via the
#   `object` lifecycle. An nsacl that is merely CONFIGURED (added to the config
#   but never "apply ns acls"-ed) has kernelstate=NOTAPPLIED and therefore has
#   ZERO effect on live traffic on the ADC. This sample never applies the ACL, so
#   creating/destroying it cannot disrupt the target box. Destroy issues a plain
#   NITRO DELETE /config/nsacl/<aclname>, fully cleaning it up.
#
#   NITRO call sequence exercised here (object lifecycle):
#     Create -> POST   /nitro/v1/config/nsacl            (aclname+aclaction+priority)
#     Read   -> GET    /nitro/v1/config/nsacl/tf_nitro_sample
#     Update -> PUT    /nitro/v1/config/nsacl/tf_nitro_sample  (only `attributes`)
#     Delete -> DELETE /nitro/v1/config/nsacl/tf_nitro_sample
#
# Attribute placement (per resource_nitro_resource.go):
#   - `attributes`               : the UPDATEABLE fields. Sent on both create and
#                                  update. Change them -> in-place PUT.
#   - `non_updateable_attributes`: ForceNew fields (RequiresReplace on change).
#                                  Sent on create (merged with attributes). The
#                                  primary_id_attribute (aclname) lives here.
#   NOTE: the workflow YAML's own `non_updateable_attributes:` list and
#   `allow_recreate:` keys are metadata only; the Go resource does not read them.
#   What is updateable vs ForceNew is decided purely by which of the two TF maps
#   you put each field in.
#
# Files required together (this sample is self-contained):
#   nitro_resource.tf                 (this file)
#   nitro_resource_workflows.yaml     (the referenced workflows_file)
#   Both must sit in the same working directory so that
#   workflows_file = "nitro_resource_workflows.yaml" resolves relative to CWD.
#
# How to run (dev_overrides; no `terraform init`):
#   export TF_CLI_CONFIG_FILE=/tmp/shard/dev.tfrc
#   export NS_URL="http://10.101.132.153/"
#   export NS_LOGIN="nsroot"
#   export NS_PASSWORD='CADS123$%^'
#   terraform validate
#   terraform plan
#   terraform apply -auto-approve
#   terraform destroy -auto-approve

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

# Manages a benign, NEVER-applied nsacl "tf_nitro_sample" through the generic
# nitro_resource passthrough using the `object` lifecycle (workflow key "nsacl").
resource "citrixadc_nitro_resource" "tf_nsacl" {
  workflows_file = "nitro_resource_workflows.yaml"
  workflow       = "nsacl"

  # UPDATEABLE fields (in-place PUT on change).
  attributes = {
    priority = "10"
  }

  # FORCE-NEW fields (change => destroy+recreate). aclname is the primary id.
  # aclaction=ALLOW on an un-applied ACL is inert (kernelstate=NOTAPPLIED).
  non_updateable_attributes = {
    aclname   = "tf_nitro_sample"
    aclaction = "ALLOW"
  }
}
