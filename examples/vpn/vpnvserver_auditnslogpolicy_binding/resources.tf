# Since the auditnslogpolicy resource is not yet available on Terraform,
# the tf_auditnslogpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add audit nslogAction tf_auditnslogaction 1.1.1.1 -loglevel NONE
# add audit nslogPolicy tf_auditnslogpolicy ns_true tf_auditnslogaction

resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vpnvserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_auditnslogpolicy_binding" "tf_bind" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = "tf_auditnslogpolicy"
  bindpoint = "REQUEST"
  priority  = 200
}
