# Since the auditnslogpolicy resource is not yet available on Terraform,
# the tf_auditnslogpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add audit nslogAction tf_auditnslogaction 1.1.1.1 -loglevel NONE
# add audit nslogPolicy tf_auditnslogpolicy ns_true tf_auditnslogaction

resource "citrixadc_aaagroup_auditnslogpolicy_binding" "tf_aaagroup_auditnslogpolicy_binding" {
  groupname = "my_group"
  policy    = "tf_auditnslogpolicy"
  priority  = 150
}
