# Since the aaapreauthenticationpolicy resource is not yet available on Terraform,
# the tf_aaapolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add aaa preauthenticationaction tf_aaaaction DENY
# add aaa preauthenticationpolicy tf_aaapolicy NS_TRUE tf_aaaaction
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vpnvserverexample"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_aaapreauthenticationpolicy_binding" "tf_binding" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = "tf_aaapolicy"
  priority  = 40
  secondary = "false"
  bindpoint = "OTHERTCP_REQUEST"
}