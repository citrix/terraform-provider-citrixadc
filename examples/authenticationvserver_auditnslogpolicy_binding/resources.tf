# Since the auditnslogpolicy resource is not yet available on Terraform,
# the tf_auditnslogpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add audit nslogAction tf_auditnslogaction 1.1.1.1 -loglevel NONE
# add audit nslogPolicy tf_auditnslogpolicy ns_true tf_auditnslogaction
resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_authenticationvserver_auditnslogpolicy_binding" "tf_bind" {
  name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy    = "tf_auditnslogpolicy"
  priority  = 90
  bindpoint = "AAA_RESPONSE"
}