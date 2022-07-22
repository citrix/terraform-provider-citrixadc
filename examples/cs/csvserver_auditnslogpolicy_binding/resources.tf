# Since the auditnslogpolicy resource is not yet available on Terraform,
# the tf_auditnslogpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add audit nslogAction tf_auditnslogaction 1.1.1.1 -loglevel NONE
# add audit nslogPolicy tf_auditnslogpolicy ns_true tf_auditnslogaction

resource "citrixadc_csvserver_auditnslogpolicy_binding" "tf_csvserver_auditnslogpolicy_binding" {
  name = citrixadc_csvserver.tf_csvserver.name
  policyname = "tf_auditnslogpolicy"
  priority = 5
}

resource "citrixadc_csvserver" "tf_csvserver" {
  name = "tf_csvserver"
  ipv46 = "10.202.11.11"
  port = 8080
  servicetype = "HTTP"
}