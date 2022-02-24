# Since the analyticsprofile resource is not yet available on Terraform,
# the new_profile profile must be created by hand(manually) in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add analyticsprofile new_profile -type tcpinsight
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_analyticsprofile_binding" "tf_bind" {
  name             = citrixadc_vpnvserver.tf_vpnvserver.name
  analyticsprofile = "new_profile"
}
